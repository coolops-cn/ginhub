package main

import (
	"flag"
	"fmt"

	"github.com/coolops-cn/ginhub/bootstrap"
	"github.com/coolops-cn/ginhub/pkg/config"
	"github.com/gin-gonic/gin"

	internalConfig "github.com/coolops-cn/ginhub/internal/config"
)

func init() {
	internalConfig.InitConfig()
}

func main() {

	// 初始化配置文件
	var env string
	flag.StringVar(&env, "env", "", "加载 .env 配置文件")
	flag.Parse()
	config.InitConfig(env)

	// 初始化日志
	bootstrap.SetupLogger()

	gin.SetMode(gin.ReleaseMode)
	// 创建 Gin
	router := gin.New()

	// 初始化数据库
	bootstrap.SetupDatabase()

	// 初始化路由
	bootstrap.SetupRouter(router)

	// 运行服务
	if err := router.Run(":" + config.Get("app.port")); err != nil {
		fmt.Println(err.Error())
	}
}
