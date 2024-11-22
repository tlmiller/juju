// Copyright 2024 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package errors

import interrors "github.com/juju/juju/internal/errors"

const (
	// AlreadyExists describes an error that occurs when a secret backend already exists.
	AlreadyExists = interrors.ConstError("secret backend already exists")

	// RefCountAlreadyExists describes an error that occurs when a secret backend reference count record already exists.
	RefCountAlreadyExists = interrors.ConstError("secret backend reference count already exists")

	// RefCountNotFound describes an error that occurs when a secret backend reference count record is not found.
	RefCountNotFound = interrors.ConstError("secret backend reference count not found")

	// NotFound describes an error that occurs when the secret backend being operated on does not exist.
	NotFound = interrors.ConstError("secret backend not found")

	// NotValid describes an error that occurs when the secret backend being operated on is not valid.
	NotValid = interrors.ConstError("secret backend not valid")

	// Forbidden describes an error that occurs when the operation is forbidden.
	Forbidden = interrors.ConstError("secret backend operation forbidden")

	// NotSupported describes an error that occurs when the secret backend is not supported.
	NotSupported = interrors.ConstError("secret backend not supported")
)
