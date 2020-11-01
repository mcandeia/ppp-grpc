package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	ppp "github.com/mcandeia/ppp-grpc/ppp"
	server "github.com/mcandeia/ppp-grpc/server"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)
var port = 10000

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port)) // could not be localhost:%d inside docker
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
    log.Println("Running...")
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	ppp.RegisterPaymentProviderServer(grpcServer, server.NewPaymentProviderServer())
	grpcServer.Serve(lis)
}