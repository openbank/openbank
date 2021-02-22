#!/usr/bin/env bash

set -euo pipefail

readonly root_dir=$(git rev-parse --show-toplevel)
pushd "$root_dir"
GOBIN=$(pwd)/bin go install \
  github.com/gunk/gunk \
  github.com/gunk/gunk/docgen \
  github.com/gunk/gunk/scopegen

PATH=$root_dir/bin:$PATH
export PATH
