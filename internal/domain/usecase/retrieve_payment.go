package usecase

import (
	"context"
	"errors"
	"log"

	"github.com/acgyiyo/payment_api_test/internal/domain/entity"
	"github.com/acgyiyo/payment_api_test/internal/domain/gateway"
)

type RetrievePayment interface {
	SearchPaymentByTransactionID(ctx context.Context, transactionID string) (*entity.PaymentResponse, error)
}

type retrievePayment struct {
	paymentStore gateway.PaymentStore
}

func NewRetrievePayment(st gateway.PaymentStore) RetrievePayment {
	return &retrievePayment{
		paymentStore: st,
	}
}

func (rp *retrievePayment) SearchPaymentByTransactionID(ctx context.Context, transactionID string) (*entity.PaymentResponse, error) {
	result, err := rp.paymentStore.SearchPaymentByTransactionID(ctx, transactionID)
	if err != nil {
		log.Print("error retrievement payment: SearchPaymentByTransactionID failed", err)
		return nil, errors.New("Error getting payment with transactioID: " + transactionID)
	}

	return convertPaymentToResponse(result), nil
}
