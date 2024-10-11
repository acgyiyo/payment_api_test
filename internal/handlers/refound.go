package handlers

import (
	"encoding/json"
	"net/http"
	"payment-api/models"
	"payment-api/services"
)

func ProcessRefund(w http.ResponseWriter, r *http.Request) {
	var refund models.RefundRequest
	json.NewDecoder(r.Body).Decode(&refund)

	result := services.BankSimulator.ProcessRefund(refund)
	models.UpdateTransaction(refund.TransactionID, result)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
