package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"self_game/config"
	"self_game/utils/logging"
	"self_game/utils/vo"
)

var (
	GameToken      = "token"
	GameTokenValue = config.Config.Cfg.Token
	logger         = logging.GetLogger()
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

	logger.Infof("reqURL=%s,responseBody=%v", c.Request.URL, string(resp))
	c.AbortWithStatusJSON(http.StatusOK, retData)
	return
}

// post请求获取请求参数
func ParsePostBody(c *gin.Context, resp interface{}) (err error) {
	// 从请求体中获取请求的数据
	rqt, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Infof("[Request Url Body] req:%s url:%s  body:%s", c.GetString("reqID"),
		c.Request.RequestURI, string(rqt))
	// 将请求数据绑定到指定的结构体中
	err = json.Unmarshal(rqt, resp)
	if err != nil {
		logger.Error(err)
	}
	logger.Info(err)
	return
}
