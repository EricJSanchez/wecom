package application

import (
	"encoding/json"
	"fmt"
	"github.com/EricJSanchez/wecom/util"
)

const (
	// 创建菜单
	createMenuAddr = "https://qyapi.weixin.qq.com/cgi-bin/menu/create?access_token=%s&agentid=%s"
	// 获取菜单
	getMenuAddr = "https://qyapi.weixin.qq.com/cgi-bin/menu/get?access_token=%s&agentid=%s"
	// 删除菜单
	deleteMenuAddr = "https://qyapi.weixin.qq.com/cgi-bin/menu/delete?access_token=%s&agentid=%s"
)

type SubButton struct {
	Type      string        `json:"type"`
	Name      string        `json:"name"`
	URL       string        `json:"url,omitempty"`
	Key       string        `json:"key,omitempty"`
	SubButton []interface{} `json:"sub_button,omitempty"`
	Pagepath  string        `json:"pagepath,omitempty"`
	Appid     string        `json:"appid,omitempty"`
}

type Button struct {
	Name      string      `json:"name"`
	SubButton []SubButton `json:"sub_button,omitempty"`
	Type      string      `json:"type,omitempty"`
	Key       string      `json:"key,omitempty"`
}

// CreateMenuOptions 创建菜单请求参数
type CreateMenuOptions struct {
	Button []Button `json:"button"`
}

// CreateMenuSchema 创建菜单响应内容
type CreateMenuSchema struct {
	util.CommonError
}

// CreateMenu 创建菜单
func (r *Client) CreateMenu(options CreateMenuOptions) (info CreateMenuSchema, err error) {
	var (
		accessToken string
		data        []byte
	)
	_ = util.Record(r.cache, createMenuAddr)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(createMenuAddr, accessToken, r.agentID), options)
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

// GetMenuOptions 获取菜单请求参数
type GetMenuOptions struct {
}

// GetMenuSchema 获取菜单响应内容
type GetMenuSchema struct {
	util.CommonError
	Button []Button `json:"button"`
}

// GetMenu 获取菜单
func (r *Client) GetMenu() (info GetMenuSchema, err error) {
	var (
		accessToken string
		data        []byte
	)
	_ = util.Record(r.cache, getMenuAddr)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.HTTPGet(fmt.Sprintf(getMenuAddr, accessToken, r.agentID))
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

// DeleteMenuOptions 删除菜单请求参数
type DeleteMenuOptions struct {
}

// DeleteMenuSchema 删除菜单响应内容
type DeleteMenuSchema struct {
	util.CommonError
}

// DeleteMenu 删除菜单
func (r *Client) DeleteMenu() (info DeleteMenuSchema, err error) {
	var (
		accessToken string
		data        []byte
	)
	_ = util.Record(r.cache, deleteMenuAddr)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.HTTPGet(fmt.Sprintf(deleteMenuAddr, accessToken, r.agentID))
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
