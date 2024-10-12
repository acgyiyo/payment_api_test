package main

import (
	"fmt"
	"log"

	"github.com/acgyiyo/payment_api_test/internal/handlers"
	"github.com/acgyiyo/payment_api_test/internal/models"
	"github.com/acgyiyo/payment_api_test/internal/services"
	"github.com/acgyiyo/payment_api_test/internal/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load the configuration from the file
	config, err := utils.LoadConfig("./internal/config")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	models.InitDB(config)

	services.InitBankSimulator()

	// Create a Gin router
	router := gin.Default()

	// Define routes
	router.POST("/payment", handlers.ProcessPayment)
	router.GET("/payment/:id", handlers.QueryPayment)
	router.POST("/refunds", handlers.ProcessRefund)

	router.GET("/healthcheck ", handlers.HealthCheck)

	// Start the server
	log.Println("Starting payment API on :8082")
	if err := router.Run(":8082"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	fmt.Println("Server started successfully!")
}
