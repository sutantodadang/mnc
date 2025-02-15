package utils

import "github.com/gin-gonic/gin"

type Response struct {
	Status string `json:"status"`
	Result any    `json:"result"`
}

func ResponseJson(c *gin.Context, code int, data any, err error) {

	if err != nil {
		c.JSON(code, gin.H{"error": err.Error()})
		return

	}

	resp := new(Response)
	resp.Status = "SUCCESS"
	resp.Result = data

	c.JSON(code, resp)
}
