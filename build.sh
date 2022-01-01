#!/bin/bash 
protoc -I libcni/ \
    -I${GOPATH}/src \
    -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    --go_out=libcni --go_opt=paths=source_relative \
    --go-grpc_out=libcni --go-grpc_opt=paths=source_relative \
    libcni/cnigrpc.proto
go build -v -o cnitool/cnitool cnitool/cnitool.go
go build -v -o cniserver/server cniserver/server.go
#go build -i -v -o cnigrpc/bin/client cnigrpc/client/client.go
