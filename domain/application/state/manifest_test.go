// Copyright 2024 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package state

import (
	"context"

	"github.com/canonical/sqlair"
	jc "github.com/juju/testing/checkers"
	gc "gopkg.in/check.v1"

	corecharm "github.com/juju/juju/core/charm"
	"github.com/juju/juju/domain/application/charm"
	schematesting "github.com/juju/juju/domain/schema/testing"
)

type manifestSuite struct {
	schematesting.ModelSuite
}

var _ = gc.Suite(&manifestSuite{})

var decodeManifestTestCases = [...]struct {
	name   string
	input  []charmManifest
	output charm.Manifest
}{
	{
		name:   "empty",
		input:  []charmManifest{},
		output: charm.Manifest{},
	},
	{
		name: "decode base",
		input: []charmManifest{
			{
				Index:        0,
				OS:           "ubuntu",
				Track:        "latest",
				Risk:         "edge",
				Architecture: "amd64",
			},
		},
		output: charm.Manifest{
			Bases: []charm.Base{
				{
					Name: "ubuntu",
					Channel: charm.Channel{
						Track: "latest",
						Risk:  charm.RiskEdge,
					},
					Architectures: []string{"amd64"},
				},
			},
		},
	},
	{
		name: "decode bases",
		input: []charmManifest{
			{
				Index:        0,
				OS:           "ubuntu",
				Track:        "latest",
				Risk:         "edge",
				Architecture: "amd64",
			},
			{
				Index:        0,
				OS:           "ubuntu",
				Track:        "latest",
				Risk:         "edge",
				Architecture: "arm64",
			},
		},
		output: charm.Manifest{
			Bases: []charm.Base{
				{
					Name: "ubuntu",
					Channel: charm.Channel{
						Track: "latest",
						Risk:  charm.RiskEdge,
					},
					Architectures: []string{"amd64", "arm64"},
				},
			},
		},
	},
}

var encodeManifestTestCases = [...]struct {
	name   string
	id     corecharm.ID
	input  charm.Manifest
	output []setCharmManifest
}{
	{
		name:   "empty",
		input:  charm.Manifest{},
		output: []setCharmManifest{},
	},
	{
		name: "no architectures",
		id:   "deadbeef",
		input: charm.Manifest{
			Bases: []charm.Base{
				{
					Name: "ubuntu",
					Channel: charm.Channel{
						Track: "22.04",
						Risk:  charm.RiskEdge,
					},
				},
			},
		},
		output: []setCharmManifest{
			{
				CharmUUID:      "deadbeef",
				Index:          0,
				OSID:           0,
				ArchitectureID: 0,
				Track:          "22.04",
				Risk:           "edge",
			},
		},
	},
	{
		name: "no architectures - multiple",
		id:   "deadbeef",
		input: charm.Manifest{
			Bases: []charm.Base{
				{
					Name: "ubuntu",
					Channel: charm.Channel{
						Track: "22.04",
						Risk:  charm.RiskEdge,
					},
				},
				{
					Name: "ubuntu",
					Channel: charm.Channel{
						Track: "24.04",
						Risk:  charm.RiskEdge,
					},
				},
			},
		},
		output: []setCharmManifest{
			{
				CharmUUID:      "deadbeef",
				Index:          0,
				NestedIndex:    0,
				OSID:           0,
				ArchitectureID: 0,
				Track:          "22.04",
				Risk:           "edge",
			},
			{
				CharmUUID:      "deadbeef",
				Index:          1,
				NestedIndex:    0,
				OSID:           0,
				ArchitectureID: 0,
				Track:          "24.04",
				Risk:           "edge",
			},
		},
	},
	{
		name: "architectures",
		id:   "deadbeef",
		input: charm.Manifest{
			Bases: []charm.Base{
				{
					Name: "ubuntu",
					Channel: charm.Channel{
						Track: "22.04",
						Risk:  charm.RiskEdge,
					},
					Architectures: []string{"amd64", "arm64"},
				},
			},
		},
		output: []setCharmManifest{
			{
				CharmUUID:      "deadbeef",
				Index:          0,
				NestedIndex:    0,
				OSID:           0,
				ArchitectureID: 0,
				Track:          "22.04",
				Risk:           "edge",
			},
			{
				CharmUUID:      "deadbeef",
				Index:          0,
				NestedIndex:    1,
				OSID:           0,
				ArchitectureID: 1,
				Track:          "22.04",
				Risk:           "edge",
			},
		},
	},
}

func (s *manifestSuite) TestDecodeManifest(c *gc.C) {
	for _, tc := range decodeManifestTestCases {
		c.Logf("Running test case %q", tc.name)

		decoded, err := decodeManifest(tc.input)
		c.Assert(err, jc.ErrorIsNil)
		c.Check(decoded, gc.DeepEquals, tc.output)
	}
}

func (s *manifestSuite) TestEncodeManifest(c *gc.C) {
	for _, tc := range encodeManifestTestCases {
		c.Logf("Running test case %q", tc.name)

		encoded, err := encodeManifest(tc.id, tc.input)
		c.Assert(err, jc.ErrorIsNil)
		c.Check(encoded, gc.DeepEquals, tc.output)
	}
}

type manifestStateSuite struct {
	schematesting.ModelSuite
}

var _ = gc.Suite(&manifestStateSuite{})

func (s *manifestStateSuite) TestManifestOS(c *gc.C) {
	type osType struct {
		ID   int    `db:"id"`
		Name string `db:"name"`
	}

	stmt := sqlair.MustPrepare(`
SELECT os.* AS &osType.* FROM os ORDER BY id;
`, osType{})

	var results []osType
	err := s.TxnRunner().Txn(context.Background(), func(ctx context.Context, tx *sqlair.TX) error {
		return tx.Query(ctx, stmt).GetAll(&results)
	})
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(results, gc.HasLen, 1)

	m := []string{
		"ubuntu",
	}

	for i, value := range m {
		c.Logf("result %d: %#v", i, value)
		result, err := encodeManifestOS(value)
		c.Assert(err, jc.ErrorIsNil)
		c.Check(result, gc.DeepEquals, results[i].ID)
	}
}

func (s *manifestStateSuite) TestManifestArchitecture(c *gc.C) {
	type archType struct {
		ID   int    `db:"id"`
		Name string `db:"name"`
	}

	stmt := sqlair.MustPrepare(`
SELECT architecture.* AS &archType.* FROM architecture ORDER BY id;
`, archType{})

	var results []archType
	err := s.TxnRunner().Txn(context.Background(), func(ctx context.Context, tx *sqlair.TX) error {
		return tx.Query(ctx, stmt).GetAll(&results)
	})
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(results, gc.HasLen, 5)

	m := []string{
		"amd64",
		"arm64",
		"ppc64el",
		"s390x",
		"riscv64",
	}

	for i, value := range m {
		c.Logf("result %d: %#v", i, value)
		result, err := encodeManifestArchitecture(value)
		c.Assert(err, jc.ErrorIsNil)
		c.Check(result, gc.DeepEquals, results[i].ID)
	}
}
