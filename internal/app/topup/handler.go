package topup

import (
	"errors"
	"mnc/internal/constants"
	"mnc/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ITopUpHandler interface {
	AddTopUp(c *gin.Context)
}

type TopUpHandler struct {
	service ITopUpService
}

// RegisterWarehouse implements IWarehouseHandler.
func (t *TopUpHandler) AddTopUp(c *gin.Context) {

	req := new(AddTopup)
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

	req.Userid = id.(string)

	data, err := t.service.AddTopUp(c, *req)
	if err != nil {
		utils.ResponseJson(c, 500, nil, err)
		return
	}

	utils.ResponseJson(c, 200, data, err)

}

func NewTopUpHandler(service ITopUpService) ITopUpHandler {
	return &TopUpHandler{
		service: service,
	}
}
