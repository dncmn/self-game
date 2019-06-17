package utils

import (
	"code.dncmn.io/self-game/utils/logging"
	"fmt"
	"github.com/liviosoares/go-watson-sdk/watson"
	"github.com/liviosoares/go-watson-sdk/watson/text_to_speech"
	"os"
)

var (
	logger = logging.GetLogger()
)

func getTextToSpeechClient() (client text_to_speech.Client, err error) {
	config := watson.Config{
		Credentials: watson.Credentials{
			Username: "05e4ae2e-a288-4c0b-a45b-1c411322d0f5",
			Password: "tpukwaROXtZA",
		},
	}

	client, err = text_to_speech.NewClient(config)
	if err != nil {
		logger.Error(err)
		return
	}
	return
}

// isSlow:表示是否是慢速
// dir_path:文件的路径  /homeworkTest/images/L1/LessonExerciseL1U1C1/Are you ok?.mp3
// 每个月有1万个免费字符
func TextToNormalSpeech(text, dir_path string, isSlow bool) (err error) {
	client, err := getTextToSpeechClient()
	if err != nil {
		logger.Error(err)
		return
	}

	if isSlow {
		text = fmt.Sprintf("<voice-transformation type=\"Custom\" rate=\"x-slow\">%s</voice-transformation>", text)
	}

	data, err := client.Synthesize(text, "en-US_AllisonVoice", "audio/mp3", "")
	if err != nil {
		logger.Error(err)
		return
	}

	// write data to file
	out, err := os.Create(fmt.Sprint(dir_path, ".mp3"))
	if err != nil {
		logger.Error(err)
		return
	}
	defer out.Close()
	_, err = out.Write(data)
	if err != nil {
		logger.Error(err)
		return
	}
	return
}
