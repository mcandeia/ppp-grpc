package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	ppp "mcandeia/grpc/ppp"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)
var port = 10000

type paymentProviderServer struct {
	ppp.UnimplementedPaymentProviderServer
}

// GetFeature returns the feature at the given point.
func (s *paymentProviderServer) GetManifest(ctx context.Context, empty *emptypb.Empty) (*ppp.Manifest, error) {
	// No feature was found, return an unnamed feature
	paymentMethods := ppp.PaymentMethod { AllowsSplit: ppp.SplitStage_onAuthorize, Name: ppp.PaymentMethodName_BankInvoice  }
	return &ppp.Manifest { PaymentMethods: []*ppp.PaymentMethod{&paymentMethods} }, nil
}
func newServer() *paymentProviderServer {
	s := &paymentProviderServer{}
	return s
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
    log.Println("Running...")
	grpcServer := grpc.NewServer()
	ppp.RegisterPaymentProviderServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}