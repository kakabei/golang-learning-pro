package models

const (
	ErrorCodeParam     = 1 // 参数错误
	ErrorCodeSys       = 2 // 系统错误
	ErrorCodeNotExists = 3 // 任务不存在
)

const (
	ErrorCodeSucc         = iota
	ErrorCodeError        = 1
	ErrorCodePramInvalid  = 2
	ErrorCodeUnmarshal    = 3
	ErrorCodeAssignFailed = 4
)
