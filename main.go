package main

import (
	"github.com/gin-gonic/gin"
	"self-game/router"
)

func main() {
	router.Router(gin.Default())
}
