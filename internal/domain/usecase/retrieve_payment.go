package usecase

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/acgyiyo/payment_api_test/internal/domain/entity"
	"github.com/acgyiyo/payment_api_test/internal/domain/gateway"
	"github.com/acgyiyo/payment_api_test/internal/infrastructure/service/audit"
	"github.com/acgyiyo/payment_api_test/internal/infrastructure/service/metric"
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
		log.Print("error retrievement payment: SearchPaymentByTransactionID failed: ", err)
		audit.AuditMsg(fmt.Sprintf("error retrievement payment: SearchPaymentByTransactionID failed: %s, tags:{%s:%+v}",
			err.Error(), "transactionID", transactionID))
		return nil, errors.New("Error getting payment with transactioID: " + transactionID)
	}

	//counting metrics
	metric.Count("SearchPaymentByTransactionID", 1, nil, 0)

	return convertPaymentToResponse(result), nil
}
