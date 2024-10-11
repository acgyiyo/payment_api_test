package models

type PaymentRequest struct {
	Amount     float64 `json:"amount"`
	CardInfo   string  `json:"card_info"`
	MerchantID string  `json:"merchant_id"`
}

type RefundRequest struct {
	TransactionID string  `json:"transaction_id"`
	Amount        float64 `json:"amount"`
}

type PaymentResponse struct {
	Status  string `json:"status"`
	TransID string `json:"transaction_id"`
	Message string `json:"message"`
}

var transactions = map[string]PaymentResponse{}

func SaveTransaction(response PaymentResponse) {
	transactions[response.TransID] = response
}

func UpdateTransaction(transID string, response PaymentResponse) {
	transactions[transID] = response
}
