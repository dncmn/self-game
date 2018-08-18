package main

import (
	"fmt"
	"self_game/dao"
	"self_game/model"
	"testing"
	"time"
)

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
