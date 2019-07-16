#!/bin/sh

go install \
	github.com/golang/protobuf/protoc-gen-go \
	github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
	github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger \
	github.com/gunk/gunk \
	|| exit 1
