package handlers

import (
	"log"
	"net/http"

	"github.com/acgyiyo/payment_api_test/internal/models"
	"github.com/acgyiyo/payment_api_test/internal/services"
	"github.com/gin-gonic/gin"
)

func ProcessRefund(c *gin.Context) {
	var refund models.RefundRequest

	// validate request
	//TODO ......

	if err := c.BindJSON(&refund); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//check if payment belongs to merchant and customer for security porpouses
	//TODO chencking....

	payment, exist := models.GetTransaction(refund.TransactionID)
	if !exist {
		log.Print("transaction not found")                              //TODO Improve
		c.JSON(http.StatusInternalServerError, "transaction not found") //TODO response error
	}

	result, err := services.BankSimulator.ProcessRefundInBank(payment)
	if err != nil {
		log.Print("process in bank failed")         //TODO IMPROVE
		c.JSON(http.StatusInternalServerError, err) //TODO response error
	}

	err = models.UpdateTransaction(result)
	if err != nil {
		log.Print("process in bank failed")         //TODO improve
		c.JSON(http.StatusInternalServerError, err) //TODO response error
	}

	//TODO notify in topic update

	c.JSON(http.StatusOK, result)
}
