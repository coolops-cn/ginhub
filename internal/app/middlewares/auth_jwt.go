package middlewares

import (
	"github.com/coolops-cn/ginhub/internal/app/models/user"
	"github.com/coolops-cn/ginhub/pkg/errors"
	"github.com/coolops-cn/ginhub/pkg/jwt"
	"github.com/coolops-cn/ginhub/pkg/response"
	"github.com/gin-gonic/gin"
)

// auth jwt验证中间件

func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从标头 Authorization:Bearer xxxxx 中获取信息，并验证 JWT 的准确性
		cliam, err := jwt.NewJWT().ParserToken(c)

		// 如果解析失败，返回
		if err != nil {
			response.WriteResponse(c, errors.ErrTokenParser, err.Error())
			return
		}

		// 解析成功，设置用户信息
		_user := user.Get(cliam.UserID)
		if _user.ID == 0 {
			response.WriteResponse(c, errors.UserNotFound, nil)
			return
		}

		// 将用户信息写入上下文
		c.Set("current_user_id", _user.GetStringID())
		c.Set("current_user_name", _user.Username)
		c.Set("current_user", _user)

		c.Next()
	}
}
