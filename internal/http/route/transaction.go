package route

import (
	"mnc/internal/app/transaction"
	"mnc/internal/http/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterTransactionRoute(app *gin.Engine, handler transaction.ITransactionHandler, middleware *middlewares.Middleware) {
	trxRoute := app.Group("/api/v1")

	trxRoute.GET("/transactions", middleware.Auth(), handler.TransactionReport)

}
