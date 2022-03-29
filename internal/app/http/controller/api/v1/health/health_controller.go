package health

import (
	"net/http"

	v1 "github.com/coolops-cn/ginhub/internal/app/http/controller/api/v1"
	"github.com/coolops-cn/ginhub/pkg/response"
	"github.com/gin-gonic/gin"
)

type HealthController struct {
	v1.BaseAPIController
}

func (hc HealthController) HealthCheck(c *gin.Context) {
	response.WriteResponse(c, nil, gin.H{
		"status": http.StatusOK,
		"data":   "Health Check OK!",
	})
}
