package user

import (
	"code.dncmn.io/self-game/cmd"
	"code.dncmn.io/self-game/dao"
	"code.dncmn.io/self-game/data"
	"code.dncmn.io/self-game/model"
)

func init() {
	cmd.User = User{}
}

type User struct {
}

func (u User) UserLogin(body data.UserLoginReq) (resp data.UserLoginResp, err error) {
	return
}

func (u User) GetUserLoginLogService(uid string, n int) (resp data.UserLoginLogResp, err error) {
	var (
		dbLogs []model.LogLogin
	)

	var (
	//courses = make([]model.UserCourse, 0)
	//user    = model.User{UID: "65fd2df7-cbf6-43a7-b746-534fc86d38a9"}
	)

	//gloDB.Model(&user).Related(&courses, "Courses")
	//
	//for _, c := range courses {
	//	fmt.Println(fmt.Sprintf("uid=%v,courseID=%v", c.UID, c.CourseID))
	//}

	resp.UID = uid
	dbLogs, err = dao.GetUserLoginLogByUIDAndLimitDao(uid, n)
	if err != nil {
		return
	}
	if len(dbLogs) == 0 {
		resp.IsLogin = false
		return
	}
	resp.UserName = dbLogs[0].UserName
	resp.IsLogin = true
	for _, l := range dbLogs {
		res := data.LoginInfo{
			LoginIP:   l.LoginIP,
			LoginTime: l.CreatedAt.Format("2006-01-02"),
		}
		resp.Logs = append(resp.Logs, res)
	}
	return
}
