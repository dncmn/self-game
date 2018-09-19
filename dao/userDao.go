package dao

import (
	"self_game/constants"
	"self_game/model"
	"self_game/utils"
)

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
