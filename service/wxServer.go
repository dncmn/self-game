package service

import (
	"self-game/config"
	"self-game/utils"
	"self-game/utils/async"
	"time"
)

var (
	IsStop = true
)

func init() {
	go async.Do(func() {
		initService()
	})
}

func initService() {
	logger.Info("sync info begin")
	for IsStop {
		logger.Infof("current_time=%v", utils.GetTimeZoneTime(config.Config.Cfg.TimeZone).Unix())
		time.Sleep(time.Second * 30)
	}
	logger.Info("sync info end")
}
