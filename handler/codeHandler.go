package handler

import (
	"github.com/boombuler/barcode/qr"
	"github.com/gin-gonic/gin"
	"self-game/utils/qrcode"
	"self-game/utils/vo"
)

func GetCodeImageHandler(c *gin.Context) {
	retData := vo.NewData()
	defer SendResponse(c, retData)
	qrc := qrcode.NewQrCode("http://www.baidu.com", 300, 300, qr.M, qr.Auto)
	ph := qrcode.GetQrCodeFullPath()
	_, _, err := qrc.Encode(ph)
	if err != nil {
		retData.Code = -100
		return
	}

	logger.Infof("qrc=%v,path=%v", qrc, ph)
	retData.Data = "/Users/mn/go/src/self-game/compoments/images/hello.jpeg"
	return
}
