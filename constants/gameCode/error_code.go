package gameCode

type ErrorCode int

// 基础验证码
const (
	RequestSuccess                  = iota + 10000 // 请求成功
	RequestFailed                                  // 请求失败
	RequestTokenError                              // token error
	RequestParamsError                             // 参数错误
	RequestLoginUserOrPasswordError                // 登录时用户名或者密码错误

)

// 用户相关 20000-30000
const (
	UserNameAlreadyExist = iota + 20000
)

// 微信相关
const (
	ErrorCodeWechatGetJSconfig               = iota + 3000 // 获取微信签名错误
	ErrorCodeWechatSendTemplateInfo                        // 发送微信模板消息错误
	ErrorCodeWechatCodeToAccessToken                       // code to accessToken错误
	ErrorCodeWechatDownloadResourceByAudioID               // 下载媒体资源报错
)
