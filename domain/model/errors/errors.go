// Copyright 2023 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package errors

import (
	"github.com/juju/errors"
)

const (
	// AgentVersionNotSupported describes an error that occurs when then agent
	// version chosen for model is not supported with respect to the currently
	// running controller.
	AgentVersionNotSupported = errors.ConstError("agent version not supported")

	// AlreadyExists describes an error that occurs when a model already exists.
	AlreadyExists = errors.ConstError("model already exists")

	// AlreadyActivated describes an error that occurs when an attempt is made
	// to activate a model that has already been activated.
	AlreadyActivated = errors.ConstError("model already activated")

	// ConstraintsNotFound describes an error that occurs when no model
	// constraints have been set for the model but they are requested.
	ConstraintsNotFound = errors.ConstError("model constraints not found")

	// ModelNamespaceNotFound describes an error that occurs when no database
	// namespace for a model exists.
	ModelNamespaceNotFound = errors.ConstError("model namespace not found")

	// NotFound describes an error that occurs when the model being operated on
	// does not exist.
	NotFound = errors.ConstError("model not found")

	// SecretBackendAlreadySet describes an error that occurs when a model's
	// secret backend has already been set.
	SecretBackendAlreadySet = errors.ConstError("secret backend already set")

	// UserNotFoundOnModel describes an error that occurs when information about
	// a user on a particular model cannot be found. This does not mean the user
	// does not exist.
	UserNotFoundOnModel = errors.ConstError("user not found on model")

	// AuthTypeDoesNotExist describes an error that occurs when an
	// authentication type does not exist.
	AuthTypeDoesNotExist = errors.ConstError("auth type not found")

	// AuthTypeNotFoundInCloud describes an error that occurs when an
	// authentication type in a particular cloud cannot be found. This does
	// not mean that the cloud or auth type does not exist.
	AuthTypeNotFoundInCloud = errors.ConstError("auth type not found in cloud")

	// NoCredentialSupplied describes an error that occurs when no credentials
	// are provided and the cloud requires non-empty credentials.
	NoCredentialSupplied = errors.ConstError("no credentials provided for cloud that require non-empty credentials")
)
