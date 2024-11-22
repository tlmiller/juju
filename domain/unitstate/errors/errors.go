// Copyright 2024 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package errors

import interrors "github.com/juju/juju/internal/errors"

const (
	// UnitNotFound describes an error that occurs when
	// the unit being operated on does not exist.
	UnitNotFound = interrors.ConstError("unit not found")
)
