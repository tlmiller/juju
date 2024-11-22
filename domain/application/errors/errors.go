// Copyright 2024 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package errors

import interrors "github.com/juju/juju/internal/errors"

const (
	// ApplicationNotFound describes an error that occurs when the application
	// being operated on does not exist.
	ApplicationNotFound = interrors.ConstError("application not found")

	// ApplicationAlreadyExists describes an error that occurs when the
	// application being created already exists.
	ApplicationAlreadyExists = interrors.ConstError("application already exists")

	// ApplicationNotAlive describes an error that occurs when trying to update an application that is not alive.
	ApplicationNotAlive = interrors.ConstError("application is not alive")

	// ApplicationHasUnits describes an error that occurs when the application
	// being deleted still has associated units.
	ApplicationHasUnits = interrors.ConstError("application has units")

	// ScalingStateInconsistent is returned by SetScalingState when the scaling state
	// is inconsistent with the application scale.
	ScalingStateInconsistent = interrors.ConstError("scaling state is inconsistent")

	// ScaleChangeInvalid is returned when an attempt is made to set an invalid application scale value.
	ScaleChangeInvalid = interrors.ConstError("scale change invalid")

	// MissingStorageDirective describes an error that occurs when expected
	// storage directives are missing.
	MissingStorageDirective = interrors.ConstError("no storage directive specified")

	// ApplicationNameNotValid describes an error when the application is
	// not valid.
	ApplicationNameNotValid = interrors.ConstError("application name not valid")

	// ApplicationIDNotValid describes an error when the application ID is
	// not valid.
	ApplicationIDNotValid = interrors.ConstError("application ID not valid")

	// UnitNotFound describes an error that occurs when the unit being operated on
	// does not exist.
	UnitNotFound = interrors.ConstError("unit not found")

	// UnitAlreadyExists describes an error that occurs when the
	// unit being created already exists.
	UnitAlreadyExists = interrors.ConstError("unit already exists")

	// UnitNotAssigned describes an error that occurs when the unit being operated on
	// is not assigned.
	UnitNotAssigned = interrors.ConstError("unit not assigned")

	// UnitHasSubordinates describes an error that occurs when trying to set a unit's life
	// to Dead but it still has subordinates.
	UnitHasSubordinates = interrors.ConstError("unit has subordinates")

	// UnitHasStorageAttachments describes an error that occurs when trying to set a unit's life
	// to Dead but it still has storage attachments.
	UnitHasStorageAttachments = interrors.ConstError("unit has storage attachments")

	// UnitIsAlive describes an error that occurs when trying to remove a unit that is still alive.
	UnitIsAlive = interrors.ConstError("unit is alive")

	// InvalidApplicationState describes an error where the application state is invalid.
	// There are missing required fields.
	InvalidApplicationState = interrors.ConstError("invalid application state")

	// CharmNotValid describes an error that occurs when the charm is not valid.
	CharmNotValid = interrors.ConstError("charm not valid")

	// CharmOriginNotValid describes an error that occurs when the charm origin is not valid.
	CharmOriginNotValid = interrors.ConstError("charm origin not valid")

	// CharmNameNotValid describes an error that occurs when attempting to get
	// a charm using an invalid name.
	CharmNameNotValid = interrors.ConstError("charm name not valid")

	// CharmSourceNotValid describes an error that occurs when attempting to get
	// a charm using an invalid charm source.
	CharmSourceNotValid = interrors.ConstError("charm source not valid")

	// CharmNotFound describes an error that occurs when a charm cannot be found.
	CharmNotFound = interrors.ConstError("charm not found")

	// LXDProfileNotFound describes an error that occurs when an LXD profile
	// cannot be found.
	LXDProfileNotFound = interrors.ConstError("LXD profile not found")

	// CharmAlreadyExists describes an error that occurs when a charm already
	// exists for the given natural key.
	CharmAlreadyExists = interrors.ConstError("charm already exists")

	// CharmRevisionNotValid describes an error that occurs when attempting to get
	// a charm using an invalid revision.
	CharmRevisionNotValid = interrors.ConstError("charm revision not valid")

	// CharmMetadataNotValid describes an error that occurs when the charm metadata
	// is not valid.
	CharmMetadataNotValid = interrors.ConstError("charm metadata not valid")

	// CharmManifestNotValid describes an error that occurs when the charm manifest
	// is not valid.
	CharmManifestNotValid = interrors.ConstError("charm manifest not valid")

	// CharmBaseNameNotValid describes an error that occurs when the charm base
	// name is not valid.
	CharmBaseNameNotValid = interrors.ConstError("charm base name not valid")

	// CharmBaseNameNotSupported describes an error that occurs when the charm
	// base name is not supported.
	CharmBaseNameNotSupported = interrors.ConstError("charm base name not supported")

	// CharmRelationKeyConflict describes an error that occurs when the charm
	// has multiple relations with the same name
	CharmRelationNameConflict = interrors.ConstError("charm relation name conflict")

	// CharmRelationReservedNameMisuse describes an error that occurs when the charm
	// relation name is a reserved name which it is not allowed to use.
	CharmRelationReservedNameMisuse = interrors.ConstError("charm relation reserved name misuse")

	// CharmRelationRoleNotValid describes an error that occurs when the charm
	// relation roles is not valid. Either it is an unknown role, or it has the
	// wrong value.
	CharmRelationRoleNotValid = interrors.ConstError("charm relation role not valid")

	// ResourceNotFound describes an error that occurs when a resource is
	// not found.
	ResourceNotFound = interrors.ConstError("resource not found")

	// UnknownResourceType describes an error where the resource type is
	// not oci-image or file.
	UnknownResourceType = interrors.ConstError("unknown resource type")

	// ResourceNameNotValid describes an error where the resource name is not
	// valid, usually because it's empty.
	ResourceNameNotValid = interrors.ConstError("resource name not valid")
)
