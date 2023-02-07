package test

import (
	"encoding/json"
	"fmt"
	"github.com/EricJSanchez/wecom"
	"github.com/EricJSanchez/wecom/cache"
	"github.com/EricJSanchez/wecom/config"
)

func Wework(names ...string) *wecom.Wecom {
	if len(names) > 0 {
		var name = ""
		var agentId = ""
		name = names[0]
		if len(names) > 1 {
			agentId = names[1]
		}
		return NewInstance(name, agentId)
	} else {
		fmt.Sprintf("企业微信查询出错！")
		return nil
	}
}

func NewInstance(name, agentId string) *wecom.Wecom {
	redisCache := cache.NewRedis(&cache.RedisOpts{
		Host:        "192.168.100.107:6379",
		Password:    "t[]******#[p",
		Database:    0,
		MaxIdle:     40,
		MaxActive:   100,
		IdleTimeout: 240,
	})

	// TODO 读取数据库
	wk := wecom.NewWecom(&config.Config{
		CorpID:         name,
		CorpSecret:     "",
		AgentID:        agentId,
		AgentSecret:    "",
		Cache:          redisCache,
		RasVersion:     0,
		RasPrivateKey:  "",
		Token:          "",
		EncodingAESKey: "",
		ContactSecret:  "zhAK******yFgHarM",
		CustomerSecret: "Zdf3ur****7d5Vp4M",
		Cookie:         "",
	})
	return wk
}

func Pr(val ...interface{}) {
	if len(val) > 1 {
		for _, v := range val {
			switch v.(type) {
			case []uint8:
				fmt.Println("[]uint8 ori: ", v)
				fmt.Printf("[]uint8 str: %s-\n\n", v)
				continue
			default:
				bytes, _ := json.MarshalIndent(v, "", "    ")
				fmt.Printf("%T : %s-\n", v, bytes)
			}
		}
	} else {
		bytes, _ := json.MarshalIndent(val, "", "    ")
		fmt.Printf("%T : %s-\n", val, bytes)
	}
}
