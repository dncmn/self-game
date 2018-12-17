package service

import (
	mp "gopkg.in/chanxuehong/wechat.v2/mp/oauth2"
	"gopkg.in/chanxuehong/wechat.v2/oauth2"
	"self-game/config"
	"self-game/constants/redisKey"
	"self-game/utils"
	"self-game/utils/async"
	"time"
)

var (
	IsStop      = true
	oath2Client = &oauth2.Client{}
	ep          mp.Endpoint
)

func init() {
	go async.Do(func() {
		initService()
	})
	ep = *mp.NewEndpoint(config.Config.Wechat.AppID, config.Config.Wechat.Secret)
	oath2Client.Endpoint = &ep

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
