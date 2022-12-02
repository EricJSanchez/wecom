package application

import (
	"encoding/json"
	"fmt"
	"github.com/EricJSanchez/wecom/util"
)

const (
	// 获取指定的应用详情
	getApplicationInfoAddr = "https://qyapi.weixin.qq.com/cgi-bin/agent/get?access_token=%s&agentid=%s"
	// 获取access_token对应的应用列表
	getApplicationListAddr = "https://qyapi.weixin.qq.com/cgi-bin/agent/list?access_token=%s"
	// 设置指定的应用详情
	setApplicationInfoAddr = "https://qyapi.weixin.qq.com/cgi-bin/agent/set?access_token=%s"
	// 设置工作台模板
	setWorkbenchTemplateAddr = "https://qyapi.weixin.qq.com/cgi-bin/agent/set_workbench_template?access_token=%s"
)

type User struct {
	Userid string `json:"userid"`
}
type AllowUserinfos struct {
	User []User `json:"user"`
}
type AllowPartys struct {
	Partyid []int `json:"partyid"`
}
type AllowTags struct {
	Tagid []int `json:"tagid"`
}

// GetApplicationInfoOptions 获取指定的应用详情请求参数
type GetApplicationInfoOptions struct {
}

// GetApplicationInfoSchema 获取指定的应用详情响应内容
type GetApplicationInfoSchema struct {
	util.CommonError
	Agentid            int            `json:"agentid"`
	Name               string         `json:"name"`
	SquareLogoURL      string         `json:"square_logo_url"`
	Description        string         `json:"description"`
	AllowUserinfos     AllowUserinfos `json:"allow_userinfos"`
	AllowPartys        AllowPartys    `json:"allow_partys"`
	AllowTags          AllowTags      `json:"allow_tags"`
	Close              int            `json:"close"`
	RedirectDomain     string         `json:"redirect_domain"`
	ReportLocationFlag int            `json:"report_location_flag"`
	Isreportenter      int            `json:"isreportenter"`
	HomeURL            string         `json:"home_url"`
}

// GetApplicationInfo 获取指定的应用详情
func (r *Client) GetApplicationInfo() (info GetApplicationInfoSchema, err error) {
	var (
		accessToken string
		data        []byte
	)
	_ = util.Record(r.cache, getApplicationInfoAddr)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.HTTPGet(fmt.Sprintf(getApplicationInfoAddr, accessToken, r.agentID))
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

type Agentlist struct {
	Agentid       int    `json:"agentid"`
	Name          string `json:"name"`
	SquareLogoURL string `json:"square_logo_url"`
}

// GetApplicationListOptions 获取access_token对应的应用列表请求参数
type GetApplicationListOptions struct {
}

// GetApplicationListSchema 获取access_token对应的应用列表响应内容
type GetApplicationListSchema struct {
	util.CommonError
	Agentlist []Agentlist `json:"agentlist"`
}

// GetApplicationList 获取access_token对应的应用列表
func (r *Client) GetApplicationList() (info GetApplicationListSchema, err error) {
	var (
		accessToken string
		data        []byte
	)
	_ = util.Record(r.cache, getApplicationListAddr)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.HTTPGet(fmt.Sprintf(getApplicationListAddr, accessToken))
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

// SetApplicationInfoOptions 设置应用详情请求参数
type SetApplicationInfoOptions struct {
	Agentid            int    `json:"agentid"`
	ReportLocationFlag int    `json:"report_location_flag"`
	LogoMediaid        string `json:"logo_mediaid"`
	Name               string `json:"name"`
	Description        string `json:"description"`
	RedirectDomain     string `json:"redirect_domain"`
	Isreportenter      int    `json:"isreportenter"`
	HomeURL            string `json:"home_url"`
}

// SetApplicationInfoSchema 设置应用详情响应内容
type SetApplicationInfoSchema struct {
	util.CommonError
}

// SetApplicationInfo 设置应用详情
func (r *Client) SetApplicationInfo(options SetApplicationInfoOptions) (info SetApplicationInfoSchema, err error) {
	var (
		accessToken string
		data        []byte
	)
	_ = util.Record(r.cache, setApplicationInfoAddr)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(setApplicationInfoAddr, accessToken), options)
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
