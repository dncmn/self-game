package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"self_game/config"
	"self_game/handler"
)

func Router(r *gin.Engine) {

	cc := r.Group("/")
	{
		cc.GET("/", handler.HandlerSignatureHandler)
	}
	user := r.Group("/user", handler.VerifyToken)
	{
		user.GET("/name", handler.GetUserNameHandler)
		user.POST("/name/:uid", handler.PostUserNameHandler)
		user.GET("/health_check", handler.ConsulHealthCheck)
	}
	r.Run(fmt.Sprintf(":%d", config.Config.Cfg.Port))

}
