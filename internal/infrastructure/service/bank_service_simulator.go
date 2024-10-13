package service

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/acgyiyo/payment_api_test/internal/domain/entity"
)

type BankServiceSimulator interface {
	ProcessPaymentInBank(payment *entity.Payment) (*entity.Payment, error)
	ProcessRefundInBank(refundReq *entity.Payment) (*entity.Payment, error)
}

type bankServiceSimulator struct{}

func NewBankServiceSimulator() BankServiceSimulator {
	return &bankServiceSimulator{}
}

func (bs *bankServiceSimulator) ProcessPaymentInBank(payment *entity.Payment) (*entity.Payment, error) {
	log.Printf("Processing Payment: %+v", payment)
	var status = "acepted"
	var message = "Payment accepted succesfully"

	if !validateCard(payment.CustomerCard) {
		return nil, errors.New("invalid Card Information")
	}

	log.Printf("Going to transfer from customer card number: %s to merchant:%s", payment.CustomerCard, payment.MerchantID)
	//simulation: validates amount and transfer form card to customer
	if rand.Intn(10) == 0 {
		status = "declined"
		message = "Payment declined due to anomalies"
	}

	var trx = fmt.Sprintf("txn-%s+%d", time.Now().Format("20060102150405"), rand.Intn(1000))

	payment.TransactionID = trx
	payment.Status = status
	payment.Message = message

	return payment, nil
}

func (bs *bankServiceSimulator) ProcessRefundInBank(refundReq *entity.Payment) (*entity.Payment, error) {
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
