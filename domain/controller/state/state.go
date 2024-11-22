// Copyright 2024 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package state

import (
	"context"

	"github.com/canonical/sqlair"

	"github.com/juju/juju/core/database"
	"github.com/juju/juju/core/model"
	"github.com/juju/juju/domain"
	interrors "github.com/juju/juju/internal/errors"
)

// State represents a type for interacting with the underlying state.
type State struct {
	*domain.StateBase
}

// NewState returns a new State for interacting with the underlying state.
func NewState(factory database.TxnRunnerFactory) *State {
	return &State{
		StateBase: domain.NewStateBase(factory),
	}
}

// ControllerModelUUID returns the model UUID of the controller model.
func (st *State) ControllerModelUUID(ctx context.Context) (model.UUID, error) {
	db, err := st.DB()
	if err != nil {
		return "", interrors.Capture(err)
	}

	var uuid controllerModelUUID
	stmt, err := st.Prepare(`
SELECT &controllerModelUUID.model_uuid
FROM   controller
`, uuid)
	if err != nil {
		return "", interrors.Errorf("preparing select controller model uuid statement %w", err)
	}
	err = db.Txn(ctx, func(ctx context.Context, tx *sqlair.TX) error {
		err := tx.Query(ctx, stmt).Get(&uuid)
		if interrors.Is(err, sqlair.ErrNoRows) {
			// This should never reasonably happen.
			return interrors.Errorf("internal error: controller model uuid not found")
		}
		return err
	})
	if err != nil {
		return "", interrors.Errorf("getting controller model uuid %w", err)
	}

	return model.UUID(uuid.UUID), nil
}
