package qrcode

import (
	"image/jpeg"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"

	"code.dncmn.io/self-game/config"
	"code.dncmn.io/self-game/utils"
	"code.dncmn.io/self-game/utils/logging"
	"github.com/tuotoo/qrcode"
	"os"
)

var (
	logger = logging.GetLogger()
)

type QrCode struct {
	URL    string
	Width  int
	Height int
	Ext    string
	Level  qr.ErrorCorrectionLevel
	Mode   qr.Encoding
}

const (
	EXT_JPG = ".jpg"
)

func NewQrCode(url string, width, height int, level qr.ErrorCorrectionLevel, mode qr.Encoding) *QrCode {
	return &QrCode{
		URL:    url,
		Width:  width,
		Height: height,
		Level:  level,
		Mode:   mode,
		Ext:    EXT_JPG,
	}
}

func GetQrCodePath() string {
	return config.Config.Code.QrCodeSavePath
}

func GetQrCodeFullPath() string {
	return config.Config.Code.RuntimeRootPath + config.Config.Code.QrCodeSavePath
}

func GetQrCodeFullUrl(name string) string {
	return config.Config.Code.PrefixUrl + "/" + GetQrCodePath() + name
}

func GetQrCodeFileName(value string) string {
	return utils.EncodeMD5(value)
}

func (q *QrCode) GetQrCodeExt() string {
	return q.Ext
}

func (q *QrCode) CheckEncode(path string) bool {
	src := path + GetQrCodeFileName(q.URL) + q.GetQrCodeExt()
	if utils.CheckNotExist(src) == true {
		return false
	}

	return true
}

func (q *QrCode) Encode(path string) (string, string, error) {
	name := GetQrCodeFileName(q.URL) + q.GetQrCodeExt()
	src := path + name
	logger.Infof("Encode:name=%v,src=%v", name, src)
	if utils.CheckNotExist(src) == true {
		code, err := qr.Encode(q.URL, q.Level, q.Mode)
		if err != nil {
			return "", "", err
		}

		code, err = barcode.Scale(code, q.Width, q.Height)
		if err != nil {
			return "", "", err
		}

		f, err := utils.MustOpen(name, path)
		if err != nil {
			return "", "", err
		}
		defer f.Close()

		err = jpeg.Encode(f, code, nil)
		if err != nil {
			return "", "", err
		}
	}
	return name, path, nil
}

func GetCodeInfo(filePath string) (content string, err error) {
	var (
		fi *os.File
	)

	fi, err = os.Open(filePath)
	if err != nil {
		logger.Error(err)
		return
	}
	defer fi.Close()
	qm, err := qrcode.Decode(fi)
	if err != nil {
		logger.Error(err)
		return
	}
	content = qm.Content
	return
}
