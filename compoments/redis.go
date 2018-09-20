package compoments

import (
	"github.com/go-redis/redis"
	"log"
	"self-game/config"
	"self-game/constants/redisKey"
	"time"
)

var redisClient *RedisInstance

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
	ri = &redisInstance
	return ri
}

//Set 设置redisKey
func (ri *RedisInstance) Set(rk *redisKey.RedisKeyInfo, args string, value interface{}, expires ...time.Duration) (err error) {
	expire := rk.Expire
	if len(expires) > 0 {
		expire = expires[0]
	}
	err = ri.RedisCli.Set(rk.GetStrKey(args), value, expire).Err()
	return
}

//Get 获取redis数据
func (ri *RedisInstance) Get(rk *redisKey.RedisKeyInfo, args string) (result string, keyExist bool, err error) {
	result, err = ri.RedisCli.Get(rk.GetStrKey(args)).Result()
	if err != nil {
		if err == redis.Nil {
			return result, false, nil
		}
		return result, false, err
	}
	return result, true, nil
}

//HSet hset设置redisKey
func (ri *RedisInstance) HSet(rk *redisKey.RedisKeyInfo, args string, field string, value interface{}) (err error) {
	err = ri.RedisCli.HSet(rk.GetStrKey(args), field, value).Err()
	return
}

//HGet hget获取redis数据
func (ri *RedisInstance) HGet(rk *redisKey.RedisKeyInfo, args string, field string) (result string, err error) {
	result, err = ri.RedisCli.HGet(rk.GetStrKey(args), field).Result()
	return
}

//Del 删除key数据
func (ri *RedisInstance) Del(rk *redisKey.RedisKeyInfo, args string) (err error) {
	err = ri.RedisCli.Del(rk.GetStrKey(args)).Err()
	return
}
