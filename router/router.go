package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"self-game/config"
	_ "self-game/dao"
	"self-game/handler"
)

func Router(r *gin.Engine) {

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

	// 权限相关
	powerGroup := cc.Group("/power")
	{
		powerGroup.GET("/list:uid")
	}

	r.Run(fmt.Sprintf(":%d", config.Config.Cfg.Port))
}
