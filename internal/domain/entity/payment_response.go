package entity

type PaymentResponse struct {
	TransactionID string  `json:"transaction_id"`
	Amount        float64 `json:"amount"`
	MerchantID    string  `json:"merchant_id"`
	CustomerCard  string  `json:"customer_card"`
}
