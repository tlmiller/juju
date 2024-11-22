// Copyright 2024 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package state

import (
	"context"

	"github.com/canonical/sqlair"

	coremachine "github.com/juju/juju/core/machine"
	machineerrors "github.com/juju/juju/domain/machine/errors"
	"github.com/juju/juju/internal/database"
	interrors "github.com/juju/juju/internal/errors"
)

// RequireMachineReboot sets the machine referenced by its UUID as requiring a
// reboot.
//
// Reboot requests are handled through the "machine_requires_reboot" table which
// contains only machine UUID for which a reboot has been requested. This
// function is idempotent.
func (st *State) RequireMachineReboot(ctx context.Context, uuid string) error {
	db, err := st.DB()
	if err != nil {
		return interrors.Capture(err)
	}
	machineUUIDParam := machineUUID{uuid}
	setRebootFlag := `INSERT INTO machine_requires_reboot (machine_uuid) VALUES ($machineUUID.uuid)`
	setRebootFlagStmt, err := sqlair.Prepare(setRebootFlag, machineUUIDParam)
	if err != nil {
		return interrors.Capture(err)
	}
	err = db.Txn(ctx, func(ctx context.Context, tx *sqlair.TX) error {
		return tx.Query(ctx, setRebootFlagStmt, machineUUIDParam).Run()
	})

	if database.IsErrConstraintPrimaryKey(err) {
		// if the same uuid is added twice, do nothing (idempotency)
		return nil
	}
	return interrors.Errorf("requiring reboot of machine %q %w", uuid, err)
}

// ClearMachineReboot removes the reboot flag of the machine referenced by its UUID if a reboot
// has previously been required.
//
// It basically removes the uuid from the "machine_requires_reboot" table if
// present. This function is idempotent.
func (st *State) ClearMachineReboot(ctx context.Context, uuid string) error {
	db, err := st.DB()
	if err != nil {
		return interrors.Capture(err)
	}
	machineUUIDParam := machineUUID{uuid}
	unsetRebootFlag := `DELETE FROM machine_requires_reboot WHERE machine_uuid = $machineUUID.uuid`
	unsetRebootFlagStmt, err := sqlair.Prepare(unsetRebootFlag, machineUUIDParam)
	if err != nil {
		return interrors.Capture(err)
	}
	err = db.Txn(ctx, func(ctx context.Context, tx *sqlair.TX) error {
		return tx.Query(ctx, unsetRebootFlagStmt, machineUUIDParam).Run()
	})
	return interrors.Errorf("cancelling reboot of machine %q %w", uuid, err)
}

// IsMachineRebootRequired checks if the specified machine requires a reboot.
//
// It queries the "machine_requires_reboot" table for the machine UUID to
// determine if a reboot is required. Returns a boolean value indicating if a
// reboot is required, and an error if any occur during the process.
func (st *State) IsMachineRebootRequired(ctx context.Context, uuid string) (bool, error) {
	db, err := st.DB()
	if err != nil {
		return false, interrors.Capture(err)
	}

	var isRebootRequired bool
	machineUUIDParam := machineUUID{uuid}
	isRebootFlag := `SELECT machine_uuid as &machineUUID.uuid  FROM machine_requires_reboot WHERE machine_uuid = $machineUUID.uuid`
	isRebootFlagStmt, err := sqlair.Prepare(isRebootFlag, machineUUIDParam)
	if err != nil {
		return false, interrors.Capture(err)
	}
	err = db.Txn(ctx, func(ctx context.Context, tx *sqlair.TX) error {
		err := tx.Query(ctx, isRebootFlagStmt, machineUUIDParam).Get(&machineUUIDParam)
		if err != nil && !interrors.Is(err, sqlair.ErrNoRows) {
			return interrors.Capture(err)
		}
		isRebootRequired = !interrors.Is(err, sqlair.ErrNoRows)
		return nil
	})

	return isRebootRequired, interrors.Errorf("requiring reboot of machine %q %w", uuid, err)
}

// ShouldRebootOrShutdown determines if a machine should reboot or shutdown
// based on its state and parent's state.
//
// The function first checks if a parent machine exists and requires a reboot.
// If so, it returns ShouldShutdown immediately.
//
// If the parent machine does not require a reboot, the function checks if the
// current machine requires a reboot. If so, it returns ShouldReboot. If neither
// the parent machine nor the current machine require a reboot, it returns
// ShouldDoNothing.
//
// The function also check if there is a grandparent machine, which is not
// supported. In this case, the function returns an
// errors.GrandParentNotSupported.
//
// The function returns any error issued through interaction with the database,
// annotated with the UUID of the machine.
func (st *State) ShouldRebootOrShutdown(ctx context.Context, uuid string) (coremachine.RebootAction, error) {
	db, err := st.DB()
	if err != nil {
		return coremachine.ShouldDoNothing, interrors.Capture(err)
	}

	// Prepare query to get parent UUID
	machineUUIDParam := machineUUID{uuid}
	getParentQuery := `SELECT machine_parent.parent_uuid as &machineUUID.uuid  FROM machine_parent WHERE machine_uuid = $machineUUID.uuid`
	getParentStmt, err := sqlair.Prepare(getParentQuery, machineUUIDParam)
	if err != nil {
		return coremachine.ShouldDoNothing, interrors.Errorf("requiring reboot action for machine %q %w", uuid, err)
	}

	// Prepare query to check if a machine requires reboot
	isRebootFlag := `SELECT machine_uuid as &machineUUID.uuid  FROM machine_requires_reboot WHERE machine_uuid = $machineUUID.uuid`
	isRebootFlagStmt, err := sqlair.Prepare(isRebootFlag, machineUUIDParam)
	if err != nil {
		return coremachine.ShouldDoNothing, interrors.Errorf("requiring reboot action for machine %q %w", uuid, err)
	}

	var parentShouldReboot, machineShouldReboot bool
	err = db.Txn(ctx, func(ctx context.Context, tx *sqlair.TX) error {
		// Get parent UUID
		var machine, parentMachine, grandParentMachine machineUUID
		err := tx.Query(ctx, getParentStmt, machineUUIDParam).Get(&parentMachine)
		if err != nil && !interrors.Is(err, sqlair.ErrNoRows) {
			return interrors.Capture(err)
		}

		// Check that there is no grandparent (it is not supported)
		err = tx.Query(ctx, getParentStmt, parentMachine).Get(&grandParentMachine)
		if err != nil && !interrors.Is(err, sqlair.ErrNoRows) {
			return interrors.Capture(err)
		}
		if err == nil {
			// Grandparent are not supported. If you get there, possible cause are:
			// - db corruption => need investigation, some parent machine have a parent themselves.
			// - design change => new requirements imply that machine can have grandparent.
			//
			// In this later case you will need to update above code to fetch
			// all grandparent is the chain, and check them for reboot. Moreover, be careful of
			// loophole: if we accept grandparent in the actual representation in DQLite, we may
			// have cycle.
			// Moreover, reboot watcher statement and implementation may need to be updated.
			return interrors.Errorf("found  %q parent of %q parent of %q %w", grandParentMachine.UUID, parentMachine.UUID, uuid, machineerrors.GrandParentNotSupported)
		}

		// Check parent reboot status
		if parentMachine.UUID != "" {
			err := tx.Query(ctx, isRebootFlagStmt, parentMachine).Get(&machine)
			if err != nil && !interrors.Is(err, sqlair.ErrNoRows) {
				return interrors.Capture(err)
			}
			parentShouldReboot = !interrors.Is(err, sqlair.ErrNoRows)
			if parentShouldReboot {
				return nil // early exit, no need to check current machine reboot, it will shutdown anyway
			}
		}

		// Check machine reboot status
		err = tx.Query(ctx, isRebootFlagStmt, machineUUIDParam).Get(&machine)
		if err != nil && !interrors.Is(err, sqlair.ErrNoRows) {
			return interrors.Capture(err)
		}
		machineShouldReboot = !interrors.Is(err, sqlair.ErrNoRows)
		return nil
	})
	if err != nil {
		return coremachine.ShouldDoNothing, interrors.Errorf("requiring reboot action for machine %q %w", uuid, err)
	}

	// Parent need reboot
	if parentShouldReboot {
		return coremachine.ShouldShutdown, nil
	}
	// Machine need reboot, with no parent or no parent requesting reboot
	if machineShouldReboot {
		return coremachine.ShouldReboot, nil
	}
	return coremachine.ShouldDoNothing, nil
}
