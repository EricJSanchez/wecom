package externalcontact

import (
	"encoding/json"
	"fmt"
	"wecom/util"
)

const (
	// ExternalcontactGetGroupChatAddr 获取配置过客户群管理的客户群列表。
	ExternalcontactGetGroupChatAddr = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/groupchat/list?access_token=%s"
	// ExternalcontactGetGroupChatDetailAddr 通过客户群ID，获取详情。包括群名、群成员列表、群成员入群时间、入群方式。（客户群是由具有客户群使用权限的成员创建的外部群）
	ExternalcontactGetGroupChatDetailAddr = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/groupchat/get?access_token=%s"
)

type OwnerFilter struct {
	UseridList []string `json:"userid_list"`
}

// ExternalcontactGetGroupChatOptions 群列表请求体
type ExternalcontactGetGroupChatOptions struct {
	StatusFilter int         `json:"status_filter"`
	OwnerFilter  OwnerFilter `json:"owner_filter"`
	Cursor       string      `json:"cursor"`
	Limit        int         `json:"limit"`
}

type GroupChatList struct {
	ChatID string `json:"chat_id"`
	Status int    `json:"status"`
}

// ExternalcontactGetGroupChatSchema 群列表返回体
type ExternalcontactGetGroupChatSchema struct {
	util.CommonError
	GroupChatList []GroupChatList `json:"group_chat_list"`
	NextCursor    string          `json:"next_cursor"`
}

//---------------------------------------------------------------

// ExternalcontactGetGroupChatDetailOptions 获取配置过客户群管理的客户群列表。
type ExternalcontactGetGroupChatDetailOptions struct {
	ChatID   string `json:"chat_id"`
	NeedName int    `json:"need_name"`
}

type Invitor struct {
	Userid string `json:"userid"`
}
type MemberList struct {
	Userid        string  `json:"userid"`
	Type          int     `json:"type"`
	JoinTime      int     `json:"join_time"`
	JoinScene     int     `json:"join_scene"`
	Invitor       Invitor `json:"invitor,omitempty"`
	GroupNickname string  `json:"group_nickname"`
	Name          string  `json:"name"`
	Unionid       string  `json:"unionid,omitempty"`
}
type AdminList struct {
	Userid string `json:"userid"`
}
type GroupChat struct {
	ChatID     string       `json:"chat_id"`
	Name       string       `json:"name"`
	Owner      string       `json:"owner"`
	CreateTime int          `json:"create_time"`
	Notice     string       `json:"notice"`
	MemberList []MemberList `json:"member_list"`
	AdminList  []AdminList  `json:"admin_list"`
}

// ExternalcontactGetGroupChatDetailSchema 获取配置过客户群管理的客户群列表返回
type ExternalcontactGetGroupChatDetailSchema struct {
	util.CommonError
	GroupChat GroupChat `json:"group_chat"`
}

// GetGroupChatList 获取群列表
func (r *Client) GetGroupChatList(options ExternalcontactGetGroupChatOptions) (info ExternalcontactGetGroupChatSchema, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(ExternalcontactGetGroupChatAddr, accessToken), options)
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

// GetGroupChatDetail 获取群列表
func (r *Client) GetGroupChatDetail(options ExternalcontactGetGroupChatDetailOptions) (info ExternalcontactGetGroupChatDetailSchema, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(ExternalcontactGetGroupChatDetailAddr, accessToken), options)
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
