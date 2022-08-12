package externalcontact

import (
	"encoding/json"
	"fmt"
	"github.com/EricJSanchez/wecom/util"
)

const (
	//获取客户群详情
	externalcontactGroupChatInfoAddr = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/groupchat/get?access_token=%s&debug=1"
	//获取客户群详情
	externalcontactGroupChatListAddr = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/groupchat/list?access_token=%s"
)

//获取客户群详情参数
type ExternalcontactGroupChatInfoOptions struct {
	ChatId   string `json:"chat_id"`   // 	 	客户群ID
	NeedName int    `json:"need_name"` // 	是否需要返回群成员的名字group_chat.member_list.name。0-不返回；1-返回。默认不返回
}

//客户群详情
type ExternalcontactGroupChatInfo struct {
	ChatId     string                           `json:"chat_id"`     //客户群ID
	Name       string                           `json:"name"`        //群名
	Owner      string                           `json:"owner"`       //群主ID
	CreateTime int                              `json:"create_time"` //群的创建时间
	Notice     string                           `json:"notice"`      //群公告
	MemberList []ExternalcontactGroupChatMember `json:"member_list"` //群成员列表
	AdminList  []ExternalcontactGroupChatAdmin  `json:"admin_list"`  //群管理员列表
}

//群成员列表
type ExternalcontactGroupChatMember struct {
	Userid        string                          `json:"userid"`         // 群成员id
	Type          int                             `json:"type"`           //成员类型1 - 企业成员 2 - 外部联系人
	Unionid       string                          `json:"unionid"`        //外部联系人在微信开放平台的唯一身份标识（微信unionid），通过此字段企业可将外部联系人与公众号/小程序用户关联起来。仅当群成员类型是微信用户（包括企业成员未添加好友），且企业或第三方服务商绑定了微信开发者ID有此字段。查看绑定方法
	JoinTime      int                             `json:"join_time"`      //入群时间
	JoinScene     int                             `json:"join_scene"`     //入群方式。1 - 由群成员邀请入群（直接邀请入群）2 - 由群成员邀请入群（通过邀请链接入群）3 - 通过扫描群二维码入群
	Invitor       ExternalcontactGroupChatInvitor `json:"invitor"`        //邀请者
	GroupNickname string                          `json:"group_nickname"` //在群里的昵称
	Name          string                          `json:"name"`           //名字。仅当 need_name = 1 时返回如果是微信用户，则返回其在微信中设置的名字如果是企业微信联系人，则返回其设置对外展示的别名或实名

}

//邀请人
type ExternalcontactGroupChatInvitor struct {
	Userid string `json:"userid"` //邀请者的userid

}

//群管理员列表
type ExternalcontactGroupChatAdmin struct {
	Userid string `json:"userid"` //群管理员userid

}

//群列表返回参数
type ExternalcontactGroupChatList struct {
	StatusFilter int                     `json:"status_filter"` //客户群跟进状态过滤。0 - 所有列表(即不过滤) 1 - 离职待继承 2 - 离职继承中 3 - 离职继承完成 默认为0
	OwnerFilter  []GroupChatOwnerFilters `json:"owner_filter"`  //群主过滤。 如果不填，表示获取应用可见范围内全部群主的数据（但是不建议这么用，如果可见范围人数超过1000人，为了防止数据包过大，会报错 81017）
	Cursor       string                  `json:"cursor"`        //用于分页查询的游标，字符串类型，由上一次调用返回，首次调用不填
	Limit        int                     `json:"limit"`         //分页，预期请求的数据量，取值范围 1 ~ 1000
}

//群主过滤
type GroupChatOwnerFilters struct {
	UseridList string `json:"userid_list"` //用户ID列表。最多100个

}

//获取群详情返回结果
type GetExternalcontactGroupChatInfoSchema struct {
	util.CommonError
	ExternalcontactGroupChatInfo ExternalcontactGroupChatInfo `json:"group_chat"`
}

//获取群详情
func (r *Client) GetExternalContactGroupChatInfo(options ExternalcontactGroupChatInfoOptions) (info GetExternalcontactGroupChatInfoSchema, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	optionJson, err := json.Marshal(options)
	if err != nil {
		return
	}
	data, err = util.HTTPPost(fmt.Sprintf(externalcontactGroupChatInfoAddr, accessToken), string(optionJson))
	if err != nil {
		fmt.Println("获取群详情 err", err)
		return
	}
	if err = json.Unmarshal(data, &info); err != nil {
		return
	}
	if info.ErrCode != 0 {
		return info, NewSDKErr(info.ErrCode, info.ErrMsg)
	}
	fmt.Println()
	fmt.Printf("获取群详情%+v", info)
	fmt.Println()
	return info, nil
}

//获取群列表返回结果
type GetExternalcontactGroupChatListSchema struct {
	util.CommonError
	ExternalcontactGroupChatList ExternalcontactGroupChatList `json:"externalcontact_group_chat_list"`
}

//获取群列表
func (r *Client) GetExternalcontactGroupChatList() (list GetExternalcontactGroupChatListSchema, err error) {
	accessToken, err := r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err := util.HTTPGet(fmt.Sprintf(externalcontactGroupChatListAddr, accessToken))
	if err != nil {
		return
	}
	if err = json.Unmarshal(data, &list); err != nil {
		return
	}
	if list.ErrCode != 0 {
		return list, NewSDKErr(list.ErrCode, list.ErrMsg)
	}
	return list, nil
}
