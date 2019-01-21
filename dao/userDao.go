package dao

import (
	"fmt"
	"self-game/constants"
	"self-game/constants/redisKey"
	"self-game/model"
	"self-game/utils"
	"time"
)

// 查找登录日志
func GetUserLoginLogByUIDAndLimitDao(uid string, limit int) (data []model.LogLogin, err error) {
	err = db.Where("uid=?", uid).Order("login_time desc").Limit(limit).Find(&data).Error
	return
}

// 保存登录日志
func InsertToUserLogin(uid, uname, loginIP, timeZone string) (err error) {
	loginLog := model.LogLogin{
		UID:       uid,
		UserName:  uname,
		LoginIP:   loginIP,
		LoginTime: utils.GetTimeZoneTime(timeZone).Unix(),
	}
	//err = db.Create(&loginLog).Error
	err = pgdb.Create(&loginLog).Error
	return
}

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
		UserName: username,
		Mobile:   mobile,
		Sex:      sex,
	}
	err = db.Create(&user).Error
	if err != nil {
		logger.Error(err)
		return
	}
	uid = user.ID
	return
}
