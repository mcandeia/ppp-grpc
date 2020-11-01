module github.com/mcandeia/ppp-grpc

go 1.15

replace github.com/mcandeia/ppp-grpc/server => ./server

require (
	github.com/golang/protobuf v1.4.3
	google.golang.org/grpc v1.33.1
	google.golang.org/protobuf v1.25.0
)
