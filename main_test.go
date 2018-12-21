package main

import (
	"fmt"
	"gopkg.in/chanxuehong/wechat.v2/mp/message/mass/mass2all"
	"net/url"
	"self-game/compoments"
	"self-game/config"
	"self-game/constants"
	"self-game/constants/redisKey"
	"self-game/model"
	"self-game/service"
	"self-game/utils"
	"self-game/utils/qrcode"
	"self-game/utils/taobaoIP"
	"strconv"
	"testing"
	"time"
)

func TestBaiDuTranslate(t *testing.T) {
	text := "hello"
	resp, err := utils.TranslaTate(text)
	if err != nil {
		t.Error(err)
		return
	}

	for _, v := range resp.TransResult {
		fmt.Println("result=", v.Dst)
	}
}

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

func TestOssPut(t *testing.T) {
	dir_path := "./config"
	err := utils.PutFilesToOSS(dir_path)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("uplode success")
}

// watson ibm  生成慢速音频
func TestTextToSpeech(t *testing.T) {
	text := "helloWorld"
	dir_path := "/Users/mn/Desktop/" + text
	err := utils.TextToNormalSpeech(text, dir_path, true)
	if err != nil {
		t.Error(err)
		t.Error(err)
		return
	}

	fmt.Println("success")
}

// 测试从url上下载资源
func TestDownLoadMP3FromURL(t *testing.T) {
	src := "https://qa-game.oss-cn-beijing.aliyuncs.com/wxtools/homeworkTest/audio/L2/LessonExercise_L2U11C1/3_speak/L2U11C1_3_2.mp3"
	_, err := utils.DownLoadFileFromUrl("/Users/mn/Desktop/output.mp3", src)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("download resource success")
}

// 根据openid获取用户的信息
func TestGetUserInfoByOpenID(t *testing.T) {
	openid := "oTVNt1dPSf0U7PLI0AytXfhZad0M"
	info, err := service.WechatGetUserInfoByOpenID(openid)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(info)
}

// 测试发送微信模板消息
func TestSendWechatTemplateInfo(t *testing.T) {
	var (
		body = service.SendTemplateRes{}
		err  error
	)
	body.OpenID = "oTVNt1bGq0r807c4p67aPOp_ooQQ"
	body.TempleteID = "oxqQSgyT5aYa2Hmv7nO03MDId5kXZfTB5Q86wR0UM5E"
	body.ActionURL = "http://www.baidu.com"
	pl := constants.GetPhoniceRemindTPL()

	info, err := service.WechatGetUserInfoByOpenID(body.OpenID)
	if err != nil {
		t.Error(err)
		return
	}

	// 更改姓名
	tmp_name := constants.InnerData{}
	tmp_name.Color = pl.Data["keyword1"].Color
	tmp_name.Value = info.Nickname
	pl.Data["keyword1"] = tmp_name

	// 更改性别
	tmp_sex := constants.InnerData{
		Color: pl.Data["keyword2"].Color,
	}
	tmp_sex.Value = strconv.Itoa(info.Sex)
	pl.Data["keyword2"] = tmp_sex

	// 更改放假说明
	tmp_notice := constants.InnerData{
		Color: pl.Data["remark"].Color,
	}
	tmp_notice.Value = fmt.Sprintf("%s 来自  %s,马上就要放假了", info.Nickname, info.Country)
	pl.Data["remark"] = tmp_notice

	body.KeyWordData = pl.Data
	err = service.WechatSendTemplateInfo(body)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("send message success")
}

// 测试微信群发消息
func TestWechatMassSendInfo(t *testing.T) {
	txt := mass2all.NewText("hello")
	err := service.WechatMassSendTextMsgByOpenID(txt)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("success")
}

// 测试微信下载资源
func TestWechatDownloadResource(t *testing.T) {
	mediaID := "UEvKhiihBcKG9Dw9Ni6bsokY3LJIPfqVIVh80HTxfeKOTdVTlSHQZCH1ry5CIUPh"

	imagePath, err := service.WechatDownImageByMediaID(mediaID)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(imagePath)

	//mediaID = "R4B686Pt7GDo0NUXsrK8qsIoywyx9Re4oMxW1OW1p-dRzLbroV0EBbTEMXI3u27E"
	//mp3Path, err := service.WechatDownAudioByMediaID(mediaID)
	//if err != nil {
	//	t.Error(err)
	//	return
	//}
	//fmt.Println("success")
	//fmt.Println(mp3Path)
}
