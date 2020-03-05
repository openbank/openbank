#!/usr/bin/env bash

set -euo pipefail

readonly root_dir=$(git rev-parse --show-toplevel)
pushd "$root_dir"
GOBIN=$(pwd)/bin go install \
  github.com/golang/protobuf/protoc-gen-go \
  github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
  github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger \
  github.com/gunk/gunk \
  github.com/gunk/gunk/docgen \
  github.com/gunk/gunk/scopegen \
  github.com/Kunde21/pulpMd

PATH=$root_dir/bin:$PATH
export PATH
