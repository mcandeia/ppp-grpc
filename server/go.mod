module github.com/mcandeia/ppp-grpc/server

go 1.15

replace github.com/mcandeia/ppp-grpc/proto => ../ppp

require (
	google.golang.org/grpc v1.33.1
	google.golang.org/protobuf v1.25.0
	github.com/mcandeia/ppp-grpc/proto v0.0.0-00010101000000-000000000000
)
