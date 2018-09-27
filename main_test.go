package main

import (
	"fmt"
	"net/url"
	"self-game/compoments"
	"self-game/config"
	"self-game/constants/redisKey"
	"self-game/model"
	"self-game/utils"
	"self-game/utils/qrcode"
	"self-game/utils/taobaoIP"
	"testing"
	"time"
)

func TestURL(t *testing.T) {
	//rs := url.QueryEscape()
	rs := "http://www.baidu.com/"

	params := url.Values{}
	params.Add("name", "manan")
	pl, err := url.Parse(rs)
	if err != nil {
		t.Error(err)
		return
	}

	//pl.Query() = params
	//fmt.Println(pl.Query())

	fmt.Println(pl.RawQuery)

	//vals, err := url.ParseQuery(rs)
	//if err != nil {
	//	t.Error(err)
	//	return
	//}
	//fmt.Println(vals)
}

func TestTimeZone(t *testing.T) {

	lc, _ := time.LoadLocation("Local")
	fmt.Println(time.Now().In(lc))
	fmt.Print(time.Now())
	fmt.Println()
	l0, _ := time.LoadLocation("UTC")
	fmt.Print(time.Now().In(l0))

	ls, _ := time.LoadLocation("Asia/Shanghai")
	fmt.Println(time.Now().In(ls))
}

func TestReadCongi(t *testing.T) {
	db := compoments.GetDB()
	data := model.LogLogin{}
	data.UID = "test003"
	data.UserName = "name004"
	data.LoginTime = time.Now().Unix()
	err := db.Create(&data).Error
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(data)
}

func TestMobileCheck(t *testing.T) {
	s := []string{"18505921256", "13489594009", "d557"}
	for _, v := range s {
		fmt.Println(utils.CheckMobileIsLegal(v))
	}
}

func TestJiaMi(t *testing.T) {
	str := "helo"
	fmt.Println(utils.EncodeMD5(str))
}

func TestGetCountryAndCity(t *testing.T) {
	ip := "219.142.86.84"
	country, city, err := taobaoIP.GetCountryAndCity(ip)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("country=", country, " city=", city)
}

func TestTimeGet(t *testing.T) {
	n := time.Now()
	fmt.Println(n.Unix(), n.UnixNano()/1e6)

	tl, _ := time.LoadLocation(config.Config.Cfg.TimeZone)
	fmt.Println(time.Now().In(tl).Format("2006-01-02 15:04:05"))

	tm, _ := time.LoadLocation("America/Los_Angeles")
	fmt.Println(time.Now().In(tm).Format("2006-01-02 15:04:05"))
}

func TestLinkRedisServer(t *testing.T) {
	redisCli := compoments.GetRedisClient()
	//err := redisCli.Set(redisKey.UserToken, "hello", 0)
	//if err != nil {
	//	t.Error(err)
	//	return
	//}

	res, ok, err := redisCli.Get(redisKey.UserToken, "936dd200-3ea6-4c0a-9943-1f2fa844470a")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(res, ok)
	fmt.Println("success")
}

func TestGetInfoFromCode(t *testing.T) {
	dir := "./compoments/runtime/qrcode/1e7e59d1df442a0971f9e846610caca0.jpg"
	info, err := qrcode.GetCodeInfo(dir)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(info)
}
