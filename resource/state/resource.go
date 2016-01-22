// Copyright 2015 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package state

import (
	"fmt"
	"io"
	"path"

	"github.com/juju/errors"
	"github.com/juju/names"
	charmresource "gopkg.in/juju/charm.v6-unstable/resource"

	"github.com/juju/juju/resource"
)

type resourcePersistence interface {
	// ListResources returns the resource data for the given service ID.
	ListResources(serviceID string) ([]resource.Resource, error)

	// StageResource adds the resource in a separate staging area
	// if the resource isn't already staged. If the resource already
	// exists then it is treated as unavailable as long as the new one
	// is staged.
	//
	// A separate staging area is necessary because we are dealing with
	// the DB and storage at the same time for the same resource in some
	// operations (e.g. SetResource).  Resources are staged in the DB,
	// added to storage, and then finalized in the DB.
	StageResource(id, serviceID string, res resource.Resource) error

	// UnstageResource ensures that the resource is removed
	// from the staging area. If it isn't in the staging area
	// then this is a noop.
	UnstageResource(id, serviceID string) error

	// SetResource stores the resource info. If the resource
	// is already staged then it is unstaged.
	SetResource(id, serviceID string, res resource.Resource) error

	// SetUnitResource stores the resource info for a unit.
	SetUnitResource(serviceID, unitID string, res resource.Resource) error
}

type resourceStorage interface {
	// PutAndCheckHash stores the content of the reader into the storage.
	PutAndCheckHash(path string, r io.Reader, length int64, hash string) error

	// Remove removes the identified data from the storage.
	Remove(path string) error

	// Get returns a reader for the resource at path. The size of the
	// data is also returned.
	Get(path string) (io.ReadCloser, int64, error)
}

type resourceState struct {
	persist resourcePersistence
	storage resourceStorage
}

// ListResources returns the resource data for the given service ID.
func (st resourceState) ListResources(serviceID string) ([]resource.Resource, error) {
	resources, err := st.persist.ListResources(serviceID)
	if err != nil {
		return nil, errors.Trace(err)
	}

	return resources, nil
}

// GetResource returns the resource data for the identified resource.
func (st resourceState) GetResource(serviceID, name string) (resource.Resource, error) {
	var res resource.Resource

	resources, err := st.ListResources(serviceID)
	if err != nil {
		return res, errors.Trace(err)
	}

	for _, res := range resources {
		if res.Name == name {
			return res, nil
		}
	}
	return res, errors.NotFoundf("resource %q", name)
}

// TODO(ericsnow) Separate setting the metadata from storing the blob?

// SetResource stores the resource in the Juju model.
func (st resourceState) SetResource(serviceID string, res resource.Resource, r io.Reader) (err error) {
	defer func() {
		if err != nil {
			logger.Tracef("error setting resource %q for service %q: %v", res.Name, serviceID, err)
		} else {
			logger.Tracef("successfully added resource %q for %q", res.Name, serviceID)
		}
	}()
	logger.Tracef("adding resource %q for service %q", res.Name, serviceID)
	if err := res.Validate(); err != nil {
		return errors.Annotate(err, "bad resource metadata")
	}
	id := res.Name
	hash := res.Fingerprint.String()

	// TODO(ericsnow) Do something else if r is nil?

	// We use a staging approach for adding the resource metadata
	// to the model. This is necessary because the resource data
	// is stored separately and adding to both should be an atomic
	// operation.

	if err := st.persist.StageResource(id, serviceID, res); err != nil {
		return errors.Trace(err)
	}

	path := storagePath(res.Name, serviceID)
	if err := st.storage.PutAndCheckHash(path, r, res.Size, hash); err != nil {
		if err := st.persist.UnstageResource(id, serviceID); err != nil {
			logger.Errorf("could not unstage resource %q (service %q): %v", res.Name, serviceID, err)
		}
		return errors.Trace(err)
	}

	if err := st.persist.SetResource(id, serviceID, res); err != nil {
		if err := st.storage.Remove(path); err != nil {
			logger.Errorf("could not remove resource %q (service %q) from storage: %v", res.Name, serviceID, err)
		}
		if err := st.persist.UnstageResource(id, serviceID); err != nil {
			logger.Errorf("could not unstage resource %q (service %q): %v", res.Name, serviceID, err)
		}
		return errors.Trace(err)
	}
	return nil
}

// SetUnitResource records the resource being used by a unit in the Juju model.
func (st resourceState) SetUnitResource(serviceID, unitID string, res resource.Resource) (err error) {
	defer func() {
		if err != nil {
			logger.Tracef("error setting resource %q for %q: %v", res.Name, unitID, err)
		} else {
			logger.Tracef("successfully added resource %q for %q", res.Name, unitID)
		}
	}()
	logger.Tracef("adding resource %q for unit %q", res.Name, unitID)
	if err := res.Validate(); err != nil {
		return errors.Annotate(err, "bad resource metadata")
	}
	if err := st.persist.SetUnitResource(serviceID, unitID, res); err != nil {
		return errors.Trace(err)
	}
	return nil
}

// OpenResource returns metadata about the resource, and a reader for
// the resource.
func (st resourceState) OpenResource(forunit names.UnitTag, serviceID, name string) (resource.Resource, io.ReadCloser, error) {
	resourceInfo, err := st.GetResource(serviceID, name)
	if err != nil {
		return resource.Resource{}, nil, errors.Trace(err)
	}
	if resourceInfo.IsPlaceholder() {
		return resource.Resource{}, nil, errors.NotFoundf("resource %q", name)
	}

	r, resSize, err := st.storage.Get(storagePath(name, serviceID))
	if err != nil {
		return resource.Resource{}, nil, errors.Trace(err)
	}
	if resSize != resourceInfo.Size {
		msg := "storage returned a size (%d) which doesn't match resource metadata (%d)"
		return resource.Resource{}, nil, errors.Errorf(msg, resSize, resourceInfo.Size)
	}

	r = resourceReader{
		reader:  r,
		persist: st.persist,
		unit:    forunit,
		service: serviceID,
		res:     resourceInfo,
	}

	return resourceInfo, r, nil
}

// storagePath returns the path used as the location where the resource
// is stored in state storage. This requires that the returned string
// be unique and that it be organized in a structured way. In this case
// we start with a top-level (the service), then under that service use
// the "resources" section. The provided ID is located under there.
func storagePath(id, serviceID string) string {
	return path.Join("service-"+serviceID, "resources", id)
}

// resourceReader tracks when a resource's bytes have been fully downloaded by the
type resourceReader struct {
	reader  io.ReadCloser
	persist resourcePersistence
	unit    names.UnitTag
	service string
	res     resource.Resource
}

func (r resourceReader) Read(p []byte) (n int, err error) {
	n, err = r.reader.Read(p)
	if err == io.EOF {
		// record that the unit is now using this version of the resource
		if err := r.persist.SetUnitResource(r.unit.Id(), r.service, r.res); err != nil {
			// what can we do here?
			rev := fmt.Sprintf("%d", r.res.Revision)
			if r.res.Origin == charmresource.OriginUpload {
				rev = r.res.Timestamp.String()
			}
			logger.Errorf("Failed to record that unit %q is using resource %q revision %v", r.res.Name, rev)
		}
	}
	return n, err
}

func (r resourceReader) Close() error {
	return r.reader.Close()
}
