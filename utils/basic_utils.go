package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
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

// sha1加密
func EncodeSha1(pwd string) (result string) {
	sha1 := sha1.New()
	sha1.Write([]byte(pwd))
	return hex.EncodeToString(sha1.Sum([]byte("")))
}

// md5加密
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

func MapToStruct(source interface{}, target interface{}) (err error) {
	tempBytes, err := json.Marshal(source)
	if err != nil {
		return
	}
	// 获取数据
	if err = json.Unmarshal(tempBytes, target); err != nil {
		return
	}
	return
}

func StructToMap(src, dst interface{}) (err error) {
	byt := make([]byte, 0)
	byt, err = json.Marshal(src)
	if err != nil {
		return
	}
	err = json.Unmarshal(byt, &dst)
	return
}

// down file from url
func DownLoadFileFromUrl(filePath, url string) (body []byte, err error) {
	// download file
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("get file data error", err)
		return
	}

	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read file data err", err)
		return
	}

	err = ioutil.WriteFile(filePath, body, 0644)
	if err != nil {
		fmt.Println("write data to file error", err)
		return
	}

	return body, nil
}
