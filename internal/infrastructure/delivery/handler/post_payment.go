package handler

import (
	"net/http"

	"github.com/acgyiyo/payment_api_test/internal/domain/entity"
	"github.com/acgyiyo/payment_api_test/internal/domain/usecase"
	"github.com/gin-gonic/gin"
)

type PostPayment interface {
	ProcessPayment(ctx *gin.Context)
}

type postPayment struct {
	usecase usecase.RegisterPayment
}

func NewPostPayment(rp usecase.RegisterPayment) PostPayment {
	return &postPayment{
		usecase: rp,
	}
}

func (p *postPayment) ProcessPayment(ctx *gin.Context) {
	var paymentReq entity.PaymentRequest
	if err := ctx.BindJSON(&paymentReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"type": "error", "message": err.Error()})
		return
	}

	//TODO validate inputs

	response, err := p.usecase.SavePayment(ctx, convertReqToPayment(paymentReq))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"type": "error", "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func convertReqToPayment(req entity.PaymentRequest) entity.Payment {
	return entity.Payment{
		Amount:       req.Amount,
		MerchantID:   req.MerchantID,
		CustomerCard: req.CustomerCard,
	}
}
