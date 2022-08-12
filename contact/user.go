package contact

import (
	"encoding/json"
	"fmt"
	"wecom/util"
)

const (
	// 创建成员
	userCreateAddr = "https://qyapi.weixin.qq.com/cgi-bin/user/create?access_token=%s"
	// 读取成员
	userGetAddr = "https://qyapi.weixin.qq.com/cgi-bin/user/get?access_token=%s&userid=%s"
	// 更新成员
	userUpdateAddr = "https://qyapi.weixin.qq.com/cgi-bin/user/update?access_token=%s"
	// 删除成员
	userDeleteAddr = "https://qyapi.weixin.qq.com/cgi-bin/user/delete?access_token=%s&userid=%s"
	// 批量删除成员
	userBatchdeleteAddr = "https://qyapi.weixin.qq.com/cgi-bin/user/batchdelete?access_token=%s"
	// 获取部门成员
	userSimplelistAddr = "https://qyapi.weixin.qq.com/cgi-bin/user/simplelist?access_token=%s&department_id=%s&fetch_child=%s"
	// 获取部门成员详情
	userListAddr = "https://qyapi.weixin.qq.com/cgi-bin/user/list?access_token=%s&department_id=%s&fetch_child=%s"
	// userid转openid
	userConvertToOpenidAddr = "https://qyapi.weixin.qq.com/cgi-bin/user/convert_to_openid?access_token=%s"
	// openid转userid
	userConvertToUseridAddr = "https://qyapi.weixin.qq.com/cgi-bin/user/convert_to_userid?access_token=%s"
	// 二次验证
	userAuthsuccAddr = "https://qyapi.weixin.qq.com/cgi-bin/user/authsucc?access_token=%s&userid=%s"
	// 邀请成员
	batchInviteAddr = "https://qyapi.weixin.qq.com/cgi-bin/batch/invite?access_token=%s"
	// 获取加入企业二维码
	corpGetJoinQrcodeAddr = "https://qyapi.weixin.qq.com/cgi-bin/corp/get_join_qrcode?access_token=%s&size_type=%d"
	// 获取企业活跃成员数
	userGetActiveStatAddr = "https://qyapi.weixin.qq.com/cgi-bin/user/get_active_stat?access_token=%s"
	// 2022-08 新版本
	userListIdAddr = "https://qyapi.weixin.qq.com/cgi-bin/user/list_id?access_token=%s"
)

// UserCreateOptions 创建成员请求参数
type UserCreateOptions struct {
	OpenUserid       string                 `json:"open_userid,omitempty"`       // 全局唯一
	Userid           string                 `json:"userid"`                      // 企业微信员工ID
	Name             string                 `json:"name,omitempty"`              // 成员名称
	Alias            string                 `json:"alias,omitempty"`             // 成员别名
	Position         string                 `json:"position,omitempty"`          // 职务信息
	Mobile           string                 `json:"mobile,omitempty"`            // 电话号码
	Gender           string                 `json:"gender,omitempty"`            // 0表示未定义，1表示男性，2表示女性
	Email            string                 `json:"email,omitempty"`             // 邮件
	Avatar           string                 `json:"avatar,omitempty"`            // 头像
	LeaderInDept     []int                  `json:"leader_in_dept,omitempty"`    // 是否是管理,表示在所在的部门内是否为上级
	IsLeaderInDept   []int                  `json:"is_leader_in_dept,omitempty"` // 是否是管理,表示在所在的部门内是否为上级
	Enable           int                    `json:"enable,omitempty"`            // 启用/禁用成员。1表示启用成员，0表示禁用成员
	Extattr          map[string]interface{} `json:"extattr,omitempty"`           // 扩展属性
	Telephone        string                 `json:"telephone,omitempty"`         // 座机
	MainDepartment   int                    `json:"main_department,omitempty"`   // 主部门
	QrCode           string                 `json:"qr_code,omitempty"`           // 员工个人二维码
	Department       []int                  `json:"department,omitempty"`        // 成员所属部门id列表
	Order            []int                  `json:"order,omitempty"`             // 部门内的排序值
	ThumbAvatar      string                 `json:"thumb_avatar,omitempty"`      // 头像缩略图url
	ExternalProfile  map[string]interface{} `json:"external_profile,omitempty"`  // 成员对外属性
	ExternalPosition string                 `json:"external_position,omitempty"` // 对外职务
	Address          string                 `json:"address,omitempty"`           // 地址
	HideMobile       int                    `json:"hide_mobile,omitempty"`
	Isleader         int                    `json:"isleader,omitempty"`
}

// UserCreate 创建成员
func (r *Client) UserCreate(options UserCreateOptions) (info util.CommonError, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(userCreateAddr, accessToken), options)
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

// UserGetOptions 读取成员请求参数
type UserGetOptions struct {
	Userid string `json:"userid"` // 成员UserID。对应管理端的帐号，企业内必须唯一。不区分大小写，长度为1~64个字节
}

// UserGetSchema 成员详情响应内容
type UserGetSchema struct {
	util.CommonError
	UserCreateOptions
	Status int `json:"status"` // 1=已激活，2=已禁用，4=未激活，5=退出企业。
}

// UserGet UserGet 读取成员 微信已经废弃此接口 2022-08， 新IP不再能调用
func (r *Client) UserGet(options UserGetOptions) (info UserGetSchema, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.HTTPGet(fmt.Sprintf(userGetAddr, accessToken, options.Userid))
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

// UserUpdateOptions 更新成员请求参数
type UserUpdateOptions struct {
	UserCreateOptions
	AvatarMediaid string `json:"avatar_mediaid,omitempty"`
}

// UserUpdate 更新成员
func (r *Client) UserUpdate(options UserUpdateOptions) (info util.CommonError, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(userUpdateAddr, accessToken), options)
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

// UserDeleteOptions 删除成员请求参数
type UserDeleteOptions struct {
	Userid string `json:"userid"` // 企业微信员工ID
}

// UserDelete 删除成员
func (r *Client) UserDelete(options UserDeleteOptions) (info util.CommonError, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.HTTPGet(fmt.Sprintf(userDeleteAddr, accessToken, options.Userid))
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

// UserBatchdeleteOptions 批量删除成员请求参数
type UserBatchdeleteOptions struct {
	Useridlist []string `json:"useridlist"`
}

// UserBatchdelete 批量删除成员
func (r *Client) UserBatchdelete(options UserBatchdeleteOptions) (info util.CommonError, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(userBatchdeleteAddr, accessToken), options)
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

// UserSimplelistOptions 获取部门成员请求参数
type UserSimplelistOptions struct {
	DepartmentId string `json:"department_id"`
	FetchChild   string `json:"fetch_child"`
}

// UserSimpleInfoSchema 部门成员信息
type UserSimpleInfoSchema struct {
	Userid     string `json:"userid"`
	Name       string `json:"name"`
	Department []int  `json:"department"`
	OpenUserid string `json:"open_userid"`
}

// UserSimplelistSchema 部门成员响应内容
type UserSimplelistSchema struct {
	util.CommonError
	Userlist []UserSimpleInfoSchema `json:"userlist"`
}

// UserSimplelist 获取部门成员
func (r *Client) UserSimplelist(options UserSimplelistOptions) (info UserSimplelistSchema, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.HTTPGet(fmt.Sprintf(userSimplelistAddr, accessToken, options.DepartmentId, options.FetchChild))
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

// UserListOptions 获取部门成员详情请求参数
type UserListOptions struct {
	DepartmentId string `json:"department_id"`
	FetchChild   string `json:"fetch_child"`
}

// UserListInfoSchema 获取部门成员详情信息
type UserListInfoSchema struct {
	UserCreateOptions
	Status int `json:"status"` // 1=已激活，2=已禁用，4=未激活，5=退出企业。
}

// UserListSchema 获取部门成员详情响应内容
type UserListSchema struct {
	util.CommonError
	Userlist []UserListInfoSchema `json:"userlist"`
}

// UserList 获取部门成员详情 微信已经废弃此接口 2022-08， 新IP不再能调用
func (r *Client) UserList(options UserListOptions) (info UserListSchema, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.HTTPGet(fmt.Sprintf(userListAddr, accessToken, options.DepartmentId, options.FetchChild))
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

// UserListIdOptions 员工列表请求参数
type UserListIdOptions struct {
	Cursor string `json:"cursor"`
	Limit  int    `json:"limit"`
}

// UserListIdDeptUser 获取员工ID列表响应子集
type UserListIdDeptUser struct {
	Userid     string `json:"userid"`
	Department int    `json:"department"`
}

// UserListIdSchema 获取员工ID列表响应
type UserListIdSchema struct {
	util.CommonError
	NextCursor string               `json:"next_cursor"`
	DeptUser   []UserListIdDeptUser `json:"dept_user"`
}

// UserListId 获取员工ID列表
func (r *Client) UserListId(options UserListIdOptions) (info UserListIdSchema, err error) {
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
	fmt.Println(string(optionJson))
	data, err = util.HTTPPost(fmt.Sprintf(userListIdAddr, accessToken), string(optionJson))
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

// UserConvertToOpenidOptions userid转openid请求参数
type UserConvertToOpenidOptions struct {
	Userid string `json:"userid"`
}

// UserConvertToOpenidSchema userid转openid响应内容
type UserConvertToOpenidSchema struct {
	util.CommonError
	Openid string `json:"openid"`
}

// UserConvertToOpenid userid转openid
func (r *Client) UserConvertToOpenid(options UserConvertToOpenidOptions) (info UserConvertToOpenidSchema, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(userConvertToOpenidAddr, accessToken), options)
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

// openid转userid请求参数
type UserConvertToUseridOptions struct {
	Openid string `json:"openid"`
}

// UserConvertToUseridSchema openid转userid响应内容
type UserConvertToUseridSchema struct {
	util.CommonError
	Userid string `json:"userid"`
}

// UserConvertToUserid openid转userid
func (r *Client) UserConvertToUserid(options UserConvertToUseridOptions) (info UserConvertToUseridSchema, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(userConvertToUseridAddr, accessToken), options)
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

// UserAuthsuccOptions 二次验证请求参数
type UserAuthsuccOptions struct {
	Userid string `json:"userid"`
}

// UserAuthsucc 二次验证
func (r *Client) UserAuthsucc(options UserAuthsuccOptions) (info util.CommonError, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.HTTPGet(fmt.Sprintf(userAuthsuccAddr, accessToken, options.Userid))
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

// BatchInviteOptions 邀请成员请求参数
type BatchInviteOptions struct {
	User  []string `json:"user"`
	Party []int    `json:"party"`
	Tag   []int    `json:"tag"`
}

// BatchInviteSchema 邀请成员响应内容
type BatchInviteSchema struct {
	util.CommonError
	Invaliduser  []string `json:"invaliduser"`
	Invalidparty []int    `json:"invalidparty"`
	Invalidtag   []int    `json:"invalidtag"`
}

// BatchInvite 邀请成员
func (r *Client) BatchInvite(options BatchInviteOptions) (info BatchInviteSchema, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(batchInviteAddr, accessToken), options)
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

// CorpGetJoinQrcodeOptions 获取加入企业二维码请求参数
type CorpGetJoinQrcodeOptions struct {
	SizeType int `json:"size_type"`
}

// CorpGetJoinQrcodeSchema 获取加入企业二维码响应内容
type CorpGetJoinQrcodeSchema struct {
	util.CommonError
	JoinQrcode int `json:"join_qrcode"`
}

// CorpGetJoinQrcode 获取加入企业二维码
func (r *Client) CorpGetJoinQrcode(options CorpGetJoinQrcodeOptions) (info CorpGetJoinQrcodeSchema, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.HTTPGet(fmt.Sprintf(corpGetJoinQrcodeAddr, accessToken, options.SizeType))
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

// UserGetActiveStatOptions 获取企业活跃成员数请求参数
type UserGetActiveStatOptions struct {
	Date string `json:"date"`
}

// UserGetActiveStatSchema 获取企业活跃成员数响应内容
type UserGetActiveStatSchema struct {
	util.CommonError
	ActiveCnt int `json:"active_cnt"`
}

// UserGetActiveStat 获取企业活跃成员数
func (r *Client) UserGetActiveStat(options UserGetActiveStatOptions) (info UserGetActiveStatSchema, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(userGetActiveStatAddr, accessToken), options)
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
