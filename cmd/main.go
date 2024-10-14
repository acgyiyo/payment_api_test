package main

import (
	"log"

	"github.com/acgyiyo/payment_api_test/internal/config"
	"github.com/acgyiyo/payment_api_test/internal/domain/usecase"
	"github.com/acgyiyo/payment_api_test/internal/infrastructure/delivery"
	"github.com/acgyiyo/payment_api_test/internal/infrastructure/delivery/handler"
	"github.com/acgyiyo/payment_api_test/internal/infrastructure/repository"
	"github.com/acgyiyo/payment_api_test/internal/infrastructure/service"
	"github.com/acgyiyo/payment_api_test/internal/infrastructure/service/metric"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load the configuration from the file
	configParams, err := config.LoadConfig("./internal/config/schemas")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	db, _ := config.InitDB(configParams)
	metric.InitDatadogStatsd()

	//init repositories and services
	paymentStore := repository.NewPaymentDataBase(db)
	bankService := service.NewBankServiceSimulator()

	//init useCases
	registerPaymentUseCase := usecase.NewRegisterPayment(paymentStore, bankService)
	retrievePaymentUseCase := usecase.NewRetrievePayment(paymentStore)
	refundPaymentUseCase := usecase.NewUpdatePayment(paymentStore, bankService)

	//init handlers
	paymentPostHandler := handler.NewPostPayment(registerPaymentUseCase)
	paymentGetHandler := handler.NewGetPayment(retrievePaymentUseCase)
	refundHandler := handler.NewRefundPayment(refundPaymentUseCase)

	//GIN routes
	router := gin.Default()
	router.POST("/login", handler.Login)
	delivery.PaymentRoutes(router, paymentPostHandler, paymentGetHandler, refundHandler)

	// Start the server //TODO refactor
	log.Println("Starting payment API on :8082")
	if err := router.Run(":8082"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
