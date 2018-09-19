package service

import "self_game/utils/logging"

var (
	logger = logging.GetLogger()
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
