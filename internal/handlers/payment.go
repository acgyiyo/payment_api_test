package handlers

import (
	"encoding/json"
	"internal/models"
	"net/http"

	"github.com/acgyiyo/payment_api_test/internal/services"
)

func ProcessPayment(w http.ResponseWriter, r *http.Request) {
	var payment models.PaymentRequest
	json.NewDecoder(r.Body).Decode(&payment)

	result := services.BankSimulator.ProcessPayment(payment)
	models.SaveTransaction(result)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
