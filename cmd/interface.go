package cmd

import (
	"code.dncmn.io/self-game/data"
)

type IUser interface {
	GetUserLoginLogService(uid string, n int) (resp data.UserLoginLogResp, err error)
	UserLogin(req data.UserLoginReq) (resp data.UserLoginResp, err error)
}
