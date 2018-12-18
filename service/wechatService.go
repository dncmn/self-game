package service

import (
	"crypto/tls"
	"encoding/json"
	"gopkg.in/chanxuehong/wechat.v2/mp/jssdk"
	"gopkg.in/chanxuehong/wechat.v2/mp/message/template"
	"gopkg.in/chanxuehong/wechat.v2/mp/user"
	"strconv"

	"gopkg.in/chanxuehong/wechat.v2/mp/core"
	"gopkg.in/chanxuehong/wechat.v2/mp/media"
	mp "gopkg.in/chanxuehong/wechat.v2/mp/oauth2"
	"gopkg.in/chanxuehong/wechat.v2/oauth2"
	"io/ioutil"
	"net/http"
	"path"
	"self-game/config"
	"self-game/constants"
	"self-game/constants/redisKey"
	"self-game/utils"
	"self-game/utils/async"
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
	go async.Do(func() {
		initService()
	})
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

// 下载音频数据
func WechatDownAudioByAudioID(audioID string) (mp3Path string, err error) {
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
	osskey := path.Join("self-game", strconv.FormatInt(utils.GetTimeZoneTime(config.Config.Cfg.TimeZone).UnixNano(), 10), mp3Path, mp3Name)
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
	AppId      string
	Ticket     string
	NonceStr   string
	TimeTagStr string
	Url        string
	Signature  string
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
