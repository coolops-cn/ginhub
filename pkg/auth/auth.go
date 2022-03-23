package auth

import (
	"errors"

	"github.com/coolops-cn/ginhub/internal/app/models/user"
	"github.com/coolops-cn/ginhub/pkg/logger"
	"github.com/gin-gonic/gin"
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

// CurrentUser 从 gin.context 中获取当前登录用户
func CurrentUser(c *gin.Context) user.User {
	userModel, ok := c.MustGet("current_user").(user.User)
	if !ok {
		logger.LogIf(errors.New("无法获取用户"))
		return user.User{}
	}
	// db is now a *DB value
	return userModel
}

// CurrentUID 从 gin.context 中获取当前登录用户 ID
func CurrentUID(c *gin.Context) string {
	return c.GetString("current_user_id")
}
