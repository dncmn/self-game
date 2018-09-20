package dao

import (
	"fmt"
	"self_game/constants"
	"self_game/constants/redisKey"
	"self_game/model"
	"self_game/utils"
	"time"
)

// userToken检验
func GetUserIDByUserToken(token string) (uid string, isExist bool, err error) {
	token, isExist, err = redisClient.Get(redisKey.UserIDByToken, token)
	return
}

// 设置userToken  key:value(uid:userToken)
func SetUserToken(uid string, expireDate time.Duration) (token string, err error) {
	token = utils.EncodeMD5(fmt.Sprintf("%v%v", uid, time.Now().Unix()))
	err = redisClient.Set(redisKey.UserToken, uid, token, expireDate)
	return
}

// 设置token--->uid
func SetUserIDByToken(uid, token string) (err error) {
	err = redisClient.Set(redisKey.UserIDByToken, token, uid)
	return
}

// 插入到数据库
func InserUserToDB(username, password, mobile, timezone string, sex constants.UserSexType) (uid string, err error) {
	user := model.User{
		UserName:     username,
		Password:     password,
		Mobile:       mobile,
		Sex:          sex,
		RegosterTime: utils.GetTimeZoneTime(timezone).Unix(),
	}
	err = db.Create(&user).Error
	if err != nil {
		logger.Error(err)
		return
	}
	uid = user.ID
	return
}
