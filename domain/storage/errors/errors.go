// Copyright 2024 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package errors

import interrors "github.com/juju/juju/internal/errors"

// These errors are used for storage pool operations.
const (
	// MissingPoolTypeError is used when a provider type is empty.
	MissingPoolTypeError = interrors.ConstError("pool provider type is empty")
	// MissingPoolNameError is used when a name is empty.
	MissingPoolNameError = interrors.ConstError("pool name is empty")
	// InvalidPoolNameError is used when a storage pool name is invalid.
	InvalidPoolNameError = interrors.ConstError("pool name is not valid")
	// PoolNotFoundError is used when a storage pool is not found.
	PoolNotFoundError = interrors.ConstError("storage pool is not found")
	// PoolAlreadyExists is used when a storage pool already exists.
	PoolAlreadyExists = interrors.ConstError("storage pool already exists")
	// ErrNoDefaultStoragePool is returned when a storage pool is required but none is specified nor available as a default.
	ErrNoDefaultStoragePool = interrors.ConstError("no storage pool specified and no default available")
)

// These errors are used for storage directives operations.
const (
	// MissingSharedStorageDirectiveError is used when a storage directive for shared storage is not provided.
	MissingSharedStorageDirectiveError = interrors.ConstError("no storage directive specified")
)
