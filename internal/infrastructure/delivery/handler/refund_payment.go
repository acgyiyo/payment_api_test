package handler

import (
	"net/http"

	"github.com/acgyiyo/payment_api_test/internal/domain/entity"
	"github.com/acgyiyo/payment_api_test/internal/domain/usecase"
	"github.com/gin-gonic/gin"
)

type RefundPayment interface {
	ProcessRefund(ctx *gin.Context)
}

type refundPayment struct {
	usecase usecase.UpdatePayment
}

func NewRefundPayment(rp usecase.UpdatePayment) RefundPayment {
	return &refundPayment{
		usecase: rp,
	}
}

func (p *refundPayment) ProcessRefund(ctx *gin.Context) {
	var refund entity.RefundRequest
	if err := ctx.BindJSON(&refund); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"type": "error", "message": err.Error()})
		return
	}

	//TODO validate inputs

	response, err := p.usecase.UpdatePayment(ctx, convertRefundToPayment(refund))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"type": "error", "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func convertRefundToPayment(req entity.RefundRequest) entity.Payment {
	return entity.Payment{
		TransactionID: req.TransactionID,
	}
}
