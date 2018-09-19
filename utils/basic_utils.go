package utils

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"regexp"
	"strings"
	"time"
)

// 检查手机号格式是否正确
func CheckMobileIsLegal(mobile string) (ok bool) {
	reg := `^1([38][0-9]|14[57]|5[^4])\d{8}$`
	rgx := regexp.MustCompile(reg)
	ok = rgx.MatchString(mobile)
	return
}

// 检查字符串是否为空
func IsStringEmpty(v string) (ok bool) {
	if strings.TrimSpace(v) == "" || len(v) == 0 {
		ok = true
	}
	return
}

// 用户注册时的md5加密
func EncodeMD5(pwd string) (result string) {
	h := md5.New()
	io.WriteString(h, pwd)
	result = hex.EncodeToString(h.Sum(nil))
	return
}

// 获取指定时区的时间
func GetTimeZoneTime(timezone string) (serverTiem time.Time) {
	zone, _ := time.LoadLocation(timezone)
	serverTiem = time.Now().In(zone)
	return
}
