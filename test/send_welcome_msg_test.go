package test

import (
	"github.com/EricJSanchez/wecom/externalcontact"
	"testing"
)

func TestSend(t *testing.T) {

	client, err := Wework("wwd384e73ed3cf5305", "1000030").GetExternalContact()
	if err != nil {
		t.Error(err)
	}

	req := externalcontact.WelComeMsgReq{
		WelcomeCode: "ccccccccc",
		Text: struct {
			Content string `json:"content"`
		}(struct{ Content string }{Content: "aaaa"}),
		Attachments: []interface{}{},
	}

	err = client.SendWelComeMsg(req)
	if err != nil {
		t.Error(err)
	}

}
