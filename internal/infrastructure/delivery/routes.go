package delivery

import (
	"github.com/acgyiyo/payment_api_test/internal/infrastructure/delivery/handler"
	"github.com/gin-gonic/gin"
)

func PaymentRoutes(route *gin.Engine, paymentPostHandler handler.PostPayment, paymentGetHanlder handler.GetPayment, refundHandler handler.RefundPayment) {
	routes := route.Group("/payment2")
	{
		routes.POST("", paymentPostHandler.ProcessPayment)
		routes.GET("/:id", paymentGetHanlder.QueryPayment)
		routes.POST("/refund", refundHandler.ProcessRefund)
	}
}
