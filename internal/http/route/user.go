package route

import (
	"mnc/internal/app/user"
	"mnc/internal/http/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoute(app *gin.Engine, handler user.IUserHandler, middleware *middlewares.Middleware) {
	userRoute := app.Group("/api/v1")

	userRoute.POST("/register", handler.RegisterUser)
	userRoute.POST("/login", handler.LoginUser)
	userRoute.PUT("/profile", middleware.Auth(), handler.UpdateUser)

}
