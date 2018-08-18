package main

import (
	"github.com/gin-gonic/gin"
	"self_game/router"
)

func main() {
	router.Router(gin.Default())
}
