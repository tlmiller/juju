// Copyright 2023 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package errors

import "github.com/juju/juju/internal/errors"

const (
	// ErrNotFound is returned when a path is not found.
	ErrNotFound = errors.ConstError("path not found")

	// ErrHashAndSizeAlreadyExists is returned when a hash already exists, but
	// the associated size is different. This should never happen, it means that
	// there is a collision in the hash function.
	ErrHashAndSizeAlreadyExists = errors.ConstError("hash exists for different file size")

	// ErrHashAlreadyExists is returned when a hash already exists.
	ErrHashAlreadyExists = errors.ConstError("hash already exists")
)
