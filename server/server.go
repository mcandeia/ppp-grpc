package server

import (
	"context"

	ppp "github.com/mcandeia/ppp-grpc/ppp"

	"google.golang.org/protobuf/types/known/emptypb"
)

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

func NewPaymentProviderServer() *paymentProviderServer {
	s := &paymentProviderServer{}
	return s
}