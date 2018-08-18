package timeid

import (
	"sync"

	"github.com/golang/glog"
	"github.com/sdming/gosnow"
	"github.com/xytd/misc"
)

var (
	timeId   int64       = 0
	timeLock *sync.Mutex = &sync.Mutex{}
	snow     *gosnow.SnowFlake
)

// 返回Id,但是时间千万不能被回退
func NewTimeId() int64 {
	timeLock.Lock()
	ret := timeId
	timeId++
	timeLock.Unlock()
	return ret
}

func NewSnowId() (uint64, error) {
	id, err := snow.Next()
	return id, err
}

//func NewIdStr() string {
//	return Int64Str(NewId())
//}

func init() {
	timeId = misc.GetTimestamp()
	if goSnow, err := gosnow.Default(); err != nil {
		glog.Fatalln("Init snowId error")
	} else {
		snow = goSnow
	}
}
