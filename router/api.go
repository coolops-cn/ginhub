package router

import (
	"github.com/coolops-cn/ginhub/internal/app/http/controller/api/v1/auth"
	"github.com/gin-gonic/gin"
)

func RegisterAPIRouter(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		authGroup := v1.Group("/auth")
		{
			suc := new(auth.SignupController)
			authGroup.POST("/signup/username/exist", suc.IsUsernameExist)
			authGroup.POST("/signup/using-password", suc.SignupUsingPassword)

			luc := new(auth.LoginController)
			authGroup.POST("/login/using-password", luc.LoginByPassword)
			authGroup.POST("/login/refresh-token", luc.RefreshToken)
		}
	}
}
