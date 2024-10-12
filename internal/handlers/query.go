package handlers

import (
	"net/http"

	"github.com/acgyiyo/payment_api_test/internal/models"
	"github.com/gin-gonic/gin"
)

func QueryPayment(c *gin.Context) {
	paymentID := c.Param("id")
	if paymentID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing payment ID in the request"})
		return
	}

	payment, exists := models.GetTransaction(paymentID)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Payment ID not found"})
		return
	}

	c.JSON(http.StatusOK, payment)
}
