package requests

import (
	"github.com/coolops-cn/ginhub/internal/app/requests/validators"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

// 重置密码请求

type ResetPasswordRequest struct {
	OldPassword        string `json:"old_password" valid:"old_password"`
	NewPassword        string `json:"new_password" valid:"new_password"`
	NewPasswordConfirm string `json:"new_password_confirm" valid:"new_password_confirm"`
}

// ValidateResetPassword 校验请求参数
func ValidateResetPassword(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"old_password":         []string{"required", "min:6"},
		"new_password":         []string{"required", "min:6"},
		"new_password_confirm": []string{"required"},
	}

	messages := govalidator.MapData{
		"old_password": []string{
			"required:旧密码是必须参数",
			"min:密码长度大于6",
		},
		"new_password": []string{
			"required:新密码是必须参数",
			"min:密码长度大于6",
		},
		"new_password_confirm": []string{
			"required:新密码确认是必须参数",
		},
	}

	//. 进行校验
	errs := validate(data, rules, messages)

	// 校验两次输入的新密码是否一致
	_data := data.(*ResetPasswordRequest)
	errs = validators.ValidatePasswordConfirm(_data.NewPassword, _data.NewPasswordConfirm, errs)

	return errs
}
