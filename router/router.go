package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"self_game/config"
	_ "self_game/dao"
	"self_game/handler"
)

func Router(r *gin.Engine) {

	cc := r.Group("/api/v1", handler.VerifyToken)
	{
		cc.GET("/", handler.HandlerSignatureHandler)

	}

	userLoginGroup := cc.Group("/anonymous")
	{
		userLoginGroup.POST("/user/register", handler.RegisterUserHandler)
		userLoginGroup.POST("/user/login", handler.UserLoginHandler)
	}

	user := cc.Group("/user", handler.VerifyUserToken)
	{
		user.GET("/name/:uid", handler.GetUserNameHandler)
		user.POST("/name", handler.PostUserNameHandler)
		user.GET("/health_check", handler.ConsulHealthCheck)
	}
	r.Run(fmt.Sprintf(":%d", config.Config.Cfg.Port))
}
