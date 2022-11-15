#!/bin/bash

# imaging-study service
cd ./api/openapi
oapi-codegen -generate "server" -package openapi ../../doc/openapi.yaml > server.gen.go
oapi-codegen -generate "types" -package openapi ../../doc/openapi.yaml > type.gen.go
oapi-codegen -generate "spec" -package openapi ../../doc/openapi.yaml > spec.gen.go
