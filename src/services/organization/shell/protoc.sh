#!/bin/bash

protoc \
    --proto_path=. \
    --go-grpc_out=paths=source_relative:/api_pb \
    --grpc-gateway_out=logtostderr=true,paths=source_relative:/api_pb \
    $(find . -name "*.proto")