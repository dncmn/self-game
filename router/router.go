package router

import (
	"github.com/gin-gonic/gin"
	"self_game/handler"
	"fmt"
	"self_game/config"
)

func Router(r *gin.Engine) {

	user := r.Group("/user",handler.VerifyToken)
	{
		user.GET("/name", handler.GetUserNameHandler)
		user.POST("/name/:uid", handler.PostUserNameHandler)
		user.GET("/health_check", handler.ConsulHealthCheck)
	}
	r.Run(fmt.Sprintf(":%d",config.Config.Cfg.Port))

}
