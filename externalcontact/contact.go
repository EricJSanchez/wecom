package externalcontact

import (
	"encoding/json"
	"fmt"
	"github.com/EricJSanchez/wecom/util"
)

const (
	// 获取企业已配置的「联系我」列表 list_contact_way
	externalcontactListContactWay = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/list_contact_way?access_token=%s"
	//获取企业已配置的「联系我」方式 get_contact_way
	externalcontactGetContactWay = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_contact_way?access_token=%s"
	//删除企业已配置的「联系我」方式 del_contact_way
	externalcontactDelContactWay = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/del_contact_way?access_token=%s"
	//配置客户联系「联系我」方式 add_contact_way
	externalcontactAddContactWay = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/add_contact_way?access_token=%s"
	//更新企业已配置的「联系我」方式 update_contact_way
	externalcontactUpdateContactWay = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/update_contact_way?access_token=%s"
)

// 获取企业已配置的「联系我」列表 参数
type ListContactWayGetOptions struct {
	StartTime int64  `json:"start_time"`
	EndTime   int64  `json:"end_time"`
	Cursor    string `json:"cursor"`
	Limit     int    `json:"limit"`
}

// 获取企业已配置的「联系我」列表 响应内容
type ListContactWayGetSchema struct {
	util.CommonError
	ContactWay []ContactWay `json:"contact_way"`
	NextCursor string       `json:"next_cursor"`
}

// ContactWay  信息
type ContactWay struct {
	ConfigId string `json:"config_id"`
}

// ListContactWayGet 批量获取「联系我」列表
func (r *Client) ListContactWayGet(options ListContactWayGetOptions) (info ListContactWayGetSchema, err error) {
	var (
		accessToken string
		data        []byte
	)
	_ = util.Record(r.cache, externalcontactListContactWay)
	accessToken, err = r.ctx.GetAccessToken()
	//fmt.Println(accessToken)
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(externalcontactListContactWay, accessToken), options)
	if err != nil {
		return
	}
	//fmt.Println(string(data))
	if err = json.Unmarshal(data, &info); err != nil {
		return
	}
	//fmt.Println(info)
	if info.ErrCode != 0 {
		return info, NewSDKErr(info.ErrCode, info.ErrMsg)
	}
	return info, nil
}

// 获取企业已配置的「联系我」列表 响应内容
type GetContactInfoSchema struct {
	util.CommonError
	ContactWay ContactWayInfo `json:"contact_way"`
}
type ContactWayInfo struct {
	ConfigId      string      `json:"config_id"`
	Type          int         `json:"type"`
	Scene         int         `json:"scene"`
	IsTemp        bool        `json:"is_temp"`
	Remark        string      `json:"remark"`
	SkipVerify    bool        `json:"skip_verify"`
	State         string      `json:"state"`
	Style         int         `json:"style"`
	QrCode        string      `json:"qr_code"`
	User          []string    `json:"user"`
	Party         []int       `json:"party"`
	ExpiresIn     int         `json:"expires_in"`
	ChatExpiresIn int         `json:"chat_expires_in"`
	Unionid       string      `json:"unionid"`
	Conclusions   Conclusions `json:"conclusions"`
}

type Conclusions struct {
	Text        Text        `json:"text"`
	Image       Image       `json:"image"`
	Link        Link        `json:"link"`
	Miniprogram Miniprogram `json:"miniprogram"`
}

type Text struct {
	Content string `json:"content"`
}

type Image struct {
	MediaId string `json:"media_id"`
	PicUrl  string `json:"pic_url"`
}
type Link struct {
	Title  string `json:"title"`
	Picurl string `json:"picurl"`
	Desc   string `json:"desc"`
	URL    string `json:"url"`
}

type Miniprogram struct {
	Title      string `json:"title"`
	PicMediaId string `json:"pic_media_id"`
	Appid      string `json:"appid"`
	Page       string `json:"page"`
}

// 获取企业已配置的「联系我」方式
func (r *Client) GetContactInfo(options ContactWay) (info GetContactInfoSchema, err error) {
	var (
		accessToken string
		data        []byte
	)
	_ = util.Record(r.cache, externalcontactGetContactWay)
	accessToken, err = r.ctx.GetAccessToken()
	//fmt.Println(accessToken)
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(externalcontactGetContactWay, accessToken), options)
	if err != nil {
		return
	}
	//fmt.Println(string(data))
	if err = json.Unmarshal(data, &info); err != nil {
		return
	}
	//fmt.Println(info)
	if info.ErrCode != 0 {
		return info, NewSDKErr(info.ErrCode, info.ErrMsg)
	}
	return info, nil
}

type ContactWaySchema struct {
	util.CommonError
}

// 删除 【联系我】 方式
func (r *Client) DelContactWay(options ContactWay) (info ContactWaySchema, err error) {
	var (
		accessToken string
		data        []byte
	)
	_ = util.Record(r.cache, externalcontactDelContactWay)
	accessToken, err = r.ctx.GetAccessToken()
	//fmt.Println(accessToken)
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(externalcontactDelContactWay, accessToken), options)
	if err != nil {
		return
	}
	//fmt.Println(string(data))
	if err = json.Unmarshal(data, &info); err != nil {
		return
	}
	//fmt.Println(info)
	if info.ErrCode != 0 {
		return info, NewSDKErr(info.ErrCode, info.ErrMsg)
	}
	return info, nil
}

type AddContactWayOptions struct {
	Type          int         `json:"type"`
	Scene         int         `json:"scene"`
	Style         int         `json:"style,omitempty"`
	Remark        string      `json:"remark"`
	SkipVerify    bool        `json:"skip_verify"`
	State         string      `json:"state,omitempty"`
	User          []string    `json:"user,omitempty"`
	Party         []int       `json:"party,omitempty"`
	IsTemp        bool        `json:"is_temp,omitempty"`
	ExpiresIn     int         `json:"expires_in,omitempty"`
	ChatExpiresIn int         `json:"chat_expires_in,omitempty"`
	Unionid       string      `json:"unionid,omitempty"`
	Conclusions   Conclusions `json:"conclusions,omitempty"`
}

type AddContactWaySchema struct {
	util.CommonError
	ConfigId string `json:"config_id"`
	QrCode   string `json:"qr_code"`
}

func (r *Client) AddContactWay(options AddContactWayOptions) (info AddContactWaySchema, err error) {
	var (
		accessToken string
		data        []byte
	)
	_ = util.Record(r.cache, externalcontactAddContactWay)
	accessToken, err = r.ctx.GetAccessToken()
	//fmt.Println(accessToken)
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(externalcontactAddContactWay, accessToken), options)
	if err != nil {
		return
	}
	//fmt.Println(string(data))
	if err = json.Unmarshal(data, &info); err != nil {
		return
	}
	//fmt.Println(info)
	if info.ErrCode != 0 {
		return info, NewSDKErr(info.ErrCode, info.ErrMsg)
	}
	return info, nil
}

type UpdateContactWayOptions struct {
	ConfigId      string      `json:"config_id"`
	Remark        string      `json:"remark"`
	SkipVerify    bool        `json:"skip_verify"`
	Style         int         `json:"style,omitempty"`
	State         string      `json:"state,omitempty"`
	User          []string    `json:"user,omitempty"`
	Party         []int       `json:"party,omitempty"`
	ExpiresIn     int         `json:"expires_in,omitempty"`
	ChatExpiresIn int         `json:"chat_expires_in,omitempty"`
	Unionid       string      `json:"unionid,omitempty"`
	Conclusions   Conclusions `json:"conclusions,omitempty"`
}

func (r *Client) UpdateContactWay(options UpdateContactWayOptions) (info ContactWaySchema, err error) {
	var (
		accessToken string
		data        []byte
	)
	_ = util.Record(r.cache, externalcontactUpdateContactWay)
	accessToken, err = r.ctx.GetAccessToken()
	//fmt.Println(accessToken)
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(externalcontactUpdateContactWay, accessToken), options)
	if err != nil {
		return
	}
	//fmt.Println(string(data))
	if err = json.Unmarshal(data, &info); err != nil {
		return
	}
	//fmt.Println(info)
	if info.ErrCode != 0 {
		return info, NewSDKErr(info.ErrCode, info.ErrMsg)
	}
	return info, nil
}
