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
	result, err := rp.bankService.ProcessPaymentInBank(&payment)
	if err != nil {
		log.Print("error processing payment: ProcessPaymentInBank failed: ", err)
		audit.AuditMsg(fmt.Sprintf("error processing payment: ProcessPaymentInBank failed: %s, tags:{%s:%+v}",
			err.Error(), "payment", payment))
		return nil, errors.New("error validating payment in bank: " + err.Error())
	}

	err = rp.paymentStore.SavePayment(ctx, result)
	if err != nil {
		log.Print("error processing payment: SavePayment failed: ", err)
		audit.AuditMsg(fmt.Sprintf("error processing payment: SavePayment failed: %s, tags:{%s:%+v}",
			err.Error(), "payment", payment))
		return nil, errors.New("error saving payment")
	}

	//counting metrics
	metric.Count("SavePayment", 1, nil, 0)

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
