package main

import (
	_ "code.dncmn.io/self-game/cmd/user"
	"code.dncmn.io/self-game/router"
	"code.dncmn.io/self-game/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
)

func main() {
	c := make(chan os.Signal)
	go func() {
		signal.Notify(c)
		fmt.Println("begin")
		router.Router(gin.Default())
	}()
	<-c                    // 当关闭服务器的时候,就进行一些释放资源的操作
	service.IsStop = false // 信号量,表示是否跳出这个循环,然后进行一系列的资源释放操作
}
