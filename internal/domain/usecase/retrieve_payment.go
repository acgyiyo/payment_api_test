package usecase

import (
	"context"
	"fmt"

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
	fmt.Print("Consultando payment in service\n")

	result, err := rp.paymentStore.SearchPaymentByTransactionID(ctx, transactionID)
	if err != nil {
		return nil, err
	}

	return convertPaymentToResponse(result), nil
}

/*
func convertPaymentToResponse(payment *entity.Payment) *entity.PaymentResponse {
	return &entity.PaymentResponse{
		TransactionID: payment.TransactionID,
		Amount:        payment.Amount,
		MerchantID:    payment.MerchantID,
		CustomerCard:  payment.CustomerCard,
	}
}
*/
