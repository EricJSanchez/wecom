package externalcontact

import (
	"encoding/json"
	"fmt"
	"github.com/EricJSanchez/wecom/util"
)

const (
	//获取客户群详情
	addMsgTemplateAddr = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/add_msg_template?access_token=%s"
)

type AddMsgTemplateOptions struct {
	ChatType       string        `json:"chat_type"`
	ExternalUserid []string      `json:"external_userid"`
	Sender         string        `json:"sender"`
	Text           Text          `json:"text,omitempty"`
	Attachments    []Attachments `json:"attachments"`
}
type BsText struct {
	Content string `json:"content"`
}
type BsImage struct {
	MediaID string `json:"media_id"`
	PicURL  string `json:"pic_url"`
}
type BsLink struct {
	Title  string `json:"title"`
	Picurl string `json:"picurl"`
	Desc   string `json:"desc"`
	URL    string `json:"url"`
}
type BsMiniprogram struct {
	Title      string `json:"title"`
	PicMediaID string `json:"pic_media_id"`
	Appid      string `json:"appid"`
	Page       string `json:"page"`
}
type BsVideo struct {
	MediaID string `json:"media_id"`
}
type BsFile struct {
	MediaID string `json:"media_id"`
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
