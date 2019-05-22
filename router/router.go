package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"self-game/config"
	"self-game/controller"
	_ "self-game/dao"
	"self-game/utils/qrcode"
)

func Router(r *gin.Engine) {
	r.StaticFS("/compoments/runtime/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))
	wxServerCheck := r.Group("/api/v2")
	{
		wxServerCheck.GET("/check", controller.HandlerSignatureHandler) // 在官网上配置服务器签名
		wxServerCheck.POST("/check", controller.HandMessagesHandler)    // 被动回复消息

	}

	cc := r.Group("/api/v1", controller.VerifyToken)
	userLoginGroup := cc.Group("/anonymous")
	{
		userLoginGroup.POST("/user/register", controller.RegisterUserHandler)
		userLoginGroup.POST("/user/login", controller.UserLoginHandler)
		userLoginGroup.GET("/login_by_uid", controller.GetUserLoginHandler)
	}

	// 用户相关
	user := cc.Group("/user", controller.VerifyUserToken)
	{
		user.GET("/name/:uid", controller.GetUserNameHandler)
		user.POST("/name", controller.PostUserNameHandler)
		user.GET("/health_check", controller.ConsulHealthCheck)
	}

	// 二维码相关
	codeGroup := cc.Group("/qrcode")
	{
		codeGroup.GET("/", controller.GetCodeImageHandler)    // 产生二维码
		codeGroup.GET("/info", controller.GetCodeInfoHandler) // 获取二维码信息---是一个url，就是从那里获取信息
	}

	// 微信相关
	wechatGroup := cc.Group("/wechat")
	{
		wechatGroup.GET("/accesstokne_by_code", controller.WechatAuthHandler)                    // code 换access_token或openid
		wechatGroup.GET("/download_media_by_audioid", controller.WechatDownloadMediaDataHandler) // 根据media_id下载音频数据
		wechatGroup.POST("/send_template_info", controller.WechatSendTemplateInfoHandler)        // 发送模板消息
		wechatGroup.GET("/jsconfig", controller.WechatGetJSConfigHandler)                        // 获取jsconfig
		wechatGroup.POST("/receive_msg", controller.WechatReceiveMsgHandler)                     // 接收各种类型的消息
	}

	// 权限相关
	powerGroup := cc.Group("/power")
	{
		powerGroup.GET("/list:uid", controller.GetUserPowerListHandler) // 列举某个用户的权限
	}

	err := r.Run(fmt.Sprintf(":%d", config.Config.Cfg.Port))
	if err != nil {
		panic(err)
	}
}
