package main

import (
	"fmt"
	"net/url"
	"self_game/dao"
	"self_game/model"
	"testing"
	"time"
)

// signature=d9687d4fa07bbefe27beade52723d3d745553ba6&echostr=6699841386932337635&timestamp=1536145698&nonce=441907582
func TestJiaMi(t *testing.T) {

	signature := ""
	echostr := ""
	timestamp := 0
	nonce := 0
	fmt.Println(signature, echostr, timestamp, nonce)
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

	pl.Query() = params
	fmt.Println(pl.Query())

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
	db := dao.GetDB()
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

//func TestInsertLevelTestConfig(t *testing.T) {
//	db := dao.GetDB()
//	fileName := "./aa.xlsx"
//	f, err := xlsx.OpenFile(fileName)
//	if err != nil {
//		t.Error(err)
//		return
//	}
//
//	for _, sheet := range f.Sheets {
//		fmt.Println("sheet name=", strings.TrimSpace(sheet.Name) == "L0")
//
//		level := 0
//		if strings.TrimSpace(sheet.Name) == "L2" {
//			level = 2
//		} else if strings.TrimSpace(sheet.Name) == "L4" {
//			level = 4
//		}
//
//		for _, row := range sheet.Rows {
//			if row.Cells[0].String() == "" || strings.HasPrefix(row.Cells[0].String(), "test") {
//				continue
//			}
//
//			cells := row.Cells
//
//			cfg := model.ConfigLevelTest{}
//			switch cells[0].String() {
//			case "1":
//				n := len(cells)
//				cfg.Level = level
//				cfg.Typ = 1
//				cfg.Index = cells[1].Value
//				cfg.Answer = cells[n-1].Value
//				cfg.ImageList = cells[n-2].Value
//				cfg.VoiceURL = cells[n-4].Value
//			case "2":
//				n := len(cells)
//				cfg.Level = level
//				cfg.Typ = 2
//				cfg.Answer = cells[n-1].Value
//				cfg.Index = cells[1].Value
//
//				if level == 0 {
//					cfg.ChoiceList = cells[4].Value
//					cfg.ImageList = cells[3].Value
//
//				} else {
//					cfg.Text = cells[4].Value
//					cfg.ChoiceList = cells[5].Value
//					cfg.ImageList = cells[7].Value
//
//				}
//			case "3":
//				n := len(cells)
//				cfg.Level = level
//				cfg.Typ = 3
//				cfg.Index = cells[1].Value
//				cfg.Answer = cells[n-1].Value
//				cfg.Text = cells[2].Value
//				cfg.VoiceURL = cells[3].Value
//			default:
//				continue
//			}
//
//			err := db.Create(&cfg).Error
//			if err != nil {
//				t.Error(err)
//				return
//			}
//
//			fmt.Println()
//		}
//	}
//
//}
