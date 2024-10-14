package usecase

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/acgyiyo/payment_api_test/internal/domain/entity"
	"github.com/acgyiyo/payment_api_test/internal/domain/gateway"
	"github.com/acgyiyo/payment_api_test/internal/infrastructure/service/audit"
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
		audit.AuditMsg(fmt.Sprintf("error updating payment: SearchPaymentByTransactionID failed: %s, tags:{%s:%+v}",
			err.Error(), "refund", refund))
		return nil, errors.New("transaction not found")
	}

	_, err = rp.bankService.ProcessRefundInBank(result)
	if err != nil {
		log.Print("error processing refund: ProcessRefundInBank failed", err)
		audit.AuditMsg(fmt.Sprintf("error processing refund: ProcessRefundInBank failed: %s, tags:{%s:%+v}",
			err.Error(), "refund", refund))
		return nil, errors.New("error processing refund in bank")
	}

	err = rp.paymentStore.UpdatePayment(ctx, result)
	if err != nil {
		log.Print("error processing refund: UpdatePayment failed", err)
		audit.AuditMsg(fmt.Sprintf("error processing refund: UpdatePayment failed: %s, tags:{%s:%+v}",
			err.Error(), "refund", refund))
		return nil, errors.New("error updating payment")
	}

	return convertPaymentToResponse(result), nil
}
