package externalcontact

import (
	"encoding/json"
	"fmt"
	"github.com/EricJSanchez/wecom/util"
)

const (
	// 获取客户列表
	externalcontactListAddr = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/list?access_token=%s&userid=%s"
	// 获取客户详情
	externalcontactGetAddr = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get?access_token=%s&external_userid=%s&cursor=%s"
	// 批量获取客户详情
	externalcontactBatchGetByUserAddr = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/batch/get_by_user?access_token=%s"
	// 修改客户备注信息
	externalcontactRemarkAddr = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/remark?access_token=%s"
	//获取配置了客户联系功能的成员列表
	externalcontactGetFollowUserListAddr = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_follow_user_list?access_token=%s"
	// 获取规则组列表
	externalcontactCustomerStrategyListAddr = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/customer_strategy/list?access_token=%s"
	// 获取规则组详情
	externalcontactCustomerStrategyGetAddr = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/customer_strategy/get?access_token=%s"
	// 获取规则组管理范围
	externalcontactCustomerStrategyGetRangeAddr = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/customer_strategy/get_range?access_token=%s"
	// 创建新的规则组
	externalcontactCustomerStrategyCreateAddr = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/customer_strategy/create?access_token=%s"
	// 编辑规则组及其管理范围
	externalcontactCustomerStrategyEditAddr = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/customer_strategy/edit?access_token=%s"
	// 删除规则组
	externalcontactCustomerStrategyDelAddr = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/customer_strategy/del?access_token=%s"
)

// ExternalcontactListOptions 获取客户列表请求参数
type ExternalcontactListOptions struct {
	Userid string `json:"userid"` // 企业成员的userid
}

// ExternalcontactListSchema 获取客户列表响应内容
type ExternalcontactListSchema struct {
	util.CommonError
	ExternalUserid []string `json:"external_userid"` // 外部联系人的userid列表
}

// ExternalcontactList 获取客户列表
func (r *Client) ExternalcontactList(options ExternalcontactListOptions) (info ExternalcontactListSchema, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.HTTPGet(fmt.Sprintf(externalcontactListAddr, accessToken, options.Userid))
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

// ExternalcontactGetOptions 请求参数
type ExternalcontactGetOptions struct {
	ExternalUserid string `json:"external_userid"` // 外部联系人的userid，注意不是企业成员的帐号
	Cursor         string `json:"cursor"`          // 上次请求返回的next_cursor
}

// ExternalContact 外部联系人信息
type ExternalContact struct {
	ExternalUserid  string                 `json:"external_userid"`
	Name            string                 `json:"name"`
	Position        string                 `json:"position"`
	Avatar          string                 `json:"avatar"`
	CorpName        string                 `json:"corp_name"`
	CorpFullName    string                 `json:"corp_full_name"`
	Type            int                    `json:"type"`
	Gender          int                    `json:"gender"`
	Unionid         string                 `json:"unionid"`
	ExternalProfile map[string]interface{} `json:"external_profile"`
}

// Tags 外部联系人标签信息
type Tags struct {
	GroupName string `json:"group_name"`
	TagName   string `json:"tag_name"`
	TagID     string `json:"tag_id,omitempty"`
	Type      int    `json:"type"`
}

// FollowUser 添加了此外部联系人的企业成员信息
type FollowUser struct {
	Userid         string   `json:"userid"`
	Remark         string   `json:"remark"`
	Description    string   `json:"description"`
	Createtime     int      `json:"createtime"`
	Tags           []Tags   `json:"tags,omitempty"`
	RemarkCorpName string   `json:"remark_corp_name,omitempty"`
	RemarkMobiles  []string `json:"remark_mobiles,omitempty"`
	OperUserid     string   `json:"oper_userid"`
	AddWay         int      `json:"add_way"`
	State          string   `json:"state,omitempty"`
}

// ExternalcontactGetSchema 获取客户详情响应内容
type ExternalcontactGetSchema struct {
	util.CommonError
	ExternalContact ExternalContact `json:"external_contact"`
	FollowUser      []FollowUser    `json:"follow_user"`
	NextCursor      string          `json:"next_cursor"`
}

// ExternalcontactGet 获取客户详情
func (r *Client) ExternalcontactGet(options ExternalcontactGetOptions) (info ExternalcontactGetSchema, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.HTTPGet(fmt.Sprintf(externalcontactGetAddr, accessToken, options.ExternalUserid, options.Cursor))
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

// 批量获取客户详情请求参数
type ExternalcontactBatchGetByUserOptions struct {
	UseridList []string `json:"userid_list"`
	Cursor     string   `json:"cursor"`
	Limit      int      `json:"limit"`
}

// FollowInfo 添加了此外部联系人的企业成员信息
type FollowInfo struct {
	Userid        string   `json:"userid"`
	Remark        string   `json:"remark"`
	Description   string   `json:"description"`
	Createtime    int      `json:"createtime"`
	TagID         []string `json:"tag_id"`
	State         string   `json:"state"`
	RemarkMobiles []string `json:"remark_mobiles"`
	OperUserid    string   `json:"oper_userid"`
	AddWay        int      `json:"add_way"`
}

// ExternalContactList 外部联系人列表信息
type ExternalContactList struct {
	ExternalContact ExternalContact `json:"external_contact,omitempty"`
	FollowInfo      FollowInfo      `json:"follow_info,omitempty"`
}

// ExternalcontactBatchGetByUserSchema 批量获取客户详情响应内容
type ExternalcontactBatchGetByUserSchema struct {
	util.CommonError
	ExternalContactList []ExternalContactList `json:"external_contact_list"`
	NextCursor          string                `json:"next_cursor"`
}

// ExternalcontactBatchGetByUser 批量获取客户详情
func (r *Client) ExternalcontactBatchGetByUser(options ExternalcontactBatchGetByUserOptions) (info ExternalcontactBatchGetByUserSchema, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(externalcontactBatchGetByUserAddr, accessToken), options)
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

// ExternalcontactRemarkOptions 修改客户备注信息请求参数
type ExternalcontactRemarkOptions struct {
	Userid           string   `json:"userid"`
	ExternalUserid   string   `json:"external_userid"`
	Remark           string   `json:"remark"`
	Description      string   `json:"description"`
	RemarkCompany    string   `json:"remark_company"`
	RemarkMobiles    []string `json:"remark_mobiles"`
	RemarkPicMediaid string   `json:"remark_pic_mediaid"`
}

// ExternalcontactRemark 修改客户备注信息
func (r *Client) ExternalcontactRemark(options ExternalcontactRemarkOptions) (info util.CommonError, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(externalcontactRemarkAddr, accessToken), options)
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

type ExternalcontactGetFollowUserListSchema struct {
	util.CommonError
	Errcode    int      `json:"errcode"`
	Errmsg     string   `json:"errmsg"`
	FollowUser []string `json:"follow_user"`
}

//externalcontactGetFollowUserListAddr 获取配置了客户联系功能的成员列表
func (r *Client) ExternalcontactGetFollowUserList() (info ExternalcontactGetFollowUserListSchema, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.HTTPGet(fmt.Sprintf(externalcontactGetFollowUserListAddr, accessToken))
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

// ExternalcontactCustomerStrategyListOptions 获取规则组列表请求参数
type ExternalcontactCustomerStrategyListOptions struct {
	Cursor string `json:"cursor"`
	Limit  int    `json:"limit"`
}

// StrategyInfo	规则组信息
type StrategyInfo struct {
	StrategyID int `json:"strategy_id"`
}

// ExternalcontactCustomerStrategyListSchema 获取规则组列表响应内容
type ExternalcontactCustomerStrategyListSchema struct {
	util.CommonError
	Errcode    int            `json:"errcode"`
	Errmsg     string         `json:"errmsg"`
	Strategy   []StrategyInfo `json:"strategy"`
	NextCursor string         `json:"next_cursor"`
}

// ExternalcontactCustomerStrategyList 获取规则组列表
func (r *Client) ExternalcontactCustomerStrategyList(options ExternalcontactCustomerStrategyListOptions) (info ExternalcontactCustomerStrategyListSchema, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(externalcontactCustomerStrategyListAddr, accessToken), options)
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

// ExternalcontactCustomerStrategyGetOptions 获取规则组详情请求参数
type ExternalcontactCustomerStrategyGetOptions struct {
	StrategyID int `json:"strategy_id"`
}

// Privilege 权限信息
type Privilege struct {
	ViewCustomerList        bool `json:"view_customer_list"`
	ViewCustomerData        bool `json:"view_customer_data"`
	ViewRoomList            bool `json:"view_room_list"`
	ContactMe               bool `json:"contact_me"`
	JoinRoom                bool `json:"join_room"`
	ShareCustomer           bool `json:"share_customer"`
	OperResignCustomer      bool `json:"oper_resign_customer"`
	OperResignGroup         bool `json:"oper_resign_group"`
	SendCustomerMsg         bool `json:"send_customer_msg"`
	EditWelcomeMsg          bool `json:"edit_welcome_msg"`
	ViewBehaviorData        bool `json:"view_behavior_data"`
	ViewRoomData            bool `json:"view_room_data"`
	SendGroupMsg            bool `json:"send_group_msg"`
	RoomDeduplication       bool `json:"room_deduplication"`
	RapidReply              bool `json:"rapid_reply"`
	OnjobCustomerTransfer   bool `json:"onjob_customer_transfer"`
	EditAntiSpamRule        bool `json:"edit_anti_spam_rule"`
	ExportCustomerList      bool `json:"export_customer_list"`
	ExportCustomerData      bool `json:"export_customer_data"`
	ExportCustomerGroupList bool `json:"export_customer_group_list"`
	ManageCustomerTag       bool `json:"manage_customer_tag"`
}

// Strategy 规则组详情
type Strategy struct {
	StrategyID   int       `json:"strategy_id"`
	ParentID     int       `json:"parent_id"`
	StrategyName string    `json:"strategy_name"`
	CreateTime   int       `json:"create_time"`
	AdminList    []string  `json:"admin_list"`
	Privilege    Privilege `json:"privilege"`
}

// ExternalcontactCustomerStrategyGetSchema 获取规则组详情响应内容
type ExternalcontactCustomerStrategyGetSchema struct {
	util.CommonError
	Strategy Strategy `json:"strategy"`
}

// ExternalcontactCustomerStrategyGet 获取规则组详情
func (r *Client) ExternalcontactCustomerStrategyGet(options ExternalcontactCustomerStrategyGetOptions) (info ExternalcontactCustomerStrategyGetSchema, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(externalcontactCustomerStrategyGetAddr, accessToken), options)
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

// ExternalcontactCustomerStrategyGetRangeOptions 获取规则组管理范围请求参数
type ExternalcontactCustomerStrategyGetRangeOptions struct {
	StrategyID int    `json:"strategy_id"`
	Cursor     string `json:"cursor"`
	Limit      int    `json:"limit"`
}

// Range 范围信息
type Range struct {
	Type    int    `json:"type"`
	Userid  string `json:"userid,omitempty"`
	Partyid int    `json:"partyid,omitempty"`
}

// ExternalcontactCustomerStrategyGetRangeSchema 获取规则组管理范围响应内容
type ExternalcontactCustomerStrategyGetRangeSchema struct {
	util.CommonError
	Errcode    int     `json:"errcode"`
	Errmsg     string  `json:"errmsg"`
	Range      []Range `json:"range"`
	NextCursor string  `json:"next_cursor"`
}

// ExternalcontactCustomerStrategyGetRange 获取规则组管理范围
func (r *Client) ExternalcontactCustomerStrategyGetRange(options ExternalcontactCustomerStrategyGetRangeOptions) (info ExternalcontactCustomerStrategyGetRangeSchema, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(externalcontactCustomerStrategyGetRangeAddr, accessToken), options)
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

// ExternalcontactCustomerStrategyCreateOptions 创建新的规则组请求参数
type ExternalcontactCustomerStrategyCreateOptions struct {
	ParentID     int       `json:"parent_id"`
	StrategyName string    `json:"strategy_name"`
	AdminList    []string  `json:"admin_list"`
	Privilege    Privilege `json:"privilege"`
	Range        []Range   `json:"range"`
}

// ExternalcontactCustomerStrategyCreateSchema 创建新的规则组响应内容
type ExternalcontactCustomerStrategyCreateSchema struct {
	util.CommonError
	StrategyID int `json:"strategy_id"`
}

// ExternalcontactCustomerStrategyCreate 创建新的规则组
func (r *Client) ExternalcontactCustomerStrategyCreate(options ExternalcontactCustomerStrategyCreateOptions) (info ExternalcontactCustomerStrategyCreateSchema, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(externalcontactCustomerStrategyCreateAddr, accessToken), options)
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
type ExternalcontactCustomerStrategyEditOptions struct {
	StrategyID   int       `json:"strategy_id"`
	StrategyName string    `json:"strategy_name"`
	AdminList    []string  `json:"admin_list"`
	Privilege    Privilege `json:"privilege"`
	RangeAdd     []Range   `json:"range_add"`
	RangeDel     []Range   `json:"range_del"`
}

// ExternalcontactCustomerStrategyEdit 编辑规则组及其管理范围
func (r *Client) ExternalcontactCustomerStrategyEdit(options ExternalcontactCustomerStrategyEditOptions) (info util.CommonError, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(externalcontactCustomerStrategyEditAddr, accessToken), options)
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
type ExternalcontactCustomerStrategyDelOptions struct {
	StrategyID int `json:"strategy_id"`
}

// ExternalcontactCustomerStrategyDel 删除规则组
func (r *Client) ExternalcontactCustomerStrategyDel(options ExternalcontactCustomerStrategyDelOptions) (info util.CommonError, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(externalcontactCustomerStrategyDelAddr, accessToken), options)
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
