#!/bin/bash

protoc -I ./organization/proto \
    --openapiv2_out=logtostderr=true:./organization/proto/haas/organization/ \
    ./organization/proto/haas/organization/organization.proto
