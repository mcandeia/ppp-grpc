package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	ppp "github.com/mcandeia/ppp-grpc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)
var port = 10000

type paymentProviderServer struct {
	ppp.UnimplementedPaymentProviderServer
}

// GetFeature returns the feature at the given point.
func (s *paymentProviderServer) GetManifest(ctx context.Context, empty *emptypb.Empty) (*ppp.Manifest, error) {
	// No feature was found, return an unnamed feature
	bankInvoice := ppp.PaymentMethod { AllowsSplit: ppp.SplitStage_onAuthorize, Name: ppp.PaymentMethodName_BankInvoice  }
	visa := ppp.PaymentMethod { AllowsSplit: ppp.SplitStage_onAuthorize, Name: ppp.PaymentMethodName_Visa }
	return &ppp.Manifest { PaymentMethods: []*ppp.PaymentMethod{&visa, &bankInvoice} }, nil
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
	reflection.Register(grpcServer)
	ppp.RegisterPaymentProviderServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}