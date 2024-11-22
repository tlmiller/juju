#!/usr/bin/env bash

set -xe

# Step 1: Run the go patch files over the target directory.
gopatch -p scripts/errors-patch/errors.patch "${1}/..."
gopatch -p scripts/errors-patch/errors-rename.patch "${1}/..."

# Step 2: Run sed replacement steps over go patched files.

# This sed step is fixing lines of the form:
# - errors.Errorf("some message" + " %w", err) to errors.Errorf("some message %w", err)
# We do this because go patch doesn't have the ability to modify strings.
find "$1" -type f -iname '*.go' -exec sed -i '' -E "s,\"(.*)\"[ ]?\+[ ]?\" \%w\",\"\1 \%w\",g" "{}" +;

# Step 3 fix up imports that have been modified by go patch.
gci  write --skip-generated -s standard -s default -s 'Prefix(github.com/juju/juju)' "${1}/."