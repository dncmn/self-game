package handler

import (
	"github.com/gin-gonic/gin"
	"self_game/config"
	"log"
	"net/http"
)


var(
	GameToken ="token"
	GameTokenValue=config.Config.Cfg.Token
)


// 验证token
func VerifyToken(c *gin.Context){
	if c.Request.Header.Get(GameToken)!=GameTokenValue{
		c.JSON(http.StatusBadRequest,gin.H{"error":"token error"})
		log.Println("token error")
		c.Abort()
		return
	}
	c.Next()
}
