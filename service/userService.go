package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"self_game/config"
	"self_game/constants"
	"self_game/dao"
	"self_game/model"
	"self_game/utils"
	"self_game/utils/taobaoIP"
	"strconv"
	"strings"
)

func GetSignatrueParams(c *gin.Context) (signature, echostr string, timestamp, nonce int, err error) {
	if signature = c.Query("signature"); strings.TrimSpace(signature) == "" {
		err = errors.New("params error")
		logger.Error(err)
		return
	}
	if echostr = c.Query("echostr"); strings.TrimSpace(signature) == "" {
		err = errors.New("params error")
		logger.Error(err)
		return
	}

	if timestamp, err = strconv.Atoi(c.Query("timestamp")); err != nil {
		err = errors.New("params error")
		logger.Error(err)
		return
	}
	if nonce, err = strconv.Atoi(c.Query("nonce")); err != nil {
		err = errors.New("params error")
		logger.Error(err)
		return
	}
	return
}

// 通过用户名查找用户信息
func GetUserByUID(uid string) (user interface{}, err error) {
	return

}

// 查询国家和城市
func UpdateUserCountryAndCity(uid, ip string) (err error) {
	country, city, err := taobaoIP.GetCountryAndCity(ip)
	if err != nil {
		logger.Errorf("ip=%v,err=%v", ip, err.Error())
		return
	}
	err = gloDB.Model(&model.User{}).Where("id=?", uid).
		Update(map[string]interface{}{"register_ip": ip, "country": country, "city": city}).Error
	logger.Infof("uid=%s,ip=%s,country=%s,city=%s", uid, ip, country, city)
	return
}

//插入玩家到数据库
func InsertUserToDB(body UserRegisterReq) (uid string, err error) {
	body.Password = utils.EncodeMD5(body.Password)
	uid, err = dao.InserUserToDB(body.UserName, body.Password, body.Mobile, config.Config.Cfg.TimeZone,
		body.Sex)
	if err != nil {
		logger.Error(err)
	}
	return
}

// 检查用户是否存在
func CheckUserIsExist(name string) (err error) {
	user := model.User{}
	err = gloDB.Model(&model.User{}).Where("user_name=?", name).First(&user).Error
	if err == nil {
		err = errors.New("user already exist")
	}
	return
}

// 检查用户注册的时候的参数
func CheckUserRegisterParams(body UserRegisterReq) (err error) {
	if utils.IsStringEmpty(body.UserName) || utils.IsStringEmpty(body.Password) || utils.IsStringEmpty(body.Mobile) {
		err = errors.New("param is null")
		logger.Errorf("userRegister:username is null:%v", err)
		return
	}

	if !utils.CheckMobileIsLegal(body.Mobile) {
		err = errors.New("mobile is not illegal")
		logger.Errorf("userRegister:mobile is wrong.err=%v", err.Error())
		return
	}

	if body.Sex < constants.UserSexTypeMale || body.Sex >= constants.UserSexTypeTotal {
		err = errors.New("user sex error")
		logger.Errorf("userRegister:sex is not legal.err=%v", err.Error())
		return
	}
	return
}
