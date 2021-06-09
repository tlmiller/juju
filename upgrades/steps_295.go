// Copyright 2020 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package upgrades

import (
	"github.com/juju/errors"
	"github.com/juju/juju/caas"
	"github.com/juju/juju/environs"
)

type UpgradeKubernetesClusterCredential interface {
	InClusterCredentialUpgrade() error
}

// stateStepsFor295 returns upgrade steps for juju 2.9.5
func stateStepsFor295() []Step {
	return []Step{
		&upgradeStep{
			description: "prepare k8s controller for in cluster credentials",
			targets:     []Target{DatabaseMaster},
			run:         controllerInClusterCredentials,
		},
	}
}

func controllerInClusterCredentials(context Context) error {
	cloudSpec, modelCfg, uuid, err := context.State().KubernetesInClusterCredentialSpec()
	if errors.IsNotFound(err) {
		return nil
	} else if err != nil {
		return errors.Trace(err)
	}

	cloudSpec.IsControllerCloud = false
	broker, err := caas.New(environs.OpenParams{
		ControllerUUID: uuid,
		Cloud:          cloudSpec,
		Config:         modelCfg,
	})
	if err != nil {
		return errors.Trace(err)
	}

	upgrader, ok := broker.(UpgradeKubernetesClusterCredential)
	if !ok {
		return errors.New("caas broker does not implement kubernetes cluster credential upgrader")
	}

	upgrader.InClusterCredentialUpgrade()

	return nil
}
