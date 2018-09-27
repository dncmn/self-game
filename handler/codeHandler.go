package handler

import (
	"github.com/boombuler/barcode/qr"
	"github.com/gin-gonic/gin"
	"self-game/config"
	"self-game/constants/gameCode"
	"self-game/utils"
	"self-game/utils/qrcode"
	"self-game/utils/vo"
)

func GetCodeImageHandler(c *gin.Context) {
	retData := vo.NewData()
	defer SendResponse(c, retData)

	var (
		err     error
		codeURL string
		name    string
		ph2     string
	)

	if codeURL = c.Query("codeURL"); utils.IsStringEmpty(codeURL) {
		retData.Code = gameCode.RequestParamsError
		return
	}

	qrc := qrcode.NewQrCode(codeURL, 300, 300, qr.M, qr.Auto)
	ph := qrcode.GetQrCodeFullPath()
	name, ph2, err = qrc.Encode(ph)
	if err != nil {
		retData.Code = -100
		return
	}
	logger.Infof("name=%v,ph2=%v", name, ph2)
	logger.Infof("runtimeRootPath=%v,qrCodeSavePath=%v", config.Config.Code.RuntimeRootPath, config.Config.Code.QrCodeSavePath)

	logger.Infof("qrc=%v,path=%v", qrc, ph)
	retData.Data = config.Config.Code.PrefixUrl + ph2 + name
	return
}
