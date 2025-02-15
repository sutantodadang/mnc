package transaction

import (
	"errors"
	"mnc/internal/constants"
	"mnc/internal/utils"

	"github.com/gin-gonic/gin"
)

type ITransactionHandler interface {
	TransactionReport(c *gin.Context)
}

type TransactionHandler struct {
	service ITransactionService
}

// TransactionReport implements ITransactionHandler.
func (o *TransactionHandler) TransactionReport(c *gin.Context) {

	id, ok := c.Get(constants.USER_ID)
	if !ok {
		utils.ResponseJson(c, 400, nil, errors.New("user id not found"))
		return
	}

	data, err := o.service.TransactionReport(c, id.(string))
	if err != nil {
		utils.ResponseJson(c, 500, nil, err)
		return
	}

	utils.ResponseJson(c, 200, data, nil)
}

func NewTransactionHandler(service ITransactionService) ITransactionHandler {
	return &TransactionHandler{
		service: service,
	}
}
