// Copyright 2024 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package service

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/juju/clock"
	"github.com/juju/collections/transform"
	"github.com/juju/errors"
	"github.com/juju/names/v5"
	"github.com/juju/version/v2"

	"github.com/juju/juju/caas"
	coreapplication "github.com/juju/juju/core/application"
	"github.com/juju/juju/core/assumes"
	"github.com/juju/juju/core/changestream"
	corecharm "github.com/juju/juju/core/charm"
	coredatabase "github.com/juju/juju/core/database"
	coreerrors "github.com/juju/juju/core/errors"
	"github.com/juju/juju/core/leadership"
	corelife "github.com/juju/juju/core/life"
	"github.com/juju/juju/core/logger"
	coremodel "github.com/juju/juju/core/model"
	"github.com/juju/juju/core/network"
	"github.com/juju/juju/core/providertracker"
	coresecrets "github.com/juju/juju/core/secrets"
	corestatus "github.com/juju/juju/core/status"
	corestorage "github.com/juju/juju/core/storage"
	coreunit "github.com/juju/juju/core/unit"
	"github.com/juju/juju/core/watcher"
	"github.com/juju/juju/core/watcher/eventsource"
	"github.com/juju/juju/domain"
	"github.com/juju/juju/domain/application"
	domaincharm "github.com/juju/juju/domain/application/charm"
	applicationerrors "github.com/juju/juju/domain/application/errors"
	"github.com/juju/juju/domain/ipaddress"
	"github.com/juju/juju/domain/life"
	"github.com/juju/juju/domain/linklayerdevice"
	domainstorage "github.com/juju/juju/domain/storage"
	"github.com/juju/juju/environs"
	internalcharm "github.com/juju/juju/internal/charm"
	internalerrors "github.com/juju/juju/internal/errors"
	interrors "github.com/juju/juju/internal/errors"
	"github.com/juju/juju/internal/storage"
)

// AtomicApplicationState describes retrieval and persistence methods for
// applications that require atomic transactions.
type AtomicApplicationState interface {
	domain.AtomicStateBase

	// GetApplicationID returns the ID for the named application, returning an
	// error satisfying [applicationerrors.ApplicationNotFound] if the
	// application is not found.
	GetApplicationID(ctx domain.AtomicContext, name string) (coreapplication.ID, error)

	// GetUnitUUID returns the UUID for the named unit, returning an error
	// satisfying [applicationerrors.UnitNotFound] if the unit doesn't exist.
	GetUnitUUID(ctx domain.AtomicContext, unitName coreunit.Name) (coreunit.UUID, error)

	// CreateApplication creates an application, returning an error satisfying
	// [applicationerrors.ApplicationAlreadyExists] if the application already exists.
	// If returns as error satisfying [applicationerrors.CharmNotFound] if the charm
	// for the application is not found.
	CreateApplication(domain.AtomicContext, string, application.AddApplicationArg) (coreapplication.ID, error)

	// AddUnits adds the specified units to the application.
	AddUnits(domain.AtomicContext, coreapplication.ID, ...application.AddUnitArg) error

	// InsertUnit insert the specified application unit, returning an error
	// satisfying [applicationerrors.UnitAlreadyExists]
	// if the unit exists.
	InsertUnit(domain.AtomicContext, coreapplication.ID, application.InsertUnitArg) error

	// UpdateUnitContainer updates the cloud container for specified unit,
	// returning an error satisfying [applicationerrors.UnitNotFoundError]
	// if the unit doesn't exist.
	UpdateUnitContainer(domain.AtomicContext, coreunit.Name, *application.CloudContainer) error

	// SetUnitPassword updates the password for the specified unit UUID.
	SetUnitPassword(domain.AtomicContext, coreunit.UUID, application.PasswordInfo) error

	// SetCloudContainerStatus saves the given cloud container status, overwriting any current status data.
	// If returns an error satisfying [applicationerrors.UnitNotFound] if the unit doesn't exist.
	SetCloudContainerStatus(domain.AtomicContext, coreunit.UUID, application.CloudContainerStatusStatusInfo) error

	// SetUnitAgentStatus saves the given unit agent status, overwriting any current status data.
	// If returns an error satisfying [applicationerrors.UnitNotFound] if the unit doesn't exist.
	SetUnitAgentStatus(domain.AtomicContext, coreunit.UUID, application.UnitAgentStatusInfo) error

	// SetUnitWorkloadStatus saves the given unit workload status, overwriting any current status data.
	// If returns an error satisfying [applicationerrors.UnitNotFound] if the unit doesn't exist.
	SetUnitWorkloadStatus(domain.AtomicContext, coreunit.UUID, application.UnitWorkloadStatusInfo) error

	// GetApplicationLife looks up the life of the specified application, returning an error
	// satisfying [applicationerrors.ApplicationNotFoundError] if the application is not found.
	GetApplicationLife(ctx domain.AtomicContext, appName string) (coreapplication.ID, life.Life, error)

	// SetApplicationLife sets the life of the specified application.
	SetApplicationLife(domain.AtomicContext, coreapplication.ID, life.Life) error

	// GetApplicationScaleState looks up the scale state of the specified
	// application, returning an error satisfying
	// [applicationerrors.ApplicationNotFound] if the application is not found.
	GetApplicationScaleState(domain.AtomicContext, coreapplication.ID) (application.ScaleState, error)

	// SetApplicationScalingState sets the scaling details for the given caas application
	// Scale is optional and is only set if not nil.
	SetApplicationScalingState(ctx domain.AtomicContext, appID coreapplication.ID, scale *int, targetScale int, scaling bool) error

	// SetDesiredApplicationScale updates the desired scale of the specified application.
	SetDesiredApplicationScale(domain.AtomicContext, coreapplication.ID, int) error

	// GetUnitLife looks up the life of the specified unit, returning an error
	// satisfying [applicationerrors.UnitNotFound] if the unit is not found.
	GetUnitLife(domain.AtomicContext, coreunit.Name) (life.Life, error)

	// SetUnitLife sets the life of the specified unit.
	SetUnitLife(domain.AtomicContext, coreunit.Name, life.Life) error

	// InitialWatchStatementUnitLife returns the initial namespace query for the application unit life watcher.
	InitialWatchStatementUnitLife(appName string) (string, eventsource.NamespaceQuery)

	// DeleteApplication deletes the specified application, returning an error
	// satisfying [applicationerrors.ApplicationNotFoundError] if the
	// application doesn't exist. If the application still has units, as error
	// satisfying [applicationerrors.ApplicationHasUnits] is returned.
	DeleteApplication(domain.AtomicContext, string) error

	// DeleteUnit deletes the specified unit.
	// If the unit's application is Dying and no
	// other references to it exist, true is returned to
	// indicate the application could be safely deleted.
	// It will fail if the unit is not Dead.
	DeleteUnit(domain.AtomicContext, coreunit.Name) (bool, error)

	// GetSecretsForUnit returns the secrets owned by the specified unit.
	GetSecretsForUnit(
		ctx domain.AtomicContext, unitName coreunit.Name,
	) ([]*coresecrets.URI, error)

	// GetSecretsForApplication returns the secrets owned by the specified application.
	GetSecretsForApplication(
		ctx domain.AtomicContext, applicationName string,
	) ([]*coresecrets.URI, error)
}

// ApplicationState describes retrieval and persistence methods for
// applications.
type ApplicationState interface {
	AtomicApplicationState

	// GetModelType returns the model type for the underlying model. If the model
	// does not exist then an error satisfying [modelerrors.NotFound] will be returned.
	GetModelType(context.Context) (coremodel.ModelType, error)

	// StorageDefaults returns the default storage sources for a model.
	StorageDefaults(context.Context) (domainstorage.StorageDefaults, error)

	// GetStoragePoolByName returns the storage pool with the specified name,
	// returning an error satisfying [storageerrors.PoolNotFoundError] if it
	// doesn't exist.
	GetStoragePoolByName(ctx context.Context, name string) (domainstorage.StoragePoolDetails, error)

	// GetUnitUUIDs returns the UUIDs for the named units in bulk, returning an
	// error satisfying [applicationerrors.UnitNotFound] if any of the units don't
	// exist.
	GetUnitUUIDs(context.Context, []coreunit.Name) ([]coreunit.UUID, error)

	// GetUnitNames gets in bulk the names for the specified unit UUIDs, returning
	// an error satisfying [applicationerrors.UnitNotFound] if any units are not
	// found.
	GetUnitNames(context.Context, []coreunit.UUID) ([]coreunit.Name, error)

	// UpsertCloudService updates the cloud service for the specified
	// application, returning an error satisfying
	// [applicationerrors.ApplicationNotFoundError] if the application doesn't
	// exist.
	UpsertCloudService(ctx context.Context, appName, providerID string, sAddrs network.SpaceAddresses) error

	// GetApplicationUnitLife returns the life values for the specified units of
	// the given application. The supplied ids may belong to a different
	// application; the application name is used to filter.
	GetApplicationUnitLife(ctx context.Context, appName string, unitUUIDs ...coreunit.UUID) (map[coreunit.UUID]life.Life, error)

	// GetCharmByApplicationID returns the charm, charm origin and charm
	// platform for the specified application ID.
	//
	// If the application does not exist, an error satisfying
	// [applicationerrors.ApplicationNotFoundError] is returned.
	// If the charm for the application does not exist, an error satisfying
	// [applicationerrors.CharmNotFoundError] is returned.
	GetCharmByApplicationID(context.Context, coreapplication.ID) (domaincharm.Charm, domaincharm.CharmOrigin, application.Platform, error)

	// GetCharmIDByApplicationName returns a charm ID by application name. It
	// returns an error if the charm can not be found by the name. This can also
	// be used as a cheap way to see if a charm exists without needing to load
	// the charm metadata.
	GetCharmIDByApplicationName(context.Context, string) (corecharm.ID, error)
}

// DeleteSecretState describes methods used by the secret deleter plugin.
type DeleteSecretState interface {
	// DeleteSecret deletes the specified secret revisions.
	// If revisions is nil the last remaining revisions are removed.
	DeleteSecret(ctx domain.AtomicContext, uri *coresecrets.URI, revs []int) error
}

const (
	// applicationSnippet is a non-compiled regexp that can be composed with
	// other snippets to form a valid application regexp.
	applicationSnippet = "(?:[a-z][a-z0-9]*(?:-[a-z0-9]*[a-z][a-z0-9]*)*)"
)

var (
	validApplication = regexp.MustCompile("^" + applicationSnippet + "$")
)

// ApplicationService provides the API for working with applications.
type ApplicationService struct {
	st     ApplicationState
	logger logger.Logger
	clock  clock.Clock

	storageRegistryGetter corestorage.ModelStorageRegistryGetter
	secretDeleter         DeleteSecretState
}

// NewApplicationService returns a new service reference wrapping the input state.
func NewApplicationService(st ApplicationState, deleteSecretState DeleteSecretState, storageRegistryGetter corestorage.ModelStorageRegistryGetter, logger logger.Logger) *ApplicationService {
	return &ApplicationService{
		st:                    st,
		logger:                logger,
		clock:                 clock.WallClock,
		storageRegistryGetter: storageRegistryGetter,
		secretDeleter:         deleteSecretState,
	}
}

// CreateApplication creates the specified application and units if required,
// returning an error satisfying [applicationerrors.ApplicationAlreadyExists]
// if the application already exists.
func (s *ApplicationService) CreateApplication(
	ctx context.Context,
	name string,
	charm internalcharm.Charm,
	origin corecharm.Origin,
	args AddApplicationArgs,
	units ...AddUnitArg,
) (coreapplication.ID, error) {
	if err := s.validateCreateApplicationParams(name, args.ReferenceName, charm, origin); err != nil {
		return "", interrors.Errorf("invalid application args %w", err)
	}

	modelType, err := s.st.GetModelType(ctx)
	if err != nil {
		return "", interrors.Errorf("getting model type %w", err)
	}
	appArg, err := s.makeCreateApplicationArgs(ctx, modelType, charm, origin, args)
	if err != nil {
		return "", interrors.Errorf("creating application args %w", err)
	}
	// We know that the charm name is valid, so we can use it as the application
	// name if that is not provided.
	if name == "" {
		// Annoyingly this should be the reference name, but that's not
		// true in the previous code. To keep compatibility, we'll use the
		// charm name.
		name = appArg.Charm.Metadata.Name
	}

	appArg.Scale = len(units)

	unitArgs := make([]application.AddUnitArg, len(units))
	for i, u := range units {
		arg := application.AddUnitArg{
			UnitName: u.UnitName,
		}
		s.addNewUnitStatusToArg(&arg.UnitStatusArg, modelType)
		unitArgs[i] = arg
	}

	var appID coreapplication.ID
	err = s.st.RunAtomic(ctx, func(ctx domain.AtomicContext) error {
		appID, err = s.st.CreateApplication(ctx, name, appArg)
		if err != nil {
			return interrors.Errorf("creating application %q %w", name, err)
		}
		return s.st.AddUnits(ctx, appID, unitArgs...)
	})
	return appID, err
}

func (s *ApplicationService) validateCreateApplicationParams(
	name, referenceName string,
	charm internalcharm.Charm,
	origin corecharm.Origin,
) error {
	if !isValidApplicationName(name) {
		return applicationerrors.ApplicationNameNotValid
	}

	// Validate that we have a valid charm and name.
	meta := charm.Meta()
	if meta == nil {
		return applicationerrors.CharmMetadataNotValid
	}

	if !isValidCharmName(meta.Name) {
		return applicationerrors.CharmNameNotValid
	} else if !isValidReferenceName(referenceName) {
		return interrors.Errorf("reference name: %w", applicationerrors.CharmNameNotValid)
	}

	// Validate the origin of the charm.
	if err := origin.Validate(); err != nil {
		return interrors.Errorf("%w: %v", applicationerrors.CharmOriginNotValid, err)
	}
	return nil
}

func (s *ApplicationService) makeCreateApplicationArgs(
	ctx context.Context,
	modelType coremodel.ModelType,
	charm internalcharm.Charm,
	origin corecharm.Origin,
	args AddApplicationArgs,
) (application.AddApplicationArg, error) {
	// TODO (stickupkid): These should be done either in the application
	// state in one transaction, or be operating on the domain/charm types.
	//TODO(storage) - insert storage directive for app

	cons := make(map[string]storage.Directive)
	for n, sc := range args.Storage {
		cons[n] = sc
	}
	if err := s.addDefaultStorageDirectives(ctx, modelType, cons, charm.Meta()); err != nil {
		return application.AddApplicationArg{}, interrors.Errorf("adding default storage directives %w", err)
	}
	if err := s.validateStorageDirectives(ctx, modelType, cons, charm); err != nil {
		return application.AddApplicationArg{}, interrors.Errorf("invalid storage directives %w", err)
	}

	// When encoding the charm, this will also validate the charm metadata,
	// when parsing it.
	ch, _, err := encodeCharm(charm)
	if err != nil {
		return application.AddApplicationArg{}, interrors.Errorf("encode charm: %w", err)
	}

	originArg, channelArg, platformArg, err := encodeCharmOrigin(origin, args.ReferenceName)
	if err != nil {
		return application.AddApplicationArg{}, interrors.Errorf("encode charm origin: %w", err)
	}

	return application.AddApplicationArg{
		Charm:    ch,
		Platform: platformArg,
		Origin:   originArg,
		Channel:  channelArg,
	}, nil
}

func (s *ApplicationService) addNewUnitStatusToArg(arg *application.UnitStatusArg, modelType coremodel.ModelType) {
	now := s.clock.Now()
	arg.AgentStatus = application.UnitAgentStatusInfo{
		StatusID: application.UnitAgentStatusAllocating,
		StatusInfo: application.StatusInfo{
			Since: now,
		},
	}
	arg.WorkloadStatus = application.UnitWorkloadStatusInfo{
		StatusID: application.UnitWorkloadStatusWaiting,
		StatusInfo: application.StatusInfo{
			Message: corestatus.MessageInstallingAgent,
			Since:   now,
		},
	}
	if modelType == coremodel.IAAS {
		arg.WorkloadStatus.Message = corestatus.MessageWaitForMachine
	}
}

func (s *ApplicationService) makeUnitStatus(in StatusParams) application.StatusInfo {
	si := application.StatusInfo{
		Message: in.Message,
		Since:   s.clock.Now(),
	}
	if in.Since != nil {
		si.Since = *in.Since
	}
	if len(in.Data) > 0 {
		si.Data = make(map[string]string)
		for k, v := range in.Data {
			if v == nil {
				continue
			}
			si.Data[k] = fmt.Sprintf("%v", v)
		}
	}
	return si
}

// ImportApplication imports the specified application and units if required,
// returning an error satisfying [applicationerrors.ApplicationAlreadyExists]
// if the application already exists.
func (s *ApplicationService) ImportApplication(
	ctx context.Context, appName string,
	charm internalcharm.Charm, origin corecharm.Origin, args AddApplicationArgs,
	units ...ImportUnitArg,
) error {
	if err := s.validateCreateApplicationParams(appName, args.ReferenceName, charm, origin); err != nil {
		return interrors.Errorf("invalid application args %w", err)
	}

	modelType, err := s.st.GetModelType(ctx)
	if err != nil {
		return interrors.Errorf("getting model type %w", err)
	}
	appArg, err := s.makeCreateApplicationArgs(ctx, modelType, charm, origin, args)
	if err != nil {
		return interrors.Errorf("creating application args %w", err)
	}
	appArg.Scale = len(units)

	unitArgs := make([]application.InsertUnitArg, len(units))
	for i, u := range units {
		arg := application.InsertUnitArg{
			UnitName: u.UnitName,
			UnitStatusArg: application.UnitStatusArg{
				AgentStatus: application.UnitAgentStatusInfo{
					StatusID:   application.MarshallUnitAgentStatus(u.AgentStatus.Status),
					StatusInfo: s.makeUnitStatus(u.AgentStatus),
				},
				WorkloadStatus: application.UnitWorkloadStatusInfo{
					StatusID:   application.MarshallUnitWorkloadStatus(u.WorkloadStatus.Status),
					StatusInfo: s.makeUnitStatus(u.WorkloadStatus),
				},
			},
		}
		if u.CloudContainer != nil {
			arg.CloudContainer = makeCloudContainerArg(u.UnitName, *u.CloudContainer)
		}
		if u.PasswordHash != nil {
			arg.Password = &application.PasswordInfo{
				PasswordHash:  *u.PasswordHash,
				HashAlgorithm: application.HashAlgorithmSHA256,
			}
		}
		unitArgs[i] = arg
	}

	err = s.st.RunAtomic(ctx, func(ctx domain.AtomicContext) error {
		appID, err := s.st.CreateApplication(ctx, appName, appArg)
		if err != nil {
			return interrors.Errorf("creating application %q %w", appName, err)
		}
		for _, arg := range unitArgs {
			if err := s.st.InsertUnit(ctx, appID, arg); err != nil {
				return interrors.Errorf("inserting unit %q %w", arg.UnitName, err)
			}
		}
		return nil
	})
	return err
}

// AddUnits adds the specified units to the application, returning an error
// satisfying [applicationerrors.ApplicationNotFoundError] if the application doesn't exist.
func (s *ApplicationService) AddUnits(ctx context.Context, name string, units ...AddUnitArg) error {
	modelType, err := s.st.GetModelType(ctx)
	if err != nil {
		return interrors.Errorf("getting model type %w", err)
	}

	args := make([]application.AddUnitArg, len(units))
	for i, u := range units {
		arg := application.AddUnitArg{
			UnitName: u.UnitName,
		}
		s.addNewUnitStatusToArg(&arg.UnitStatusArg, modelType)
		args[i] = arg
	}

	err = s.st.RunAtomic(ctx, func(ctx domain.AtomicContext) error {
		appID, err := s.st.GetApplicationID(ctx, name)
		if err != nil {
			return interrors.Capture(err)
		}
		return s.st.AddUnits(ctx, appID, args...)
	})
	return interrors.Errorf("adding units to application %q %w", name, err)
}

// GetUnitUUIDs returns the UUIDs for the named units in bulk, returning an error
// satisfying [applicationerrors.UnitNotFound] if any of the units don't exist.
func (s *ApplicationService) GetUnitUUIDs(ctx context.Context, unitNames []coreunit.Name) ([]coreunit.UUID, error) {
	uuids, err := s.st.GetUnitUUIDs(ctx, unitNames)
	if err != nil {
		return nil, internalerrors.Errorf("failed to get unit UUIDs: %w", err)
	}
	return uuids, nil
}

// GetUnitUUID returns the UUID for the named unit, returning an error
// satisfying [applicationerrors.UnitNotFound] if the unit doesn't exist.
func (s *ApplicationService) GetUnitUUID(ctx context.Context, unitName coreunit.Name) (coreunit.UUID, error) {
	uuids, err := s.GetUnitUUIDs(ctx, []coreunit.Name{unitName})
	if err != nil {
		return "", err
	}
	return uuids[0], nil
}

// GetUnitNames gets in bulk the names for the specified unit UUIDs, returning an
// error satisfying [applicationerrors.UnitNotFound] if any units are not found.
func (s *ApplicationService) GetUnitNames(ctx context.Context, unitUUIDs []coreunit.UUID) ([]coreunit.Name, error) {
	names, err := s.st.GetUnitNames(ctx, unitUUIDs)
	if err != nil {
		return nil, internalerrors.Errorf("failed to get unit names: %w", err)
	}
	return names, nil
}

// GetUnitLife looks up the life of the specified unit, returning an error
// satisfying [applicationerrors.UnitNotFoundError] if the unit is not found.
func (s *ApplicationService) GetUnitLife(ctx context.Context, unitName coreunit.Name) (corelife.Value, error) {
	var result corelife.Value
	err := s.st.RunAtomic(ctx, func(ctx domain.AtomicContext) error {
		unitLife, err := s.st.GetUnitLife(ctx, unitName)
		result = unitLife.Value()
		return interrors.Errorf("getting life for %q %w", unitName, err)
	})
	return result, interrors.Capture(err)
}

// DeleteUnit deletes the specified unit.
// TODO(units) - rework when dual write is refactored
// This method is called (mostly during cleanup) after a unit
// has been removed from mongo. The mongo calls are
// DestroyMaybeRemove, DestroyWithForce, RemoveWithForce.
func (s *ApplicationService) DeleteUnit(ctx context.Context, unitName coreunit.Name) error {
	err := s.st.RunAtomic(ctx, func(ctx domain.AtomicContext) error {
		return s.deleteUnit(ctx, unitName)
	})
	if err != nil {
		return interrors.Errorf("deleting unit %q %w", unitName, err)
	}
	return nil
}

func (s *ApplicationService) deleteUnit(ctx domain.AtomicContext, unitName coreunit.Name) error {
	// Get unit owned secrets.
	uris, err := s.st.GetSecretsForUnit(ctx, unitName)
	if err != nil {
		return interrors.Errorf("getting unit owned secrets for %q %w", unitName, err)
	}
	// Delete unit owned secrets.
	for _, uri := range uris {
		s.logger.Debugf("deleting unit %q secret: %s", unitName, uri.ID)
		err := s.secretDeleter.DeleteSecret(ctx, uri, nil)
		if err != nil {
			return interrors.Errorf("deleting secret %q %w", uri, err)
		}
	}

	err = s.ensureUnitDead(ctx, unitName)
	if interrors.Is(err, applicationerrors.UnitNotFound) {
		return nil
	}
	if err != nil {
		return interrors.Capture(err)
	}

	isLast, err := s.st.DeleteUnit(ctx, unitName)
	if err != nil {
		return interrors.Errorf("deleting unit %q %w", unitName, err)
	}
	if isLast {
		// TODO(units): schedule application cleanup
		_ = isLast
	}
	return nil
}

// DestroyUnit prepares a unit for removal from the model
// returning an error  satisfying [applicationerrors.UnitNotFoundError]
// if the unit doesn't exist.
func (s *ApplicationService) DestroyUnit(ctx context.Context, unitName coreunit.Name) error {
	// For now, all we do is advance the unit's life to Dying.
	err := s.st.RunAtomic(ctx, func(ctx domain.AtomicContext) error {
		return s.st.SetUnitLife(ctx, unitName, life.Dying)
	})
	return interrors.Errorf("destroying unit %q %w", unitName, err)
}

// EnsureUnitDead is called by the unit agent just before it terminates.
// TODO(units): revisit his existing logic ported from mongo
// Note: the agent only calls this method once it gets notification
// that the unit has become dead, so there's strictly no need to call
// this method as the unit is already dead.
// This method is also called during cleanup from various cleanup jobs.
// If the unit is not found, an error satisfying [applicationerrors.UnitNotFound]
// is returned.
func (s *ApplicationService) EnsureUnitDead(ctx context.Context, unitName coreunit.Name, leadershipRevoker leadership.Revoker) error {
	err := s.st.RunAtomic(ctx, func(ctx domain.AtomicContext) error {
		return s.ensureUnitDead(ctx, unitName)
	})
	if interrors.Is(err, applicationerrors.UnitNotFound) {
		return nil
	}
	if err == nil {
		appName, _ := names.UnitApplication(unitName.String())
		if err := leadershipRevoker.RevokeLeadership(appName, unitName); err != nil && !interrors.Is(err, leadership.ErrClaimNotHeld) {
			s.logger.Warningf("cannot revoke lease for dead unit %q", unitName)
		}
	}
	return interrors.Errorf("ensuring unit %q is dead %w", unitName, err)
}

func (s *ApplicationService) ensureUnitDead(ctx domain.AtomicContext, unitName coreunit.Name) (err error) {
	unitLife, err := s.st.GetUnitLife(ctx, unitName)
	if err != nil {
		return interrors.Capture(err)
	}
	if unitLife == life.Dead {
		return nil
	}
	// TODO(units) - check for subordinates and storage attachments
	// For IAAS units, we need to do additional checks - these are still done in mongo.
	// If a unit still has subordinates, return applicationerrors.UnitHasSubordinates.
	// If a unit still has storage attachments, return applicationerrors.UnitHasStorageAttachments.
	err = s.st.SetUnitLife(ctx, unitName, life.Dead)
	return interrors.Errorf("ensuring unit %q is dead %w", unitName, err)
}

// RemoveUnit is called by the deployer worker and caas application provisioner worker to
// remove from the model units which have transitioned to dead.
// TODO(units): revisit his existing logic ported from mongo
// Note: the callers of this method only do so after the unit has become dead, so
// there's strictly no need to call ensureUnitDead before removing.
// If the unit is still alive, an error satisfying [applicationerrors.UnitIsAlive]
// is returned. If the unit is not found, an error satisfying
// [applicationerrors.UnitNotFound] is returned.
func (s *ApplicationService) RemoveUnit(ctx context.Context, unitName coreunit.Name, leadershipRevoker leadership.Revoker) error {
	err := s.st.RunAtomic(ctx, func(ctx domain.AtomicContext) error {
		unitLife, err := s.st.GetUnitLife(ctx, unitName)
		if err != nil {
			return interrors.Capture(err)
		}
		if unitLife == life.Alive {
			return interrors.Errorf("cannot remove unit %q: %w", unitName, applicationerrors.UnitIsAlive)
		}
		err = s.deleteUnit(ctx, unitName)
		return interrors.Errorf("deleting unit %q %w", unitName, err)
	})
	if err != nil {
		return interrors.Errorf("removing unit %q %w", unitName, err)
	}
	appName, _ := names.UnitApplication(unitName.String())
	if err := leadershipRevoker.RevokeLeadership(appName, unitName); err != nil && !interrors.Is(err, leadership.ErrClaimNotHeld) {
		s.logger.Warningf("cannot revoke lease for dead unit %q", unitName)
	}
	return nil
}

func makeCloudContainerArg(unitName coreunit.Name, cloudContainer CloudContainerParams) *application.CloudContainer {
	result := &application.CloudContainer{
		ProviderId: cloudContainer.ProviderId,
		Ports:      cloudContainer.Ports,
	}
	if cloudContainer.Address != nil {
		// TODO(units) - handle the cloudContainer.Address space ID
		// For k8s we'll initially create a /32 subnet off the container address
		// and add that to the default space.
		result.Address = &application.ContainerAddress{
			// For cloud containers, the device is a placeholder without
			// a MAC address and once inserted, not updated. It just exists
			// to tie the address to the net node corresponding to the
			// cloud container.
			Device: application.ContainerDevice{
				Name:              fmt.Sprintf("placeholder for %q cloud container", unitName),
				DeviceTypeID:      linklayerdevice.DeviceTypeUnknown,
				VirtualPortTypeID: linklayerdevice.NonVirtualPortType,
			},
			Value:       cloudContainer.Address.Value,
			AddressType: ipaddress.MarshallAddressType(cloudContainer.Address.AddressType()),
			Scope:       ipaddress.MarshallScope(cloudContainer.Address.Scope),
			Origin:      ipaddress.MarshallOrigin(network.OriginProvider),
			ConfigType:  ipaddress.MarshallConfigType(network.ConfigDHCP),
		}
		if cloudContainer.AddressOrigin != nil {
			result.Address.Origin = ipaddress.MarshallOrigin(*cloudContainer.AddressOrigin)
		}
	}
	return result
}

// RegisterCAASUnit creates or updates the specified application unit in a caas model,
// returning an error satisfying [applicationerrors.ApplicationNotFoundError]
// if the application doesn't exist. If the unit life is Dead, an error
// satisfying [applicationerrors.UnitAlreadyExists] is returned.
func (s *ApplicationService) RegisterCAASUnit(ctx context.Context, appName string, args RegisterCAASUnitParams) error {
	if args.PasswordHash == "" {
		return interrors.Errorf("password hash %w", coreerrors.NotValid)
	}
	if args.ProviderId == "" {
		return interrors.Errorf("provider id %w", coreerrors.NotValid)
	}
	if !args.OrderedScale {
		return interrors.Errorf("registering CAAS units not supported without ordered unit IDs %w", coreerrors.NotImplemented)
	}
	if args.UnitName == "" {
		return interrors.Errorf("missing unit name %w", coreerrors.NotValid)
	}

	cloudContainerParams := CloudContainerParams{
		ProviderId: args.ProviderId,
		Ports:      args.Ports,
	}
	if args.Address != nil {
		addr := network.NewSpaceAddress(*args.Address, network.WithScope(network.ScopeMachineLocal))
		cloudContainerParams.Address = &addr
		origin := network.OriginProvider
		cloudContainerParams.AddressOrigin = &origin
	}

	cloudContainer := makeCloudContainerArg(args.UnitName, cloudContainerParams)
	err := s.st.RunAtomic(ctx, func(ctx domain.AtomicContext) error {
		appID, err := s.st.GetApplicationID(ctx, appName)
		if err != nil {
			return interrors.Capture(err)
		}
		unitLife, err := s.st.GetUnitLife(ctx, args.UnitName)
		if interrors.Is(err, applicationerrors.UnitNotFound) {
			arg := application.InsertUnitArg{
				UnitName: args.UnitName,
				Password: &application.PasswordInfo{
					PasswordHash:  args.PasswordHash,
					HashAlgorithm: application.HashAlgorithmSHA256,
				},
				CloudContainer: cloudContainer,
			}
			s.addNewUnitStatusToArg(&arg.UnitStatusArg, coremodel.CAAS)
			return s.insertCAASUnit(ctx, appID, args.OrderedId, arg)
		}
		if unitLife == life.Dead {
			return interrors.Errorf("dead unit %q already exists%w", args.UnitName, errors.Hide(applicationerrors.UnitAlreadyExists))
		}
		if err := s.st.UpdateUnitContainer(ctx, args.UnitName, cloudContainer); err != nil {
			return interrors.Errorf("updating unit %q %w", args.UnitName, err)
		}

		// We want to transition to using unit UUID instead of name.
		unitUUID, err := s.st.GetUnitUUID(ctx, args.UnitName)
		if err != nil {
			return interrors.Capture(err)
		}
		return s.st.SetUnitPassword(ctx, unitUUID, application.PasswordInfo{
			PasswordHash:  args.PasswordHash,
			HashAlgorithm: application.HashAlgorithmSHA256,
		})
	})
	return interrors.Errorf("saving caas unit %q %w", args.UnitName, err)
}

func (s *ApplicationService) insertCAASUnit(
	ctx domain.AtomicContext, appID coreapplication.ID, orderedID int, arg application.InsertUnitArg,
) error {
	appScale, err := s.st.GetApplicationScaleState(ctx, appID)
	if err != nil {
		return interrors.Errorf("getting application scale state for app %q %w", appID, err)
	}
	if orderedID >= appScale.Scale ||
		(appScale.Scaling && orderedID >= appScale.ScaleTarget) {
		return interrors.Errorf("unrequired unit %s is not assigned%w", arg.UnitName, errors.Hide(applicationerrors.UnitNotAssigned))
	}
	return s.st.InsertUnit(ctx, appID, arg)
}

// UpdateCAASUnit updates the specified CAAS unit, returning an error
// satisfying applicationerrors.ApplicationNotAlive if the unit's
// application is not alive.
func (s *ApplicationService) UpdateCAASUnit(ctx context.Context, unitName coreunit.Name, params UpdateCAASUnitParams) error {
	var cloudContainer *application.CloudContainer
	if params.ProviderId != nil {
		cloudContainerParams := CloudContainerParams{
			ProviderId: *params.ProviderId,
			Ports:      params.Ports,
		}
		if params.Address != nil {
			addr := network.NewSpaceAddress(*params.Address, network.WithScope(network.ScopeMachineLocal))
			cloudContainerParams.Address = &addr
			origin := network.OriginProvider
			cloudContainerParams.AddressOrigin = &origin
		}
		cloudContainer = makeCloudContainerArg(unitName, cloudContainerParams)
	}
	appName, err := names.UnitApplication(unitName.String())
	if err != nil {
		return interrors.Capture(err)
	}
	err = s.st.RunAtomic(ctx, func(ctx domain.AtomicContext) error {
		_, appLife, err := s.st.GetApplicationLife(ctx, appName)
		if err != nil {
			return interrors.Errorf("getting application %q life: %w", appName, err)
		}
		if appLife != life.Alive {
			return interrors.Errorf("application %q is not alive%w", appName, errors.Hide(applicationerrors.ApplicationNotAlive))
		}

		if cloudContainer != nil {
			if err := s.st.UpdateUnitContainer(ctx, unitName, cloudContainer); err != nil {
				return interrors.Errorf("updating cloud container %q %w", unitName, err)
			}
		}
		// We want to transition to using unit UUID instead of name.
		unitUUID, err := s.st.GetUnitUUID(ctx, unitName)
		if err != nil {
			return interrors.Capture(err)
		}
		now := time.Now()
		since := func(in *time.Time) time.Time {
			if in != nil {
				return *in
			}
			return now
		}
		if params.AgentStatus != nil {
			if err := s.st.SetUnitAgentStatus(ctx, unitUUID, application.UnitAgentStatusInfo{
				StatusID: application.MarshallUnitAgentStatus(params.AgentStatus.Status),
				StatusInfo: application.StatusInfo{
					Message: params.AgentStatus.Message,
					Data: transform.Map(
						params.AgentStatus.Data, func(k string, v any) (string, string) { return k, fmt.Sprint(v) }),
					Since: since(params.AgentStatus.Since),
				},
			}); err != nil {
				return interrors.Errorf("saving unit %q agent status  %w", unitName, err)
			}
		}
		if params.WorkloadStatus != nil {
			if err := s.st.SetUnitWorkloadStatus(ctx, unitUUID, application.UnitWorkloadStatusInfo{
				StatusID: application.MarshallUnitWorkloadStatus(params.WorkloadStatus.Status),
				StatusInfo: application.StatusInfo{
					Message: params.WorkloadStatus.Message,
					Data: transform.Map(
						params.WorkloadStatus.Data, func(k string, v any) (string, string) { return k, fmt.Sprint(v) }),
					Since: since(params.WorkloadStatus.Since),
				},
			}); err != nil {
				return interrors.Errorf("saving unit %q workload status  %w", unitName, err)
			}
		}
		if params.CloudContainerStatus != nil {
			if err := s.st.SetCloudContainerStatus(ctx, unitUUID, application.CloudContainerStatusStatusInfo{
				StatusID: application.MarshallCloudContainerStatus(params.CloudContainerStatus.Status),
				StatusInfo: application.StatusInfo{
					Message: params.CloudContainerStatus.Message,
					Data: transform.Map(
						params.CloudContainerStatus.Data, func(k string, v any) (string, string) { return k, fmt.Sprint(v) }),
					Since: since(params.CloudContainerStatus.Since),
				},
			}); err != nil {
				return interrors.Errorf("saving unit %q cloud container status  %w", unitName, err)
			}
		}
		return nil
	})
	return interrors.Errorf("updating caas unit %q %w", unitName, err)
}

// SetUnitPassword updates the password for the specified unit, returning an error
// satisfying [applicationerrors.NotNotFound] if the unit doesn't exist.
func (s *ApplicationService) SetUnitPassword(ctx context.Context, unitName coreunit.Name, password string) error {
	return s.st.RunAtomic(ctx, func(ctx domain.AtomicContext) error {
		unitUUID, err := s.st.GetUnitUUID(ctx, unitName)
		if err != nil {
			return interrors.Capture(err)
		}
		return s.st.SetUnitPassword(ctx, unitUUID, application.PasswordInfo{
			PasswordHash:  password,
			HashAlgorithm: application.HashAlgorithmSHA256,
		})
	})
}

// DeleteApplication deletes the specified application, returning an error
// satisfying [applicationerrors.ApplicationNotFoundError] if the application doesn't exist.
// If the application still has units, as error satisfying [applicationerrors.ApplicationHasUnits]
// is returned.
func (s *ApplicationService) DeleteApplication(ctx context.Context, name string) error {
	var cleanups []func(context.Context)
	err := s.st.RunAtomic(ctx, func(ctx domain.AtomicContext) error {
		var err error
		cleanups, err = s.deleteApplication(ctx, name)
		return interrors.Capture(err)
	})
	if err != nil {
		return interrors.Errorf("deleting application %q %w", name, err)
	}
	for _, cleanup := range cleanups {
		cleanup(ctx)
	}
	return nil
}

func (s *ApplicationService) deleteApplication(ctx domain.AtomicContext, name string) ([]func(context.Context), error) {
	// Get app owned secrets.
	uris, err := s.st.GetSecretsForApplication(ctx, name)
	if err != nil {
		return nil, interrors.Errorf("getting application owned secrets for %q %w", name, err)
	}
	// Delete app owned secrets.
	for _, uri := range uris {
		s.logger.Debugf("deleting application %q secret: %s", name, uri.ID)
		err := s.secretDeleter.DeleteSecret(ctx, uri, nil)
		if err != nil {
			return nil, interrors.Errorf("deleting secret %q %w", uri, err)
		}
	}

	err = s.st.DeleteApplication(ctx, name)
	return nil, interrors.Errorf("deleting application %q %w", name, err)
}

// DestroyApplication prepares an application for removal from the model
// returning an error  satisfying [applicationerrors.ApplicationNotFoundError]
// if the application doesn't exist.
func (s *ApplicationService) DestroyApplication(ctx context.Context, appName string) error {
	// For now, all we do is advance the application's life to Dying.
	err := s.st.RunAtomic(ctx, func(ctx domain.AtomicContext) error {
		appID, err := s.st.GetApplicationID(ctx, appName)
		if interrors.Is(err, applicationerrors.ApplicationNotFound) {
			return nil
		}
		if err != nil {
			return interrors.Capture(err)
		}
		return s.st.SetApplicationLife(ctx, appID, life.Dying)
	})
	return interrors.Errorf("destroying application %q %w", appName, err)
}

// EnsureApplicationDead is called by the cleanup worker if a mongo
// destroy operation sets the application to dead.
// TODO(units): remove when everything is in dqlite.
func (s *ApplicationService) EnsureApplicationDead(ctx context.Context, appName string) error {
	err := s.st.RunAtomic(ctx, func(ctx domain.AtomicContext) error {
		appID, err := s.st.GetApplicationID(ctx, appName)
		if interrors.Is(err, applicationerrors.ApplicationNotFound) {
			return nil
		}
		if err != nil {
			return interrors.Capture(err)
		}
		return s.st.SetApplicationLife(ctx, appID, life.Dead)
	})
	return interrors.Errorf("setting application %q life to Dead %w", appName, err)
}

// UpdateApplicationCharm sets a new charm for the application, validating that aspects such
// as storage are still viable with the new charm.
func (s *ApplicationService) UpdateApplicationCharm(ctx context.Context, name string, params UpdateCharmParams) error {
	//TODO(storage) - update charm and storage directive for app
	return nil
}

// GetApplicationIDByName returns a application ID by application name. It
// returns an error if the application can not be found by the name.
//
// Returns [applicationerrors.ApplicationNameNotValid] if the name is not valid,
// and [applicationerrors.ApplicationNotFound] if the application is not found.
func (s *ApplicationService) GetApplicationIDByName(ctx context.Context, name string) (coreapplication.ID, error) {
	if !isValidApplicationName(name) {
		return "", applicationerrors.ApplicationNameNotValid
	}

	var id coreapplication.ID
	err := s.st.RunAtomic(ctx, func(ctx domain.AtomicContext) error {
		appID, err := s.st.GetApplicationID(ctx, name)
		if err != nil {
			return interrors.Capture(err)
		}
		id = appID
		return nil
	})
	return id, interrors.Capture(err)
}

// GetCharmIDByApplicationName returns a charm ID by application name. It
// returns an error if the charm can not be found by the name. This can also be
// used as a cheap way to see if a charm exists without needing to load the
// charm metadata.
//
// Returns [applicationerrors.ApplicationNameNotValid] if the name is not valid,
// and [applicationerrors.CharmNotFound] if the charm is not found.
func (s *ApplicationService) GetCharmIDByApplicationName(ctx context.Context, name string) (corecharm.ID, error) {
	if !isValidApplicationName(name) {
		return "", applicationerrors.ApplicationNameNotValid
	}

	return s.st.GetCharmIDByApplicationName(ctx, name)
}

// GetCharmByApplicationID returns the charm for the specified application
// ID.
//
// If the application does not exist, an error satisfying
// [applicationerrors.ApplicationNotFound] is returned. If the charm for the
// application does not exist, an error satisfying
// [applicationerrors.CharmNotFound is returned. If the application name is not
// valid, an error satisfying [applicationerrors.ApplicationNameNotValid] is
// returned.
func (s *ApplicationService) GetCharmByApplicationID(ctx context.Context, id coreapplication.ID) (
	internalcharm.Charm,
	domaincharm.CharmOrigin,
	application.Platform,
	error,
) {
	if err := id.Validate(); err != nil {
		return nil, domaincharm.CharmOrigin{}, application.Platform{}, internalerrors.Errorf("application ID: %w%w", err).Add(applicationerrors.ApplicationIDNotValid)
	}

	charm, origin, platform, err := s.st.GetCharmByApplicationID(ctx, id)
	if err != nil {
		return nil, origin, platform, interrors.Capture(err)
	}

	// The charm needs to be decoded into the internalcharm.Charm type.

	metadata, err := decodeMetadata(charm.Metadata)
	if err != nil {
		return nil, origin, platform, interrors.Capture(err)
	}

	manifest, err := decodeManifest(charm.Manifest)
	if err != nil {
		return nil, origin, platform, interrors.Capture(err)
	}

	actions, err := decodeActions(charm.Actions)
	if err != nil {
		return nil, origin, platform, interrors.Capture(err)
	}

	config, err := decodeConfig(charm.Config)
	if err != nil {
		return nil, origin, platform, interrors.Capture(err)
	}

	lxdProfile, err := decodeLXDProfile(charm.LXDProfile)
	if err != nil {
		return nil, origin, platform, interrors.Capture(err)
	}

	return internalcharm.NewCharmBase(
		&metadata,
		&manifest,
		&config,
		&actions,
		&lxdProfile,
	), origin, platform, nil
}

// addDefaultStorageDirectives fills in default values, replacing any empty/missing values
// in the specified directives.
func (s *ApplicationService) addDefaultStorageDirectives(ctx context.Context, modelType coremodel.ModelType, allDirectives map[string]storage.Directive, charmMeta *internalcharm.Meta) error {
	defaults, err := s.st.StorageDefaults(ctx)
	if err != nil {
		return interrors.Errorf("getting storage defaults %w", err)
	}
	return domainstorage.StorageDirectivesWithDefaults(charmMeta.Storage, modelType, defaults, allDirectives)
}

func (s *ApplicationService) validateStorageDirectives(ctx context.Context, modelType coremodel.ModelType, allDirectives map[string]storage.Directive, charm internalcharm.Charm) error {
	registry, err := s.storageRegistryGetter.GetStorageRegistry(ctx)
	if err != nil {
		return interrors.Capture(err)
	}

	validator, err := domainstorage.NewStorageDirectivesValidator(modelType, registry, s.st)
	if err != nil {
		return interrors.Capture(err)
	}
	err = validator.ValidateStorageDirectivesAgainstCharm(ctx, allDirectives, charm)
	if err != nil {
		return interrors.Capture(err)
	}
	// Ensure all stores have directives specified. Defaults should have
	// been set by this point, if the user didn't specify any.
	for name, charmStorage := range charm.Meta().Storage {
		if _, ok := allDirectives[name]; !ok && charmStorage.CountMin > 0 {
			return interrors.Errorf("%w for store %q", applicationerrors.MissingStorageDirective, name)
		}
	}
	return nil
}

// UpdateCloudService updates the cloud service for the specified application, returning an error
// satisfying [applicationerrors.ApplicationNotFoundError] if the application doesn't exist.
func (s *ApplicationService) UpdateCloudService(ctx context.Context, appName, providerID string, sAddrs network.SpaceAddresses) error {
	return s.st.UpsertCloudService(ctx, appName, providerID, sAddrs)
}

// Broker provides access to the k8s cluster to guery the scale
// of a specified application.
type Broker interface {
	Application(string, caas.DeploymentType) caas.Application
}

// CAASUnitTerminating should be called by the CAASUnitTerminationWorker when
// the agent receives a signal to exit. UnitTerminating will return how
// the agent should shutdown.
// We pass in a CAAS broker to get app details from the k8s cluster - we will probably
// make it a service attribute once more use cases emerge.
func (s *ApplicationService) CAASUnitTerminating(ctx context.Context, appName string, unitNum int, broker Broker) (bool, error) {
	// TODO(sidecar): handle deployment other than statefulset
	deploymentType := caas.DeploymentStateful
	restart := true

	switch deploymentType {
	case caas.DeploymentStateful:
		caasApp := broker.Application(appName, caas.DeploymentStateful)
		appState, err := caasApp.State()
		if err != nil {
			return false, interrors.Capture(err)
		}
		var scaleInfo application.ScaleState
		err = s.st.RunAtomic(ctx, func(ctx domain.AtomicContext) error {
			appID, err := s.st.GetApplicationID(ctx, appName)
			if err != nil {
				return interrors.Capture(err)
			}
			scaleInfo, err = s.st.GetApplicationScaleState(ctx, appID)
			return interrors.Capture(err)
		})
		if err != nil {
			return false, interrors.Capture(err)
		}
		if unitNum >= scaleInfo.Scale || unitNum >= appState.DesiredReplicas {
			restart = false
		}
	case caas.DeploymentStateless, caas.DeploymentDaemon:
		// Both handled the same way.
		restart = true
	default:
		return false, interrors.Errorf("unknown deployment type %w", coreerrors.NotSupported)
	}
	return restart, nil
}

// GetApplicationLife looks up the life of the specified application, returning
// an error satisfying [applicationerrors.ApplicationNotFoundError] if the
// application is not found.
func (s *ApplicationService) GetApplicationLife(ctx context.Context, appName string) (corelife.Value, error) {
	var result corelife.Value
	err := s.st.RunAtomic(ctx, func(ctx domain.AtomicContext) error {
		_, appLife, err := s.st.GetApplicationLife(ctx, appName)
		result = appLife.Value()
		return interrors.Errorf("getting life for %q %w", appName, err)
	})
	return result, interrors.Capture(err)
}

// SetApplicationScale sets the application's desired scale value, returning an error
// satisfying [applicationerrors.ApplicationNotFound] if the application is not found.
// This is used on CAAS models.
func (s *ApplicationService) SetApplicationScale(ctx context.Context, appName string, scale int) error {
	if scale < 0 {
		return interrors.Errorf("application scale %d not valid%w", scale, errors.Hide(applicationerrors.ScaleChangeInvalid))
	}
	err := s.st.RunAtomic(ctx, func(ctx domain.AtomicContext) error {
		appID, err := s.st.GetApplicationID(ctx, appName)
		if err != nil {
			return interrors.Capture(err)
		}
		appScale, err := s.st.GetApplicationScaleState(ctx, appID)
		if err != nil {
			return interrors.Errorf("getting application scale state for app %q %w", appID, err)
		}
		s.logger.Tracef(
			"SetScale DesiredScale %v -> %v", appScale.Scale, scale,
		)
		return s.st.SetDesiredApplicationScale(ctx, appID, scale)
	})
	return interrors.Errorf("setting scale for application %q %w", appName, err)
}

// GetApplicationScale returns the desired scale of an application, returning an error
// satisfying [applicationerrors.ApplicationNotFoundError] if the application doesn't exist.
// This is used on CAAS models.
func (s *ApplicationService) GetApplicationScale(ctx context.Context, appName string) (int, error) {
	_, scale, err := s.getApplicationScaleAndID(ctx, appName)
	return scale, interrors.Capture(err)
}

func (s *ApplicationService) getApplicationScaleAndID(ctx context.Context, appName string) (coreapplication.ID, int, error) {
	var (
		scaleState application.ScaleState
		appID      coreapplication.ID
	)
	err := s.st.RunAtomic(ctx, func(ctx domain.AtomicContext) error {
		var err error
		appID, err = s.st.GetApplicationID(ctx, appName)
		if err != nil {
			return interrors.Capture(err)
		}

		scaleState, err = s.st.GetApplicationScaleState(ctx, appID)
		return interrors.Errorf("getting scaling state for %q %w", appName, err)
	})
	return appID, scaleState.Scale, interrors.Capture(err)
}

// ChangeApplicationScale alters the existing scale by the provided change amount, returning the new amount.
// It returns an error satisfying [applicationerrors.ApplicationNotFoundError] if the application
// doesn't exist.
// This is used on CAAS models.
func (s *ApplicationService) ChangeApplicationScale(ctx context.Context, appName string, scaleChange int) (int, error) {
	var newScale int
	err := s.st.RunAtomic(ctx, func(ctx domain.AtomicContext) error {
		appID, err := s.st.GetApplicationID(ctx, appName)
		if err != nil {
			return interrors.Capture(err)
		}
		currentScaleState, err := s.st.GetApplicationScaleState(ctx, appID)
		if err != nil {
			return interrors.Errorf("getting current scale state for %q %w", appName, err)
		}

		newScale = currentScaleState.Scale + scaleChange
		s.logger.Tracef("ChangeScale DesiredScale %v, scaleChange %v, newScale %v", currentScaleState.Scale, scaleChange, newScale)
		if newScale < 0 {
			newScale = currentScaleState.Scale
			return interrors.Errorf(
				"%w: cannot remove more units than currently exist", applicationerrors.ScaleChangeInvalid)

		}
		err = s.st.SetDesiredApplicationScale(ctx, appID, newScale)
		return interrors.Errorf("changing scaling state for %q %w", appName, err)
	})
	return newScale, interrors.Errorf("changing scale for %q %w", appName, err)
}

// SetApplicationScalingState updates the scale state of an application, returning an error
// satisfying [applicationerrors.ApplicationNotFoundError] if the application doesn't exist.
// This is used on CAAS models.
func (s *ApplicationService) SetApplicationScalingState(ctx context.Context, appName string, scaleTarget int, scaling bool) error {
	err := s.st.RunAtomic(ctx, func(ctx domain.AtomicContext) error {
		appID, appLife, err := s.st.GetApplicationLife(ctx, appName)
		if err != nil {
			return interrors.Errorf("getting life for %q %w", appName, err)
		}
		currentScaleState, err := s.st.GetApplicationScaleState(ctx, appID)
		if err != nil {
			return interrors.Errorf("getting current scale state for %q %w", appName, err)
		}

		var scale *int
		if scaling {
			switch appLife {
			case life.Alive:
				// if starting a scale, ensure we are scaling to the same target.
				if !currentScaleState.Scaling && currentScaleState.Scale != scaleTarget {
					return applicationerrors.ScalingStateInconsistent
				}
			case life.Dying, life.Dead:
				// force scale to the scale target when dying/dead.
				scale = &scaleTarget
			}
		}
		err = s.st.SetApplicationScalingState(ctx, appID, scale, scaleTarget, scaling)
		return interrors.Errorf("updating scaling state for %q %w", appName, err)
	})
	return interrors.Errorf("setting scale for %q %w", appName, err)

}

// GetApplicationScalingState returns the scale state of an application, returning an error
// satisfying [applicationerrors.ApplicationNotFoundError] if the application doesn't exist.
// This is used on CAAS models.
func (s *ApplicationService) GetApplicationScalingState(ctx context.Context, appName string) (ScalingState, error) {
	var scaleState application.ScaleState
	err := s.st.RunAtomic(ctx, func(ctx domain.AtomicContext) error {
		appID, err := s.st.GetApplicationID(ctx, appName)
		if err != nil {
			return interrors.Capture(err)
		}
		scaleState, err = s.st.GetApplicationScaleState(ctx, appID)
		return interrors.Errorf("getting scaling state for %q %w", appName, err)
	})
	return ScalingState{
		ScaleTarget: scaleState.ScaleTarget,
		Scaling:     scaleState.Scaling,
	}, interrors.Capture(err)
}

// AgentVersionGetter is responsible for retrieving the agent version for a
// given model.
type AgentVersionGetter interface {
	// GetModelTargetAgentVersion returns the agent version for the specified
	// model.
	GetModelTargetAgentVersion(context.Context, coremodel.UUID) (version.Number, error)
}

// Provider defines the interface for interacting with the underlying model
// provider.
type Provider interface {
	environs.SupportedFeatureEnumerator
}

// ProviderApplicationService defines a service for interacting with the underlying
// model state.
type ProviderApplicationService struct {
	ApplicationService

	modelID            coremodel.UUID
	agentVersionGetter AgentVersionGetter
	provider           providertracker.ProviderGetter[Provider]
}

// NewProviderApplicationService returns a new Service for interacting with a models state.
func NewProviderApplicationService(
	st ApplicationState, deleteSecretState DeleteSecretState,
	modelID coremodel.UUID,
	agentVersionGetter AgentVersionGetter,
	provider providertracker.ProviderGetter[Provider],
	storageRegistryGetter corestorage.ModelStorageRegistryGetter,
	logger logger.Logger,
) *ProviderApplicationService {
	service := NewApplicationService(st, deleteSecretState, storageRegistryGetter, logger)

	return &ProviderApplicationService{
		ApplicationService: *service,
		modelID:            modelID,
		agentVersionGetter: agentVersionGetter,
		provider:           provider,
	}
}

// GetSupportedFeatures returns the set of features that the model makes
// available for charms to use.
// If the agent version cannot be found, an error satisfying
// [modelerrors.NotFound] will be returned.
func (s *ProviderApplicationService) GetSupportedFeatures(ctx context.Context) (assumes.FeatureSet, error) {
	agentVersion, err := s.agentVersionGetter.GetModelTargetAgentVersion(ctx, s.modelID)
	if err != nil {
		return assumes.FeatureSet{}, err
	}

	var fs assumes.FeatureSet
	fs.Add(assumes.Feature{
		Name:        "juju",
		Description: assumes.UserFriendlyFeatureDescriptions["juju"],
		Version:     &agentVersion,
	})

	provider, err := s.provider(ctx)
	if interrors.Is(err, coreerrors.NotSupported) {
		return fs, nil
	} else if err != nil {
		return fs, err
	}

	envFs, err := provider.SupportedFeatures()
	if err != nil {
		return fs, interrors.Errorf("enumerating features supported by environment: %w", err)
	}

	fs.Merge(envFs)

	return fs, nil
}

// WatchableApplicationService provides the API for working with applications and the
// ability to create watchers.
type WatchableApplicationService struct {
	ProviderApplicationService
	watcherFactory WatcherFactory
}

// NewWatchableApplicationService returns a new service reference wrapping the input state.
func NewWatchableApplicationService(
	st ApplicationState, deleteSecretState DeleteSecretState,
	watcherFactory WatcherFactory,
	modelID coremodel.UUID,
	agentVersionGetter AgentVersionGetter,
	provider providertracker.ProviderGetter[Provider],
	storageRegistryGetter corestorage.ModelStorageRegistryGetter,
	logger logger.Logger,
) *WatchableApplicationService {
	service := NewProviderApplicationService(st, deleteSecretState, modelID, agentVersionGetter, provider, storageRegistryGetter, logger)

	return &WatchableApplicationService{
		ProviderApplicationService: *service,
		watcherFactory:             watcherFactory,
	}
}

// WatchApplicationUnitLife returns a watcher that observes changes to the life of any units if an application.
func (s *WatchableApplicationService) WatchApplicationUnitLife(appName string) (watcher.StringsWatcher, error) {
	lifeGetter := func(ctx context.Context, db coredatabase.TxnRunner, ids []string) (map[string]life.Life, error) {
		unitUUIDs, err := transform.SliceOrErr(ids, coreunit.ParseID)
		if err != nil {
			return nil, err
		}
		unitLifes, err := s.st.GetApplicationUnitLife(ctx, appName, unitUUIDs...)
		if err != nil {
			return nil, err
		}
		result := make(map[string]life.Life, len(unitLifes))
		for unitUUID, life := range unitLifes {
			result[unitUUID.String()] = life
		}
		return result, nil
	}
	lifeMapper := domain.LifeStringsWatcherMapperFunc(s.logger, lifeGetter)

	table, query := s.st.InitialWatchStatementUnitLife(appName)
	return s.watcherFactory.NewNamespaceMapperWatcher(table, changestream.All, query, lifeMapper)
}

// WatchApplicationScale returns a watcher that observes changes to an application's scale.
func (s *WatchableApplicationService) WatchApplicationScale(ctx context.Context, appName string) (watcher.NotifyWatcher, error) {
	appID, currentScale, err := s.getApplicationScaleAndID(ctx, appName)
	if err != nil {
		return nil, interrors.Capture(err)
	}

	mask := changestream.Create | changestream.Update
	mapper := func(ctx context.Context, db coredatabase.TxnRunner, changes []changestream.ChangeEvent) ([]changestream.ChangeEvent, error) {
		newScale, err := s.GetApplicationScale(ctx, appName)
		if err != nil {
			return nil, interrors.Capture(err)
		}
		// Only dispatch if the scale has changed.
		if newScale != currentScale {
			currentScale = newScale
			return changes, nil
		}
		return nil, nil
	}
	return s.watcherFactory.NewValueMapperWatcher("application_scale", appID.String(), mask, mapper)
}

// isValidApplicationName returns whether name is a valid application name.
func isValidApplicationName(name string) bool {
	return validApplication.MatchString(name)
}

// isValidReferenceName returns whether name is a valid reference name.
// This ensures that the reference name is both a valid application name
// and a valid charm name.
func isValidReferenceName(name string) bool {
	return isValidApplicationName(name) && isValidCharmName(name)
}
