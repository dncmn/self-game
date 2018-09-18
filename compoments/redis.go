package compoments

import (
	"github.com/go-redis/redis"
	"log"
	"self_game/config"
	"self_game/utils/logging"
	"time"
)

var redisClient *RedisInstance
var logs = logging.GetLogger()

type RedisInstance struct {
	RedisCli *redis.Client
}

// 获取redisClient instance
func GetRedisClient() *RedisInstance {
	return redisClient
}

// 初始化redisClient
func init() {
	redisClient = initRedisClient()
}
func initRedisClient() (ri *RedisInstance) {
	var (
		redisInstance RedisInstance
		redisCli      *redis.Client
		pong          string
		err           error
	)
	redisConfig := config.Config.Redis
	logs.Info(redisConfig)
	redisCli = redis.NewClient(&redis.Options{
		Addr:         redisConfig.Host,
		Password:     redisConfig.Password,
		DB:           redisConfig.DB,
		MaxRetries:   3,
		IdleTimeout:  5 * time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	})

	if pong, err = redisCli.Ping().Result(); err != nil {
		log.Fatal(err)
	}
	log.Print("[redis ping success] ", pong)

	// 每分钟的心跳
	go func() {
		for range time.Tick(1 * time.Minute) {
			if _, err = redisCli.Ping().Result(); err == nil {
				log.Print("[Redis PING Success]")
			} else {
				log.Print("[Redis PING Error] ", err)
			}
		}
	}()
	redisInstance.RedisCli = redisCli
	return ri
}
