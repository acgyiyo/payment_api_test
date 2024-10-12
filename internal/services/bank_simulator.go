package services

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/acgyiyo/payment_api_test/internal/models"
)

type BankSimulatorService struct{}

var BankSimulator BankSimulatorService

func InitBankSimulator() {
	rand.Seed(time.Now().UnixNano())
}

func (bs BankSimulatorService) ProcessPaymentInBank(paymentReq models.PaymentRequest) (*models.Payment, error) {
	log.Printf("Processing Payment: %+v", paymentReq)
	var status = "acepted"
	var message = "Payment accepted succesfully"

	if !validateCard(paymentReq.CustomerCard) {
		return nil, errors.New("invalid Card Information")
	}

	log.Printf("Going to transfer from customer card number: %s to merchant:%s", paymentReq.CustomerCard, paymentReq.MerchantID)
	//simulation: validates amount and transfer form card to customer
	if rand.Intn(10) == 0 {
		status = "declined"
		message = "Payment declined due to anomalies"
	}

	var trx = fmt.Sprintf("txn-%s+%d", time.Now().Format("20060102150405"), rand.Intn(1000))

	return &models.Payment{
		TransactionID: trx,
		Status:        status,
		Message:       message,
	}, nil
}

func (bs BankSimulatorService) ProcessRefundInBank(refundReq *models.Payment) (*models.Payment, error) {
	status := "refunded"
	log.Printf("Going to refund payment: %s in Bank", refundReq.TransactionID)

	//validate if the merchant and the customer have restrictions to do the refund
	//doing validations...
	//fail return err

	refundReq.Status = status
	refundReq.Message = "Refund processed by bank"

	return refundReq, nil
}

func validateCard(cardInfo string) bool {
	return len(cardInfo) >= 13
}
