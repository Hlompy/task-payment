package httpserver

import (
	"context"

	"github.com/Hlompy/task-payment/internal/app/domain"
)

type PaymentService interface {
	ProcessPayment(ctx context.Context, intent domain.PaymentIntent) (*domain.Payment, error)
	GetPayment(ctx context.Context, merchantID, paymentID string) (*domain.Payment, error)
}

type MerchantService interface {
	Authenticate(ctx context.Context, apiKey string) (*domain.Merchant, error)
}
