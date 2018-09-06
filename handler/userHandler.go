package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"self_game/config"
	"self_game/service"
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
	fmt.Println("hello")

	c.JSON(http.StatusOK, gin.H{
		"name": config.Config.Cfg.Port,
		"env":  config.Config.Env.ENV,
	})
}

func ConsulHealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func PostUserNameHandler(c *gin.Context) {
	name := c.Param("uid")

	type Res struct {
		Name string `json:"name"`
	}

	var res Res

	err := c.BindJSON(&res)
	if err != nil {
		c.JSON(http.StatusAccepted, gin.H{"error": err})
		return
	}
	fmt.Println(name)
	fmt.Println(name)
	fmt.Println(name)
	fmt.Println(name)
	fmt.Println(name)

	c.JSON(http.StatusOK, gin.H{"name": res.Name})
}
