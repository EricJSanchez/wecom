package externalcontact

import (
	"encoding/json"
	"fmt"
	"github.com/EricJSanchez/wecom/util"
)

const send_welcome_msg = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/send_welcome_msg?access_token=%s"

type WelComeMsgReq struct {
	WelcomeCode string `json:"welcome_code"`
	Text        struct {
		Content string `json:"content"`
	} `json:"text"`
	Attachments []interface{} `json:"attachments"`
}

// SendWelComeMsg 发送欢迎语
func (r *Client) SendWelComeMsg(options WelComeMsgReq) (err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(send_welcome_msg, accessToken), options)
	if err != nil {
		return
	}
	var commError util.CommonError
	if err = json.Unmarshal(data, &commError); err != nil {
		return
	}
	if commError.ErrCode != 0 {
		return NewSDKErr(commError.ErrCode, commError.ErrMsg)
	}
	return nil
}
