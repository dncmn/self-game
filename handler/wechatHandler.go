package handler

import (
	"github.com/gin-gonic/gin"
	"self-game/constants/gameCode"
	"self-game/service"
	"self-game/utils"
	"self-game/utils/vo"
)

func WechatAuthHandler(c *gin.Context) {
	retData := vo.NewData()
	defer SendResponse(c, retData)
	var (
		err  error
		code string
		resp interface{}
	)
	if code = c.Query("code"); utils.IsStringEmpty(code) {
		retData.Code = gameCode.RequestParamsError
		retData.Message = "param error"
		logger.Error(err)
		return
	}

	resp, err = service.WechatCodeToUserTokenService(code)
	if err != nil {
		retData.Code = gameCode.RequestParamsError
		retData.Message = "code to accessToken error"
		logger.Error(err)
		return
	}
	retData.Data = resp
	return
}
