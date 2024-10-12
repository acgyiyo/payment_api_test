package models

import (
	"gorm.io/gorm"
)

type PaymentRequest struct {
	Amount     float64 `json:"amount"`
	MerchantID string  `json:"merchant_id"`
	//CustomerID   string  `json:"customer_id"`
	CustomerCard string `json:"customer_card"`
}

type RefundRequest struct {
	TransactionID string `json:"transaction_id"`
	//MerchantID    string `json:"merchant_id"`
	//CustomerCard  string `json:"customer_card"`
}

type Merchant struct {
	gorm.Model
	Name     string
	Payments []Payment `gorm:"foreignKey:MerchantID"`
}

type Payment struct {
	gorm.Model
	TransactionID string   `gorm:"uniqueIndex;not null" json:"transaction_id"`
	Status        string   `json:"status"`
	Amount        float64  `json:"amount"`
	Message       string   `json:"message"`
	MerchantID    string   `json:"merchant"`
	Merchant      Merchant `gorm:"foreignKey:MerchantID;associationForeignKey:ID"`
}

func SaveTransaction(response *Payment) error {
	if err := DB.Create(response).Error; err != nil {
		return err
	}
	return nil
}

func GetTransaction(transactionID string) (*Payment, bool) {
	var transaction Payment
	result := DB.Where("transaction_id = ?", transactionID).First(&transaction)
	return &transaction, result.Error == nil
}

func UpdateTransaction(payment *Payment) error {
	if err := DB.Save(payment).Error; err != nil {
		return err
	}
	return nil
}
