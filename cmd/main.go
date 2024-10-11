package main

import (
	"log"
	"net/http"

	"github.com/acgyiyo/payment_api_test/internal/services"
	//"./internal/handlers"
)

func main() {
	services.InitBankSimulator()

	http.HandleFunc("/payments", handlers.ProcessPayment)
	http.HandleFunc("/payments/", handlers.QueryPayment)
	http.HandleFunc("/refunds", handlers.ProcessRefund)

	log.Println("Starting payment API on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
