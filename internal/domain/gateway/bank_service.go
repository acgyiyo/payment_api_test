package gateway

import (
	"github.com/acgyiyo/payment_api_test/internal/domain/entity"
)

type BankService interface {
	ProcessPaymentInBank(payment *entity.Payment) (*entity.Payment, error)
	ProcessRefundInBank(refundReq *entity.Payment) (*entity.Payment, error)
}
