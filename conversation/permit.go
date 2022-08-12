package conversation

import (
	"encoding/json"
	"fmt"
	"wecom/contact"
	"wecom/util"
)

const (
	// 获取企业开启会话内容存档的成员列表
	msgAuditGetPermitUserListAddr = "https://qyapi.weixin.qq.com/cgi-bin/msgaudit/get_permit_user_list?access_token=%s"
	// 获取会话中外部成员的同意情况
	msgAuditCheckSingleAgreeAddr = "https://qyapi.weixin.qq.com/cgi-bin/msgaudit/check_single_agree?access_token=%s"
)

// MsgAuditGetPermitUserListOptions 获取企业开启会话内容存档的成员列表请求参数
type MsgAuditGetPermitUserListOptions struct {
	Type int `json:"type"` // 拉取对应版本的开启成员列表。1表示办公版；2表示服务版；3表示企业版。非必填，不填写的时候返回全量成员列表。
}

// MsgAuditGetPermitUserListSchema 获取企业开启会话内容存档的成员列表响应内容
type MsgAuditGetPermitUserListSchema struct {
	util.CommonError
	Ids []string `json:"ids"`
}

// MsgAuditGetPermitUserList 获取开启了存档的用户
func (r *Client) MsgAuditGetPermitUserList(options MsgAuditGetPermitUserListOptions) (info MsgAuditGetPermitUserListSchema, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(msgAuditGetPermitUserListAddr, accessToken), options)
	if err != nil {
		return
	}
	if err = json.Unmarshal(data, &info); err != nil {
		return
	}
	if info.ErrCode != 0 {
		return info, contact.NewSDKErr(info.ErrCode, info.ErrMsg)
	}
	return info, nil
}

type AgreeInfo struct {
	Userid          string `json:"userid"`
	Exteranalopenid string `json:"exteranalopenid"`
}

// MsgCheckSingleAgreeOptions 获取会话中外部成员的同意情况请求参数
type MsgCheckSingleAgreeOptions struct {
	AgreeInfo []AgreeInfo `json:"info"`
}

type Agreeinfo struct {
	StatusChangeTime int    `json:"status_change_time"`
	Userid           string `json:"userid"`
	Exteranalopenid  string `json:"exteranalopenid"`
	AgreeStatus      string `json:"agree_status"`
}

// MsgCheckSingleAgreeSchema 获取会话中外部成员的同意情况响应内容
type MsgCheckSingleAgreeSchema struct {
	util.CommonError
	Agreeinfo []Agreeinfo `json:"agreeinfo"`
}

// MsgCheckSingleAgreeList 获取会话中外部成员的同意情况
func (r *Client) MsgCheckSingleAgreeList(options MsgCheckSingleAgreeOptions) (info MsgCheckSingleAgreeSchema, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(msgAuditCheckSingleAgreeAddr, accessToken), options)
	if err != nil {
		return
	}
	if err = json.Unmarshal(data, &info); err != nil {
		return
	}
	if info.ErrCode != 0 {
		return info, contact.NewSDKErr(info.ErrCode, info.ErrMsg)
	}
	return info, nil
}
