package data

type UserLoginReq struct {
	Name string `json:"name"`
	Pwd  string `json:"pwd"`
}

type UserLoginResp struct {
	Status bool `json:"status"`
}

type UserLoginLogResp struct {
	UID      string      `json:"uid"`
	UserName string      `json:"user_name"`
	IsLogin  bool        `json:"is_login"`
	Logs     []LoginInfo `json:"logs"`
}

type LoginInfo struct {
	LoginTime interface{} `json:"login_time"`
	LoginIP   string      `json:"login_ip"`
}
