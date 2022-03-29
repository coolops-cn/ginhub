package middlewares

import (
	"github.com/gin-gonic/gin"

	uuid "github.com/satori/go.uuid"
)

// 为请求添加 request id

const (
	// XRequestIDKey defines X-Request-ID key string.
	XRequestIDKey = "X-Request-ID"
)

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 检查 Header 中是否存在 XRequestIDKey
		rid := c.GetHeader(XRequestIDKey)

		if rid == "" {
			var err error
			rid = uuid.Must(uuid.NewV4(), err).String()
			c.Request.Header.Set(XRequestIDKey, rid)
			c.Set(XRequestIDKey, rid)
		}

		// 设置 XRequestIDKey header
		c.Writer.Header().Set(XRequestIDKey, rid)
		c.Next()
	}
}
