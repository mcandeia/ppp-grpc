#!/bin/bash

brew update
brew install golang
brew install protobuf

export GO111MODULE=on  # Enable module mode
go get github.com/golang/protobuf/proto google.golang.org/protobuf/cmd/protoc-gen-go \
         google.golang.org/grpc/cmd/protoc-gen-go-grpc


## Setup Manually
export PATH="$PATH:$(go env GOPATH)/bin"
