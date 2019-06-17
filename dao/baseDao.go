package dao

import (
	"code.dncmn.io/self-game/compoments"
	"code.dncmn.io/self-game/utils/logging"
	"github.com/jinzhu/gorm"
)

var (
	logger      = logging.GetLogger()
	redisClient *compoments.RedisInstance
	db          *gorm.DB
	pgdb        *gorm.DB
)

func init() {
	db = compoments.GetDB()
	redisClient = compoments.GetRedisClient()
	//pgdb = compoments.GetPGDB()
}
