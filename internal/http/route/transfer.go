package route

import (
	"mnc/internal/app/transfer"
	"mnc/internal/http/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterTransferRoute(app *gin.Engine, handler transfer.ITransferHandler, middleware *middlewares.Middleware) {
	transferRoute := app.Group("/api/v1")

	transferRoute.POST("/transfer", middleware.Auth(), handler.MakeTransfer)

}
