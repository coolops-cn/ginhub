package config

import (
	"os"

	"github.com/coolops-cn/ginhub/pkg/helpers"
	"github.com/spf13/cast"
	viperlib "github.com/spf13/viper"
)

const defaultEnv = ".env"

var viper *viperlib.Viper

// ConfigFunc 动态加载配置信息
type ConfigFunc func() map[string]interface{}

// ConfigFuncs 先加载此数组
var ConfigFuncs map[string]ConfigFunc

func init() {
	viper = viperlib.New()

	// 设置类型
	viper.SetConfigType("env")

	// 设置环境变量的路径，相对main.go
	viper.AddConfigPath(".")

	// 设置环境变量前缀
	viper.SetEnvPrefix("GinHub")

	// 读取环境变量
	viper.AutomaticEnv()

	ConfigFuncs = make(map[string]ConfigFunc)
}

// InitConfig 初始化配置文件
func InitConfig(env string) {
	// 加载环境变量
	loadEnv(env)

	// 注册配置信息
	registerConfig()
}

// 加载配置信息
func registerConfig() {
	for name, fn := range ConfigFuncs {
		viper.Set(name, fn())
	}
}

// 加载环境变量
func loadEnv(suffix string) {
	// 默认加载.env环境变量
	envPath := defaultEnv

	if len(suffix) > 0 {
		newEnv := ".env" + suffix

		// 如果文件存在，加载新的文件
		if _, err := os.Stat(newEnv); err == nil {
			envPath = newEnv
		}
	}

	// 加载env
	viper.SetConfigName(envPath)

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	// 监听环境变量配置文件，实时加载
	viper.WatchConfig()
}

// Env 读取环境变量，支持默认值
func Env(envName string, defaultValue ...interface{}) interface{} {
	if len(defaultValue) > 0 {
		return internalGet(envName, defaultValue[0])
	}
	return internalGet(envName)
}

// Add 添加环境变量
func Add(name string, fn ConfigFunc) {
	ConfigFuncs[name] = fn
}

// Get 获取环境变量
func Get(name string, defaultValue ...interface{}) string {
	return GetString(name, defaultValue)
}

func internalGet(path string, defaultValue ...interface{}) interface{} {
	// config 或环境变量不存在的情况
	if !viper.IsSet(path) || helpers.Empty(viper.Get(path)) {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
	}
	return viper.Get(path)
}

// GetString 获取 String 类型的配置信息
func GetString(name string, defaultValue ...interface{}) string {
	return cast.ToString(internalGet(name, defaultValue))
}

// GetInt 获取 Int 类型的配置信息
func GetInt(path string, defaultValue ...interface{}) int {
	return cast.ToInt(internalGet(path, defaultValue...))
}

// GetFloat64 获取 float64 类型的配置信息
func GetFloat64(path string, defaultValue ...interface{}) float64 {
	return cast.ToFloat64(internalGet(path, defaultValue...))
}

// GetInt64 获取 Int64 类型的配置信息
func GetInt64(path string, defaultValue ...interface{}) int64 {
	return cast.ToInt64(internalGet(path, defaultValue...))
}

// GetUint 获取 Uint 类型的配置信息
func GetUint(path string, defaultValue ...interface{}) uint {
	return cast.ToUint(internalGet(path, defaultValue...))
}

// GetBool 获取 Bool 类型的配置信息
func GetBool(path string, defaultValue ...interface{}) bool {
	return cast.ToBool(internalGet(path, defaultValue...))
}

// GetStringMapString 获取结构数据
func GetStringMapString(path string) map[string]string {
	return viper.GetStringMapString(path)
}
