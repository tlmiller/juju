// Copyright 2023 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package service

import (
	"context"

	"github.com/juju/errors"

	"github.com/juju/juju/cloud"
	"github.com/juju/juju/core/changestream"
	"github.com/juju/juju/core/watcher"
)

// WatcherFactory instances return a watcher for a specified credential UUID,
type WatcherFactory interface {
	NewValueWatcher(
		namespace, uuid string, changeMask changestream.ChangeType,
	) (watcher.NotifyWatcher, error)
}

// State describes retrieval and persistence methods for storage.
type State interface {
	ProviderState

	// UpsertCloud persists the input cloud entity.
	UpsertCloud(context.Context, string, cloud.Cloud) error

	// DeleteCloud deletes the input cloud entity.
	DeleteCloud(context.Context, string) error

	// ListClouds returns the clouds matching the optional filter.
	ListClouds(context.Context) ([]cloud.Cloud, error)
}

// Service provides the API for working with clouds.
type Service struct {
	st State
}

// NewService returns a new service reference wrapping the input state.
func NewService(st State) *Service {
	return &Service{
		st: st,
	}
}

// UpsertCloud inserts or updates the specified cloud.
func (s *Service) UpsertCloud(ctx context.Context, userName string, cloud cloud.Cloud) error {
	err := s.st.UpsertCloud(ctx, userName, cloud)
	return errors.Annotatef(err, "updating cloud %q", cloud.Name)
}

// DeleteCloud removes the specified cloud.
func (s *Service) DeleteCloud(ctx context.Context, name string) error {
	err := s.st.DeleteCloud(ctx, name)
	return errors.Annotatef(err, "deleting cloud %q", name)
}

// ListAll returns all the clouds.
func (s *Service) ListAll(ctx context.Context) ([]cloud.Cloud, error) {
	all, err := s.st.ListClouds(ctx)
	return all, errors.Trace(err)
}

// Cloud returns the named cloud.
func (s *Service) Cloud(ctx context.Context, name string) (*cloud.Cloud, error) {
	cloud, err := s.st.Cloud(ctx, name)
	return cloud, errors.Trace(err)
}

// WatchableService defines a service for interacting with the underlying state
// and the ability to create watchers.
type WatchableService struct {
	Service
	watcherFactory WatcherFactory
}

// NewWatchableService returns a new service reference wrapping the
// input state and watcher factory.
func NewWatchableService(st State, watcherFactory WatcherFactory) *WatchableService {
	return &WatchableService{
		Service: Service{
			st: st,
		},
		watcherFactory: watcherFactory,
	}
}

// WatchCloud returns a watcher that observes changes to the specified cloud.
func (s *WatchableService) WatchCloud(ctx context.Context, name string) (watcher.NotifyWatcher, error) {
	return s.st.WatchCloud(ctx, s.watcherFactory.NewValueWatcher, name)
}
