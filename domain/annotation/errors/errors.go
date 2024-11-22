// Copyright 2023 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package errors

import interrors "github.com/juju/juju/internal/errors"

const (
	// UnknownKind is raised when the Kind of an ID provided to the annotations
	// state layer is not recognized
	UnknownKind = interrors.ConstError("unknown kind")
)
