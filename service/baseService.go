package service

import (
	"self-game/compoments"
	"self-game/constants"
	"self-game/utils/logging"
)

var (
	logger = logging.GetLogger()
	gloDB  = compoments.GetDB()
	//pgDB        = compoments.GetPGDB()
	redisClient = compoments.GetRedisClient()
)

// write request and response struct

type PostUserRequest struct {
	Name         string `json:"name"`
	EnglishScore int    `json:"english_score"`
}
type PostUserResponse struct {
	UID          string `json:"uid"`
	Name         string `json:"name"`
	EnglishScore int    `json:"english_score"`
	ChineseScore int    `json:"chinese_score"`
}

// 用户登陆
type UserLoginReq struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type UserLoginResp struct {
	UID      string `json:"uid"`
	UserName string `json:"user_name"`
	Country  string `json:"country"`
	City     string `json:"city"`
	Mobile   string `json:"mobile"`
	Token    string `json:"token"`
}

// 用户注册
type UserRegisterReq struct {
	UserName string                `json:"user_name"`
	Password string                `json:"password"`
	Mobile   string                `json:"mobile"`
	Sex      constants.UserSexType `json:"sex"`
}

type UserRegisterResp struct {
	UID         string `json:"uid"`
	RegisterTme string `json:"register_tme"`
}
