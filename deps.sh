#!/usr/bin/env bash

set -euo pipefail

go install \
	github.com/golang/protobuf/protoc-gen-go \
	github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
	github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger \
	github.com/gunk/gunk \
	github.com/gunk/gunk/docgen \
	github.com/Kunde21/pulpMd
