#!/usr/bin/env bash

set -euo pipefail

readonly golang_version=1.13.0

docker run --rm --workdir /src --volume "$(pwd)":/src golang:"$golang_version"-buster \
	/bin/bash -c "./deps.sh && gunk generate ./..."
