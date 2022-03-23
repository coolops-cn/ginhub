package bootstrap

import (
	"net/http"
	"strings"

	"github.com/coolops-cn/ginhub/internal/app/middlewares"
	"github.com/coolops-cn/ginhub/router"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	// 注册全局中间件
	registerGlobalMiddleware(r)

	// 注册 API 路由
	router.RegisterAPIRouter(r)

	// 注册 404 页面
	setNoFoundHandler(r)
}

func registerGlobalMiddleware(r *gin.Engine) {
	for _, m := range middlewares.Middleware {
		r.Use(m)
	}
}

func setNoFoundHandler(r *gin.Engine) {
	r.NoRoute(func(c *gin.Context) {
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			c.String(http.StatusNotFound, "404页面")
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"err_code":    404,
				"err_message": "路由未定义,请检查",
			})
		}
	})
}
