package utils

import (
	"code.dncmn.io/self-game/utils/snapHttp"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func translateUnicodeToChinese(src string) (context string, err error) {
	sUnicodeDev := strings.Split(src, "\\u")
	for _, v := range sUnicodeDev {
		if len(v) < 1 {
			continue
		}
		temp, err1 := strconv.ParseInt(v, 16, 32)
		if err1 != nil {
			err = err1
			fmt.Println(err)
			return
		}
		context += fmt.Sprintf("%c", temp)
	}
	return
}

func TranslaTate(text string) (resp BaiDuResp, err error) {
	salt := IntRange(1, 100000000)
	appid := "20181027000226126"
	secret := "vyciEmBE6MGscQHHL95R"
	total := fmt.Sprint(appid, text, salt, secret)
	sign := EncodeMD5(total)

	tl := fmt.Sprintf("http://api.fanyi.baidu.com/api/trans/vip/translate?"+
		"q=%s&from=en&to=zh&appid=20181027000226126&salt=%v&sign=%v", text, salt, sign)

	sh := snapHttp.SnapHttp{}
	err = sh.GetJson(tl, &resp)
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}

func IntRange(min, max int) int {
	if min >= max {
		return min
	}
	rand.Seed(time.Now().UnixNano()) // 设置随机种子
	v := rand.Int63n(int64(max - min))
	return min + int(v)
}

type BaiDuResp struct {
	From        string `json:"from"`
	To          string `json:"to"`
	TransResult []struct {
		Dst string `json:"dst"`
		Src string `json:"src"`
	} `json:"trans_result"`
}
