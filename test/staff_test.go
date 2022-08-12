package test

import (
	"fmt"
	"github.com/EricJSanchez/wecom"
	"github.com/EricJSanchez/wecom/cache"
	"github.com/EricJSanchez/wecom/config"
	"github.com/EricJSanchez/wecom/contact"
	"testing"
)

func NewInstance(name, agentId string) *wecom.Wework {
	redisCache := cache.NewRedis(&cache.RedisOpts{
		Host:        "192.168.0.0:6379",
		Password:    "111111",
		Database:    4,
		MaxIdle:     40,
		MaxActive:   100,
		IdleTimeout: 240,
	})

	// TODO 读取数据库
	wk := wecom.NewWework(&config.Config{
		CorpID:         name,
		CorpSecret:     "",
		AgentID:        agentId,
		AgentSecret:    "",
		Cache:          redisCache,
		RasVersion:     0,
		RasPrivateKey:  "",
		Token:          "",
		EncodingAESKey: "",
		ContactSecret:  "***secret***",
		CustomerSecret: "",
	})
	return wk
}

func Wework(names ...string) *wecom.Wework {
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
	weCom, err := Wework("wwe4f9*******36d6").GetContact()
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
