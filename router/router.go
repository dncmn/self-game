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
		wxServerCheck.GET("/", handler.HandlerSignatureHandler)

	}

	cc := r.Group("/api/v1", handler.VerifyToken)
	userLoginGroup := cc.Group("/anonymous")
	{
		userLoginGroup.POST("/user/register", handler.RegisterUserHandler)
		userLoginGroup.POST("/user/login", handler.UserLoginHandler)
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
		codeGroup.GET("/", handler.GetCodeImageHandler) // 产生二维码
	}

	// 权限相关
	powerGroup := cc.Group("/power")
	{
		powerGroup.GET("/list:uid")
	}

	r.Run(fmt.Sprintf(":%d", config.Config.Cfg.Port))
}
