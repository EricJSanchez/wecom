package externalcontact

import (
	"encoding/json"
	"fmt"
	"github.com/EricJSanchez/wecom/util"
)

const (
	// 获取待分配的离职成员列表
	externalContactGetUnassignedList = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_unassigned_list?access_token=%s"
	//分配离职成员的客户
	externalContactResignedTransferCustomer = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/resigned/transfer_customer?access_token=%s"
	//查询客户接替状态
	externalContactResignedTransferResult = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/resigned/transfer_result?access_token=%s"
	//分配离职成员的客户群
	externalContactGroupChatTransfer = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/groupchat/transfer?access_token=%s"

	//分配在职成员的客户
	externalContactTransferCustomer = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/transfer_customer?access_token=?access_token=%s"
	//查询客户接替状态
	externalContactTransferResult = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/transfer_result?access_token=%s"
	//查询客户接替状态
	externalContactGroupChatOnJobTransfer = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/groupchat/onjob_transfer??access_token=%s"
)

//ExternalContactGetUnassignedListOptions
type ExternalContactGetUnassignedListOptions struct {
	PageId   int    `json:"page_id"`   // 	分页查询，要查询页号，从0开始
	PageSize int    `json:"page_size"` //	每次返回的最大记录数，默认为1000，最大值为1000
	Cursor   string `json:"cursor"`    //分页查询游标，字符串类型，适用于数据量较大的情况，如果使用该参数则无需填写page_id，该参数由上一次调用返回
}

type ExternalContactGetUnassignedListSchema struct {
	util.CommonError
	IsLast                           bool                               `json:"is_last"` // 	是否是最后一条记录
	ExternalContactGetUnassignedInfo []ExternalContactGetUnassignedInfo `json:"info"`
	NextCursor                       string                             `json:"next_cursor"` //分页查询游标,已经查完则返回空("")，使用page_id作为查询参数时不返回
}
type ExternalContactGetUnassignedInfo struct {
	HandoverUserid string `json:"handover_userid"` //离职成员的userid
	ExternalUserid string `json:"external_userid"` // 	外部联系人userid
	DimissionTime  int    `json:"dimission_time"`  //成员离职时间
}

/**
当page_id为1，page_size为100时，表示取第101到第200条记录。
page_id和page_size参数仅适用于记录数小于五万条的情况,即 page_id*page_size < 50000；
如果记录数大于五万，则需要使用cursor参数。
*/

//GetUnassignedList 获取待分配的离职成员列表
func (r *Client) GetUnassignedList(options ExternalContactGetUnassignedListOptions) (info ExternalContactGetUnassignedListSchema, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(externalContactGetUnassignedList, accessToken), options)
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

type ExternalContactResignedTransferCustomerOptions struct {
	HandoverUserid string   `json:"handover_userid"` //	原跟进成员的userid
	TakeoverUserid string   `json:"takeover_userid"` //	接替成员的userid
	ExternalUserid []string `json:"external_userid"` //	客户的external_userid列表，最多一次转移100个客户
}

type ExternalContactTransferCustomerSchema struct {
	util.CommonError
	ExternalContactTransferCustomerInfo []ExternalContactTransferCustomerInfo `json:"customer"`
}

type ExternalContactTransferCustomerInfo struct {
	ExternalUserid string `json:"external_userid"` // 	外部联系人userid
	ErrCode        int    `json:"errcode"`         //对此客户进行分配的结果, 具体可参考全局错误码, 0表示开始分配流程,待24小时后自动接替,并不代表最终分配成功
}

// ResignedTransferCustomer 分配离职成员的客户
func (r *Client) ResignedTransferCustomer(options ExternalContactResignedTransferCustomerOptions) (info ExternalContactTransferCustomerSchema, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(externalContactResignedTransferCustomer, accessToken), options)
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

type ExternalContactResignedTransferResultOptions struct {
	HandoverUserid string `json:"handover_userid"` //	原跟进成员的userid
	TakeoverUserid string `json:"takeover_userid"` //	接替成员的userid
	Cursor         string `json:"cursor"`          //分页查询游标，字符串类型，适用于数据量较大的情况，如果使用该参数则无需填写page_id，该参数由上一次调用返回
}

type ExternalContactResignedTransferResultSchema struct {
	util.CommonError
	ExternalContactResignedTransferResultInfo []ExternalContactResignedTransferResultInfo `json:"customer"`
	NextCursor                                string                                      `json:"next_cursor"` //分页查询游标,已经查完则返回空("")，使用page_id作为查询参数时不返回
}

type ExternalContactResignedTransferResultInfo struct {
	ExternalUserid string `json:"external_userid"` // 	外部联系人userid
	Status         int    `json:"status"`          //接替状态， 1-接替完毕 2-等待接替 3-客户拒绝 4-接替成员客户达到上限
	TakeoverTime   int    `json:"takeover_time"`   //	接替客户的时间，如果是等待接替状态，则为未来的自动接替时间
}

//ResignedTransferResult 查询客户接替状态
func (r *Client) ResignedTransferResult(options ExternalContactResignedTransferResultOptions) (info ExternalContactResignedTransferResultSchema, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(externalContactResignedTransferResult, accessToken), options)
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

type ExternalContactGroupChatTransferOptions struct {
	ChatIdList []string `json:"chat_id_list"` //需要转群主的客户群ID列表。取值范围： 1 ~ 100
	NewOwner   string   `json:"new_owner"`    //新群主ID
}

type ExternalContactGroupChatTransferSchema struct {
	util.CommonError
	ExternalContactFailedChatList []ExternalContactFailedChatList `json:"failed_chat_list"`
}

type ExternalContactFailedChatList struct {
	util.CommonError
	ChatId string `json:"chat_id"` //新群主ID
}

/**
群主离职了的客户群，才可继承
继承给的新群主，必须是配置了客户联系功能的成员
继承给的新群主，必须有设置实名
继承给的新群主，必须有激活企业微信
同一个人的群，限制每天最多分配300个给新群主
*/

//GroupChatTransfer 分配离职成员的客户群
func (r *Client) GroupChatTransfer(options ExternalContactGroupChatTransferOptions) (info ExternalContactGroupChatTransferSchema, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(externalContactGroupChatTransfer, accessToken), options)
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

type ExternalContactTransferCustomerOptions struct {
	HandoverUserid     string   `json:"handover_userid"`      //	原跟进成员的userid
	TakeoverUserid     string   `json:"takeover_userid"`      //	接替成员的userid
	ExternalUserid     []string `json:"external_userid"`      //	客户的external_userid列表，最多一次转移100个客户 external_userid必须是handover_userid的客户（即配置了客户联系功能的成员所添加的联系人）。
	TransferSuccessMsg string   `json:"transfer_success_msg"` //移成功后发给客户的消息，最多200个字符，不填则使用默认文案
}

//为保障客户服务体验，90个自然日内，在职成员的每位客户仅可被转接2次。

//TransferCustomer 分配在职成员的客户
func (r *Client) TransferCustomer(options ExternalContactTransferCustomerOptions) (info ExternalContactTransferCustomerSchema, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(externalContactTransferCustomer, accessToken), options)
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

//TransferResult 查询客户接替状态 ---在职状态
func (r *Client) TransferResult(options ExternalContactResignedTransferResultOptions) (info ExternalContactResignedTransferResultSchema, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(externalContactTransferResult, accessToken), options)
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

//GroupChatOnJobTransfer 分配在职成员的客户群
func (r *Client) GroupChatOnJobTransfer(options ExternalContactGroupChatTransferOptions) (info ExternalContactGroupChatTransferSchema, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(externalContactGroupChatOnJobTransfer, accessToken), options)
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
