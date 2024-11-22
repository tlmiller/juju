// Copyright 2024 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package errors

import interrors "github.com/juju/juju/internal/errors"

const (
	// SpaceAlreadyExists is returned when a space already exists.
	SpaceAlreadyExists = interrors.ConstError("space already exists")

	// SpaceNotFound is returned when a space is not found.
	SpaceNotFound = interrors.ConstError("space not found")

	// SubnetNotFound is returned when a subnet is not found.
	SubnetNotFound = interrors.ConstError("subnet not found")

	// SpaceNameNotValid is returned when a space name is not valid.
	SpaceNameNotValid = interrors.ConstError("space name is not valid")

	// AvailabilityZoneNotFound is returned when an availability zone is
	// not found.
	AvailabilityZoneNotFound = interrors.ConstError("availability zone not found")
)
