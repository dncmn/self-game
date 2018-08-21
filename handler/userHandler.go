package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"self_game/config"
)

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
