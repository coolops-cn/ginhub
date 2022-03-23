package cmd

import (
	"github.com/coolops-cn/ginhub/bootstrap"
	"github.com/coolops-cn/ginhub/pkg/config"
	"github.com/coolops-cn/ginhub/pkg/console"
	"github.com/coolops-cn/ginhub/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

// 运行服务

var ServeCommand = &cobra.Command{
	Use:   "serve",
	Short: "Start WEB Server",
	Run:   runWeb,
	Args:  cobra.NoArgs,
}

// runWeb 运行服务
func runWeb(cmd *cobra.Command, args []string) {
	gin.SetMode(gin.ReleaseMode)

	// 初始化 Gin 实例
	router := gin.New()

	// 初始化路由
	bootstrap.SetupRouter(router)

	// 运行服务
	err := router.Run(":" + config.Get("app.port"))
	if err != nil {
		logger.ErrorString("CMD", "serve", err.Error())
		console.Exit("Unable to start server, error:" + err.Error())
	}
}
