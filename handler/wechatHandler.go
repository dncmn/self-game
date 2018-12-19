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
		retData.Code = gameCode.ErrorCodeWechatCodeToAccessToken
		retData.Message = "code to accessToken error"
		logger.Error(err)
		return
	}
	retData.Data = resp
	return
}

// 根据audioID下载音频
func WechatDownloadMediaDataHandler(c *gin.Context) {
	retData := vo.NewData()
	defer SendResponse(c, retData)
	var (
		err     error
		mediaID string
		mp3Path string
		resp    interface{}
	)
	if mediaID = c.Query("media_id"); utils.IsStringEmpty(mediaID) {
		retData.Code = gameCode.RequestParamsError
		logger.Error(errors.New("param error"))
		return
	}
	mp3Path, err = service.WechatDownAudioByAudioID(mediaID)
	if err != nil {
		retData.Code = gameCode.ErrorCodeWechatDownloadResourceByAudioID
		logger.Error(errors.New("param error"))
		return
	}
	resp, err = service.WechatUploadAudioToOSS(mp3Path)
	if err != nil {
		retData.Code = gameCode.RequestParamsError
		logger.Error(errors.New("upload to oss error"))
		return
	}

	retData.Code = gameCode.RequestSuccess
	retData.Data = resp
	retData.Message = "request success"
	return
}

func WechatSendTemplateInfoHandler(c *gin.Context) {
	retData := vo.NewData()
	defer SendResponse(c, retData)
	var (
		body service.SendTemplateRes
		resp interface{}
		err  error
	)
	if err = ParsePostBody(c, &body); err != nil {
		retData.Code = gameCode.RequestParamsError
		logger.Error("param error")
		return
	}
	if err = service.WechatSendTemplateInfo(body); err != nil {
		retData.Code = gameCode.ErrorCodeWechatSendTemplateInfo
		retData.Message = err.Error()
		logger.Error(err)
		return
	}

	// 测试根据openID获取用户信息
	resp, err = service.WechatGetUserInfoByOpenID(body.OpenID)
	if err != nil {
		retData.Code = gameCode.RequestParamsError
		retData.Message = err.Error()
		logger.Error(err)
		return
	}
	retData.Code = gameCode.RequestSuccess
	retData.Data = resp
	logger.Infof("openid=%v,userInfo=%v", body.OpenID, resp)
	return
}

func WechatGetJSConfigHandler(c *gin.Context) {
	retData := vo.NewData()
	defer SendResponse(c, retData)
	var (
		baseURL string
		resp    interface{}
		err     error
	)
	if baseURL = c.Query("baseURL"); utils.IsStringEmpty(baseURL) {
		retData.Code = gameCode.RequestParamsError
		retData.Message = "param error"
		logger.Error(err)
		return
	}

	if resp, err = service.WechatGetJSConfig(baseURL); err != nil {
		retData.Code = gameCode.ErrorCodeWechatGetJSconfig
		logger.Error(err)
		retData.Message = err.Error()
		return
	}
	retData.Data = resp
	retData.Code = gameCode.RequestSuccess
	logger.Infof("get jsconfig signature.baseURL=%v,signatureInfo=%v", baseURL, resp)
	return
}

func WechatReceiveMsgHandler(c *gin.Context) {
	retData := vo.NewData()
	defer SendResponse(c, retData)
	var (
		body service.ReceiveMsgReq
		err  error
	)
	if err = ParsePostBody(c, &body); err != nil {
		retData.Code = gameCode.RequestParamsError
		return
	}

	switch body.MsgType {
	case "text": // 文本类型的消息
	case "image": // 图片类型的消息
	case "voice": // 声音类型的消息
	default:
		logger.Error(errors.New("undefined message type"))
		retData.Code = gameCode.RequestParamsError
		return
	}

	retData.Code = gameCode.RequestSuccess
	retData.Data = "request success"
	logger.Info(body)
	return
}

func HandMessagesHandler(c *gin.Context) {
	retData := vo.NewData()
	defer SendResponse(c, retData)
	var (
		body       interface{}
		content    = make([]byte, 0)
		convertMap = make(map[string]string)
		finalBody  service.XMLReq
		err        error
	)

	// 获取xml中的请求体内容
	if content, err = ParsePostXMLBody(c, &body); err != nil {
		retData.Code = gameCode.RequestParamsError
		retData.Message = err.Error()
		return
	}
	// 将xml转换为map
	convertMap, err = utils.XmlToMap(content)
	if err != nil {
		retData.Code = gameCode.RequestParamsError
		retData.Message = err.Error()
		return
	}
	// map to struct
	err = utils.StructToMap(convertMap, &finalBody)
	if err != nil {
		retData.Code = gameCode.RequestParamsError
		retData.Message = err.Error()
		return
	}

	retData.Code = gameCode.RequestSuccess
	retData.Message = "request success"
	retData.Data = finalBody
	logger.Infof("finalBody=%s", finalBody)
	return
}
