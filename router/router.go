package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"self-game/config"
	_ "self-game/dao"
	"self-game/handler"
	"self-game/utils/qrcode"
)

func Router(r *gin.Engine) {
	r.StaticFS("/compoments/runtime/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))
	wxServerCheck := r.Group("/api/v2")
	{
		wxServerCheck.GET("/check", handler.HandlerSignatureHandler)

	}

	cc := r.Group("/api/v1", handler.VerifyToken)
	userLoginGroup := cc.Group("/anonymous")
	{
		userLoginGroup.POST("/user/register", handler.RegisterUserHandler)
		userLoginGroup.POST("/user/login", handler.UserLoginHandler)
		userLoginGroup.GET("/login_by_uid", handler.GetUserLoginHandler)
	}

	// 用户相关
	user := cc.Group("/user", handler.VerifyUserToken)
	{
		user.GET("/name/:uid", handler.GetUserNameHandler)
		user.POST("/name", handler.PostUserNameHandler)
		user.GET("/health_check", handler.ConsulHealthCheck)
	}

	// 二维码相关
	codeGroup := cc.Group("/qrcode")
	{
		codeGroup.GET("/", handler.GetCodeImageHandler)    // 产生二维码
		codeGroup.GET("/info", handler.GetCodeInfoHandler) // 获取二维码信息---是一个url，就是从那里获取信息
	}

	// 微信相关
	wechatGroup := cc.Group("/wechat")
	{
		wechatGroup.GET("/accesstokne_by_code", handler.WechatAuthHandler)                    // code 换access_token或openid
		wechatGroup.GET("/download_media_by_audioid", handler.WechatDownloadMediaDataHandler) // 根据media_id下载音频数据
		wechatGroup.POST("/send_template_info", handler.WechatSendTemplateInfoHandler)        // 发送模板消息
		wechatGroup.GET("/jsconfig", handler.WechatGetJSConfigHandler)                        // 获取jsconfig
		wechatGroup.POST("/receive_msg", handler.WechatReceiveMsgHandler)                     // 接收各种类型的消息
	}

	// 权限相关
	powerGroup := cc.Group("/power")
	{
		powerGroup.GET("/list:uid")
	}

	err := r.Run(fmt.Sprintf(":%d", config.Config.Cfg.Port))
	if err != nil {
		panic(err)
	}
}
