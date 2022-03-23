package response

import (
	"net/http"

	"github.com/coolops-cn/ginhub/pkg/errors"
	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details"`
}

func WriteResponse(c *gin.Context, err *errors.Error, data interface{}) {
	if err != nil {
		c.JSON(err.StatusCode(), ErrorResponse{
			Code:    err.Code(),
			Message: err.Msg(),
			Details: data,
		})
		return
	}
	c.JSON(http.StatusOK, data)
}
