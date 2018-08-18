package misc

import (
	"net/url"
)

// 验证参数是否存在
func Validate(v url.Values, params ...string) bool {
	for _, param := range params {
		if v.Get(param) == "" {
			return false
		}
	}
	return true
}
