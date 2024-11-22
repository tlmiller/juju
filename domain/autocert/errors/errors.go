// Copyright 2024 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package errors

import interrors "github.com/juju/juju/internal/errors"

const (
	// NotFound describes an error that occurs when a cert cannot be found.
	NotFound = interrors.ConstError("cert not found")
)
