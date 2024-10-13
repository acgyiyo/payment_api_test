package entity

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	TransactionID string   `gorm:"uniqueIndex;not null" json:"transaction_id"`
	Status        string   `json:"status"`
	Amount        float64  `json:"amount"`
	Message       string   `json:"message"`
	CustomerCard  string   `json:"customer_card"`
	MerchantID    string   `json:"merchant"`
	Merchant      Merchant `gorm:"foreignKey:MerchantID;associationForeignKey:ID"`
}
