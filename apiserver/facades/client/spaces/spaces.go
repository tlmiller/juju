// Copyright 2015 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package spaces

import (
	"context"
	"fmt"

	"github.com/juju/collections/set"
	"github.com/juju/errors"
	"github.com/juju/names/v6"

	"github.com/juju/juju/apiserver/common/networkingcommon"
	apiservererrors "github.com/juju/juju/apiserver/errors"
	"github.com/juju/juju/apiserver/facade"
	"github.com/juju/juju/controller"
	corelogger "github.com/juju/juju/core/logger"
	"github.com/juju/juju/core/network"
	"github.com/juju/juju/core/permission"
	"github.com/juju/juju/rpc/params"
)

// ControllerConfigService is an interface that provides controller
// configuration.
type ControllerConfigService interface {
	ControllerConfig(context.Context) (controller.Config, error)
}

// NetworkService is the interface that is used to interact with the
// network spaces/subnets.
type NetworkService interface {
	// AddSpace creates and returns a new space.
	AddSpace(ctx context.Context, space network.SpaceInfo) (network.Id, error)
	// Space returns a space from state that matches the input ID. If the space
	// is not found, an error is returned satisfying
	// [github.com/juju/juju/domain/network/errors.SpaceNotFound].
	Space(ctx context.Context, uuid string) (*network.SpaceInfo, error)
	// SpaceByName returns a space from state that matches the input name. If
	// the space is not found, an error is returned satisfying
	// [github.com/juju/juju/domain/network/errors.SpaceNotFound].
	SpaceByName(ctx context.Context, name string) (*network.SpaceInfo, error)
	// GetAllSpaces returns all spaces for the model.
	GetAllSpaces(ctx context.Context) (network.SpaceInfos, error)
	// UpdateSpace updates the space name identified by the passed uuid. If
	// the space is not found, an error is returned satisfying
	// [github.com/juju/juju/domain/network/errors.SpaceNotFound].
	UpdateSpace(ctx context.Context, uuid string, name string) error
	// RemoveSpace deletes a space identified by its uuid. If the space is not
	// found, an error is returned satisfying
	// [github.com/juju/juju/domain/network/errors.SpaceNotFound].
	RemoveSpace(ctx context.Context, uuid string) error
	// ReloadSpaces loads spaces and subnets from the provider into state.
	ReloadSpaces(ctx context.Context) error
	// GetAllSubnets returns all the subnets for the model.
	GetAllSubnets(ctx context.Context) (network.SubnetInfos, error)
	// SubnetsByCIDR returns the subnets matching the input CIDRs.
	SubnetsByCIDR(ctx context.Context, cidrs ...string) ([]network.SubnetInfo, error)
	// Subnet returns the subnet identified by the input UUID,
	// or an error if it is not found.
	Subnet(ctx context.Context, uuid string) (*network.SubnetInfo, error)
	// UpdateSubnet updates the spaceUUID of the subnet identified by the input
	// UUID.
	UpdateSubnet(ctx context.Context, uuid, spaceUUID string) error
	// SupportsSpaces returns whether the current environment supports spaces.
	SupportsSpaces(ctx context.Context) (bool, error)
	// SupportsSpaceDiscovery returns whether the current environment supports
	// discovering spaces from the provider.
	SupportsSpaceDiscovery(ctx context.Context) (bool, error)
}

// API provides the spaces API facade for version 6.
type API struct {
	controllerConfigService ControllerConfigService

	networkService NetworkService
	backing        Backing
	resources      facade.Resources
	auth           facade.Authorizer

	check  BlockChecker
	logger corelogger.Logger
}

type apiConfig struct {
	NetworkService          NetworkService
	ControllerConfigService ControllerConfigService
	Backing                 Backing
	Check                   BlockChecker
	Resources               facade.Resources
	Authorizer              facade.Authorizer
	logger                  corelogger.Logger
}

// newAPIWithBacking creates a new server-side Spaces API facade with
// the given Backing.
func newAPIWithBacking(cfg apiConfig) (*API, error) {
	// Only clients can access the Spaces facade.
	if !cfg.Authorizer.AuthClient() {
		return nil, apiservererrors.ErrPerm
	}

	return &API{
		networkService:          cfg.NetworkService,
		controllerConfigService: cfg.ControllerConfigService,
		auth:                    cfg.Authorizer,
		backing:                 cfg.Backing,
		resources:               cfg.Resources,
		check:                   cfg.Check,
		logger:                  cfg.logger,
	}, nil
}

// CreateSpaces creates a new Juju network space, associating the
// specified subnets with it (optional; can be empty).
func (api *API) CreateSpaces(ctx context.Context, args params.CreateSpacesParams) (results params.ErrorResults, err error) {
	err = api.auth.HasPermission(ctx, permission.AdminAccess, api.backing.ModelTag())
	if err != nil {
		return results, err
	}
	if err := api.check.ChangeAllowed(ctx); err != nil {
		return results, errors.Trace(err)
	}
	if err = api.checkSupportsSpaces(ctx); err != nil {
		return results, apiservererrors.ServerError(errors.Trace(err))
	}

	results.Results = make([]params.ErrorResult, len(args.Spaces))

	for i, space := range args.Spaces {
		err := api.createOneSpace(ctx, space)
		if err == nil {
			continue
		}
		results.Results[i].Error = apiservererrors.ServerError(errors.Trace(err))
	}

	return results, nil
}

// createOneSpace creates one new Juju network space, associating the
// specified subnets with it (optional; can be empty).
func (api *API) createOneSpace(ctx context.Context, args params.CreateSpaceParams) error {
	// Validate the args, assemble information for api.backing.AddSpaces
	spaceTag, err := names.ParseSpaceTag(args.SpaceTag)
	if err != nil {
		return errors.Trace(err)
	}

	for _, cidr := range args.CIDRs {
		if !network.IsValidCIDR(cidr) {
			return errors.New(fmt.Sprintf("%q is not a valid CIDR", cidr))
		}
	}

	subnets, err := api.networkService.SubnetsByCIDR(ctx, args.CIDRs...)
	if err != nil {
		return err
	}

	// Add the validated space.
	spaceInfo := network.SpaceInfo{
		Name:       network.SpaceName(spaceTag.Id()),
		ProviderId: network.Id(args.ProviderId),
		Subnets:    subnets,
	}
	_, err = api.networkService.AddSpace(ctx, spaceInfo)
	if err != nil {
		return errors.Trace(err)
	}
	return nil
}

// ListSpaces lists all the available spaces and their associated subnets.
func (api *API) ListSpaces(ctx context.Context) (results params.ListSpacesResults, err error) {
	err = api.auth.HasPermission(ctx, permission.ReadAccess, api.backing.ModelTag())
	if err != nil {
		return results, err
	}

	err = api.checkSupportsSpaces(ctx)
	if err != nil {
		return results, apiservererrors.ServerError(errors.Trace(err))
	}

	spaces, err := api.networkService.GetAllSpaces(ctx)
	if err != nil {
		return results, errors.Trace(err)
	}

	results.Results = make([]params.Space, len(spaces))
	for i, space := range spaces {
		result := params.Space{}
		result.Id = space.ID
		result.Name = string(space.Name)

		if err != nil {
			err = errors.Annotatef(err, "fetching spaces")
			result.Error = apiservererrors.ServerError(err)
			results.Results[i] = result
			continue
		}
		subnets := space.Subnets

		result.Subnets = make([]params.Subnet, len(subnets))
		for i, subnet := range subnets {
			result.Subnets[i] = networkingcommon.SubnetInfoToParamsSubnet(subnet)
		}
		results.Results[i] = result
	}
	return results, nil
}

// ShowSpace shows the spaces for a set of given entities.
func (api *API) ShowSpace(ctx context.Context, entities params.Entities) (params.ShowSpaceResults, error) {
	err := api.auth.HasPermission(ctx, permission.ReadAccess, api.backing.ModelTag())
	if err != nil {
		return params.ShowSpaceResults{}, err
	}

	err = api.checkSupportsSpaces(ctx)
	if err != nil {
		return params.ShowSpaceResults{}, apiservererrors.ServerError(errors.Trace(err))
	}

	// Retrieve the list of all spaces, needed for the bindings.
	allSpaces, err := api.networkService.GetAllSpaces(ctx)
	if err != nil {
		return params.ShowSpaceResults{}, apiservererrors.ServerError(errors.Trace(err))
	}
	// Retrieve the list of all subnets, needed for the machine spaces.
	allSubnets, err := api.networkService.GetAllSubnets(ctx)
	if err != nil {
		return params.ShowSpaceResults{}, apiservererrors.ServerError(errors.Trace(err))
	}

	results := make([]params.ShowSpaceResult, len(entities.Entities))
	for i, entity := range entities.Entities {
		spaceName, err := names.ParseSpaceTag(entity.Tag)
		if err != nil {
			results[i].Error = apiservererrors.ServerError(errors.Trace(err))
			continue
		}
		var result params.ShowSpaceResult
		space, err := api.networkService.SpaceByName(ctx, spaceName.Id())
		if err != nil {
			newErr := errors.Annotatef(err, "fetching space %q", spaceName)
			results[i].Error = apiservererrors.ServerError(newErr)
			continue
		}
		result.Space.Name = string(space.Name)
		result.Space.Id = space.ID
		subnets := space.Subnets

		result.Space.Subnets = make([]params.Subnet, len(subnets))
		for i, subnet := range subnets {
			result.Space.Subnets[i] = networkingcommon.SubnetInfoToParamsSubnet(subnet)
		}

		// TODO(nvinuesa): This logic should be implemented in the
		// network service once we finish migrating applications to
		// dqlite.
		applications, err := api.applicationsBoundToSpace(space.ID, allSpaces)
		if err != nil {
			newErr := errors.Annotatef(err, "fetching applications")
			results[i].Error = apiservererrors.ServerError(newErr)
			continue
		}
		result.Applications = applications

		// TODO(nvinuesa): This logic should be implemented in the
		// network service once we finish migrating machines to
		// dqlite.
		machineCount, err := api.getMachineCountBySpaceID(space.ID, allSubnets)
		if err != nil {
			newErr := errors.Annotatef(err, "fetching machine count")
			results[i].Error = apiservererrors.ServerError(newErr)
			continue
		}

		result.MachineCount = machineCount
		results[i] = result
	}

	return params.ShowSpaceResults{Results: results}, err
}

// ReloadSpaces refreshes spaces from substrate
func (api *API) ReloadSpaces(ctx context.Context) error {
	err := api.auth.HasPermission(ctx, permission.AdminAccess, api.backing.ModelTag())
	if err != nil {
		return errors.Trace(err)
	}
	if err := api.check.ChangeAllowed(ctx); err != nil {
		return errors.Trace(err)
	}
	return errors.Trace(api.networkService.ReloadSpaces(ctx))
}

// checkSupportsSpaces checks if the environment implements NetworkingEnviron
// and also if it supports spaces. If we don't support spaces, an
// [errors.NotSupported] error will be returned.
func (api *API) checkSupportsSpaces(ctx context.Context) error {
	if supported, err := api.networkService.SupportsSpaces(ctx); err != nil {
		return errors.Trace(err)
	} else if !supported {
		return errors.NotSupportedf("spaces")
	}
	return nil
}

func (api *API) getMachineCountBySpaceID(spaceID string, allSubnets network.SubnetInfos) (int, error) {
	var count int
	machines, err := api.backing.AllMachines()
	if err != nil {
		return 0, errors.Trace(err)
	}
	for _, machine := range machines {
		spacesSet, err := machine.AllSpaces(allSubnets)
		if err != nil {
			return 0, errors.Trace(err)
		}
		if spacesSet.Contains(spaceID) {
			count++
		}
	}
	return count, nil
}

func (api *API) applicationsBoundToSpace(spaceID string, allSpaces network.SpaceInfos) ([]string, error) {
	allBindings, err := api.backing.AllEndpointBindings()
	if err != nil {
		return nil, errors.Trace(err)
	}

	applications := set.NewStrings()
	for app, bindings := range allBindings {
		for _, boundSpace := range bindings.Map() {
			if boundSpace == spaceID {
				applications.Add(app)
				break
			}
		}
	}
	return applications.SortedValues(), nil
}

// ensureSpacesAreMutable checks that the current user
// is allowed to edit the Space topology.
func (api *API) ensureSpacesAreMutable(ctx context.Context) error {
	err := api.auth.HasPermission(ctx, permission.AdminAccess, api.backing.ModelTag())
	if err != nil {
		return err
	}
	if err := api.check.ChangeAllowed(ctx); err != nil {
		return errors.Trace(err)
	}
	if err = api.ensureSpacesNotProviderSourced(ctx); err != nil {
		return apiservererrors.ServerError(errors.Trace(err))
	}
	return nil
}

// ensureSpacesNotProviderSourced checks if the environment implements
// NetworkingEnviron and also if it supports provider spaces.
// An error is returned if it is the provider and not the Juju operator
// that determines the space topology.
func (api *API) ensureSpacesNotProviderSourced(ctx context.Context) error {
	supported, err := api.networkService.SupportsSpaceDiscovery(ctx)
	if err != nil {
		return errors.Trace(err)
	} else if supported {
		return errors.NotSupportedf("modifying provider-sourced spaces")
	}
	return nil
}
