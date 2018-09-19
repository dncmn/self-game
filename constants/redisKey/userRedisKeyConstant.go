package redisKey

import (
	"fmt"
	"time"
)

type RedisKeyInfo struct {
	Project Project
	Module  RedisModule
	Key     string
	Expire  time.Duration
}
type Project string
type RedisModule string

const (
	Company = "HappySelf"
)
const (
	mUser RedisModule = "user"
	mBot  RedisModule = "bot"
)
const (
	//ProjectApi self_game项目
	ProjectApi Project = "self_game"
)

var (
	// UserToken 用户令牌 根据用户id 来获取token: UserToken + uid
	UserToken = &RedisKeyInfo{
		Key:    "userToken",
		Expire: 0,
	}
	// UserIDByToken 用户令牌 根据用户Token获取用户ID UserIdByToken + token
	UserIDByToken = &RedisKeyInfo{
		Key:    "userTokenV2:",
		Expire: 0,
	}
	// SphinxUserBotIDByUID sphinx 用户的机器人ID
	SphinxUserBotIDByUID = &RedisKeyInfo{
		Key:    "SphinxUserBotID",
		Expire: 0,
		Module: mUser,
	}
)

const (
	// FieldAttrGuide 用户属性标记
	FieldAttrGuide = "guide"
)

const splitChar = ":"

// GetStrKey 获取redisKey
func (rk *RedisKeyInfo) GetStrKey(arg string) (key string) {
	if rk.Module != "" {
		project := rk.Project
		if project == "" {
			project = ProjectApi
		}
		key = fmt.Sprintf("%s:%s:%s:%s", Company, project, rk.Module, rk.Key)
		if arg != "" {
			key = key + splitChar + arg
		}
	} else {
		key = rk.Key + arg
	}
	return
}
