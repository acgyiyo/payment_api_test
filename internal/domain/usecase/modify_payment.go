package usecase

import (
	"context"
	"fmt"
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
	fmt.Print("Actualizando payment in service\n")

	result, err := rp.paymentStore.SearchPaymentByTransactionID(ctx, refund.TransactionID)
	if err != nil {
		log.Print("transaction not found") //TODO Improve
		return nil, err
		//ctx.JSON(http.StatusInternalServerError, "transaction not found") //TODO response error
	}

	_, err = rp.bankService.ProcessRefundInBank(result)
	if err != nil {
		log.Print("error processing refund: ProcessRefundInBank failed", err)
		return nil, err
	}

	err = rp.paymentStore.UpdatePayment(ctx, result)
	if err != nil {
		return nil, err
	}

	return convertPaymentToResponse(result), nil
}
