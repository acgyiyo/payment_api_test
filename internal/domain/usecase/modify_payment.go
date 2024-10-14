package usecase

import (
	"context"
	"errors"
	"log"

	"github.com/acgyiyo/payment_api_test/internal/domain/entity"
	"github.com/acgyiyo/payment_api_test/internal/domain/gateway"
)

type UpdatePayment interface {
	UpdatePayment(ctx context.Context, payment entity.Payment) (*entity.PaymentResponse, error)
}

type updatePayment struct {
	paymentStore gateway.PaymentStore
	bankService  gateway.BankService
}

func NewUpdatePayment(st gateway.PaymentStore, bs gateway.BankService) UpdatePayment {
	return &updatePayment{
		paymentStore: st,
		bankService:  bs,
	}
}

func (rp *updatePayment) UpdatePayment(ctx context.Context, refund entity.Payment) (*entity.PaymentResponse, error) {
	result, err := rp.paymentStore.SearchPaymentByTransactionID(ctx, refund.TransactionID)
	if err != nil {
		log.Print("error updating payment: SearchPaymentByTransactionID failed: ", err)
		return nil, errors.New("transaction not found")
	}

	_, err = rp.bankService.ProcessRefundInBank(result)
	if err != nil {
		log.Print("error processing refund: ProcessRefundInBank failed", err)
		return nil, errors.New("error processing refund in bank")
	}

	err = rp.paymentStore.UpdatePayment(ctx, result)
	if err != nil {
		log.Print("error processing refund: UpdatePayment failed", err)
		return nil, errors.New("error updating payment")
	}

	return convertPaymentToResponse(result), nil
}
