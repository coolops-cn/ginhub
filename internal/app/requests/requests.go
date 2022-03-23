package requests

import (
	"fmt"

	"github.com/coolops-cn/ginhub/pkg/errors"
	"github.com/coolops-cn/ginhub/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

// 定义验证函数类型
type ValidateFunc func(interface{}, *gin.Context) map[string][]string

func Validate(data interface{}, c *gin.Context, handler ValidateFunc) bool {
	// 1. 解析请求，支持 JSON 数据、表单请求和 URL Query
	if err := c.ShouldBind(data); err != nil {
		response.WriteResponse(c, errors.InvalidParams, nil)
		fmt.Println(err.Error())
		return false
	}

	// 2. 表单验证
	errs := handler(data, c)

	// 3. 判断验证是否通过
	if len(errs) > 0 {
		response.WriteResponse(c, errors.UnprocessableEntity, errs)
		return false
	}

	return true
}

func validate(data interface{}, rules govalidator.MapData, messages govalidator.MapData) map[string][]string {
	// 配置选项
	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		TagIdentifier: "valid", // 模型中的 Struct 标签标识符
		Messages:      messages,
	}

	// 开始验证
	return govalidator.New(opts).ValidateStruct()
}
