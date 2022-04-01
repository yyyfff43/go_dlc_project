/*
* @File : redis
* @Describe :
* @Author: yangfan@zongheng.com
* @Date : 2022/4/1 11:23
* @Software: GoLand
 */

package redis

import (
	"fmt"
	"go_dlc_project/setting"

	"github.com/go-redis/redis"
)

var (
	client *redis.Client
	Nil    = redis.Nil
)

// Init 初始化连接
func Init(cfg *setting.RedisConfig) (err error) {
	client = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password:     cfg.Password, // no password set
		DB:           cfg.DB,       // use default DB
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConns,
	})

	_, err = client.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func Close() {
	_ = client.Close()
}
