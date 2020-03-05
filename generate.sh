#!/usr/bin/env bash

set -euo pipefail

source ./install_deps.sh
gunk generate ./...
