package handler

import (
	"net/http"

	"github.com/acgyiyo/payment_api_test/internal/infrastructure/service/auth"
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	MerchantID string `json:"merchant_id"`
	Password   string `json:"password"`
}

func Login(c *gin.Context) {
	var loginReq LoginRequest

	// Bind the JSON request body to the LoginRequest struct
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// In real systems, you should verify password and merchant details.
	if loginReq.MerchantID == "merchant123" && loginReq.Password == "password" {
		token, err := auth.GenerateJWT(loginReq.MerchantID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})

	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid merchant credentials"})
	}
}
