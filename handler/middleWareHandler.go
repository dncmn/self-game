package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"self_game/config"
	"self_game/utils/vo"
)

var (
	GameToken      = "token"
	GameTokenValue = config.Config.Cfg.Token
)

// 验证token
func VerifyToken(c *gin.Context) {
	if c.Request.Header.Get(GameToken) != GameTokenValue {
		c.JSON(http.StatusBadRequest, gin.H{"error": "token error"})
		log.Println("token error")
		c.Abort()
		return
	}
	c.Next()
}

// 发送响应体
func SendResponse(c *gin.Context, retData *vo.Data) {
	resp, err := json.Marshal(retData)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, retData)
		return
	}

	fmt.Println("reqURL=", c.Request.URL, ", responseBody:", string(resp))
	c.AbortWithStatusJSON(http.StatusOK, retData)
	return
}
