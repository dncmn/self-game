package router

import (
	"github.com/gin-gonic/gin"
	"self_game/handler"
)

func Router(r *gin.Engine) {

	user := r.Group("/user")
	r.Use()
	user.GET("/name", handler.GetUserNameHandler)
	user.POST("/name/:uid", handler.PostUserNameHandler)
	user.GET("/health_check", handler.ConsulHealthCheck)
	r.Run()

}
