package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"self-game/config"
	"self-game/constants"
	"self-game/dao"
	"self-game/model"
	"self-game/utils"
	"self-game/utils/taobaoIP"
	"strings"
)

type UserLoginLogResp struct {
	UID      string      `json:"uid"`
	UserName string      `json:"user_name"`
	IsLogin  bool        `json:"is_login"`
	Logs     []loginInfo `json:"logs"`
}

type loginInfo struct {
	LoginTime interface{} `json:"login_time"`
	LoginIP   string      `json:"login_ip"`
}

func GetUserLoginLogService(uid string, n int) (resp UserLoginLogResp, err error) {
	var (
		dbLogs []model.LogLogin
	)

	resp.UID = uid
	dbLogs, err = dao.GetUserLoginLogByUIDAndLimitDao(uid, n)
	if err != nil {
		logger.Error(err)
		return
	}
	if len(dbLogs) == 0 {
		resp.IsLogin = false
		return
	}
	resp.UserName = dbLogs[0].UserName
	resp.IsLogin = true
	for _, l := range dbLogs {
		res := loginInfo{
			LoginIP:   l.LoginIP,
			LoginTime: l.CreatedAt.Format(config.Config.Cfg.TimeModelStr),
		}
		resp.Logs = append(resp.Logs, res)
	}
	return
}

func GetSignatrueParams(c *gin.Context) (signature, echostr string, timestamp, nonce string, err error) {
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

	if timestamp = c.Query("timestamp"); strings.TrimSpace(timestamp) == "" {
		err = errors.New("params error")
		logger.Error(err)
		return
	}
	if nonce = c.Query("nonce"); strings.TrimSpace(nonce) == "" {
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

// 通过userToken获取uid
func GetUIDByUserToken(token string) (uid string, err error) {

	var (
		ok bool
	)
	uid, ok, err = dao.GetUserIDByUserToken(token)
	if err != nil {
		logger.Errorf("token=%v,err=%v", token, err.Error())
		return
	}
	if !ok {
		err = errors.New("uid not found")
		logger.Errorf("token=%v,err=%v", token, err.Error())
		return
	}
	return
}

// 检查用户信息
func CheckUserExist(name, password string) (user model.User, err error) {

	password = utils.EncodeMD5(password)

	err = gloDB.Model(&model.User{}).Where("user_name=? and password=?", name, password).
		First(&user).Error
	if err != nil {
		err = errors.New("user login error")
		logger.Errorf("username=%v,login error.err=%v", err.Error())
		return
	}

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
