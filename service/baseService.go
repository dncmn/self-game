package service

import (
	"self_game/compoments"
	"self_game/constants"
	"self_game/utils/logging"
)

var (
	logger = logging.GetLogger()
	gloDB  = compoments.GetDB()
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

// 用户注册
type UserRegisterReq struct {
	UserName string                `json:"user_name"`
	Password string                `json:"password"`
	Mobile   string                `json:"mobile"`
	Sex      constants.UserSexType `json:"sex"`
}

type UserRegisterResp struct {
}
