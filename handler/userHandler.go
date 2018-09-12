package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"self_game/config"
	"self_game/service"
	"self_game/utils/vo"
	"strings"
)

func HandlerSignatureHandler(c *gin.Context) {
	var (
		signature string
		echostr   string
		timestamp int
		nonce     int
		err       error
	)

	if signature, echostr, timestamp, nonce, err = service.GetSignatrueParams(c); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(signature, echostr, timestamp, nonce)

	c.JSON(http.StatusOK, gin.H{
		"ok": true,
	})

}

func GetUserNameHandler(c *gin.Context) {
	retData := &vo.Data{}
	defer SendResponse(c, retData)
	fmt.Println("hello")
	var (
		uid string
	)
	uid = c.Param("uid")
	if strings.TrimSpace(uid) == "12345" {
		retData.Code = -101
		retData.Data = "param error"
		return
	}

	retData.Data = map[string]interface{}{
		"name": config.Config.Cfg.Port,
		"env":  config.Config.Env.ENV,
	}
	retData.Code = 1
	return
}

func ConsulHealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func PostUserNameHandler(c *gin.Context) {
	retData := vo.NewData()
	defer SendResponse(c, retData)

	var (
		res  service.PostUserRequest
		resp service.PostUserResponse
		err  error
		uid  string
	)

	if err = ParsePostBody(c, &res); err != nil {
		retData.Data = err.Error()
		retData.Code = -101
		return
	}
	fmt.Println("name=", res.Name, "english_score=", res.EnglishScore)

	if strings.TrimSpace(res.Name) == "" {
		retData.Data = "params error"
		retData.Code = -101
		fmt.Println("dadadadfa")
		fmt.Println("dadadadfa")
		fmt.Println("dadadadfa")
		fmt.Println("dadadadfa")
		fmt.Println("dadadadfa")
		fmt.Println("dadadadfa")
		return
	}

	resp.UID = uid
	resp.Name = res.Name
	resp.ChineseScore = res.EnglishScore + 2
	resp.EnglishScore = res.EnglishScore
	retData.Code = 1
	retData.Data = map[string]interface{}{
		"user_info": resp,
	}
	return
}
