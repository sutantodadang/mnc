package payment

import (
	"errors"
	"mnc/internal/constants"
	"mnc/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type IPaymentHandler interface {
	MakePayment(c *gin.Context)
}

type PaymentHandler struct {
	service IPaymentService
}

// MakePayment implements IPaymentHandler.
func (r *PaymentHandler) MakePayment(c *gin.Context) {

	req := new(MakePaymentRequest)

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

	data, err := r.service.MakePayment(c, *req)
	if err != nil {
		utils.ResponseJson(c, 500, nil, err)
		return
	}

	utils.ResponseJson(c, 201, data, err)
}

func NewPaymentHandler(service IPaymentService) IPaymentHandler {
	return &PaymentHandler{
		service: service,
	}
}
