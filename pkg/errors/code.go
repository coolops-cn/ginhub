package errors

// 公共错误代码
var (
	Success             = NewError(0, "成功")
	ServerError         = NewError(100001, "服务内部错误")
	InvalidParams       = NewError(100002, "入参错误")
	NotFound            = NewError(100003, "找不到")
	TooManyRequests     = NewError(100004, "请求太多")
	UserNotLogin        = NewError(100005, "用户未登录")
	UnprocessableEntity = NewError(100006, "请求验证不通过")

	// Token 相关
	ErrTokenExpired           = NewError(1000010, "令牌已过期")
	ErrTokenExpiredMaxRefresh = NewError(1000011, "令牌已过最大刷新时间")
	ErrTokenMalformed         = NewError(1000012, "请求令牌格式有误")
	ErrTokenInvalid           = NewError(1000013, "请求令牌无效")
	ErrTokenRefresh           = NewError(1000014, "令牌刷新失败")
	ErrHeaderEmpty            = NewError(1000015, "需要认证才能访问！")
	ErrHeaderMalformed        = NewError(1000016, "请求头中 Authorization 格式有误")

	// 用户相关
	UserSignupFailed      = NewError(200001, "用户创建失败")
	UserLoginUnauthorized = NewError(200002, "登录校验失败")
)
