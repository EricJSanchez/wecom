package externalcontact

import (
	"encoding/json"
	"fmt"
	"github.com/EricJSanchez/wecom/util"
	"github.com/spf13/cast"
)

const (
	//获取企业标签库
	externalContactGetTagListAddr = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_corp_tag_list?access_token=%s"
	//添加企业标签以及标签组
	externalContactAddTagListAddr = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/add_corp_tag?access_token=%s"
	//编辑企业标签以及标签组
	externalContactEditTagListAddr = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/edit_corp_tag?access_token=%s"
	//删除企业标签以及标签组
	externalContactDelTagListAddr = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/del_corp_tag?access_token=%s"
	//给用户添加或者删除企业标签
	externalContactMarkTagListAddr = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/mark_tag?access_token=%s"
)

// ExternalcontactGetTagListOptions 获取企业标签库请求参数
type ExternalcontactGetTagListOptions struct {
	TagID   []string `json:"tag_id"`   // 要查询的标签id
	GroupID []string `json:"group_id"` // 要查询的标签组id，返回该标签组以及其下的所有标签信息
}

type Tag struct {
	ID         string `json:"id"`          // 标签id
	Name       string `json:"name"`        // 标签名称
	CreateTime int    `json:"create_time"` // 标签创建时间
	Order      int    `json:"order"`       // 标签排序的次序值，order值大的排序靠前。有效的值范围是[0, 2^32)
	Deleted    bool   `json:"deleted"`     // 标签是否已经被删除，只在指定tag_id/group_id进行查询时返回
}

type TagGroupChild struct {
	GroupID    string `json:"group_id"`    // 标签组id
	GroupName  string `json:"group_name"`  // 标签组名称
	CreateTime int    `json:"create_time"` // 标签组创建时间
	Order      int    `json:"order"`       // 标签组排序的次序值，order值大的排序靠前。有效的值范围是[0, 2^32)
	Deleted    bool   `json:"deleted"`     // 标签组是否已经被删除，只在指定tag_id进行查询时返回
	Tag        []Tag  `json:"tag"`         // 标签组内的标签列表
}

// ExternalcontactGetTagListSchema  获取企业标签库响应内容
type ExternalcontactGetTagListSchema struct {
	util.CommonError
	TagGroup []TagGroupChild `json:"tag_group"` // 标签组列表
}

// ExternalcontactAddTagWithGroupListOptions  添加企业标签
type ExternalcontactAddTagWithGroupListOptions struct {
	GroupID   string                             `json:"group_id,omitempty"`   // 要查询的标签组id，返回该标签组以及其下的所有标签信息
	GroupName string                             `json:"group_name,omitempty"` // 标签组名称
	Order     int                                `json:"order"`                // 标签组排序的次序值，order值大的排序靠前。有效的值范围是[0, 2^32)
	Tag       []ExternalcontactAddTagListOptions `json:"tag"`
}

// ExternalcontactAddTagListOptions  添加企业标签
type ExternalcontactAddTagAndGroupListOptions struct {
	GroupName string                             `json:"group_name"` // 标签组名称
	Order     int                                `json:"order"`      // 标签组排序的次序值，order值大的排序靠前。有效的值范围是[0, 2^32)
	Tag       []ExternalcontactAddTagListOptions `json:"tag"`
}

type ExternalcontactAddTagListOptions struct {
	TagName string `json:"name"`  //添加的标签名称，最长为30个字符
	Order   int    `json:"order"` // 标签组排序的次序值，order值大的排序靠前。有效的值范围是[0, 2^32)
}

type ExternalcontactAddTagAndGroupListSchema struct {
	util.CommonError
	TagGroup TagGroup `json:"tag_group"` // 标签组列表
}
type TagGroup struct {
	GroupID    string   `json:"group_id"`    // 标签组id
	GroupName  string   `json:"group_name"`  // 标签组名称
	CreateTime int      `json:"create_time"` // 标签组创建时间
	Tag        []AddTag `json:"tag"`         // 标签组内的标签列表
}

type AddTag struct {
	ID         string `json:"id"`          // 标签id
	Name       string `json:"name"`        // 标签名称
	CreateTime int    `json:"create_time"` // 标签创建时间
	Order      int    `json:"order"`       // 标签排序的次序值，order值大的排序靠前。有效的值范围是[0, 2^32)
}

type ExternalcontactEditTagAndGroupListOptions struct {
	ID    string `json:"id"`    // 标签/标签组id
	Name  string `json:"name"`  // 标签/标签组名称
	Order int    `json:"order"` // 标签/标签组排序的次序值，order值大的排序靠前。有效的值范围是[0, 2^32)
}

type ExternalcontactDelTagAndGroupListOptions struct {
	TagID   []string `json:"tag_id"`   // 要查询的标签id
	GroupID []string `json:"group_id"` // 要查询的标签组id，返回该标签组以及其下的所有标签信息
}

// ExternalcontactMarkTagListOptions 请确保external_userid是userid的外部联系人。 add_tag和remove_tag不可同时为空。 同一个标签组下现已支持多个标签
type ExternalcontactMarkTagListOptions struct {
	UserId         string   `json:"userid"`          // 添加外部联系人的userid
	ExternalUserid string   `json:"external_userid"` // 外部联系人userid
	AddTag         []string `json:"add_tag"`         // 要标记的标签列表
	RemoveTag      []string `json:"remove_tag"`      //要移除的标签列表
}

// GetTagList 获取企业标签库列表
func (r *Client) GetTagList(options ExternalcontactGetTagListOptions) (info ExternalcontactGetTagListSchema, err error) {
	var (
		accessToken string
		data        []byte
	)
	_ = util.Record(r.cache, externalContactGetTagListAddr)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(externalContactGetTagListAddr, accessToken), options)
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

// AddTagList 添加企业客户标签
func (r *Client) AddTagList(options ExternalcontactAddTagWithGroupListOptions) (info ExternalcontactAddTagAndGroupListSchema, err error) {
	var (
		accessToken string
		data        []byte
	)
	_ = util.Record(r.cache, externalContactAddTagListAddr)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(externalContactAddTagListAddr, accessToken), options)
	if err != nil {
		return
	}
	fmt.Println(cast.ToString(data))
	if err = json.Unmarshal(data, &info); err != nil {
		return
	}
	if info.ErrCode != 0 {
		return info, NewSDKErr(info.ErrCode, info.ErrMsg)
	}
	return info, nil
}

// AddTagListAndGroup 添加企业客户标签和标签组
func (r *Client) AddTagListAndGroup(options ExternalcontactAddTagAndGroupListOptions) (info ExternalcontactAddTagAndGroupListSchema, err error) {
	var (
		accessToken string
		data        []byte
	)
	_ = util.Record(r.cache, externalContactAddTagListAddr)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(externalContactAddTagListAddr, accessToken), options)
	if err != nil {
		return
	}
	fmt.Println(cast.ToString(data))
	if err = json.Unmarshal(data, &info); err != nil {
		return
	}
	if info.ErrCode != 0 {
		return info, NewSDKErr(info.ErrCode, info.ErrMsg)
	}
	return info, nil
}

// EditTagListOrGroup 修改企业客户标签
func (r *Client) EditTagListOrGroup(options ExternalcontactEditTagAndGroupListOptions) (info util.CommonError, err error) {
	var (
		accessToken string
		data        []byte
	)
	_ = util.Record(r.cache, externalContactEditTagListAddr)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(externalContactEditTagListAddr, accessToken), options)
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

func (r *Client) DeleteTagListOrGroup(options ExternalcontactDelTagAndGroupListOptions) (info util.CommonError, err error) {
	var (
		accessToken string
		data        []byte
	)
	_ = util.Record(r.cache, externalContactDelTagListAddr)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(externalContactDelTagListAddr, accessToken), options)
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

func (r *Client) MarkTagListExternalContact(options ExternalcontactMarkTagListOptions) (info util.CommonError, err error) {
	var (
		accessToken string
		data        []byte
	)
	_ = util.Record(r.cache, externalContactMarkTagListAddr)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(externalContactMarkTagListAddr, accessToken), options)
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
