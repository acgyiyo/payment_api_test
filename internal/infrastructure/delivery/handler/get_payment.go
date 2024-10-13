package handler

import (
	"net/http"

	"github.com/acgyiyo/payment_api_test/internal/domain/usecase"
	"github.com/gin-gonic/gin"
)

type GetPayment interface {
	QueryPayment(ctx *gin.Context)
}

type getPayment struct {
	usecase usecase.RetrievePayment
}

func NewGetPayment(rp usecase.RetrievePayment) GetPayment {
	return &getPayment{
		usecase: rp,
	}
}

func (p *getPayment) QueryPayment(ctx *gin.Context) {
	transactionID := ctx.Param("id")
	if transactionID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing transactionID in the request"})
		return
	}

	//TODO validate inputs, numeric regex, etc

	response, err := p.usecase.SearchPaymentByTransactionID(ctx, transactionID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"type": "error", "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (p *getPayment) QueryPaymentByMerchat(ctx *gin.Context) {
	//TODO to be implemented
	//return payment[]
}
