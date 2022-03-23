package auth

import (
	v1 "github.com/coolops-cn/ginhub/internal/app/http/controller/api/v1"
	"github.com/coolops-cn/ginhub/internal/app/requests"
	"github.com/coolops-cn/ginhub/pkg/auth"
	"github.com/coolops-cn/ginhub/pkg/errors"
	"github.com/coolops-cn/ginhub/pkg/response"
	"github.com/gin-gonic/gin"
)

// 重置密码
type PasswordController struct {
	v1.BaseAPIController
}

// ResetPassword 重置密码
func (p *PasswordController) ResetPassword(c *gin.Context) {

	request := requests.ResetPasswordRequest{}

	if ok := requests.Validate(&request, c, requests.ValidateResetPassword); !ok {
		return
	}

	// 校验旧密码是否正确
	_user := auth.CurrentUser(c)
	_newUser, err := auth.Attempt(_user.Username, request.OldPassword)
	if err != nil {
		response.WriteResponse(c, errors.UserPasswordError, nil)
		return
	}
	// 更新密码
	_newUser.Password = request.NewPassword
	_newUser.Save()
	response.WriteResponse(c, nil, gin.H{
		"data": "用户密码修改成功",
	})
}
