package anysdk

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"net"
	"regexp"
	"sort"
)

// CheckRemoteIpFunc 检查函数生成器,检测订单输入远程IP是否属于认证IP
func CheckRemoteIPFunc(allowAddrs []string) func(addr string) bool {

	return func(addr string) bool {
		// check ip style
		ip := addr
		remoteIP := net.ParseIP(addr)
		if remoteIP == nil {
			return false
		}

		ip = remoteIP.String()
		check := false
		for _, v := range allowAddrs {
			if v == ip {
				check = true
				break
			}
		}
		return check
	}
}

// CheckPaySignFunc 检查函数生成器,检测订单签名
func CheckPaySignFunc(privateKey, enhancedKey string) func(param map[string]string) (signCheck bool, enhancedSignCheck bool) {
	if privateKey == "" {
		return nil
	}
	return func(param map[string]string) (signCheck bool, enhancedSignCheck bool) {
		// check sign field
		sign, ok := param["sign"]
		if !ok || sign == "" {
			return
		}
		// merge params values by sorted keys
		keysList := make([]string, 0, 20)
		for k, v := range param {
			// ignore empty value and sign field
			if v != "" && k != "sign" {
				keysList = append(keysList, k)
			}
		}
		sort.Strings(keysList)

		// merging
		var mergedValues bytes.Buffer
		for _, k := range keysList {
			mergedValues.WriteString(param[k])
		}
		mergedValueStr := mergedValues.String()
		// calculate md5 sum for merged values append with private key
		calcSign := fmt.Sprintf("%x", md5.Sum([]byte(
			fmt.Sprintf("%x", md5.Sum([]byte(mergedValueStr)))+privateKey)))
		// sign check passed
		if calcSign == sign {
			signCheck = true
		}

		// enhanced_sign check
		if enhancedKey == "" {
			return
		}
		enhancedSign, ok := param["enhanced_sign"]
		if !ok {
			return
		}
		// remove enhancedKey from values str
		enhancedKeyRegexp := regexp.MustCompile(enhancedSign)
		mergedValueStr = enhancedKeyRegexp.ReplaceAllString(mergedValueStr, "")
		calcEnhancedSign := fmt.Sprintf("%x", md5.Sum([]byte(
			fmt.Sprintf("%x", md5.Sum([]byte(mergedValueStr)))+enhancedKey)))
		if calcEnhancedSign != enhancedSign {
			return
		}
		enhancedSignCheck = true
		return
	}
}

// GenOrderSign 生成订单的签名
func GenOrderSign(privateKey, enhancedKey string, param map[string]string) (sign string, enhancedSign string) {
	// merge params values by sorted keys
	keysList := make([]string, 0, 20)
	for k := range param {
		keysList = append(keysList, k)
	}
	sort.Strings(keysList)
	var mergedValues bytes.Buffer
	for _, k := range keysList {
		mergedValues.WriteString(param[k])
	}
	mergedValueStr := mergedValues.Bytes()
	sign = fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%x", md5.Sum(mergedValueStr))+privateKey)))
	enhancedSign = fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%x", md5.Sum(mergedValueStr))+enhancedKey)))
	return
}
