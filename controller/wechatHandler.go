package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/chanxuehong/wechat.v2/mp/user"
	"self-game/constants/gameCode"
	"self-game/service"
	"self-game/utils"
	"self-game/utils/async"
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
	mp3Path, err = service.WechatDownAudioByMediaID(mediaID)
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

/*
	获取认证的时候,注意参数的顺序。如果顺序不对，即使程序没有报错,前端也会无法进行录音操作
*/
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
	var (
		body      interface{}
		content   = make([]byte, 0)
		finalBody service.XMLReq
		err       error
		userInfo  *user.UserInfo
	)

	// 获取xml中的请求体内容
	if content, err = ParsePostXMLBody(c, &body); err != nil {
		retData.Code = gameCode.RequestParamsError
		retData.Message = err.Error()
		return
	}
	// xml bytes to struct
	err = utils.XmlByteToStruct(content, &finalBody)
	if err != nil {
		retData.Code = gameCode.RequestParamsError
		retData.Message = err.Error()
		logger.Error(err)
		return
	}
	logger.Infof("finalBody=%v", finalBody.FromUserName)

	// 获取发送者的消息
	userInfo, err = service.WechatGetUserInfoByOpenID(finalBody.FromUserName)
	if err != nil {
		retData.Code = gameCode.RequestParamsError
		retData.Message = err.Error()
		logger.Error(err)
		return
	}

	// 对消息类型判断
	cnt := "文本消息"
	switch true {
	case finalBody.MsgType == "text":
		if finalBody.FromUserName == "oTVNt1dPSf0U7PLI0AytXfhZad0M" {
			cnt = "我是大咪咪"
		} else {
			cnt = "文本消息:原来你说的是" + finalBody.Content
		}
	case finalBody.MsgType == "image":
		cnt = "图片消息"
	case finalBody.MsgType == "voice":
		cnt = fmt.Sprint("声音消息:刚刚你说的是", finalBody.Recognition)
	case finalBody.MsgType == "video":
		cnt = "小视频消息"
	case finalBody.MsgType == "event": // 关注、取消关注事件      当用户关注公众号的时候,会进行一系列操作。然后取消公众号的时候,然后再进行一些删除数据的操作
		if finalBody.Event == "subscribe" { // 关注公众号事件:主要进行一些数据的初始化操作
			cnt = "Welcome to here!"
		} else { // 取消关注:感觉是进行一些数据数据的删除，状态的修改操作。
			cnt = "see you later!"
		}
	default:
		cnt = "未识别类型"
	}

	cnt = fmt.Sprint(userInfo.Nickname, " send msg:content:", cnt)

	// 记录发送消息的日志
	go async.Do(func() {
		err = service.WechatLogUserSendMstToWechat(finalBody)
		if err != nil {
			logger.Error(err)
			return
		}
		logger.Infof("log msg success:openid=%s,msg_type=%s,msg_send_tme=%v",
			finalBody.FromUserName, finalBody.MsgType, finalBody.CreateTime)
	})
	xmlStr := fmt.Sprintf("<xml><ToUserName><![CDATA[%s]]></ToUserName><FromUserName><![CDATA[%s]]></FromUserName><CreateTime>%v</CreateTime><MsgType><![CDATA[text]]></MsgType><Content><![CDATA[%s]]></Content></xml><MsgId>%s</MsgId>",
		finalBody.FromUserName, finalBody.ToUserName, finalBody.CreateTime, cnt, finalBody.MsgId)
	if finalBody.MsgType == "image" { // 简单的设置
		xmlStr = fmt.Sprintf("<xml><ToUserName><![CDATA[%s]]></ToUserName><FromUserName><![CDATA[%s]]></FromUserName><CreateTime>%v</CreateTime><MsgType><![CDATA[image]]></MsgType><Image><MediaId><![CDATA[%s]]></MediaId></Image></xml>",
			finalBody.FromUserName, finalBody.ToUserName, finalBody.CreateTime, finalBody.MediaId)
	}
	c.Data(200, "", []byte(xmlStr))
	return
}
