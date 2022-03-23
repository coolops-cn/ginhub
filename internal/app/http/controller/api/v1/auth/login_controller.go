package auth

import (
	"fmt"

	v1 "github.com/coolops-cn/ginhub/internal/app/http/controller/api/v1"
	"github.com/coolops-cn/ginhub/internal/app/requests"
	"github.com/coolops-cn/ginhub/internal/pkg/auth"
	"github.com/coolops-cn/ginhub/pkg/errors"
	"github.com/coolops-cn/ginhub/pkg/jwt"
	"github.com/coolops-cn/ginhub/pkg/response"
	"github.com/gin-gonic/gin"
)

// 登录相关控制器

type LoginController struct {
	v1.BaseAPIController
}

func (lc *LoginController) LoginByPassword(c *gin.Context) {

	request := requests.LoginByPasswordRequest{}

	if ok := requests.Validate(&request, c, requests.ValidateLoginByPassword); !ok {
		return
	}

	// 校验用户名密码是否正确
	_user, err := auth.Attempt(request.Username, request.Password)

	if err != nil {
		response.WriteResponse(c, errors.UserLoginUnauthorized, nil)
		fmt.Println("111111111111111", err.Error())
	} else {
		token := jwt.NewJWT().IssueToken(_user.GetStringID(), _user.Username)
		response.WriteResponse(c, nil, gin.H{
			"token": token,
		})
	}
}

// RefreshToken 刷新 Access Token
func (lc *LoginController) RefreshToken(c *gin.Context) {

	token, err := jwt.NewJWT().RefreshToken(c)

	if err != nil {
		response.WriteResponse(c, errors.ErrTokenRefresh, nil)
	} else {
		response.WriteResponse(c, nil, gin.H{
			"token": token,
		})
	}
}
