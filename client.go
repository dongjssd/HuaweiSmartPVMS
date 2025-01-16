/**
    @author: dongjs
    @date: 2025/1/10
    @description:
**/

package HuaweiSmartPVMS

import (
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
	"time"
)

type Client struct {
	userName    string
	password    string
	redisClient *redis.Client    //用于存储token
	redisSync   *redsync.Redsync //用于多线程锁
}

var redisSyncCustomOptions = []redsync.Option{
	redsync.WithExpiry(2 * time.Minute),
	redsync.WithTries(600), // 2分钟，每200毫秒重试1次，重试 2 * 60 * (1000/200) 次
	redsync.WithRetryDelay(200 * time.Millisecond),
}

// 初始化访问客户端
func InitClient(userName, password string, redisClient *redis.Client) (*Client, error) {
	client := Client{
		userName:    userName,
		password:    password,
		redisClient: redisClient,
	}
	if redisClient != nil {
		redisSync := initRedisSync(redisClient)
		client.redisSync = redisSync
	} else {
		return nil, errors.New("请传入redis client")
	}
	return &client, nil
}

// initRedisSync creates and returns a new Redsync instance from given Redis connection pools.
func initRedisSync(redisClient *redis.Client) *redsync.Redsync {
	pool := goredis.NewPool(redisClient)
	return redsync.New(pool)
}
