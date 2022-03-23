package requests

import (
	"github.com/coolops-cn/ginhub/internal/app/requests/validators"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

// 请求验证

type SignupUsernameExistRequest struct {
	Username string `json:"username,omitempty" valid:"username"`
}

// ValidateSignupUsernameExist 校验注册用户名
func ValidateSignupUsernameExist(data interface{}, c *gin.Context) map[string][]string {

	// 自定义规则
	rules := govalidator.MapData{
		"username": []string{"requierd"},
	}

	// 自定义错误消息提示
	messages := govalidator.MapData{
		"username": []string{
			"required:用户名为必填项,参数名为 username",
		},
	}

	// 校验
	return validate(data, rules, messages)
}

// 注册
type SignupUsingPasswordRequest struct {
	Username        string `json:"username" valid:"username"`
	Password        string `json:"password,omitempty" valid:"password"`
	PasswordConfirm string `json:"password_confirm,omitempty" valid:"password_confirm"`
}

func ValidateSignupUsingPassword(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"username":         []string{"required", "alpha_num", "between:3,20", "not_exists:users,username"},
		"password":         []string{"required", "min:6"},
		"password_confirm": []string{"required"},
	}

	messages := govalidator.MapData{
		"username": []string{
			"required:用户名为必填项",
			"alpha_num:用户名格式错误，只允许数字和英文",
			"between:用户名长度需在 3~20 之间",
		},
		"password": []string{
			"required:密码为必填项",
			"min:密码长度需大于 6",
		},
		"password_confirm": []string{
			"required:确认密码框为必填项",
		},
	}

	errs := validate(data, rules, messages)

	_data := data.(*SignupUsingPasswordRequest)
	errs = validators.ValidatePasswordConfirm(_data.Password, _data.PasswordConfirm, errs)

	return errs
}
