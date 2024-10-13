package gateway

import (
	"context"

	"github.com/acgyiyo/payment_api_test/internal/domain/entity"
)

type PaymentStore interface {
	SavePayment(ctx context.Context, payment *entity.Payment) error
	SearchPaymentByTransactionID(ctx context.Context, transactionID string) (*entity.Payment, error)
	UpdatePayment(ctx context.Context, payment *entity.Payment) error
}
