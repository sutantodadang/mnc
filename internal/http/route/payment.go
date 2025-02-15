package route

import (
	"mnc/internal/app/payment"
	"mnc/internal/http/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterPaymentRoute(app *gin.Engine, handler payment.IPaymentHandler, middleware *middlewares.Middleware) {
	paymentRoute := app.Group("/api/v1")

	paymentRoute.POST("/pay", middleware.Auth(), handler.MakePayment)

}
