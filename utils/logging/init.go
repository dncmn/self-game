package logging

import (
	"github.com/snaplingo-org/sllog"
)

var (
	logging     *sllog.Sllogger
	gormLogging *sllog.GormLogger
)

func init() {
	sllog.Init(sllog.FileLogConfig{
		Path:     "/var/log/game-server",
		Filename: "/var/log/game-server/self.log",
		MaxLines: 1000000,
		Maxsize:  1 << 28, //256 MB
		Daily:    true,
		MaxDays:  3,
		Rotate:   true,
		Level:    sllog.DebugLevel,
	})

	logging = sllog.GetLogger()
	gormLogging = sllog.GetGormLogger()
	logging.Info("logger success init.")
}

func GetLogger() *sllog.Sllogger {
	return logging
}

func GetGormLogger() *sllog.GormLogger {
	return gormLogging
}
