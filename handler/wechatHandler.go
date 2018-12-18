package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"self-game/constants/gameCode"
	"self-game/service"
	"self-game/utils"
	"self-game/utils/vo"
)

// code exchange accessToken
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

// 根据audioID下载音频
func WechatDownloadMediaDataHandler(c *gin.Context){
	retData := vo.NewData()
	defer SendResponse(c, retData)
	var(
		err error
		mediaID string
		mp3Path string
		resp interface{}
	)
	if mediaID=c.Query("media_id");utils.IsStringEmpty(mediaID){
		retData.Code=gameCode.RequestParamsError
		logger.Error(errors.New("param error"))
		return
	}
	mp3Path,err=service.WechatDownAudioByAudioID(mediaID)
	if err!=nil{
		retData.Code=gameCode.RequestParamsError
		logger.Error(errors.New("param error"))
		return
	}
	resp,err=service.WechatUploadAudioToOSS(mp3Path)
	if err!=nil{
		retData.Code=gameCode.RequestParamsError
		logger.Error(errors.New("upload to oss error"))
		return
	}

	retData.Code=gameCode.RequestSuccess
	retData.Data=resp
	retData.Message="request success"
	return
}

func WechatSendTemplateInfoHandler(c *gin.Context){
	retData:=vo.NewData()
	defer SendResponse(c, retData)
	var(
		body service.SendTemplateRes
		resp interface{}
		err error
	)
	if err=ParsePostBody(c,&body);err!=nil{
		retData.Code=gameCode.RequestParamsError
		logger.Error("param error")
		return
	}
	if err=service.WechatSendTemplateInfo(body);err!=nil{
		retData.Code=gameCode.RequestParamsError
		retData.Message=err.Error()
		logger.Error(err)
		return
	}

	// 测试根据openID获取用户信息
	resp,err=service.WechatGetUserInfoByOpenID(body.OpenID)
	if err!=nil{
		retData.Code=gameCode.RequestParamsError
		retData.Message=err.Error()
		logger.Error(err)
		return
	}
	retData.Code=gameCode.RequestSuccess
	retData.Data=resp
	logger.Infof("openid=%v,userInfo=%v",body.OpenID,resp)
	return
}