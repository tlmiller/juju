#!/usr/bin/env bash

set -xe

gopatch -p scripts/errors-patch/errors.patch "${1}/..."
find "$1" -type f -iname '*.go' -exec sed -i '' -E "s,\"(.*)\"[ ]?\+[ ]?\" \%w\",\"\1 \%w\",g" "{}" +;