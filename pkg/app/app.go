package app

import (
	"time"

	"github.com/coolops-cn/ginhub/pkg/config"
)

func IsLocal() bool {
	return config.Get("app.env") == "local"
}

func IsTest() bool {
	return config.Get("app.env") == "test"
}

func IsProd() bool {
	return config.Get("app.env") == "prod"
}

// TimenowInTimezone 获取当前时间，支持时区
func TimenowInTimezone() time.Time {
	chinaTimezone, _ := time.LoadLocation(config.GetString("app.timezone"))
	return time.Now().In(chinaTimezone)
}
