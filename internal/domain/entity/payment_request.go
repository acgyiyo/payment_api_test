package entity

type PaymentRequest struct {
	Amount       float64 `json:"amount"`
	MerchantID   string  `json:"merchant_id"`
	CustomerCard string  `json:"customer_card"`
}
