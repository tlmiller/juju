// Copyright 2015 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

// +build go1.3

package lxdclient_test

import (
	"io/ioutil"
	"path/filepath"

	"github.com/juju/errors"
	jc "github.com/juju/testing/checkers"
	"github.com/lxc/lxd"
	gc "gopkg.in/check.v1"
	goyaml "gopkg.in/yaml.v2"

	"github.com/juju/juju/container/lxd/lxdclient"
)

var (
	_ = gc.Suite(&configSuite{})
	_ = gc.Suite(&configFunctionalSuite{})
)

type configBaseSuite struct {
	lxdclient.BaseSuite

	remote lxdclient.Remote
}

func (s *configBaseSuite) SetUpTest(c *gc.C) {
	s.BaseSuite.SetUpTest(c)

	s.remote = lxdclient.Remote{
		Name: "my-remote",
		Host: "some-host",
		Cert: s.Cert,
	}
}

type configSuite struct {
	configBaseSuite
}

func (s *configSuite) TestWithDefaultsOkay(c *gc.C) {
	cfg := lxdclient.Config{
		Namespace: "my-ns",
		Dirname:   "some-dir",
		Filename:  "config.yaml",
		Remote:    s.remote,
	}
	updated, err := cfg.WithDefaults()
	c.Assert(err, jc.ErrorIsNil)

	c.Check(updated, jc.DeepEquals, cfg)
}

func (s *configSuite) TestWithDefaultsMissingDirname(c *gc.C) {
	cfg := lxdclient.Config{
		Namespace: "my-ns",
		Dirname:   "",
		Filename:  "config.yaml",
		Remote:    s.remote,
	}
	updated, err := cfg.WithDefaults()
	c.Assert(err, jc.ErrorIsNil)

	c.Check(updated, jc.DeepEquals, lxdclient.Config{
		Namespace: "my-ns",
		// TODO(ericsnow)  This will change on Windows once the LXD
		// code is cross-platform.
		Dirname:  "/.config/lxc", // IsolationSuite sets $HOME to "".
		Filename: "config.yaml",
		Remote:   s.remote,
	})
}

func (s *configSuite) TestWithDefaultsFilename(c *gc.C) {
	cfg := lxdclient.Config{
		Namespace: "my-ns",
		Dirname:   "some-dir",
		Filename:  "",
		Remote:    s.remote,
	}
	updated, err := cfg.WithDefaults()
	c.Assert(err, jc.ErrorIsNil)

	c.Check(updated, jc.DeepEquals, lxdclient.Config{
		Namespace: "my-ns",
		Dirname:   "some-dir",
		Filename:  "config.yml",
		Remote:    s.remote,
	})
}

func (s *configSuite) TestWithDefaultsMissingRemote(c *gc.C) {
	cfg := lxdclient.Config{
		Namespace: "my-ns",
		Dirname:   "some-dir",
		Filename:  "config.yaml",
	}
	updated, err := cfg.WithDefaults()
	c.Assert(err, jc.ErrorIsNil)

	c.Check(updated, jc.DeepEquals, lxdclient.Config{
		Namespace: "my-ns",
		Dirname:   "some-dir",
		Filename:  "config.yaml",
		Remote:    lxdclient.Local,
	})
}

func (s *configSuite) TestValidateOkay(c *gc.C) {
	cfg := lxdclient.Config{
		Namespace: "my-ns",
		Dirname:   "some-dir",
		Filename:  "config.yaml",
		Remote:    s.remote,
	}
	err := cfg.Validate()

	c.Check(err, jc.ErrorIsNil)
}

func (s *configSuite) TestValidateOnlyRemote(c *gc.C) {
	cfg := lxdclient.Config{
		Namespace: "",
		Dirname:   "",
		Filename:  "",
		Remote:    s.remote,
	}
	err := cfg.Validate()

	c.Check(err, jc.ErrorIsNil)
}

func (s *configSuite) TestValidateMissingRemote(c *gc.C) {
	cfg := lxdclient.Config{
		Namespace: "my-ns",
		Dirname:   "some-dir",
		Filename:  "config.yaml",
	}
	err := cfg.Validate()

	c.Check(err, jc.Satisfies, errors.IsNotValid)
}

func (s *configSuite) TestValidateZeroValue(c *gc.C) {
	var cfg lxdclient.Config
	err := cfg.Validate()

	c.Check(err, jc.Satisfies, errors.IsNotValid)
}

func (s *configSuite) TestWriteOkay(c *gc.C) {
	c.Skip("not implemented yet")
	// TODO(ericsnow) Finish!
}

func (s *configSuite) TestWriteRemoteAlreadySet(c *gc.C) {
	c.Skip("not implemented yet")
	// TODO(ericsnow) Finish!
}

func (s *configSuite) TestWriteInvalid(c *gc.C) {
	var cfg lxdclient.Config
	err := cfg.Write()

	c.Check(err, jc.Satisfies, errors.IsNotValid)
}

func (s *configSuite) TestAsNonLocalOkay(c *gc.C) {
	// TODO(ericsnow) Finish!
}

func (s *configSuite) TestAsNonLocalNoop(c *gc.C) {
	cfg := lxdclient.Config{
		Namespace: "my-ns",
		Dirname:   "some-dir",
		Filename:  "config.yml",
		Remote:    s.remote,
	}
	nonlocal, err := cfg.AsNonLocal()
	c.Assert(err, jc.ErrorIsNil)

	c.Check(nonlocal, jc.DeepEquals, cfg)
}

type configFunctionalSuite struct {
	configBaseSuite
}

func (s *configFunctionalSuite) SetUpTest(c *gc.C) {
	s.configBaseSuite.SetUpTest(c)

	origConfigDir := lxd.ConfigDir
	s.AddCleanup(func(c *gc.C) {
		lxd.ConfigDir = origConfigDir
	})
	// TODO(ericsnow) Add a cleanup func to remove any added certs.
}

func (s *configFunctionalSuite) TestWrite(c *gc.C) {
	dirname := c.MkDir()
	cfg := lxdclient.Config{
		Namespace: "my-ns",
		Dirname:   dirname,
		Filename:  "config.yml",
		Remote:    s.remote,
	}
	err := cfg.Write()
	c.Assert(err, jc.ErrorIsNil)

	checkFiles(c, cfg)
}

func (s *configFunctionalSuite) TestAsNonLocal(c *gc.C) {
	if !lxdRunningLocally() {
		c.Skip("LXD not running locally")
	}

	cfg := lxdclient.Config{
		Namespace: "my-ns",
		Dirname:   "some-dir",
		Filename:  "config.yml",
		Remote:    lxdclient.Local,
	}
	nonlocal, err := cfg.AsNonLocal()
	c.Assert(err, jc.ErrorIsNil)

	checkValidRemote(c, &nonlocal.Remote)
	c.Check(nonlocal, jc.DeepEquals, lxdclient.Config{
		Namespace: "my-ns",
		Dirname:   "some-dir",
		Filename:  "config.yml",
		Remote: lxdclient.Remote{
			Name: lxdclient.Local.Name,
			Host: nonlocal.Remote.Host,
			Cert: nonlocal.Remote.Cert,
		},
	})
	// TODO(ericsnow) Check that the server has the certs.
}

func lxdRunningLocally() bool {
	origConfigDir := lxd.ConfigDir
	defer func() {
		lxd.ConfigDir = origConfigDir
	}()

	_, err := lxdclient.Connect(lxdclient.Config{
		Namespace: "my-ns",
		Dirname:   "some-dir",
		Filename:  "config.yml",
		Remote:    lxdclient.Local,
	})
	return err == nil
}

func checkFiles(c *gc.C, cfg lxdclient.Config) {
	var certificate lxdclient.Cert
	if cfg.Remote.Cert != nil {
		certificate = *cfg.Remote.Cert
	}

	filename := filepath.Join(cfg.Dirname, "client.crt")
	c.Logf("reading cert PEM from %q", filename)
	certPEM, err := ioutil.ReadFile(filename)
	c.Assert(err, jc.ErrorIsNil)
	c.Check(string(certPEM), gc.Equals, string(certificate.CertPEM))

	filename = filepath.Join(cfg.Dirname, "client.key")
	c.Logf("reading key PEM from %q", filename)
	keyPEM, err := ioutil.ReadFile(filename)
	c.Assert(err, jc.ErrorIsNil)
	c.Check(string(keyPEM), gc.Equals, string(certificate.KeyPEM))

	filename = filepath.Join(cfg.Dirname, cfg.Filename)
	c.Logf("reading config from %q", filename)
	configData, err := ioutil.ReadFile(filename)
	c.Assert(err, jc.ErrorIsNil)
	var config lxd.Config
	err = goyaml.Unmarshal(configData, &config)
	c.Assert(err, jc.ErrorIsNil)
	c.Check(config, jc.DeepEquals, lxd.Config{
		DefaultRemote: "local",
		Remotes: map[string]lxd.RemoteConfig{
			// TODO(ericsnow) Use the following once we switch to a newer LXD.
			//"local": lxd.LocalRemote,
			"local": config.Remotes["local"],
			cfg.Remote.Name: lxd.RemoteConfig{
				Addr:   "https://" + cfg.Remote.Host + ":8443",
				Public: false,
			},
		},
		// TODO(ericsnow) Use the following once we switch to a newer LXD.
		//Aliases: nil,
	})
}
