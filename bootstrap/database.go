package bootstrap

import (
	"errors"
	"fmt"
	"time"

	"github.com/coolops-cn/ginhub/internal/app/models/user"
	"github.com/coolops-cn/ginhub/pkg/config"
	"github.com/coolops-cn/ginhub/pkg/database"
	"github.com/coolops-cn/ginhub/pkg/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 初始化数据库

func SetupDatabase() {
	var dbConfig gorm.Dialector

	switch config.Get("database.type") {
	case "mysql":
		dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&multiStatements=true&loc=Local",
			config.Get("database.username"),
			config.Get("database.password"),
			config.Get("database.host"),
			config.Get("database.port"),
			config.Get("database.database"),
			config.Get("database.charset"),
		)

		dbConfig = mysql.New(mysql.Config{
			DSN: dsn,
		})
	default:
		panic(errors.New("不支持的数据库类型"))
	}

	// 连接数据库
	database.Connect(dbConfig, logger.NewGormLogger())

	// 设置最大连接数
	database.SQLDB.SetMaxOpenConns(config.GetInt("database.max_open_connections"))
	// 设置最大空闲连接数
	database.SQLDB.SetMaxIdleConns(config.GetInt("database.max_idle_connections"))
	// 设置每个链接的过期时间
	database.SQLDB.SetConnMaxLifetime(time.Duration(config.GetInt("database.max_life_seconds")) * time.Second)

	// 数据库自动迁移
	database.DB.AutoMigrate(&user.User{})
}
