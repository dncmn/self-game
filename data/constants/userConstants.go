package constants

import (
	//"code.dncmn.io/self-game/config"
	"time"
)

// 用户性别
type UserSexType int8

const (
	UserSexTypeMale UserSexType = iota + 1
	_
	_
	UserSexTypeTotal
)

// 解锁类型
type UnlockCourseType int8

const (
	UnlockCourseByFee UnlockCourseType = iota + 1 // 付费解锁
	UnlockCourseByGM                              // 人工解锁(gm后台解锁)
)

// wechat constants
const (
	WechatDownloadAmrLocalAddr = "/tmp" // wechat constants
)

type InnerData struct {
	Value string `json:"value"`
	Color string `json:"color"`
}

// wechat send template info
type Tpl struct {
	Data map[string]InnerData `json:"data"`
}

func GetPhoniceRemindTPL() (pl *Tpl) {
	pl = &Tpl{}
	data := make(map[string]InnerData)
	data["first"] = InnerData{
		Value: `snaplingo-manan-notice`,
		Color: "#173177",
	}
	data["keyword1"] = InnerData{
		Value: "马楠",
		Color: "#173177",
	}

	data["keyword2"] = InnerData{
		Value: "男",
		Color: "#61b2a7",
	}
	data["keyword3"] = InnerData{
		Value: time.Now().Format("2006-01-02 15:04:05"),
		Color: "#173177",
	}
	data["remark"] = InnerData{
		Value: "you can you up,no can no bi bi",
		Color: "#173177",
	}
	pl.Data = data
	return
}
