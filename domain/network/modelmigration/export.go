// Copyright 2024 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package modelmigration

import (
	"context"

	"github.com/juju/description/v5"
	"github.com/juju/errors"
	"github.com/juju/loggo"

	"github.com/juju/juju/core/modelmigration"
	"github.com/juju/juju/core/network"
	"github.com/juju/juju/domain/network/service"
	"github.com/juju/juju/domain/network/state"
)

var logger = loggo.GetLogger("juju.migration.")

// RegisterExport registers the export operations with the given coordinator.
func RegisterExport(coordinator Coordinator) {
	coordinator.Add(&exportOperation{})
}

// ExportService provides a subset of the network domain
// service methods needed for spaces/subnets export.
type ExportService interface {
	GetAllSpaces(ctx context.Context) (network.SpaceInfos, error)
	GetAllSubnets(ctx context.Context) (network.SubnetInfos, error)
}

// exportOperation describes a way to execute a migration for
// exporting external controllers.
type exportOperation struct {
	modelmigration.BaseOperation

	exportService ExportService
}

// Setup implements Operation.
func (e *exportOperation) Setup(scope modelmigration.Scope) error {
	e.exportService = service.NewService(
		state.NewState(scope.ModelDB()),
		logger,
	)
	return nil
}

// Execute the migration export, which adds the spaces and subnets to the model.
func (e *exportOperation) Execute(ctx context.Context, model description.Model) error {
	spaces, err := e.exportService.GetAllSpaces(ctx)
	if err != nil {
		return errors.Trace(err)
	}
	for _, space := range spaces {
		// We do not export the alpha space because it is created by default
		// with the new model. This is OK, because it is immutable.
		// Any subnets added to the space will still be exported.
		if space.ID == network.AlphaSpaceId {
			continue
		}

		model.AddSpace(description.SpaceArgs{
			Id:         space.ID,
			Name:       string(space.Name),
			ProviderID: string(space.ProviderId),
		})
	}

	// Export subnets.
	subnets, err := e.exportService.GetAllSubnets(ctx)
	if err != nil {
		return errors.Trace(err)
	}
	for _, subnet := range subnets {
		args := description.SubnetArgs{
			ID:                string(subnet.ID),
			CIDR:              subnet.CIDR,
			ProviderId:        string(subnet.ProviderId),
			ProviderSpaceId:   string(subnet.ProviderSpaceId),
			ProviderNetworkId: string(subnet.ProviderNetworkId),
			VLANTag:           subnet.VLANTag,
			SpaceID:           subnet.SpaceID,
			SpaceName:         subnet.SpaceName,
			AvailabilityZones: subnet.AvailabilityZones,
			FanLocalUnderlay:  subnet.FanLocalUnderlay(),
			FanOverlay:        subnet.FanOverlay(),
		}
		model.AddSubnet(args)
	}

	return nil
}
