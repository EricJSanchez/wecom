package test

import (
	"github.com/EricJSanchez/wecom/cache"
	"github.com/EricJSanchez/wecom/util"
	"testing"
)

// 获取应用管理员列表
func TestRecord(t *testing.T) {
	redisCache := cache.NewRedis(&cache.RedisOpts{
		Host:        "192.168.100.107:6379",
		Password:    "111111",
		Database:    0,
		MaxIdle:     40,
		MaxActive:   100,
		IdleTimeout: 240,
	})
	err := util.Record(redisCache, "http://a.qq.com/path/of/url2?token=tokenstring")
	Pr(err)
}
