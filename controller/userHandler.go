package controller

import (
	"errors"
	"fmt"
	"github.com/dncmn/bitset"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"self-game/config"
	"self-game/constants/gameCode"
	"self-game/dao"
	"self-game/model"
	"self-game/service"
	"self-game/utils"
	"self-game/utils/async"
	"self-game/utils/vo"
	"strconv"
	"strings"
)

// 查找用户登录记录(近n次)
func GetUserLoginHandler(c *gin.Context) {
	retData := vo.NewData()
	defer SendResponse(c, retData)
	var (
		uid string
		n   int
		err error
		res interface{}
	)
	if uid = c.Query("uid"); utils.IsStringEmpty(uid) {
		err = errors.New("params error")
		retData.Data = err
		retData.Code = gameCode.RequestParamsError
		logger.Error(err)
		return
	}
	if n, err = strconv.Atoi(c.Query("look_count")); err != nil {
		err = errors.New("params error")
		retData.Code = gameCode.RequestParamsError
		retData.Data = err
		logger.Error(err)
		return
	}

	if res, err = service.GetUserLoginLogService(uid, n); err != nil {
		err = errors.New("params error")
		retData.Code = gameCode.RequestParamsError
		retData.Data = err
		logger.Error(err)
		return
	}
	retData.Data = res
	retData.Code = gameCode.RequestSuccess
	return
}

// 用户登录
func UserLoginHandler(c *gin.Context) {
	retData := vo.NewData()
	defer SendResponse(c, retData)
	var (
		body  service.UserLoginReq
		resp  service.UserLoginResp
		token string
		err   error
		user  model.User
	)

	if err = ParsePostBody(c, &body); err != nil {
		retData.Code = gameCode.RequestLoginUserOrPasswordError
		logger.Error(errors.New("name or password error"))
		return
	}

	if user, err = service.CheckUserExist(body.UserName, body.Password); err != nil {
		retData.Code = gameCode.RequestLoginUserOrPasswordError
		logger.Error(errors.New("name or password error"))
		return
	}

	// 设置 uid-->token
	token, err = dao.SetUserToken(user.ID, 0)
	if err != nil {
		retData.Code = gameCode.RequestParamsError
		logger.Errorf("user_login:save userToken error:uid[%v] err=[%v]", user.ID, err)
		return
	}

	// 设置token-->uid
	err = dao.SetUserIDByToken(user.ID, token)
	if err != nil {
		retData.Code = gameCode.RequestParamsError
		logger.Errorf("user_login:save userToken error:uid[%v] err=[%v]", user.ID, err)
		return
	}

	logger.Infof("createUserToken:uid[%v],token[%v]", user.ID, token)
	resp.UserName = user.UserName
	resp.UID = user.ID
	resp.Mobile = user.Mobile
	resp.City = user.City
	resp.Country = user.Country
	resp.Token = token
	retData.Code = gameCode.RequestSuccess
	retData.Data = resp

	// 记录登录日志
	go async.Do(func() {
		ip := c.ClientIP()
		err = dao.InsertToUserLogin(user.ID, user.UserName, ip, config.Config.Cfg.TimeZone)
		if err != nil {
			logger.Errorf("save login log error: uid=%v,err=%v", user.ID, err)
		}
	})
	return
}

// 用户注册
func RegisterUserHandler(c *gin.Context) {
	retData := &vo.Data{}
	defer SendResponse(c, retData)

	var (
		err          error
		requestBody  service.UserRegisterReq
		registerResp service.UserRegisterResp
		uid          string
	)

	if err = ParsePostBody(c, &requestBody); err != nil {
		logger.Errorf("uname=%v,err=%v", requestBody.UserName, err.Error())
		retData.Code = gameCode.RequestParamsError
		return
	}
	if err = service.CheckUserRegisterParams(requestBody); err != nil {
		retData.Code = gameCode.RequestParamsError
		logger.Error(err)
		return
	}
	// 检查该用户是否存在
	err = service.CheckUserIsExist(requestBody.UserName)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			retData.Code = gameCode.UserNameAlreadyExist
			logger.Error(err)
			return
		}
	}

	// save user to db
	if uid, err = service.InsertUserToDB(requestBody); err != nil {
		retData.Code = gameCode.RequestParamsError
		return
	}

	// 更改玩家的ip,国家和城市
	go async.Do(func() {
		ip := c.ClientIP()
		err = service.UpdateUserCountryAndCity(uid, ip)
		if err != nil {
			logger.Errorf("ip=%v,err=%v", ip, err)
			return
		}
	})

	registerResp.UID = uid
	registerResp.RegisterTme = utils.GetTimeZoneTime(config.Config.Cfg.TimeZone).Format("2006-01-02 15:04:05")
	retData.Data = registerResp
	retData.Code = gameCode.RequestSuccess
	logger.Infof("userRegister:%v", requestBody)
	return
}

func HandlerSignatureHandler(c *gin.Context) {
	var (
		signature string
		echostr   string
		timestamp string
		nonce     string
		err       error
	)

	if signature, echostr, timestamp, nonce, err = service.GetSignatrueParams(c); err != nil {
		fmt.Println(err)
		return
	}

	logger.Infof("signature=%s,echostr=%s,nonce=%s,nonce=%s",
		signature, echostr, timestamp, nonce)
	ok := service.WechatCheckServer(timestamp, nonce, signature)
	if ok {
		c.Data(200, "", []byte(echostr))
		return
	}
	logger.Info("request success")
}

func GetUserNameHandler(c *gin.Context) {
	retData := &vo.Data{}
	defer SendResponse(c, retData)
	fmt.Println("hello")
	var (
		uid string
	)
	uid = c.Param("uid")
	if strings.TrimSpace(uid) == "12345" {
		retData.Code = -101
		retData.Data = "param error"
		return
	}

	retData.Data = map[string]interface{}{
		"name": config.Config.Cfg.Port,
		"env":  config.Config.Env.ENV,
	}
	retData.Code = 1
	return
}

func ConsulHealthCheck(c *gin.Context) {

	retData := vo.NewData()
	defer SendResponse(c, retData)

	var (
		b = &bitset.BitSet{}
	)
	b.Set(3)
	retData.Data = b.String()
	retData.Code = gameCode.RequestSuccess
	return
}

func PostUserNameHandler(c *gin.Context) {
	retData := vo.NewData()
	defer SendResponse(c, retData)

	var (
		res  service.PostUserRequest
		resp service.PostUserResponse
		err  error
		uid  string
	)

	if err = ParsePostBody(c, &res); err != nil {
		retData.Data = err.Error()
		retData.Code = -101
		logger.Error("param error")
		return
	}
	fmt.Println("name=", res.Name, "english_score=", res.EnglishScore)

	if strings.TrimSpace(res.Name) == "" {
		retData.Data = "params error"
		retData.Code = -101
		logger.Error("param error")
		return
	}

	resp.UID = uid
	resp.Name = res.Name
	resp.ChineseScore = res.EnglishScore + 2
	resp.EnglishScore = res.EnglishScore
	retData.Code = 1
	retData.Data = map[string]interface{}{
		"user_info": resp,
	}
	logger.Infof("userName=%v,score=%v,responseBody=%v", res.Name, res.EnglishScore, resp)
	return
}
