package errno

var (
	//公共错误码
	OK                  = &Errno{Code: 0, Message: "成功"}
	InternalServerError = &Errno{Code: 10001, Message: "内部服务错误"}
	ErrBind             = &Errno{Code: 10002, Message: "请求body绑定到结构体时错误"}

	Errvalidation = &Errno{Code: 20001, Message: "验证失败"}
	ErrDatabase   = &Errno{Code: 20002, Message: "数据库失败"}
	ErrToken      = &Errno{Code: 20003, Message: "登录token错误"}

	//用户模块 错误码
	ErrUserNotFound      = &Errno{Code: 20102, Message: "没有找到这个用户"}
	ErrPasswordNil       = &Errno{Code: 20103, Message: "密码为空"}
	ErrEncrypt           = &Errno{Code: 20104, Message: "密码加密错误"}
	ErrTokenInvalid      = &Errno{Code: 20105, Message: "token失效"}
	ErrPasswordIncorrect = &Errno{Code: 20106, Message: "密码错误"}
)
