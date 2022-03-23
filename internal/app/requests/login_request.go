package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

// 用户登录

type LoginByPasswordRequest struct {
	Username string `json:"username,omitempty" valid:"username"`
	Password string `json:"password,omitempty" valid:"password"`
}

func ValidateLoginByPassword(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"username": []string{"required", "min:3"},
		"password": []string{"required", "min:6"},
	}

	messages := govalidator.MapData{
		"username": []string{
			"required:登录用户名必须",
			"min:用户名长度需大于3",
		},
		"password": []string{
			"required:密码为必填项",
			"min:密码长度需大于6",
		},
	}

	errs := validate(data, rules, messages)

	return errs
}
