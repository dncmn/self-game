package service

import (
	"fmt"
	"self-game/config"
	"self-game/utils/async"
	"time"
)

var (
	LocalMap = "init"
)

func init() {
	go async.Do(func() {
		err := initConfigInfo()
		if err != nil {
			logger.Error(err)
			return
		}
	})
}

// 打算在这个方法里面进行一些更改配置文件的操作。
/*
	比如更改配置文件,每隔10h从数据库中检索有没有添加新的课程,然后添加进去,这时候应该添加一个锁。避免在写数据的时候,有读数据的操作。
*/
func initConfigInfo() (err error) {
	lc, _ := time.LoadLocation("Asia/Shanghai")
	ts := time.NewTicker(time.Hour * 10)
	defer ts.Stop()
	for {
		select {
		case <-ts.C:
			now := time.Now().In(lc)
			logger.Infof("before.LocalMap=%v", LocalMap)
			LocalMap = fmt.Sprint(LocalMap, now.Unix())
			logger.Infof("start to sync info.time=%v.value=%v", now.Format(config.Config.Cfg.TimeModelStr), LocalMap)
		}
	}
}
