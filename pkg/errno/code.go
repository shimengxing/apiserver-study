package errno

var (
	//公共错误码
	OK                  = &Errno{Code: 0, Message: "成功"}
	InternalServerError = &Errno{Code: 10001, Message: "内部服务错误"}
	ErrBind             = &Errno{Code: 10002, Message: "请求body绑定到结构体时错误"}

	//用户模块 错误码
	ErrUserNotFound = &Errno{Code: 20102, Message: "没有找到这个用户"}
)
