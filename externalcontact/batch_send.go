package externalcontact

import (
	"encoding/json"
	"fmt"
	"github.com/EricJSanchez/wecom/util"
)

const (
	//获取客户群详情
	addMsgTemplateAddr = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/add_msg_template?access_token=%s"
	msgSendResultAddr  = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_groupmsg_send_result?access_token=%s"
)

type AddMsgTemplateOptions struct {
	ChatType       string        `json:"chat_type"`
	ExternalUserid []string      `json:"external_userid"`
	Sender         string        `json:"sender,omitempty"`
	Text           BsText        `json:"text,omitempty"`
	Attachments    []interface{} `json:"attachments,omitempty"`
}
type BsText struct {
	Content string `json:"content,omitempty"`
}
type BsImage struct {
	MediaID string `json:"media_id,omitempty"`
	PicURL  string `json:"pic_url,omitempty"`
}
type BsLink struct {
	Title  string `json:"title,omitempty"`
	Picurl string `json:"picurl,omitempty"`
	Desc   string `json:"desc,omitempty"`
	URL    string `json:"url,omitempty"`
}
type BsMiniprogram struct {
	Title      string `json:"title,omitempty"`
	PicMediaID string `json:"pic_media_id,omitempty"`
	Appid      string `json:"appid,omitempty"`
	Page       string `json:"page,omitempty"`
}
type BsVideo struct {
	MediaID string `json:"media_id,omitempty"`
}
type BsFile struct {
	MediaID string `json:"media_id,omitempty"`
}
type Attachments struct {
	Msgtype     string        `json:"msgtype"`
	Image       BsImage       `json:"image,omitempty"`
	Link        BsLink        `json:"link,omitempty"`
	Miniprogram BsMiniprogram `json:"miniprogram,omitempty"`
	Video       BsVideo       `json:"video,omitempty"`
	File        BsFile        `json:"file,omitempty"`
}

type AddMsgTemplateSchema struct {
	util.CommonError
	FailList []string `json:"fail_list"`
	Msgid    string   `json:"msgid"`
}

func (r *Client) AddMsgTemplate(options AddMsgTemplateOptions) (info AddMsgTemplateSchema, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	optionJson, err := json.Marshal(options)
	if err != nil {
		return
	}
	data, err = util.HTTPPost(fmt.Sprintf(addMsgTemplateAddr, accessToken), string(optionJson))
	if err != nil {
		return
	}
	if err = json.Unmarshal(data, &info); err != nil {
		return
	}
	if info.ErrCode != 0 {
		return info, NewSDKErr(info.ErrCode, info.ErrMsg)
	}
	return info, nil
}

type MsgSendResultOptions struct {
	Msgid  string `json:"msgid"`
	Userid string `json:"userid"`
	Limit  int    `json:"limit"`
	Cursor string `json:"cursor"`
}

type MsgSendResultSchema struct {
	util.CommonError
	NextCursor string     `json:"next_cursor"`
	SendList   []SendList `json:"send_list"`
}
type SendList struct {
	ExternalUserid string `json:"external_userid"`
	ChatID         string `json:"chat_id"`
	Userid         string `json:"userid"`
	Status         int    `json:"status"`
	SendTime       int    `json:"send_time"`
}

func (r *Client) MsgSendResult(options MsgSendResultOptions) (info MsgSendResultSchema, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	optionJson, err := json.Marshal(options)
	if err != nil {
		return
	}
	data, err = util.HTTPPost(fmt.Sprintf(msgSendResultAddr, accessToken), string(optionJson))
	if err != nil {
		return
	}
	if err = json.Unmarshal(data, &info); err != nil {
		return
	}
	if info.ErrCode != 0 {
		return info, NewSDKErr(info.ErrCode, info.ErrMsg)
	}
	return info, nil
}
