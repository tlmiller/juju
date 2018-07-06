// Copyright 2018 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package params

//MachineSeriesUpgradeStatus is the current status a machine series upgrade
type MachineSeriesUpgradeStatus string

const (
	MachineSeriesUpgradeNotStarted    MachineSeriesUpgradeStatus = "NotStarted"
	MachineSeriesUpgradeStarted       MachineSeriesUpgradeStatus = "Started"
	MachineSeriesUpgradeUnitsRunning  MachineSeriesUpgradeStatus = "UnitsRunning"
	MachineSeriesUpgradeJujuComplete  MachineSeriesUpgradeStatus = "JujuComplete"
	MachineSeriesUpgradeAgentsStopped MachineSeriesUpgradeStatus = "AgentsStopped"
	MachineSeriesUpgradeComplete      MachineSeriesUpgradeStatus = "Complete"
)

//UnitSeriesUpgradeStatus is the current status of an upgrade
type UnitSeriesUpgradeStatus string

const (
	UnitNotStarted UnitSeriesUpgradeStatus = "NotStarted"
	UnitStarted    UnitSeriesUpgradeStatus = "Started"
	UnitErrored    UnitSeriesUpgradeStatus = "Errored"
	UnitCompleted  UnitSeriesUpgradeStatus = "Completed"
)
