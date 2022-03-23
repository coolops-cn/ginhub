package auth

import (
	"errors"

	"github.com/coolops-cn/ginhub/internal/app/models/user"
)

// Attempt 尝试登录
func Attempt(name string, password string) (user.User, error) {
	_user := user.GetByMulti(name)
	if _user.ID == 0 {
		return user.User{}, errors.New("账号不存在")
	}

	if !_user.ComparePassword(password) {
		return user.User{}, errors.New("密码错误")
	}

	return _user, nil
}
