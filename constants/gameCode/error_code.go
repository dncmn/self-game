package gameCode

type ErrorCode int

// 基础验证码
const (
	RequestSuccess     = iota + 10000         // 请求成功
	RequestFailed                             // 请求失败
	RequestParamsError int            = 10001 // 参数错误

)

// 用户相关 20000-30000
const (
	UserNameAlreadyExist = iota + 20000
)
