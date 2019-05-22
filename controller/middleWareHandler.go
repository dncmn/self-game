package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"self-game/config"
	"self-game/constants/gameCode"
	"self-game/service"
	"self-game/utils"
	"self-game/utils/logging"
	"self-game/utils/vo"
)

var (
	GameToken           = "token"
	UserToken           = "UserToken"
	GameTokenValue      = config.Config.Cfg.Token
	logger              = logging.GetLogger()
	ContextKeyAppUID    = "appuid"
	ContextKeyUserToken = "userToken"
)

// 验证userToken
func VerifyUserToken(c *gin.Context) {
	var (
		userToken string
		uid       string
		err       error
		retData   = vo.NewData()
	)

	if userToken = c.Request.Header.Get(UserToken); utils.IsStringEmpty(userToken) {
		retData.Code = gameCode.RequestTokenError
		c.JSON(http.StatusBadRequest, retData)
		logger.Error("token error")
		c.Abort()
		return
	}

	uid, err = service.GetUIDByUserToken(userToken)
	if err != nil {
		retData.Code = gameCode.RequestTokenError
		c.JSON(http.StatusBadRequest, retData)
		logger.Error("token error")
		c.Abort()
		return
	}
	logger.Infof("set uid %v", uid)
	c.Set(ContextKeyAppUID, uid)
	c.Set(ContextKeyUserToken, userToken)
	c.Next()
}

// 验证token
func VerifyToken(c *gin.Context) {
	retData := vo.NewData()
	if c.Request.Header.Get(GameToken) != GameTokenValue {
		retData.Code = gameCode.RequestTokenError
		c.JSON(http.StatusBadRequest, retData)
		logger.Error("token error")
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

	logger.Infof("response:reqURL=%s,responseBody=%v", c.Request.URL, string(resp))
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
	return
}

// post请求获取请求参数
func ParsePostXMLBody(c *gin.Context, resp interface{}) (content []byte, err error) {
	// 从请求体中获取请求的数据
	content, err = ioutil.ReadAll(c.Request.Body)
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Infof("[Request Url Body] req:%s url:%s  body:%s", c.GetString("reqID"),
		c.Request.RequestURI, string(content))
	// 将请求数据绑定到指定的结构体中
	return
}
