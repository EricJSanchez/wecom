package externalcontact

import (
	"encoding/json"
	"fmt"
	"github.com/EricJSanchez/wecom/util"
)

const (
	//获取规则组列表
	externalcontactMomentStrategyListAddr = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/moment_strategy/list?access_token=%s"
	//获取规则组详情
	externalcontactMomentStrategyGetAddr = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/moment_strategy/get?access_token=%s"
	//获取规则组管理范围
	externalcontactMomentStrategyGetRangeAddr = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/moment_strategy/get_range?access_token=%s"
	//创建新的规则组
	externalcontactMomentStrategyCreateAddr = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/moment_strategy/create?access_token=%s"
	//编辑规则组及其管理范围
	externalcontactMomentStrategyEditAddr = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/moment_strategy/edit?access_token=%s"
	//删除规则组
	externalcontactMomentStrategyDelAddr = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/moment_strategy/del?access_token=%s"
)

// ExternalcontactMomentStrategyListOptions 获取规则组列表请求参数
type ExternalcontactMomentStrategyListOptions struct {
	Cursor string `json:"cursor"`
	Limit  int    `json:"limit"`
}

// ExternalcontactCustomerStrategyListSchema 获取规则组列表响应内容
type ExternalcontactMomentStrategyListSchema struct {
	util.CommonError
	Errcode    int            `json:"errcode"`
	Errmsg     string         `json:"errmsg"`
	Strategy   []StrategyInfo `json:"strategy"`
	NextCursor string         `json:"next_cursor"`
}

// ExternalcontactMomentStrategyList 获取规则组列表
func (r *Client) ExternalcontactMomentStrategyList(options ExternalcontactMomentStrategyListOptions) (info ExternalcontactMomentStrategyListSchema, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(externalcontactMomentStrategyListAddr, accessToken), options)
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

// ExternalcontactMomentStrategyGetOptions 获取规则组详情请求参数
type ExternalcontactMomentStrategyGetOptions struct {
	StrategyID int `json:"strategy_id"`
}

// ExternalcontactMomentStrategyGetSchema 获取规则组详情响应内容
type ExternalcontactMomentStrategyGetSchema struct {
	util.CommonError
	MomentStrategy MomentStrategy `json:"strategy"`
}

// MomentStrategy 规则组详情
type MomentStrategy struct {
	StrategyID      int             `json:"strategy_id"`
	ParentID        int             `json:"parent_id"`
	StrategyName    string          `json:"strategy_name"`
	CreateTime      int             `json:"create_time"`
	AdminList       []string        `json:"admin_list"`
	MomentPrivilege MomentPrivilege `json:"privilege"`
}

type MomentPrivilege struct {
	ViewMomentList           bool `json:"view_moment_list"`
	SendMoment               bool `json:"send_moment"`
	ManageMomentCoverAndSign bool `json:"manage_moment_cover_and_sign"`
}

// ExternalcontactCustomerStrategyGet 获取规则组详情
func (r *Client) ExternalcontactMomentStrategyGet(options ExternalcontactMomentStrategyGetOptions) (info ExternalcontactMomentStrategyGetSchema, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(externalcontactMomentStrategyGetAddr, accessToken), options)
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

// ExternalcontactMomentStrategyGetRangeOptions 获取规则组管理范围请求参数
type ExternalcontactMomentStrategyGetRangeOptions struct {
	StrategyID int    `json:"strategy_id"`
	Cursor     string `json:"cursor"`
	Limit      int    `json:"limit"`
}

// ExternalcontactMomentStrategyGetRangeSchema 获取规则组管理范围响应内容
type ExternalcontactMomentStrategyGetRangeSchema struct {
	util.CommonError
	Errcode    int     `json:"errcode"`
	Errmsg     string  `json:"errmsg"`
	Range      []Range `json:"range"`
	NextCursor string  `json:"next_cursor"`
}

// ExternalcontactMomentStrategyGetRange 获取规则组管理范围
func (r *Client) ExternalcontactMomentStrategyGetRange(options ExternalcontactMomentStrategyGetRangeOptions) (info ExternalcontactMomentStrategyGetRangeSchema, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(externalcontactMomentStrategyGetRangeAddr, accessToken), options)
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

// ExternalcontactMomentStrategyCreateOptions 创建新的规则组请求参数
type ExternalcontactMomentStrategyCreateOptions struct {
	ParentID     int             `json:"parent_id"`
	StrategyName string          `json:"strategy_name"`
	AdminList    []string        `json:"admin_list"`
	Privilege    MomentPrivilege `json:"privilege"`
	Range        []Range         `json:"range"`
}

// ExternalcontactCustomerStrategyCreateSchema 创建新的规则组响应内容
type ExternalcontactMomentStrategyCreateSchema struct {
	util.CommonError
	StrategyID int `json:"strategy_id"`
}

// ExternalcontactCustomerStrategyCreate 创建新的规则组
func (r *Client) ExternalcontactMomentStrategyCreate(options ExternalcontactMomentStrategyCreateOptions) (info ExternalcontactMomentStrategyCreateSchema, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	fmt.Println("options", options)
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(externalcontactMomentStrategyCreateAddr, accessToken), options)
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

// ExternalcontactCustomerStrategyEditOptions 编辑规则组及其管理范围请求参数
type ExternalcontactMomentStrategyEditOptions struct {
	StrategyID   int             `json:"strategy_id"`
	StrategyName string          `json:"strategy_name"`
	AdminList    []string        `json:"admin_list"`
	Privilege    MomentPrivilege `json:"privilege"`
	RangeAdd     []Range         `json:"range_add"`
	RangeDel     []Range         `json:"range_del"`
}

// ExternalcontactCustomerStrategyEdit 编辑规则组及其管理范围
func (r *Client) ExternalcontactMomentStrategyEdit(options ExternalcontactMomentStrategyEditOptions) (info util.CommonError, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(externalcontactMomentStrategyEditAddr, accessToken), options)
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

// ExternalcontactCustomerStrategyDelOptions 删除规则组请求参数
type ExternalcontactMomentStrategyDelOptions struct {
	StrategyID int `json:"strategy_id"`
}

// ExternalcontactCustomerStrategyDel 删除规则组
func (r *Client) ExternalcontactMomentStrategyDel(options ExternalcontactMomentStrategyDelOptions) (info util.CommonError, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(externalcontactMomentStrategyDelAddr, accessToken), options)
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
