package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"strings"
)

func GetSignatrueParams(c *gin.Context) (signature, echostr string, timestamp, nonce int, err error) {
	if signature = c.Query("signature"); strings.TrimSpace(signature) == "" {
		err = errors.New("params error")
		log.Println(err)
		return
	}
	if echostr = c.Query("echostr"); strings.TrimSpace(signature) == "" {
		err = errors.New("params error")
		log.Println(err)
		return
	}

	if timestamp, err = strconv.Atoi(c.Query("timestamp")); err != nil {
		err = errors.New("params error")
		log.Println(err)
		return
	}
	if nonce, err = strconv.Atoi(c.Query("nonce")); err != nil {
		err = errors.New("params error")
		log.Println(err)
		return
	}
	return
}

// 通过用户名查找用户信息
func GetUserByUID(uid string) (user interface{}, err error) {
	return

}
