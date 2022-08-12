package externalcontact

import (
	"encoding/json"
	"fmt"
	"github.com/EricJSanchez/wecom/util"
)

const (
	// 获取「联系客户统计」数据
	externalcontactGetUserBehaviorDataAddr = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_user_behavior_data?access_token=%s"
	// 获取「群聊数据统计」数据 - 按群主聚合的方式
	externalcontactGroupchatStatisticAddr = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/groupchat/statistic?access_token=%s"
	// 获取「群聊数据统计」数据 - 按自然日聚合的方式
	externalcontactGroupchatStatisticGroupByDayAddr = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/groupchat/statistic_group_by_day?access_token=%s"
)

// ExternalcontactGetUserBehaviorDataOptions 获取「联系客户统计」数据请求参数
type ExternalcontactGetUserBehaviorDataOptions struct {
	Userid    []string `json:"userid"`
	Partyid   []int    `json:"partyid"`
	StartTime int      `json:"start_time"`
	EndTime   int      `json:"end_time"`
}

// BehaviorData 行为数据详情
type BehaviorData struct {
	StatTime            int     `json:"stat_time"`
	ChatCnt             int     `json:"chat_cnt"`
	MessageCnt          int     `json:"message_cnt"`
	ReplyPercentage     float64 `json:"reply_percentage"`
	AvgReplyTime        int     `json:"avg_reply_time"`
	NegativeFeedbackCnt int     `json:"negative_feedback_cnt"`
	NewApplyCnt         int     `json:"new_apply_cnt"`
	NewContactCnt       int     `json:"new_contact_cnt"`
}

// ExternalcontactGetUserBehaviorDataSchema 获取「联系客户统计」数据响应内容
type ExternalcontactGetUserBehaviorDataSchema struct {
	util.CommonError
	BehaviorData []BehaviorData `json:"behavior_data"`
}

// ExternalcontactGetUserBehaviorData 获取「联系客户统计」数据
func (r *Client) ExternalcontactGetUserBehaviorData(options ExternalcontactGetUserBehaviorDataOptions) (info ExternalcontactGetUserBehaviorDataSchema, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(externalcontactGetUserBehaviorDataAddr, accessToken), options)
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

// OwnerFilterChild 群主过滤信息
type OwnerFilterChild struct {
	UseridList []string `json:"userid_list"`
}

// ExternalcontactGroupchatStatisticOptions 获取「群聊数据统计」数据 - 按群主聚合的方式
type ExternalcontactGroupchatStatisticOptions struct {
	DayBeginTime int              `json:"day_begin_time"`
	DayEndTime   int              `json:"day_end_time"`
	OwnerFilter  OwnerFilterChild `json:"owner_filter"`
	OrderBy      int              `json:"order_by"`
	OrderAsc     int              `json:"order_asc"`
	Offset       int              `json:"offset"`
	Limit        int              `json:"limit"`
}

// Data 记录数据详情
type Data struct {
	NewChatCnt            int `json:"new_chat_cnt"`
	ChatTotal             int `json:"chat_total"`
	ChatHasMsg            int `json:"chat_has_msg"`
	NewMemberCnt          int `json:"new_member_cnt"`
	MemberTotal           int `json:"member_total"`
	MemberHasMsg          int `json:"member_has_msg"`
	MsgTotal              int `json:"msg_total"`
	MigrateTraineeChatCnt int `json:"migrate_trainee_chat_cnt"`
}

// Items 记录列表。表示某个群主所拥有的客户群的统计数据
type Items struct {
	Owner string `json:"owner"`
	Data  Data   `json:"data"`
}

// ExternalcontactGroupchatStatisticSchema 获取「群聊数据统计」数据 - 按群主聚合的方式
type ExternalcontactGroupchatStatisticSchema struct {
	util.CommonError
	Total      int     `json:"total"`
	NextOffset int     `json:"next_offset"`
	Items      []Items `json:"items"`
}

// ExternalcontactGroupchatStatistic 获取「群聊数据统计」数据 - 按群主聚合的方式
func (r *Client) ExternalcontactGroupchatStatistic(options ExternalcontactGroupchatStatisticOptions) (info ExternalcontactGroupchatStatisticSchema, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(externalcontactGroupchatStatisticAddr, accessToken), options)
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

// ExternalcontactGroupchatStatisticGroupByDayOptions 获取「群聊数据统计」数据 - 按自然日聚合的方式
type ExternalcontactGroupchatStatisticGroupByDayOptions struct {
	DayBeginTime int         `json:"day_begin_time"`
	DayEndTime   int         `json:"day_end_time"`
	OwnerFilter  OwnerFilter `json:"owner_filter"`
}

// 日期统计记录详情
type DateData struct {
	NewChatCnt            int `json:"new_chat_cnt"`
	ChatTotal             int `json:"chat_total"`
	ChatHasMsg            int `json:"chat_has_msg"`
	NewMemberCnt          int `json:"new_member_cnt"`
	MemberTotal           int `json:"member_total"`
	MemberHasMsg          int `json:"member_has_msg"`
	MsgTotal              int `json:"msg_total"`
	MigrateTraineeChatCnt int `json:"migrate_trainee_chat_cnt"`
}

// DateItems 记录列表。表示某个自然日客户群的统计数据
type DateItems struct {
	StatTime int      `json:"stat_time"`
	Data     DateData `json:"data"`
}

// ExternalcontactGroupchatStatisticGroupByDaySchema 获取「群聊数据统计」数据 - 按自然日聚合的方式
type ExternalcontactGroupchatStatisticGroupByDaySchema struct {
	util.CommonError
	Errcode int         `json:"errcode"`
	Errmsg  string      `json:"errmsg"`
	Items   []DateItems `json:"items"`
}

// ExternalcontactGroupchatStatisticGroupByDay 获取「群聊数据统计」数据 - 按自然日聚合的方式
func (r *Client) ExternalcontactGroupchatStatisticGroupByDay(options ExternalcontactGroupchatStatisticGroupByDayOptions) (info ExternalcontactGroupchatStatisticGroupByDaySchema, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(externalcontactGroupchatStatisticGroupByDayAddr, accessToken), options)
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
