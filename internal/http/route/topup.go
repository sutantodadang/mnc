package route

import (
	"mnc/internal/app/topup"
	"mnc/internal/http/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterTopUpRoute(app *gin.Engine, handler topup.ITopUpHandler, middleware *middlewares.Middleware) {
	topUpRoute := app.Group("/api/v1")

	topUpRoute.POST("/topup", middleware.Auth(), handler.AddTopUp)

}
