// Copyright 2024 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package errors

import interrors "github.com/juju/juju/internal/errors"

const (
	// PermissionDenied describes an error that occurs when the secret being operated on
	// does not have the required authorisation set for the caller.
	PermissionDenied = interrors.ConstError("permission denied")

	// SecretNotFound describes an error that occurs when the secret being operated on
	// does not exist.
	SecretNotFound = interrors.ConstError("secret not found")

	// SecretIsNotLocal describes an error that occurs when a secret is not from the current model.
	SecretIsNotLocal = interrors.ConstError("secret is from a different model")

	// SecretLabelAlreadyExists describes an error that occurs when there's already a secret label for
	// a specified secret owner.
	SecretLabelAlreadyExists = interrors.ConstError("secret label already exists")

	// SecretRevisionNotFound describes an error that occurs when the secret revision being operated on
	// does not exist.
	SecretRevisionNotFound = interrors.ConstError("secret revision not found")

	// SecretConsumerNotFound describes an error that occurs when the secret consumer being operated on is not found.
	SecretConsumerNotFound = interrors.ConstError("secret consumer not found")

	// AutoPruneNotSupported describes an error that occurs when a charm secret tries to set auto prune on a secret.
	AutoPruneNotSupported = interrors.ConstError("charm secrets do not support auto prune")

	// InvalidSecretPermissionChange describes an error that occurs when an attempt is made to update a secret permission
	// and the scope or subject type is changed.
	InvalidSecretPermissionChange = interrors.ConstError("cannot change a secret permission scope or subject type")

	// SecretAccessScopeNotFound describes an error that occurs when the secret access scope
	// being operated on does not exist.
	SecretAccessScopeNotFound = interrors.ConstError("secret access scope not found")

	// MissingSecretBackendID describes an error that occurs when importing a secret and the backend doesn't exist.
	MissingSecretBackendID = interrors.ConstError("missing secret backend id")
)
