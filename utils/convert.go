package utils

import (
	"code.dncmn.io/self-game/data/constants"
	"code.dncmn.io/self-game/utils/bash"
	"code.dncmn.io/self-game/utils/logging"
	"fmt"
	"io/ioutil"
	"path"
	"strings"
)

var logs = logging.GetLogger()

var (
	fpath = "ffmpeg "
)

// audio to mp3
const (
	TAmrToMp3 = "%s -y -i %s -ar 16000 %s"
	TPcmToMp3 = "%s -y -f s16le -ac 1 -i %s -ar 16000 %s"
	TWavToMp3 = "%s -y -f s16le -ac 1 -i %s -ar 16000 %s"
)

func audioBytesToMp3(filePath, mp3Path string, bytes []byte) (err error) {
	if err = ioutil.WriteFile(filePath, bytes, 0666); err != nil {
		return
	}

	// 写入文件
	ext := path.Ext(filePath)
	switch {
	case ext == ".amr":
		cmdStr := fmt.Sprintf(TAmrToMp3,
			fpath, filePath, mp3Path)
		_, err = bash.Bash(cmdStr)
	case ext == ".pcm":
		cmdStr := fmt.Sprintf(TPcmToMp3,
			fpath, filePath, mp3Path)
		_, err = bash.Bash(cmdStr)
	case ext == ".wav":
		cmdStr := fmt.Sprintf(TWavToMp3,
			fpath, filePath, mp3Path)
		_, err = bash.Bash(cmdStr)
	case ext == ".mp3":
		cmdStr := fmt.Sprintf(TPcmToMp3, fpath, filePath, mp3Path)
		_, err = bash.Bash(cmdStr)
	}

	if err != nil {
		logs.Error(err)
		return
	}
	//if err = os.Remove(filePath); err != nil {
	//	logs.Error(err)
	//	return
	//}
	return
}
func AudioBytesToMp3(fileName string, bytes []byte) (mp3Path, mp3Name string, err error) {
	filePath := path.Join(constants.WechatDownloadAmrLocalAddr, path.Base(fileName))
	mp3Path = strings.Replace(filePath, path.Ext(filePath), ".mp3", -1)
	mp3Name = path.Base(strings.Replace(fileName, path.Ext(fileName), ".mp3", -1))
	err = audioBytesToMp3(filePath, mp3Path, bytes)
	return
}
