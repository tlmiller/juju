// Copyright 2023 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package errors

import interrors "github.com/juju/juju/internal/errors"

const (
	// MachineNotFound describes an error that occurs when the machine being
	// operated on does not exist.
	MachineNotFound = interrors.ConstError("machine not found")

	// AvailabilityZoneNotFound describes an error that occurs when the required
	// availability zone does not exist.
	AvailabilityZoneNotFound = interrors.ConstError("availability zone not found")

	// NotProvisioned describes an error that occurs when the machine being
	// operated on is not provisioned yet.
	NotProvisioned = interrors.ConstError("machine not provisioned")

	// StatusNotSet describes an error that occurs when the status of a machine
	// or a cloud instance is not set yet.
	StatusNotSet = interrors.ConstError("status not set")

	// InvalidStatus describes a status that is not valid
	InvalidStatus = interrors.ConstError("invalid status")

	// GrandParentNotSupported describes an error that occurs when the operation
	// found a grandparent machine, as it is not currently supported.
	GrandParentNotSupported = interrors.ConstError("grandparent machine are not supported currently")

	// MachineAlreadyExists describes an error that occurs when creating a
	// machine if a machine with the same name already exists.
	MachineAlreadyExists = interrors.ConstError("machine already exists")

	// MachineHasNoParent describes an error that occurs when a machine has no
	// parent.
	MachineHasNoParent = interrors.ConstError("machine has no parent")

	// MachineCloudInstanceAlreadyExists describes an error that occurs
	// when adding cloud instance on a machine that already exists.
	MachineCloudInstanceAlreadyExists = interrors.ConstError("machine cloud instance already exists")
)
