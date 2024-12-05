// Copyright 2024 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package state

import (
	"context"
	"database/sql"

	"github.com/canonical/sqlair"
	"github.com/juju/collections/transform"

	"github.com/juju/juju/core/instance"
	"github.com/juju/juju/core/machine"
	"github.com/juju/juju/core/status"
	"github.com/juju/juju/domain"
	machineerrors "github.com/juju/juju/domain/machine/errors"
	networkerrors "github.com/juju/juju/domain/network/errors"
	"github.com/juju/juju/internal/database"
	"github.com/juju/juju/internal/errors"
)

// HardwareCharacteristics returns the hardware characteristics struct with
// data retrieved from the machine cloud instance table.
func (st *State) HardwareCharacteristics(
	ctx context.Context,
	machineUUID string,
) (*instance.HardwareCharacteristics, error) {
	db, err := st.DB()
	if err != nil {
		return nil, errors.Capture(err)
	}
	query := `
SELECT    &instanceDataResult.*
FROM      v_hardware_characteristics AS v
WHERE     v.machine_uuid = $instanceDataResult.machine_uuid`
	machineUUIDQuery := instanceDataResult{
		MachineUUID: machineUUID,
	}
	stmt, err := st.Prepare(query, machineUUIDQuery)
	if err != nil {
		return nil, errors.Errorf("preparing retrieve hardware characteristics statement %w", err)
	}

	var row instanceDataResult
	if err := db.Txn(ctx, func(ctx context.Context, tx *sqlair.TX) error {
		err := tx.Query(ctx, stmt, machineUUIDQuery).Get(&row)
		if errors.Is(err, sql.ErrNoRows) {
			return errors.Errorf("machine: %q %w", machineUUID, machineerrors.NotProvisioned)
		}
		return errors.Errorf("querying machine cloud instance for machine %q %w", machineUUID, err)
	}); err != nil {
		return nil, errors.Capture(err)
	}
	return row.toHardwareCharacteristics(), nil
}

// AvailabilityZone returns the availability zone for the specified machine.
// If no hardware characteristics are set for the machine, it returns
// [machineerrors.AvailabilityZoneNotFound].
func (st *State) AvailabilityZone(
	ctx context.Context,
	machineUUID string,
) (string, error) {
	db, err := st.DB()
	if err != nil {
		return "", errors.Capture(err)
	}

	query := `
SELECT    &instanceDataResult.availability_zone_name
FROM      v_hardware_characteristics AS v
WHERE     v.machine_uuid = $instanceDataResult.machine_uuid`
	machineUUIDQuery := instanceDataResult{
		MachineUUID: machineUUID,
	}
	stmt, err := st.Prepare(query, machineUUIDQuery)
	if err != nil {
		return "", errors.Errorf("preparing retrieve hardware characteristics statement %w", err)
	}

	var row instanceDataResult
	if err := db.Txn(ctx, func(ctx context.Context, tx *sqlair.TX) error {
		err := tx.Query(ctx, stmt, machineUUIDQuery).Get(&row)
		if errors.Is(err, sql.ErrNoRows) {
			return errors.Errorf("machine cloud instance for machine %q %w", machineUUID, machineerrors.AvailabilityZoneNotFound)
		}
		return errors.Errorf("querying machine cloud instance for machine %q %w", machineUUID, err)
	}); err != nil {
		return "", errors.Capture(err)
	}
	if row.AvailabilityZone == nil {
		return "", nil
	}
	return *row.AvailabilityZone, nil
}

// SetMachineCloudInstance sets an entry in the machine cloud instance table
// along with the instance tags and the link to a lxd profile if any.
func (st *State) SetMachineCloudInstance(
	ctx context.Context,
	machineUUID string,
	instanceID instance.Id,
	displayName string,
	hardwareCharacteristics *instance.HardwareCharacteristics,
) error {
	db, err := st.DB()
	if err != nil {
		return errors.Capture(err)
	}

	setInstanceData := `
INSERT INTO machine_cloud_instance (*)
VALUES ($instanceData.*)
`
	setInstanceDataStmt, err := st.Prepare(setInstanceData, instanceData{})
	if err != nil {
		return errors.Capture(err)
	}

	setInstanceTags := `
INSERT INTO instance_tag (*)
VALUES ($instanceTag.*)
`
	setInstanceTagStmt, err := st.Prepare(setInstanceTags, instanceTag{})
	if err != nil {
		return errors.Capture(err)
	}

	azName := availabilityZoneName{}
	if hardwareCharacteristics != nil && hardwareCharacteristics.AvailabilityZone != nil {
		az := *hardwareCharacteristics.AvailabilityZone
		azName = availabilityZoneName{Name: az}
	}
	retrieveAZUUID := `
SELECT &availabilityZoneName.uuid
FROM   availability_zone
WHERE  availability_zone.name = $availabilityZoneName.name
`
	retrieveAZUUIDStmt, err := st.Prepare(retrieveAZUUID, azName)
	if err != nil {
		return errors.Capture(err)
	}

	return db.Txn(ctx, func(ctx context.Context, tx *sqlair.TX) error {
		instanceData := instanceData{
			MachineUUID: machineUUID,
			InstanceID:  instanceID.String(),
			DisplayName: displayName,
		}
		if hardwareCharacteristics != nil {
			instanceData.Arch = hardwareCharacteristics.Arch
			instanceData.Mem = hardwareCharacteristics.Mem
			instanceData.RootDisk = hardwareCharacteristics.RootDisk
			instanceData.RootDiskSource = hardwareCharacteristics.RootDiskSource
			instanceData.CPUCores = hardwareCharacteristics.CpuCores
			instanceData.CPUPower = hardwareCharacteristics.CpuPower
			instanceData.VirtType = hardwareCharacteristics.VirtType
			if hardwareCharacteristics.AvailabilityZone != nil && *hardwareCharacteristics.AvailabilityZone != "" {
				azUUID := availabilityZoneName{}
				if err := tx.Query(ctx, retrieveAZUUIDStmt, azName).Get(&azUUID); err != nil {
					if errors.Is(err, sql.ErrNoRows) {
						return errors.Errorf("%w %q for machine %q", networkerrors.AvailabilityZoneNotFound, *hardwareCharacteristics.AvailabilityZone, machineUUID)
					}
					return errors.Errorf("cannot retrieve availability zone %q for machine uuid %q: %w", *hardwareCharacteristics.AvailabilityZone, machineUUID, err)
				}
				instanceData.AvailabilityZoneUUID = &azUUID.UUID
			}
		}
		if err := tx.Query(ctx, setInstanceDataStmt, instanceData).Run(); err != nil {
			if database.IsErrConstraintPrimaryKey(err) {
				return errors.Errorf("%w for machine %q", machineerrors.MachineCloudInstanceAlreadyExists, machineUUID)
			}
			return errors.Errorf("inserting machine cloud instance for machine %q %w", machineUUID, err)
		}
		if instanceTags := tagsFromHardwareCharacteristics(machineUUID, hardwareCharacteristics); len(instanceTags) > 0 {
			if err := tx.Query(ctx, setInstanceTagStmt, instanceTags).Run(); err != nil {
				return errors.Errorf("inserting instance tags for machine %q %w", machineUUID, err)
			}
		}
		return nil
	})
}

// DeleteMachineCloudInstance removes an entry in the machine cloud instance
// table along with the instance tags and the link to a lxd profile if any, as
// well as any associated status data.
func (st *State) DeleteMachineCloudInstance(
	ctx context.Context,
	mUUID string,
) error {
	db, err := st.DB()
	if err != nil {
		return errors.Capture(err)
	}

	// Prepare query for deleting machine cloud instance.
	deleteInstanceQuery := `
DELETE FROM machine_cloud_instance
WHERE machine_uuid=$machineUUID.uuid
`
	machineUUIDParam := machineUUID{
		UUID: mUUID,
	}
	deleteInstanceStmt, err := st.Prepare(deleteInstanceQuery, machineUUIDParam)
	if err != nil {
		return errors.Capture(err)
	}

	// Prepare query for deleting instance tags.
	deleteInstanceTagsQuery := `
DELETE FROM instance_tag
WHERE machine_uuid=$machineUUID.uuid
`
	deleteInstanceTagStmt, err := st.Prepare(deleteInstanceTagsQuery, machineUUIDParam)
	if err != nil {
		return errors.Capture(err)
	}

	// Prepare query for deleting cloud instance status.
	deleteInstanceStatusQuery := `DELETE FROM machine_cloud_instance_status WHERE machine_uuid=$machineUUID.uuid`
	deleteInstanceStatusStmt, err := st.Prepare(deleteInstanceStatusQuery, machineUUIDParam)
	if err != nil {
		return errors.Capture(err)
	}

	// Prepare query for deleting cloud instance status data.
	deleteInstanceStatusDataQuery := `DELETE FROM machine_cloud_instance_status_data WHERE machine_uuid=$machineUUID.uuid`
	deleteInstanceStatusDataStmt, err := st.Prepare(deleteInstanceStatusDataQuery, machineUUIDParam)
	if err != nil {
		return errors.Capture(err)
	}

	return db.Txn(ctx, func(ctx context.Context, tx *sqlair.TX) error {
		// Delete the machine cloud instance status data. No need to return
		// error if no status data is set for the instance while deleting.
		if err := tx.Query(ctx, deleteInstanceStatusDataStmt, machineUUIDParam).Run(); err != nil && !errors.Is(err, sqlair.ErrNoRows) {
			return errors.Errorf("deleting machine cloud instance status data for machine %q %w", mUUID, domain.CoerceError(err))
		}

		// Delete the machine cloud instance status. No need to return error if
		// no status is set for the instance while deleting.
		if err := tx.Query(ctx, deleteInstanceStatusStmt, machineUUIDParam).Run(); err != nil && !errors.Is(err, sqlair.ErrNoRows) {
			return errors.Errorf("deleting machine cloud instance status for machine %q %w", mUUID, domain.CoerceError(err))
		}

		// Delete the machine cloud instance.
		if err := tx.Query(ctx, deleteInstanceStmt, machineUUIDParam).Run(); err != nil {
			return errors.Errorf("deleting machine cloud instance for machine %q: %w", mUUID, domain.CoerceError(err))
		}

		// Delete the machine cloud instance tags.
		if err := tx.Query(ctx, deleteInstanceTagStmt, machineUUIDParam).Run(); err != nil {
			return errors.Errorf("deleting instance tags for machine %q %w", mUUID, domain.CoerceError(err))
		}
		return nil
	})
}

// InstanceID returns the cloud specific instance id for this machine.
// If the machine is not provisioned, it returns a
// [machineerrors.NotProvisionedError].
func (st *State) InstanceID(ctx context.Context, mUUID string) (string, error) {
	db, err := st.DB()
	if err != nil {
		return "", errors.Capture(err)
	}

	mUUIDParam := machineUUID{UUID: mUUID}
	query := `
SELECT &instanceID.instance_id
FROM   machine_cloud_instance
WHERE  machine_uuid = $machineUUID.uuid;`
	queryStmt, err := st.Prepare(query, mUUIDParam, instanceID{})
	if err != nil {
		return "", errors.Capture(err)
	}

	var instanceId string
	err = db.Txn(ctx, func(ctx context.Context, tx *sqlair.TX) error {
		var result instanceID
		err := tx.Query(ctx, queryStmt, mUUIDParam).Get(&result)
		if err != nil {
			if errors.Is(err, sqlair.ErrNoRows) {
				return errors.Errorf("machine: %q %w", mUUID, machineerrors.NotProvisioned)
			}
			return errors.Errorf("querying instance for machine %q %w", mUUID, err)
		}

		instanceId = result.ID
		return nil
	})
	if err != nil {
		return "", errors.Capture(err)
	}
	return instanceId, nil
}

// InstanceIDAndName returns the cloud specific instance ID and display name for
// this machine.
// If the machine is not provisioned, it returns a
// [machineerrors.NotProvisionedError].
func (st *State) InstanceIDAndName(ctx context.Context, mUUID string) (string, string, error) {
	db, err := st.DB()
	if err != nil {
		return "", "", errors.Capture(err)
	}

	mUUIDParam := machineUUID{UUID: mUUID}
	query := `
SELECT &instanceIDAndDisplayName.*
FROM   machine_cloud_instance
WHERE  machine_uuid = $machineUUID.uuid;`
	queryStmt, err := st.Prepare(query, mUUIDParam, instanceIDAndDisplayName{})
	if err != nil {
		return "", "", errors.Capture(err)
	}

	var (
		instanceID, instanceName string
	)
	err = db.Txn(ctx, func(ctx context.Context, tx *sqlair.TX) error {
		var result instanceIDAndDisplayName
		err := tx.Query(ctx, queryStmt, mUUIDParam).Get(&result)
		if err != nil {
			if errors.Is(err, sqlair.ErrNoRows) {
				return errors.Errorf("machine: %q %w", mUUID, machineerrors.NotProvisioned)
			}
			return errors.Errorf("querying display name for machine %q %w", mUUID, err)
		}

		instanceID = result.ID
		instanceName = result.Name
		return nil
	})
	if err != nil {
		return "", "", errors.Capture(err)
	}
	return instanceID, instanceName, nil
}

// GetInstanceStatus returns the cloud specific instance status for the given
// machine.
// It returns NotFound if the machine does not exist.
// It returns a StatusNotSet if the instance status is not set.
// Idempotent.
func (st *State) GetInstanceStatus(ctx context.Context, mName machine.Name) (status.StatusInfo, error) {
	db, err := st.DB()
	if err != nil {
		return status.StatusInfo{}, errors.Capture(err)
	}

	// Prepare query for machine uuid (to be used in
	// machine_cloud_instance_status and machine_cloud_instance_status_data
	// tables)
	machineNameParam := machineName{Name: mName}
	machineUUID := machineUUID{}
	uuidQuery := `SELECT uuid AS &machineUUID.* FROM machine WHERE name = $machineName.name`
	uuidQueryStmt, err := st.Prepare(uuidQuery, machineNameParam, machineUUID)
	if err != nil {
		return status.StatusInfo{}, errors.Capture(err)
	}

	// Prepare query for combined machine cloud instance status and the status
	// data (to get them both in one transaction, as this a a relatively
	// frequent retrieval).
	machineStatusParam := machineStatusWithData{}
	statusCombinedQuery := `
SELECT &machineStatusWithData.*
FROM machine_cloud_instance_status AS st
LEFT JOIN machine_cloud_instance_status_data AS st_data
ON st.machine_uuid = st_data.machine_uuid
WHERE st.machine_uuid = $machineUUID.uuid`
	statusCombinedQueryStmt, err := st.Prepare(statusCombinedQuery, machineUUID, machineStatusParam)
	if err != nil {
		return status.StatusInfo{}, errors.Capture(err)
	}

	var instanceStatusWithData machineStatusData
	err = db.Txn(ctx, func(ctx context.Context, tx *sqlair.TX) error {
		// Query for the machine uuid
		err := tx.Query(ctx, uuidQueryStmt, machineNameParam).Get(&machineUUID)
		if err != nil {
			if errors.Is(err, sqlair.ErrNoRows) {
				return machineerrors.MachineNotFound
			}
			return errors.Errorf("querying uuid for machine %q %w", mName, err)
		}

		// Query for the machine cloud instance status and status data combined
		err = tx.Query(ctx, statusCombinedQueryStmt, machineUUID).GetAll(&instanceStatusWithData)
		if err != nil {
			if errors.Is(err, sqlair.ErrNoRows) {
				return errors.Errorf("machine: %q %w", mName, machineerrors.StatusNotSet)
			}
			return errors.Errorf("querying machine status and status data for machine %q %w", mName, err)
		}

		return nil
	})

	if err != nil {
		return status.StatusInfo{}, errors.Capture(err)
	}

	// Transform the status data slice into a status.Data map.
	statusDataResult := transform.SliceToMap(instanceStatusWithData, func(d machineStatusWithData) (string, interface{}) {
		return d.Key, d.Data
	})

	instanceStatus := status.StatusInfo{
		Message: instanceStatusWithData[0].Message,
		Since:   instanceStatusWithData[0].Updated,
		Data:    statusDataResult,
	}

	// Convert the internal status id from the (machine_cloud_instance_status_value table)
	// into the core status.Status type.
	instanceStatus.Status = instanceStatusWithData[0].toCoreInstanceStatusValue()

	return instanceStatus, nil
}

// SetInstanceStatus sets the cloud specific instance status for this
// machine.
// It returns [machineerrors.MachineNotFound] if the machine does not exist.
func (st *State) SetInstanceStatus(ctx context.Context, mName machine.Name, newStatus status.StatusInfo) error {
	db, err := st.DB()
	if err != nil {
		return errors.Capture(err)
	}

	// Prepare the new status to be set.
	instanceStatus := machineStatusWithData{}

	instanceStatus.StatusID = fromCoreInstanceStatusValue(newStatus.Status)
	instanceStatus.Message = newStatus.Message
	instanceStatus.Updated = newStatus.Since
	instanceStatusData := transform.MapToSlice(newStatus.Data, func(key string, value interface{}) []machineStatusWithData {
		return []machineStatusWithData{{Key: key, Data: value.(string)}}
	})

	// Prepare query for machine uuid
	machineNameParam := machineName{Name: mName}
	mUUID := machineUUID{}
	queryMachine := `SELECT uuid AS &machineUUID.* FROM machine WHERE name = $machineName.name`
	queryMachineStmt, err := st.Prepare(queryMachine, machineNameParam, mUUID)
	if err != nil {
		return errors.Capture(err)
	}

	// Prepare query for setting the machine cloud instance status
	statusQuery := `
INSERT INTO machine_cloud_instance_status (machine_uuid, status_id, message, updated_at)
VALUES ($machineUUID.uuid, $machineStatusWithData.status_id, $machineStatusWithData.message, $machineStatusWithData.updated_at)
  ON CONFLICT (machine_uuid)
  DO UPDATE SET status_id = excluded.status_id, message = excluded.message, updated_at = excluded.updated_at
`
	statusQueryStmt, err := st.Prepare(statusQuery, mUUID, instanceStatus)
	if err != nil {
		return errors.Capture(err)
	}

	// Prepare query for setting the machine cloud instance status data
	statusDataQuery := `
INSERT INTO machine_cloud_instance_status_data (machine_uuid, "key", data)
VALUES ($machineUUID.uuid, $machineStatusWithData.key, $machineStatusWithData.data)
  ON CONFLICT (machine_uuid, "key") DO UPDATE SET data = excluded.data
`
	statusDataQueryStmt, err := st.Prepare(statusDataQuery, mUUID, instanceStatus)
	if err != nil {
		return errors.Capture(err)
	}

	return db.Txn(ctx, func(ctx context.Context, tx *sqlair.TX) error {
		// Query for the machine uuid
		err := tx.Query(ctx, queryMachineStmt, machineNameParam).Get(&mUUID)
		if err != nil {
			if errors.Is(err, sqlair.ErrNoRows) {
				return machineerrors.MachineNotFound
			}
			return errors.Errorf("querying uuid for machine %q %w", mName, err)
		}

		// Query for setting the machine cloud instance status
		err = tx.Query(ctx, statusQueryStmt, mUUID, instanceStatus).Run()
		if err != nil {
			return errors.Errorf("setting machine status for machine %q %w", mName, err)
		}

		// Query for setting the machine cloud instance status data if
		// instanceStatusData is not empty.
		if len(instanceStatusData) > 0 {
			err = tx.Query(ctx, statusDataQueryStmt, mUUID, instanceStatusData).Run()
			if err != nil {
				return errors.Errorf("setting machine status data for machine %q %w", mName, err)
			}
		}
		return nil
	})
}
