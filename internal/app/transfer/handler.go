package transfer

import (
	"errors"
	"mnc/internal/constants"
	"mnc/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ITransferHandler interface {
	MakeTransfer(c *gin.Context)
}

type TransferHandler struct {
	service ITransferService
}

// MakeTransfer implements ITransferHandler.
func (p *TransferHandler) MakeTransfer(c *gin.Context) {

	req := new(MakeTransferRequest)
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

	req.UserId = id.(string)

	data, err := p.service.MakeTransfer(c, *req)
	if err != nil {
		utils.ResponseJson(c, 500, nil, err)
		return
	}

	utils.ResponseJson(c, 201, data, err)
}

func NewTransferHandler(service ITransferService) ITransferHandler {
	return &TransferHandler{
		service: service,
	}
}
