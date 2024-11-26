// Copyright 2024 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package service

import (
	"context"

	"github.com/juju/testing"
	jc "github.com/juju/testing/checkers"
	gc "gopkg.in/check.v1"

	corecredential "github.com/juju/juju/core/credential"
	coremodel "github.com/juju/juju/core/model"
	modeltesting "github.com/juju/juju/core/model/testing"
	corestatus "github.com/juju/juju/core/status"
	"github.com/juju/juju/domain/model"
	modelerrors "github.com/juju/juju/domain/model/errors"
	"github.com/juju/juju/internal/uuid"
)

type dummyModelState struct {
	models map[coremodel.UUID]model.ReadOnlyModelCreationArgs
	setID  coremodel.UUID

	modelState map[coremodel.UUID]model.ModelState
}

func (d *dummyModelState) Create(ctx context.Context, args model.ReadOnlyModelCreationArgs) error {
	if d.setID != coremodel.UUID("") {
		return modelerrors.AlreadyExists
	}
	d.models[args.UUID] = args
	d.setID = args.UUID
	return nil
}

func (d *dummyModelState) GetModel(ctx context.Context, id coremodel.UUID) (coremodel.Model, error) {
	args, exists := d.models[id]
	if !exists {
		return coremodel.Model{}, modelerrors.NotFound
	}

	return coremodel.Model{
		UUID:         args.UUID,
		Name:         args.Name,
		ModelType:    args.Type,
		AgentVersion: args.AgentVersion,
		Cloud:        args.Cloud,
		CloudType:    args.CloudType,
		CloudRegion:  args.CloudRegion,
		Credential: corecredential.Key{
			Name:  args.CredentialName,
			Owner: args.CredentialOwner,
			Cloud: args.Cloud,
		},
		OwnerName: args.CredentialOwner,
	}, nil
}

func (d *dummyModelState) Model(ctx context.Context) (coremodel.ReadOnlyModel, error) {
	if d.setID == coremodel.UUID("") {
		return coremodel.ReadOnlyModel{}, modelerrors.NotFound
	}

	args := d.models[d.setID]
	return coremodel.ReadOnlyModel{
		UUID:            args.UUID,
		AgentVersion:    args.AgentVersion,
		ControllerUUID:  args.ControllerUUID,
		Name:            args.Name,
		Type:            args.Type,
		Cloud:           args.Cloud,
		CloudType:       args.CloudType,
		CloudRegion:     args.CloudRegion,
		CredentialOwner: args.CredentialOwner,
		CredentialName:  args.CredentialName,
	}, nil
}

func (d *dummyModelState) Delete(ctx context.Context, modelUUID coremodel.UUID) error {
	delete(d.models, modelUUID)
	return nil
}

func (d *dummyModelState) GetModelState(_ context.Context, modelUUID coremodel.UUID) (model.ModelState, error) {
	mState, ok := d.modelState[modelUUID]
	if !ok {
		return model.ModelState{}, modelerrors.NotFound
	}
	return mState, nil
}

type modelServiceSuite struct {
	testing.IsolationSuite

	state          *dummyModelState
	controllerUUID uuid.UUID
}

var _ = gc.Suite(&modelServiceSuite{})

func (s *modelServiceSuite) SetUpTest(c *gc.C) {
	s.state = &dummyModelState{
		models:     map[coremodel.UUID]model.ReadOnlyModelCreationArgs{},
		modelState: map[coremodel.UUID]model.ModelState{},
	}

	s.controllerUUID = uuid.MustNewUUID()
}

func (s *modelServiceSuite) TestModelCreation(c *gc.C) {
	id := modeltesting.GenModelUUID(c)
	svc := NewModelService(id, s.state, s.state)

	s.state.models[id] = model.ReadOnlyModelCreationArgs{
		UUID:        id,
		Name:        "my-awesome-model",
		Cloud:       "aws",
		CloudType:   "ec2",
		CloudRegion: "myregion",
		Type:        coremodel.IAAS,
	}
	err := svc.CreateModel(context.Background(), s.controllerUUID)
	c.Assert(err, jc.ErrorIsNil)

	readonlyVal, err := svc.GetModelInfo(context.Background())
	c.Check(err, jc.ErrorIsNil)
	c.Check(readonlyVal, gc.Equals, coremodel.ReadOnlyModel{
		UUID:           id,
		ControllerUUID: s.controllerUUID,
		Name:           "my-awesome-model",
		Cloud:          "aws",
		CloudType:      "ec2",
		CloudRegion:    "myregion",
		Type:           coremodel.IAAS,
	})
}

func (s *modelServiceSuite) TestModelDeletion(c *gc.C) {
	id := modeltesting.GenModelUUID(c)
	svc := NewModelService(id, s.state, s.state)

	s.state.models[id] = model.ReadOnlyModelCreationArgs{
		UUID:        id,
		Name:        "my-awesome-model",
		Cloud:       "aws",
		CloudType:   "ec2",
		CloudRegion: "myregion",
		Type:        coremodel.IAAS,
	}
	err := svc.CreateModel(context.Background(), s.controllerUUID)
	c.Assert(err, jc.ErrorIsNil)

	err = svc.DeleteModel(context.Background())
	c.Assert(err, jc.ErrorIsNil)

	_, exists := s.state.models[id]
	c.Assert(exists, jc.IsFalse)
}

func (s *modelServiceSuite) TestStatusSuspended(c *gc.C) {
	id := modeltesting.GenModelUUID(c)
	svc := NewModelService(id, s.state, s.state)

	s.state.setID = id
	s.state.modelState[id] = model.ModelState{
		HasInvalidCloudCredential:    true,
		InvalidCloudCredentialReason: "invalid credential",
	}

	status, err := svc.Status(context.Background())
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(status, gc.DeepEquals, model.StatusInfo{
		Status:  corestatus.Suspended,
		Message: "suspended since cloud credential is not valid",
		Reason:  "invalid credential",
	})
}

func (s *modelServiceSuite) TestStatusDestroying(c *gc.C) {
	id := modeltesting.GenModelUUID(c)
	svc := NewModelService(id, s.state, s.state)

	s.state.setID = id
	s.state.modelState[id] = model.ModelState{
		Destroying: true,
	}

	status, err := svc.Status(context.Background())
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(status, gc.DeepEquals, model.StatusInfo{
		Status:  corestatus.Destroying,
		Message: "the model is being destroyed",
	})
}

func (s *modelServiceSuite) TestStatusBusy(c *gc.C) {
	id := modeltesting.GenModelUUID(c)
	svc := NewModelService(id, s.state, s.state)

	s.state.setID = id
	s.state.modelState[id] = model.ModelState{
		Migrating: true,
	}

	status, err := svc.Status(context.Background())
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(status, gc.DeepEquals, model.StatusInfo{
		Status:  corestatus.Busy,
		Message: "the model is being migrated",
	})
}

func (s *modelServiceSuite) TestStatus(c *gc.C) {
	id := modeltesting.GenModelUUID(c)
	svc := NewModelService(id, s.state, s.state)

	s.state.setID = id
	s.state.modelState[id] = model.ModelState{}

	status, err := svc.Status(context.Background())
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(status, gc.DeepEquals, model.StatusInfo{
		Status: corestatus.Available,
	})
}

func (s *modelServiceSuite) TestStatusFaildModelNotFound(c *gc.C) {
	id := modeltesting.GenModelUUID(c)
	svc := NewModelService(id, s.state, s.state)

	_, err := svc.Status(context.Background())
	c.Assert(err, jc.ErrorIs, modelerrors.NotFound)
}
