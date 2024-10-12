package handlers

import (
	"log"
	"net/http"

	"github.com/acgyiyo/payment_api_test/internal/models"
	"github.com/gin-gonic/gin"

	"github.com/acgyiyo/payment_api_test/internal/services"
)

func ProcessPayment(c *gin.Context) {
	var paymentReq models.PaymentRequest
	if err := c.BindJSON(&paymentReq); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	result, err := services.BankSimulator.ProcessPaymentInBank(paymentReq)
	if err != nil {
		log.Print("error processing payment: ProcessPaymentInBank failed", err)
	}

	result.Amount = paymentReq.Amount
	result.MerchantID = paymentReq.MerchantID
	err = models.SaveTransaction(result)
	if err != nil {
		log.Printf("error processing payment: SaveTransaction failed: %s", err.Error())
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, result)

}

func errorResponse(msg error) gin.H {
	return gin.H{"type": "error", "message": msg.Error()}
}
