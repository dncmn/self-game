package controller

import (
	"code.dncmn.io/self-game/utils/vo"
	"github.com/gin-gonic/gin"
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
