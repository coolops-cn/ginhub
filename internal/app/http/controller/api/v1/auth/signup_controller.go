package auth

import (
	v1 "github.com/coolops-cn/ginhub/internal/app/http/controller/api/v1"
	"github.com/coolops-cn/ginhub/internal/app/models/user"
	"github.com/coolops-cn/ginhub/internal/app/requests"
	"github.com/coolops-cn/ginhub/pkg/errors"
	"github.com/coolops-cn/ginhub/pkg/jwt"
	"github.com/coolops-cn/ginhub/pkg/response"
	"github.com/gin-gonic/gin"
)

// 用户注册Controller

type SignupController struct {
	v1.BaseAPIController
}

// IsUsernameExist 判断用户是否存在
func (sc *SignupController) IsUsernameExist(c *gin.Context) {

	request := requests.SignupUsernameExistRequest{}

	if ok := requests.Validate(&request, c, requests.ValidateSignupUsernameExist); !ok {
		return
	}
	// 检查数据库并返回响应
	response.WriteResponse(c, nil, gin.H{
		"exist": user.IsUsernameExist(request.Username),
	})
}

func (sc *SignupController) SignupUsingPassword(c *gin.Context) {
	request := requests.SignupUsingPasswordRequest{}

	if ok := requests.Validate(&request, c, requests.ValidateSignupUsingPassword); !ok {
		return
	}

	// 创建数据
	_user := user.User{
		Username: request.Username,
		Password: request.Password,
	}

	_user.Create()

	if _user.ID > 0 {
		token := jwt.NewJWT().IssueToken(_user.GetStringID(), _user.Username)
		response.WriteResponse(c, nil, gin.H{
			"data":  _user,
			"token": token,
		})
	} else {
		response.WriteResponse(c, errors.UserSignupFailed, nil)
	}
}
