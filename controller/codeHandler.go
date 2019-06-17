package controller

import (
	"code.dncmn.io/self-game/config"
	"code.dncmn.io/self-game/constants/gameCode"
	"code.dncmn.io/self-game/utils"
	"code.dncmn.io/self-game/utils/qrcode"
	"code.dncmn.io/self-game/utils/vo"
	"github.com/boombuler/barcode/qr"
	"github.com/gin-gonic/gin"
	"net/url"
)

func GetCodeImageHandler(c *gin.Context) {
	retData := vo.NewData()
	defer SendResponse(c, retData)

	var (
		err     error
		codeURL string
		name    string
		ph2     string
		rawURL  *url.URL
		uid     string
		age     string
	)

	if codeURL = c.Query("codeURL"); utils.IsStringEmpty(codeURL) {
		retData.Code = gameCode.RequestParamsError
		return
	}

	if uid = c.Query("uid"); utils.IsStringEmpty(uid) {
		retData.Code = gameCode.RequestParamsError
		return
	}
	if age = c.Query("age"); utils.IsStringEmpty(age) {
		retData.Code = gameCode.RequestParamsError
		return
	}

	rawURL, err = url.Parse(codeURL)
	if err != nil {
		retData.Code = -100
		logger.Error(err)
		return
	}
	params := url.Values{}
	params.Add("uid", uid)
	params.Add("age", age)
	params.Add("token", config.Config.Cfg.Token)
	rawURL.RawQuery = params.Encode()

	qrc := qrcode.NewQrCode(rawURL.String(), 300, 300, qr.M, qr.Auto)
	ph := qrcode.GetQrCodeFullPath()
	name, ph2, err = qrc.Encode(ph)
	if err != nil {
		retData.Code = -100
		logger.Error(err)
		return
	}
	logger.Infof("name=%v,ph2=%v", name, ph2)
	logger.Infof("runtimeRootPath=%v,qrCodeSavePath=%v", config.Config.Code.RuntimeRootPath, config.Config.Code.QrCodeSavePath)

	logger.Infof("qrc=%v,path=%v", qrc, ph)

	retData.Data = map[string]interface{}{
		"codeURL": config.Config.Code.PrefixUrl + ph2 + name,
		"destURL": rawURL.String(),
	}
	return
}

func GetCodeInfoHandler(c *gin.Context) {
	retData := vo.NewData()
	defer SendResponse(c, retData)

	var (
		uid string
		age string
	)

	if uid = c.Query("uid"); utils.IsStringEmpty(uid) {
		retData.Code = gameCode.RequestParamsError
		return
	}
	if age = c.Query("age"); utils.IsStringEmpty(age) {
		retData.Code = gameCode.RequestParamsError
		return
	}

	retData.Data = map[string]interface{}{
		"age": age,
		"uid": uid,
	}
	return
}
