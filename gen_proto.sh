#!/bin/bash


if [ -z "$1" ]; then
    echo -e ">>> Provide a proto file..."
    exit 1
fi

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative "$1"