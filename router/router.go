package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"self_game/config"
	_ "self_game/dao"
	"self_game/handler"
)

func Router(r *gin.Engine) {

	cc := r.Group("/api/v1")
	{
		cc.GET("/", handler.HandlerSignatureHandler)
	}
	user := cc.Group("/user", handler.VerifyToken)
	{
		user.GET("/name/:uid", handler.GetUserNameHandler)
		user.POST("/name", handler.PostUserNameHandler)
		user.GET("/health_check", handler.ConsulHealthCheck)
	}
	r.Run(fmt.Sprintf(":%d", config.Config.Cfg.Port))
}
