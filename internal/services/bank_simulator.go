package services

import (
	"math/rand"
	"time"

	"github.com/acgyiyo/internal/models"
)

type BankSimulatorService struct{}

var BankSimulator BankSimulatorService

func InitBankSimulator() {
	rand.Seed(time.Now().UnixNano())
}

func (bs BankSimulatorService) ProcessPayment(payment models.PaymentRequest) models.PaymentResponse {
	success := rand.Intn(2) == 1
	status := "success"
	if !success {
		status = "failure"
	}
	return models.PaymentResponse{
		Status:  status,
		TransID: "txn-" + string(rand.Intn(1000)),
		Message: "Processed by bank simulator",
	}
}

func (bs BankSimulatorService) ProcessRefund(refund models.RefundRequest) models.PaymentResponse {
	success := rand.Intn(2) == 1
	status := "refunded"
	if !success {
		status = "refund failed"
	}
	return models.PaymentResponse{
		Status:  status,
		TransID: refund.TransactionID,
		Message: "Refund processed by bank simulator",
	}
}
