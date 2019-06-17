package service

import (
	"code.dncmn.io/self-game/model"
	"crypto/tls"
	"encoding/json"
	"gopkg.in/chanxuehong/wechat.v2/mp/jssdk"
	"gopkg.in/chanxuehong/wechat.v2/mp/message/mass/mass2all"
	"gopkg.in/chanxuehong/wechat.v2/mp/message/template"
	"gopkg.in/chanxuehong/wechat.v2/mp/user"
	"sort"
	"strconv"
	"strings"

	"code.dncmn.io/self-game/config"
	"code.dncmn.io/self-game/constants"
	"code.dncmn.io/self-game/constants/redisKey"
	"code.dncmn.io/self-game/utils"
	"gopkg.in/chanxuehong/wechat.v2/mp/core"
	"gopkg.in/chanxuehong/wechat.v2/mp/media"
	mp "gopkg.in/chanxuehong/wechat.v2/mp/oauth2"
	"gopkg.in/chanxuehong/wechat.v2/oauth2"
	"io/ioutil"
	"net/http"
	"path"
	"time"
)

var (
	IsStop       = true
	oath2Client  = &oauth2.Client{}
	coreClient   = &core.Client{}
	ep           mp.Endpoint
	ticketClient jssdk.TicketServer
)

func init() {
	//go async.Do(func() {
	//	initService()
	//})
	ep = *mp.NewEndpoint(config.Config.Wechat.AppID, config.Config.Wechat.Secret)
	oath2Client.Endpoint = &ep

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	basicClient := &http.Client{Transport: tr}
	srv := core.NewDefaultAccessTokenServer(config.Config.Wechat.AppID, config.Config.Wechat.Secret, basicClient)
	coreClient.AccessTokenServer = srv
	coreClient.HttpClient = basicClient
	ticketClient = jssdk.NewDefaultTicketServer(coreClient)
}

func initService() {
	logger.Info("sync info begin")
	for IsStop {
		logger.Infof("current_time=%v", utils.GetTimeZoneTime(config.Config.Cfg.TimeZone).Unix())
		time.Sleep(time.Second * 30)
	}
	logger.Info("sync info end")

}

type WxCodeToTokenResp struct {
	Code        string `json:"code"`
	AccessToken string `json:"access_token"`
	OpenID      string `json:"open_id"`
}

// code 换token
func WechatCodeToUserTokenService(code string) (resp WxCodeToTokenResp, err error) {
	var (
		token *oauth2.Token
	)

	token, err = oath2Client.ExchangeToken(code)
	if err != nil {
		logger.Error(err)
		return
	}

	resp.Code = code
	resp.OpenID = token.OpenId
	resp.AccessToken = token.AccessToken
	err = setUserAccessToken(token)
	if err != nil {
		logger.Error(err)
		return
	}
	return
}

// set userAccessToken and refreshToken to redis
func setUserAccessToken(token *oauth2.Token) (err error) {
	err = redisClient.Set(redisKey.UserAccessToken, token.OpenId, token.AccessToken, time.Duration(token.ExpiresIn)*time.Second)
	if err != nil {
		logger.Error(err)
		return
	}
	err = redisClient.Set(redisKey.UserAccessFreshToken, token.OpenId, token.RefreshToken, time.Duration(token.ExpiresIn)*time.Second)
	return
}

// 下载图片资源
func WechatDownImageByMediaID(mediaID string) (imgPath string, err error) {
	fileName := utils.ReFileName(".jpg")
	imgPath = path.Join(constants.WechatDownloadAmrLocalAddr, fileName)

	_, err = media.Download(coreClient, mediaID, imgPath)
	if err != nil {
		logger.Error(err)
		return
	}
	return
}

// 下载音频数据
func WechatDownAudioByMediaID(audioID string) (mp3Path string, err error) {
	var (
		fileName string
		amrBytes = make([]byte, 0)
		amrPath  string
		mp3Name  string
	)
	fileName = utils.ReFileName(".amr")
	amrPath = path.Join(constants.WechatDownloadAmrLocalAddr, fileName)
	_, err = media.Download(coreClient, audioID, amrPath)
	if err != nil {
		logger.Error(err)
		return
	}
	amrBytes, err = ioutil.ReadFile(amrPath)
	if err != nil {
		logger.Error(err)
		return
	}

	// amr to mp3
	mp3Path, _, err = utils.AudioBytesToMp3(amrPath, amrBytes)
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Infof("audioID=%s,mp3Path=%v,mp3Name=%v\n", audioID, mp3Path, mp3Name)
	return
}

// 上传到oss
func WechatUploadAudioToOSS(mp3Path string) (resource_url string, err error) {
	mp3Name := path.Base(mp3Path)
	osskey := path.Join("code.dncmn.io/self-game", strconv.FormatInt(utils.GetTimeZoneTime(config.Config.Cfg.TimeZone).UnixNano(), 10), mp3Path, mp3Name)
	resource_url, err = utils.PutObject(mp3Path, osskey)
	if err != nil {
		logger.Error(err)
		return
	}
	return
}

// 根据openid获取用户信息
func WechatGetUserInfoByOpenID(openid string) (userInfo *user.UserInfo, err error) {
	userInfo, err = user.Get(coreClient, openid, "")
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Infof("openid=%s,username=%v,country=%v,userinfo=%v",
		userInfo.OpenId, userInfo.Nickname, userInfo.Country, userInfo)
	return
}

type SendTemplateRes struct {
	OpenID      string      `json:"openID"`
	TempleteID  string      `json:"templeteId"`
	KeyWordData interface{} `json:"keyWordData"`
	ActionURL   string      `json:"actionURL"`
}

// 发送模板消息
func WechatSendTemplateInfo(body SendTemplateRes) (err error) {
	marshData, err := json.Marshal(body.KeyWordData)
	if err != nil {
		logger.Errorf("marshal templdate data error %v", err)
		return
	}
	res := &template.TemplateMessage{}
	res.ToUser = body.OpenID
	res.TemplateId = body.TempleteID
	res.URL = body.ActionURL
	res.Data = marshData

	_, err = template.Send(coreClient, res)
	if err != nil {
		logger.Error(err)
		return
	}
	return
}

type SignatureResp struct {
	AppId      string `json:"app_id"`
	Ticket     string `json:"ticket"`
	NonceStr   string `json:"nonce_str"`
	TimeTagStr string `json:"time_tag_str"`
	Url        string `json:"url"`
	Signature  string `json:"signature"`
}

// get jsconfig
func WechatGetJSConfig(baseURL string) (resp SignatureResp, err error) {
	var ticket string
	ticket, err = ticketClient.Ticket()
	if err != nil {
		return
	}
	v := utils.IntRange(100000000000000, 999999999999999)
	nonceStr := strconv.FormatInt(int64(v), 10)

	timeTag := time.Now().UnixNano()
	timeTagStr := strconv.FormatInt(timeTag, 10)
	resp.Signature = jssdk.WXConfigSign(ticket, nonceStr, timeTagStr, baseURL)
	resp.AppId = config.Config.Wechat.AppID
	resp.NonceStr = nonceStr
	resp.Ticket = ticket
	resp.TimeTagStr = timeTagStr
	return
}

// 配置wechet服务器的时候做的验证
func WechatCheckServer(timestamp, nonce, signature string) (success bool) {
	list := []string{
		config.Config.Wechat.Token, timestamp, nonce,
	}
	sort.Strings(list)
	totalStr := strings.Join(list, "")
	encodeTotalStr := utils.EncodeSha1(totalStr)
	logger.Infof("totalStr=%s,md5Str=%s,signature=%s", totalStr, encodeTotalStr, signature)

	if encodeTotalStr == signature {
		return true
	}
	return
}

// 微信公众号内接收消息的请求体
type ReceiveMsgReq struct {
	URL          string `json:"URL"`
	ToUserName   string `json:"ToUserName"`        // 消息的接收者
	FromUserName string `json:"FromUserName"`      // 消息的发送者
	CreateTime   int    `json:"CreateTime"`        // 消息的创建时间
	MsgType      string `json:"MsgType"`           // 消息类型 // text,image,voice
	MsgID        int    `json:"MsgId"`             // messageID
	Content      string `json:"Content,omitempty"` // 文本消息：文本消息的内容
	PicUrl       string `json:"PicUrl,omitempty"`  // 图片消息：图片的url
	MediaId      string `json:"MedisId,omitempty"` // 语音消息: mediaID
	Format       string `json:"Format,omitempty"`  // 语音消息: 音频格式化类型mp3/wav/amr等
}

// 微信群发信息
func WechatMassSendTextMsgByOpenID(message interface{}) (err error) {
	res, err := mass2all.Send(coreClient, message)
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Infof("res.MsgDataId=%v,res.MsgId=%v", res.MsgDataId, res.MsgId)
	return
}

type XMLReq struct {
	XMLName      interface{} `xml:"xml"`
	ToUserName   string      `xml:"ToUserName"`   // 发送者的openID
	FromUserName string      `xml:"FromUserName"` // 公众号的名字
	CreateTime   string      `xml:"CreateTime"`   // 发送消息的时间戳:1545396918
	MsgType      string      `xml:"MsgType"`      // 消息的类型 test/audio/image
	Content      string      `xml:"Content"`      // 文本消息类型的内容
	MsgId        string      `xml:"MsgId"`        // 每个消息唯一的标识id
	PicUrl       string      `xml:"PicUrl"`       // 图片的url
	MediaId      string      `xml:"MediaId"`      // 媒介id(可以根据这个来下载资源)
	Format       string      `xml:"Format"`       // 音频的格式amr格式的
	Recognition  string      `xml:"Recognition"`  // 音频的翻译
	Event        string      `xml:"Event"`        // 关注/取消关注事件
	EventKey     interface{} `xml:"EventKey"`     // 绑定的时候传的参数,不知道这是啥
}

// 记录用户向公众号发送消息的日志
func WechatLogUserSendMstToWechat(body XMLReq) (err error) {
	if body.MsgType != "text" && body.MsgType != "image" {
		return
	}

	l := model.LogUserSendMsgToWechat{
		OpenID:          body.FromUserName,
		MsgType:         body.MsgType,
		CreateTimeStamp: utils.GetTimeZoneTime(config.Config.Cfg.TimeZone).Format(config.Config.Cfg.TimeModelStr),
		MsgID:           body.MsgId,
	}

	if body.MsgType == "text" {
		l.Content = body.Content
	} else {
		l.Content = body.PicUrl
		l.MediaID = body.MediaId
	}
	err = gloDB.Create(&l).Error
	if err != nil {
		logger.Error(err)
		return
	}
	return
}
