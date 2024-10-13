package usecase

import (
	"context"
	"fmt"
	"log"

	"github.com/acgyiyo/payment_api_test/internal/domain/entity"
	"github.com/acgyiyo/payment_api_test/internal/domain/gateway"
)

type RegisterPayment interface {
	SavePayment(ctx context.Context, payment entity.Payment) (*entity.PaymentResponse, error)
}

type registerPayment struct {
	paymentStore gateway.PaymentStore
	bankService  gateway.BankService
}

func NewRegisterPayment(st gateway.PaymentStore, bs gateway.BankService) RegisterPayment {
	return &registerPayment{
		paymentStore: st,
		bankService:  bs,
	}
}

func (rp *registerPayment) SavePayment(ctx context.Context, payment entity.Payment) (*entity.PaymentResponse, error) {
	fmt.Print("registrando payment in service\n")

	result, err := rp.bankService.ProcessPaymentInBank(&payment)
	if err != nil {
		log.Print("error processing payment: ProcessPaymentInBank failed", err)
	}

	err = rp.paymentStore.SavePayment(ctx, result)
	if err != nil {
		return nil, err
	}

	return convertPaymentToResponse(result), nil
}

func convertPaymentToResponse(payment *entity.Payment) *entity.PaymentResponse {
	return &entity.PaymentResponse{
		TransactionID: payment.TransactionID,
		Amount:        payment.Amount,
		MerchantID:    payment.MerchantID,
		CustomerCard:  payment.CustomerCard,
	}
}
