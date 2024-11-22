// Copyright 2024 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package errors

import interrors "github.com/juju/juju/internal/errors"

const (
	// PublicKeyAlreadyExists indicates that the authorised key already
	// exists for the specified user.
	PublicKeyAlreadyExists = interrors.ConstError("public key already exists")

	// ImportSubjectNotFound indicates that when importing public keys for a
	// subject the source of the public keys has told us that this subject
	// does not exist.
	ImportSubjectNotFound = interrors.ConstError("import subject not found")

	// InvalidPublicKey indicates a problem with a public key where it
	// was unable to be understood.
	InvalidPublicKey = interrors.ConstError("invalid public key")

	// ReservedCommentViolation indicates that a key contains a comment that is
	// reserved within the Juju system and cannot be used.
	ReservedCommentViolation = interrors.ConstError("key contains a reserved comment")

	// UnknownImportSource indicates that an import operation cannot occur
	// because the source of the information is unknown.
	UnknownImportSource = interrors.ConstError("unknown import source")
)
