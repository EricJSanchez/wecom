package contact

import (
	"encoding/json"
	"fmt"
	"github.com/EricJSanchez/wecom/util"
)

const (
	// 创建部门
	departmentCreateAddr = "https://qyapi.weixin.qq.com/cgi-bin/department/create?access_token=%s"
	// 更新部门
	departmentUpdateAddr = "https://qyapi.weixin.qq.com/cgi-bin/department/update?access_token=%s"
	// 删除部门
	departmentDeleteAddr = "https://qyapi.weixin.qq.com/cgi-bin/department/delete?access_token=%s&id=%s"
	// 获取部门列表
	departmentListAddr = "https://qyapi.weixin.qq.com/cgi-bin/department/list?access_token=%s&id=%s"
)

// DepartmentCreateOptions 创建部门请求参数
type DepartmentCreateOptions struct {
	Id       int    `json:"id,omitempty"`      // 部门id，32位整型，指定时必须大于1。若不填该参数，将自动生成id
	Name     string `json:"name"`              // 部门名称。同一个层级的部门名称不能重复。长度限制为1~32个字符，字符不能包括\:*?”<>｜
	NameEn   string `json:"name_en,omitempty"` // 英文名称。同一个层级的部门名称不能重复。需要在管理后台开启多语言支持才能生效。长度限制为1~32个字符，字符不能包括\:*?”<>｜
	Parentid int    `json:"parentid"`          // 父部门id，32位整型
	Order    int    `json:"order,omitempty"`   // 在父部门中的次序值。order值大的排序靠前。有效的值范围是[0, 2^32)
}

// DepartmentCreateSchema 添加创建部门响应内容
type DepartmentCreateSchema struct {
	util.CommonError
	Id int `json:"id"` // 新创建的部门ID
}

// DepartmentCreate 创建部门
func (r *Client) DepartmentCreate(options DepartmentCreateOptions) (info DepartmentCreateSchema, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	//fmt.Println(options)
	data, err = util.PostJSON(fmt.Sprintf(departmentCreateAddr, accessToken), options)
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

// DepartmentUpdateOptions 更新部门请求参数
type DepartmentUpdateOptions struct {
	Id       int    `json:"id"`                // 部门id，32位整型，指定时必须大于1。若不填该参数，将自动生成id
	Name     string `json:"name"`              // 部门名称。同一个层级的部门名称不能重复。长度限制为1~32个字符，字符不能包括\:*?”<>｜
	NameEn   string `json:"name_en,omitempty"` // 英文名称。同一个层级的部门名称不能重复。需要在管理后台开启多语言支持才能生效。长度限制为1~32个字符，字符不能包括\:*?”<>｜
	Parentid int    `json:"parentid"`          // 父部门id，32位整型
	Order    int    `json:"order,omitempty"`   // 在父部门中的次序值。order值大的排序靠前。有效的值范围是[0, 2^32)
}

// DepartmentUpdate 更新部门
func (r *Client) DepartmentUpdate(options DepartmentUpdateOptions) (info util.CommonError, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(departmentUpdateAddr, accessToken), options)
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

// DepartmentDeleteOptions 删除部门请求参数
type DepartmentDeleteOptions struct {
	Id string `json:"id"` // 部门id。（注：不能删除根部门；不能删除含有子部门、成员的部门）
}

// DepartmentDelete 删除客服账号
func (r *Client) DepartmentDelete(options DepartmentDeleteOptions) (info util.CommonError, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.HTTPGet(fmt.Sprintf(departmentDeleteAddr, accessToken, options.Id))
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

// DepartmentListOptions 获取部门列表请求参数
type DepartmentListOptions struct {
	Id string `json:"id"` // 部门id。获取指定部门及其下的子部门（以及及子部门的子部门等等，递归）。 如果不填，默认获取全量组织架构
}

// DepartmentInfoSchema 部门详情
type DepartmentInfoSchema struct {
	Id       int    `json:"id"`       // 部门id，32位整型，指定时必须大于1。若不填该参数，将自动生成id
	Name     string `json:"name"`     // 部门名称。同一个层级的部门名称不能重复。长度限制为1~32个字符，字符不能包括\:*?”<>｜
	NameEn   string `json:"media_id"` // 英文名称。同一个层级的部门名称不能重复。需要在管理后台开启多语言支持才能生效。长度限制为1~32个字符，字符不能包括\:*?”<>｜
	Parentid int    `json:"parentid"` // 父部门id，32位整型
	Order    int    `json:"order"`    // 在父部门中的次序值。order值大的排序靠前。有效的值范围是[0, 2^32)
}

// DepartmentListSchema 获取部门列表响应内容
type DepartmentListSchema struct {
	util.CommonError
	DepartmentList []DepartmentInfoSchema `json:"department"` // 部门列表
}

// DepartmentList 获取部门列表
func (r *Client) DepartmentList(options DepartmentListOptions) (info DepartmentListSchema, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.HTTPGet(fmt.Sprintf(departmentListAddr, accessToken, options.Id))
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
