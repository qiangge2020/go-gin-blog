package gredis

import (
	"github.com/qiangge2020/go-gin-blog/pkg/setting"
	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

func Setup() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     setting.RedisSetting.Host,
		Password: setting.RedisSetting.Password, // 没有密码，默认值
		DB:       0,                             // 默认DB 0
	})
}
