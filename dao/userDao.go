package dao

import (
	"fmt"
	"self_game/constants"
	"self_game/constants/redisKey"
	"self_game/model"
	"self_game/utils"
	"time"
)

// 设置userToken
func SetUserToken(uid string, expireDate time.Duration) (token string, err error) {
	token = utils.EncodeMD5(fmt.Sprintf("%v%v", uid, time.Now().Unix()))
	err = redisClient.Set(redisKey.UserToken, uid, token, expireDate)
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
