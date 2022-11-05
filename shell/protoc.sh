#!/bin/bash

protoc -I . \
    --go_out=/api_pb --go_opt=paths=source_relative \
    --go-grpc_out=/api_pb --go-grpc_opt=paths=source_relative \
    --grpc-gateway_out=logtostderr=true,paths=source_relative:/api_pb \
    $(find haas -name "*.proto")

protoc -I . \
    --include_imports \
    --include_source_info \
    --descriptor_set_out=/api_pb/api_descriptor.pb \
    $(find haas -name "*.proto")

protoc -I . \
    --openapiv2_out=logtostderr=true,allow_merge=true,merge_file_name=haas:/doc \
    $(find haas -name "*.proto")