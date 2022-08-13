package test

import (
	"fmt"
	"github.com/EricJSanchez/wecom"
	"github.com/EricJSanchez/wecom/cache"
	"github.com/EricJSanchez/wecom/config"
	"github.com/EricJSanchez/wecom/contact"
	"testing"
)

func NewInstance(name, agentId string) *wecom.Wecom {
	redisCache := cache.NewRedis(&cache.RedisOpts{
		Host:        "127.0.0.1:6379",
		Password:    "111111",
		Database:    4,
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
		ContactSecret:  "YNW************Z0",
		CustomerSecret: "",
	})
	return wk
}

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

func TestStaffUserId(t *testing.T) {
	weCom, err := Wework("*****").GetContact()
	if err != nil {
		t.Error(err)
		return
	}
	userList, err := weCom.UserListId(contact.UserListIdOptions{
		Cursor: "",
		Limit:  100,
	})
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(userList)
}
