#!/bin/bash
protoc -I pkg/types/v1/ \
    -I${GOPATH}/src/github.com/protocolbuffers/protobuf/src/ \
    --go_out=pkg/types/v1/ --go_opt=paths=source_relative \
    --go-grpc_out=pkg/types/v1/ --go-grpc_opt=paths=source_relative \
    pkg/types/v1/cni.proto
