package handler

import (
	"github.com/gin-gonic/gin"
	"self-game/utils/vo"
)

func GetUserPowerListHandler(c *gin.Context) {
	retData := vo.NewData()
	defer SendResponse(c, retData)

	var (
		uid       string
		poserList interface{}
	)

	logger.Infof("userPowerList:uid[%v],powerInfo=%v", uid, poserList)
	return
}
