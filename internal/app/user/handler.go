package user

import (
	"errors"
	"mnc/internal/constants"
	"mnc/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type IUserHandler interface {
	RegisterUser(c *gin.Context)
	LoginUser(c *gin.Context)
	UpdateUser(c *gin.Context)
}

type UserHandler struct {
	service IUserService
}

// UpdateUser implements IUserHandler.
func (u *UserHandler) UpdateUser(c *gin.Context) {

	req := new(UpdateUserRequest)

	if err := c.ShouldBindJSON(&req); err != nil {

		if errs, ok := err.(validator.ValidationErrors); ok {
			c.JSON(400, gin.H{"error": utils.NewValidationError(errs)})
			return
		}

		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	id, ok := c.Get(constants.USER_ID)
	if !ok {
		utils.ResponseJson(c, 400, nil, errors.New("user id not found"))
		return
	}

	req.UserID = id.(string)

	data, err := u.service.UpdateUser(c, *req)
	if err != nil {
		utils.ResponseJson(c, 500, nil, err)
		return
	}

	utils.ResponseJson(c, 200, data, err)
}

// LoginUser implements IUserHandler.
func (u *UserHandler) LoginUser(c *gin.Context) {
	req := new(LoginUserRequest)

	if err := c.ShouldBindJSON(&req); err != nil {

		if errs, ok := err.(validator.ValidationErrors); ok {
			c.JSON(400, gin.H{"error": utils.NewValidationError(errs)})
			return
		}

		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	data, err := u.service.LoginUser(c, *req)
	if err != nil {
		utils.ResponseJson(c, 500, nil, err)
		return
	}

	utils.ResponseJson(c, 200, data, err)

}

// RegisterUser implements IUserHandler.
func (u *UserHandler) RegisterUser(c *gin.Context) {

	req := new(RegisterUserRequest)

	if err := c.ShouldBindJSON(&req); err != nil {

		if errs, ok := err.(validator.ValidationErrors); ok {
			c.JSON(400, gin.H{"error": utils.NewValidationError(errs)})
			return
		}

		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	data, err := u.service.RegisterUser(c, *req)
	if err != nil {
		utils.ResponseJson(c, 500, nil, err)
		return
	}

	utils.ResponseJson(c, 201, data, err)
}

func NewUserHandler(service IUserService) IUserHandler {
	return &UserHandler{
		service: service,
	}
}
