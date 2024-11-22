// Copyright 2023 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package errors

import interrors "github.com/juju/juju/internal/errors"

const (
	// AlreadyStarted states that the upgrade could not be started.
	// This error occurs when the upgrade is already in progress.
	AlreadyStarted = interrors.ConstError("upgrade already started")
	// AlreadyExists states that an upgrade operation has already been created.
	// This error can occur when an upgrade is created.
	AlreadyExists = interrors.ConstError("upgrade already exists")
	// NotFound states that an upgrade operation cannot be found where one is
	// expected.
	NotFound = interrors.ConstError("upgrade not found")
)
