package async

import (
	"runtime/debug"

	"self-game/utils/logging"
)

var logs = logging.GetLogger()

//Do 异步运行
func Do(f func()) {
	defer LogRecover()
	f()
}

// LogRecover 处理单个请求错误消息
func LogRecover() {
	if err := recover(); err != nil {
		logs.Errorf("[Recover FATAL] %v", err)
		logs.Errorf("[FATAL Stack] %s", string(debug.Stack()))
	}
}
