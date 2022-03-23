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
	ErrTokenExpired           = NewError(100010, "令牌已过期")
	ErrTokenExpiredMaxRefresh = NewError(100011, "令牌已过最大刷新时间")
	ErrTokenMalformed         = NewError(100012, "请求令牌格式有误")
	ErrTokenInvalid           = NewError(100013, "请求令牌无效")
	ErrTokenRefresh           = NewError(100014, "令牌刷新失败")
	ErrHeaderEmpty            = NewError(100015, "需要认证才能访问！")
	ErrHeaderMalformed        = NewError(100016, "请求头中 Authorization 格式有误")
	ErrTokenParser            = NewError(100017, "Token解析失败")

	// 用户相关
	UserSignupFailed      = NewError(200001, "用户创建失败")
	UserLoginUnauthorized = NewError(200002, "登录校验失败")
	UserNotFound          = NewError(200003, "用户信息不存在")
	UserPasswordError     = NewError(200004, "用户密码错误")
)
