package platform

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/EricJSanchez/wecom/util"
	"github.com/spf13/cast"
	"time"
)

const (
	deptCacheUrl                 = "https://work.weixin.qq.com/wework_admin/contacts/party/cache?lang=zh_CN&f=json&ajax=1&timeZoneInfo[zone_offset]=-8&random=0.%d"
	deptStaffListUrl             = "https://work.weixin.qq.com/wework_admin/contacts/getDepartMember?lang=zh_CN&f=json&ajax=1&timeZoneInfo[zone_offset]=-8&random=0.%d&action=getpartycontacts&partyid=%s&page=%d&limit=%d&joinstatus=0&fetchchild=1&preFetch=false&use_corp_cache=0&_d2st=a2300212"
	searchStaffListUrl           = "https://work.weixin.qq.com/wework_admin/search_contact?lang=zh_CN&f=json&ajax=1&timeZoneInfo[zone_offset]=-8&random=0.%d"
	getRoleListUrl               = "https://work.weixin.qq.com/wework_admin/profile/role/getRoleList?lang=zh_CN&f=json&ajax=1&timeZoneInfo[zone_offset]=-8&random=0.%d&fill_manage_scope=1&_d2st=a650927"
	saveMemberUrl                = "https://work.weixin.qq.com/wework_admin/contacts/saveMember?method=update&lang=zh_CN&f=json&ajax=1&timeZoneInfo[zone_offset]=-8&random=0.%d"
	getSingleMemberUrl           = "https://work.weixin.qq.com/wework_admin/contacts/getSingleMember?id=%s&lang=zh_CN&f=json&ajax=1&timeZoneInfo[zone_offset]=-8&random=0.%d&_d2st=a2762875"
	getCorpEncryptDataAppInfoUrl = "https://work.weixin.qq.com/wework_admin/financial/getCorpEncryptDataAppInfo?lang=zh_CN&f=json&ajax=1&timeZoneInfo[zone_offset]=-8&random=0.%d&flag=0&_d2st=a2045480"
	getApplyList                 = "https://work.weixin.qq.com/wework_admin/getApplyList?lang=zh_CN&f=json&ajax=1&timeZoneInfo[zone_offset]=-8&random=0.%d&limit=%d&start=%d&status=0&oper=1&_d2st=a8452539"
)

var commonPlatForm = map[string]string{
	"referer": "https://work.weixin.qq.com/wework_admin/frame",
	"origin":  "https://work.weixin.qq.com",
}

// deptCacheUrl 返回信息
type DeptCacheRes struct {
	Data Data `json:"data"`
}
type Data struct {
	Cacheversion Cacheversion `json:"cacheversion"`
	PartyList    PartyList    `json:"party_list"`
}
type Cacheversion struct {
	CacheVersion     int    `json:"cache_version"`
	DbVersion        string `json:"db_version"`
	AcctCacheVersion int    `json:"acct_cache_version"`
	AcctDbVersion    string `json:"acct_db_version"`
}
type PartyList struct {
	List []List `json:"list"`
}
type List struct {
	Partyid        string       `json:"partyid"`
	OpenapiPartyid string       `json:"openapi_partyid"`
	Name           string       `json:"name"`
	Parentid       interface{}  `json:"parentid"`
	Authority      int          `json:"authority"`
	Islocked       bool         `json:"islocked"`
	DisplayOrder   int          `json:"display_order"`
	LanguageList   LanguageList `json:"language_list"`
	Pinyin         string       `json:"pinyin"`
	Py             string       `json:"py"`
}
type LanguageList struct {
	Info []interface{} `json:"info"`
}

// GetDeptCache 获取部门列表
func (r *Client) GetDeptCache() (deptCache DeptCacheRes, err error) {
	//var accessToken string
	cookie := r.ctx.Config.Cookie
	var header = commonPlatForm
	header["cookie"] = cookie
	if cookie == "" {
		err = errors.New("cookie 缺失")
		return
	}
	rspOrigin, err := util.PostFormEncodeWithHeader(fmt.Sprintf(deptCacheUrl, time.Now().UnixMilli()), map[string]string{
		"_d2st": "",
	}, header)
	err = json.Unmarshal(rspOrigin, &deptCache)
	if len(deptCache.Data.PartyList.List) == 0 {
		err = errors.New("请求出错:" + string(rspOrigin))
		return
	}
	if err != nil {
		return
	}
	return
}

type DeptStaffDataRes struct {
	Data DeptStaffData `json:"data"`
}
type DispOrder struct {
	DepartID   string        `json:"depart_id"`
	DispOrder  int           `json:"disp_order"`
	LeaderRank int           `json:"leader_rank"`
	Pathids    []interface{} `json:"pathids"`
	PathNames  []interface{} `json:"path_names"`
}
type ExternalWxfinder struct {
}
type Superoirs struct {
	Vids []interface{} `json:"vids"`
}
type Tag struct {
	List []interface{} `json:"list"`
}
type Extattr struct {
	Attrlist []interface{} `json:"attrlist"`
}
type ContactListChild struct {
	CorpID              string           `json:"corp_id"`
	Vid                 string           `json:"vid"`
	Name                string           `json:"name"`
	Mobile              string           `json:"mobile"`
	Email               string           `json:"email"`
	Gender              int              `json:"gender"`
	DepartIds           []string         `json:"depart_ids"`
	Position            string           `json:"position"`
	Wechat              string           `json:"wechat"`
	Avatar              string           `json:"avatar"`
	Account             string           `json:"account"`
	ExtTel              string           `json:"ext_tel"`
	DisableStat         int              `json:"disable_stat"`
	IdentityStat        int              `json:"identity_stat"`
	ManageStat          int              `json:"manage_stat"`
	BindStat            int              `json:"bind_stat"`
	EnglishName         string           `json:"english_name"`
	LoginStat           int              `json:"login_stat"`
	Alias               string           `json:"alias"`
	ActiveBiz           bool             `json:"active_biz"`
	Domainid            int              `json:"domainid"`
	WxqyUserid          string           `json:"wxqy_userid"`
	Gid                 interface{}      `json:"gid"`
	Pinyin              string           `json:"pinyin"`
	MainpartyID         interface{}      `json:"mainparty_id"`
	DisableBiz          bool             `json:"disable_biz"`
	CountryCode         string           `json:"country_code"`
	WxIDHash            string           `json:"wx_id_hash"`
	DepartNames         []interface{}    `json:"depart_names"`
	HideMobile          bool             `json:"hide_mobile"`
	DispOrder           []DispOrder      `json:"disp_order"`
	Acctid              string           `json:"acctid"`
	Attrs               []interface{}    `json:"attrs"`
	IsQuit              bool             `json:"is_quit"`
	IsWwBizmailVip      bool             `json:"is_ww_bizmail_vip"`
	IsWwBizmail         bool             `json:"is_ww_bizmail"`
	IsJoinQyh           bool             `json:"is_join_qyh"`
	PstnExtensionNumber string           `json:"pstn_extension_number"`
	ExternalAttrs       []interface{}    `json:"external_attrs"`
	UserQuitTime        int              `json:"user_quit_time"`
	IsReadyJoinAgain    bool             `json:"is_ready_join_again"`
	XcxCorpAddress      string           `json:"xcx_corp_address"`
	ExternalCorpInfo    string           `json:"external_corp_info"`
	ExternalWxfinder    ExternalWxfinder `json:"external_wxfinder"`
	PersonalEmail       string           `json:"personal_email"`
	Superoirs           Superoirs        `json:"superoirs"`
	BizMail             string           `json:"biz_mail"`
	ID                  string           `json:"id"`
	Uin                 string           `json:"uin"`
	Username            string           `json:"username"`
	Imgid               string           `json:"imgid"`
	PartyList           []string         `json:"party_list"`
	JoinStatus          string           `json:"JoinStatus"`
	Tag                 Tag              `json:"tag"`
	Extattr             Extattr          `json:"extattr"`
}

type ContactList struct {
	List []ContactListChild `json:"list"`
}
type DeptStaffData struct {
	ContactList         ContactList `json:"contact_list"`
	NextPageContactList ContactList `json:"next_page_contact_list"`
	MemberCount         int         `json:"member_count"`
	DisableCnt          int         `json:"disable_cnt"`
	PageCount           int         `json:"page_count"`
	Status              string      `json:"status"`
}

// GetDeptStaff 获取员工列表,page 0 代表第一页
func (r *Client) GetDeptStaff(partyId string, page, size int) (deptStaff DeptStaffDataRes, err error) {
	//var accessToken string
	cookie := r.ctx.Config.Cookie
	var header = commonPlatForm
	header["cookie"] = cookie
	if cookie == "" {
		err = errors.New("cookie 缺失")
		return
	}
	uri := fmt.Sprintf(deptStaffListUrl, time.Now().UnixMilli(), partyId, page, size)
	rspOrigin, err := util.GetWithHeader(uri, header)
	err = json.Unmarshal(rspOrigin, &deptStaff)
	if deptStaff.Data.PageCount == 0 {
		err = errors.New("请求出错:" + string(rspOrigin))
		return
	}
	if err != nil {
		return
	}
	return
}

type SearchStaffRes struct {
	Data []SearchStaffData `json:"data"`
}
type SdDispOrder struct {
	DepartID  string        `json:"depart_id"`
	DispOrder int           `json:"disp_order"`
	Pathids   []interface{} `json:"pathids"`
	PathNames []interface{} `json:"path_names"`
}
type SdSuperoirs struct {
	Vids []interface{} `json:"vids"`
}
type SearchStaffData struct {
	CorpID           string        `json:"corp_id"`
	Vid              string        `json:"vid"`
	Name             string        `json:"name"`
	Mobile           string        `json:"mobile"`
	Gender           int           `json:"gender"`
	DepartIds        []string      `json:"depart_ids"`
	Position         string        `json:"position,omitempty"`
	Avatar           string        `json:"avatar"`
	Account          string        `json:"account,omitempty"`
	ExtTel           string        `json:"ext_tel,omitempty"`
	DisableStat      int           `json:"disable_stat"`
	IdentityStat     int           `json:"identity_stat"`
	DeleteStat       int           `json:"delete_stat"`
	ManageStat       int           `json:"manage_stat"`
	JoinStat         int           `json:"join_stat"`
	BindStat         int           `json:"bind_stat"`
	EnglishName      string        `json:"english_name,omitempty"`
	LoginStat        int           `json:"login_stat"`
	Alias            string        `json:"alias,omitempty"`
	ActiveBiz        bool          `json:"active_biz"`
	WxqyUserid       string        `json:"wxqy_userid,omitempty"`
	Gid              interface{}   `json:"gid"`
	Pinyin           string        `json:"pinyin,omitempty"`
	MainpartyID      string        `json:"mainparty_id"`
	OldUin           int           `json:"old_uin,omitempty"`
	DisableBiz       bool          `json:"disable_biz"`
	CountryCode      string        `json:"country_code"`
	WxIDHash         string        `json:"wx_id_hash"`
	DepartNames      []string      `json:"depart_names"`
	HideMobile       bool          `json:"hide_mobile"`
	DispOrder        []SdDispOrder `json:"disp_order"`
	Acctid           string        `json:"acctid"`
	Attrs            []interface{} `json:"attrs"`
	IsQuit           bool          `json:"is_quit"`
	IsWwBizmailVip   bool          `json:"is_ww_bizmail_vip"`
	IsWwBizmail      bool          `json:"is_ww_bizmail"`
	IsJoinQyh        bool          `json:"is_join_qyh"`
	AcctidStat       int           `json:"acctid_stat"`
	ExternalAttrs    []interface{} `json:"external_attrs"`
	UserQuitTime     int           `json:"user_quit_time"`
	IsReadyJoinAgain bool          `json:"is_ready_join_again"`
	IsRealname       bool          `json:"is_realname"`
	Superoirs        SdSuperoirs   `json:"superoirs"`
	BizMail          string        `json:"biz_mail"`
}

// SearchStaff  搜索获取员工列表
func (r *Client) SearchStaff(key string) (searchStaff SearchStaffRes, err error) {
	//var accessToken string
	cookie := r.ctx.Config.Cookie
	var header = commonPlatForm
	header["cookie"] = cookie
	if cookie == "" {
		err = errors.New("cookie 缺失")
		return
	}
	tmpBody := map[string]string{
		"keyword":            key,
		"force_manage":       "1",
		"searchLeavingStaff": "false",
		"_d2st":              "",
	}
	uri := fmt.Sprintf(searchStaffListUrl, time.Now().UnixMilli())
	rspOrigin, err := util.PostFormEncodeWithHeader(uri, tmpBody, header)
	if err != nil {
		fmt.Println("拉取企业微信员工信息出错 ", err)
		fmt.Println("拉取企业微信员工信息 ", cast.ToString(rspOrigin))
	}
	err = json.Unmarshal(rspOrigin, &searchStaff)
	if err != nil {
		fmt.Println("解析出错 ", err)
		fmt.Println("解析信息出错", cast.ToString(rspOrigin))
	}
	return
}

type GetRoleListRes struct {
	Data GetRoleListData `json:"data"`
}
type GetRoleListData struct {
	RoleList GetRoleListRoleList `json:"role_list"`
}
type GetRoleListRoleList struct {
	Item []GetRoleListItem `json:"item"`
}
type GetRoleListItem struct {
	RoleID    int                  `json:"role_id"`
	RoleName  string               `json:"role_name"`
	CorpID    string               `json:"corp_id"`
	AdminList GetRoleListAdminList `json:"admin_list"`
	RoleType  string               `json:"role_type"`
}
type GetRoleListAdminList struct {
	Item []GetRoleListItemChild `json:"item"`
}
type GetRoleListItemChild struct {
	ID         string `json:"id"`
	Flags      int    `json:"flags"`
	Name       string `json:"name"`
	Logo       string `json:"logo"`
	CreateTime int    `json:"create_time"`
}

// GetRoleList  获取管理员列表
func (r *Client) GetRoleList() (ret GetRoleListRes, err error) {
	cookie := r.ctx.Config.Cookie
	var header = commonPlatForm
	header["cookie"] = cookie
	if cookie == "" {
		err = errors.New("cookie 缺失")
		return
	}
	uri := fmt.Sprintf(getRoleListUrl, time.Now().UnixMilli())
	rspOrigin, err := util.GetWithHeader(uri, header)
	err = json.Unmarshal(rspOrigin, &ret)
	if len(ret.Data.RoleList.Item) == 0 {
		err = errors.New("请求出错:" + string(rspOrigin))
		return
	}
	if err != nil {
		return
	}
	return
}

type SaveMemberDispOrder struct {
	DepartID   string `json:"depart_id"`
	DispOrder  string `json:"disp_order"`
	IsTop      string `json:"is_top"`
	LeaderRank string `json:"leader_rank"`
}

type SaveMemberExternalAttr struct {
	FieldValue string `json:"field_value"`
	FieldId    string `json:"field_id"`
	FieldName  string `json:"field_name"`
	FieldType  string `json:"field_type"`
}

type SaveMemberSchema struct {
	JoinStatus                          string                   `json:"JoinStatus"`
	Account                             string                   `json:"account"`
	Acctid                              string                   `json:"acctid"`
	AcctidStat                          string                   `json:"acctid_stat"`
	ActiveBiz                           string                   `json:"active_biz"`
	Alias                               string                   `json:"alias"`
	Avatar                              string                   `json:"avatar"`
	BIsQymailGray                       string                   `json:"b_is_qymail_gray"`
	BindStat                            string                   `json:"bind_stat"`
	BizMail                             string                   `json:"biz_mail"`
	CountryCode                         string                   `json:"country_code"`
	DeleteStat                          string                   `json:"delete_stat"`
	DisableBiz                          string                   `json:"disable_biz"`
	DisableStat                         string                   `json:"disable_stat"`
	DispOrder                           []SaveMemberDispOrder    `json:"disp_order"`
	Domain                              string                   `json:"domain"`
	Domainid                            string                   `json:"domainid"`
	Email                               string                   `json:"email"`
	EnglishName                         string                   `json:"english_name"`
	ExtTel                              string                   `json:"ext_tel"`
	ExternJobTitle                      string                   `json:"extern_job_title"`
	ExternPosition                      string                   `json:"extern_position"`
	ExternPositionInfoBSynInnerPosition string                   `json:"extern_position_info[b_syn_inner_position]"`
	ExternPositionInfoExternPosition    string                   `json:"extern_position_info[extern_position]"`
	ExternalAttrs                       []SaveMemberExternalAttr `json:"external_attrs"`
	ExternalCorpInfo                    string                   `json:"external_corp_info"`
	ExternalWxfinder                    string                   `json:"external_wxfinder[invisible]"`
	Gender                              string                   `json:"gender"`
	GenderStr                           string                   `json:"gender_str"`
	HideMobile                          string                   `json:"hide_mobile"`
	IdentityStat                        string                   `json:"identity_stat"`
	IgnoreAbnormalMobile                string                   `json:"ignore_abnormal_mobile"`
	Imgid                               string                   `json:"imgid"`
	IsSearchListShow                    string                   `json:"isSearchListShow"`
	IsJoinQyh                           string                   `json:"is_join_qyh"`
	IsQuit                              string                   `json:"is_quit"`
	IsReadyJoinAgain                    string                   `json:"is_ready_join_again"`
	IsWwBizmail                         string                   `json:"is_ww_bizmail"`
	IsWwBizmailVip                      string                   `json:"is_ww_bizmail_vip"`
	JoinStat                            string                   `json:"join_stat"`
	LoginStat                           string                   `json:"login_stat"`
	MainpartyID                         string                   `json:"mainparty_id"`
	ManageStat                          string                   `json:"manage_stat"`
	Mobile                              string                   `json:"mobile"`
	ModelType                           string                   `json:"model_type"`
	Name                                string                   `json:"name"`
	Nickname                            string                   `json:"nickname"`
	PartyList                           []string                 `json:"party_list"`
	Partyid                             string                   `json:"partyid"`
	PersonalEmail                       string                   `json:"personal_email"`
	Pinyin                              string                   `json:"pinyin"`
	Position                            string                   `json:"position"`
	PstnExtensionNumber                 string                   `json:"pstn_extension_number"`
	Realname                            string                   `json:"realname"`
	Uin                                 string                   `json:"uin"`
	UserQuitTime                        string                   `json:"user_quit_time"`
	Username                            string                   `json:"username"`
	Vid                                 string                   `json:"vid"`
	VidBindGid                          string                   `json:"vid_bind_gid"`
	Wechat                              string                   `json:"wechat"`
	WxIDHash                            string                   `json:"wx_id_hash"`
	WxNickName                          string                   `json:"wx_nick_name"`
	XcxCorpAddress                      string                   `json:"xcx_corp_address"`
}

type SaveMemberRes struct {
	Data SaveMemberData `json:"data"`
}

type SaveMemberData struct {
	CorpID string `json:"corp_id"`
	Acctid string `json:"acctid"`
	Name   string `json:"name"`
}

// SaveMember  保存用户信息
func (r *Client) SaveMember(info SaveMemberSchema) (res SaveMemberRes, err error) {
	//var accessToken string
	cookie := r.ctx.Config.Cookie
	var header = commonPlatForm
	header["cookie"] = cookie
	if cookie == "" {
		err = errors.New("cookie 缺失")
		return
	}
	var tmpBodyMap = make(map[string]interface{}, 1)
	var tmpBodyStruct = make([]util.BodyStruct, 0)
	tJson, err := json.Marshal(info)
	if err != nil {
		return
	}
	err = json.Unmarshal(tJson, &tmpBodyMap)
	if err != nil {
		return
	}
	// partyList SaveMemberDispOrder SaveMemberExternalAttr
	for k, v := range tmpBodyMap {
		if k == "disp_order" || k == "external_attrs" || k == "party_list" {
			continue
		}
		tmpBodyStruct = append(tmpBodyStruct, util.BodyStruct{
			Key: k,
			Val: cast.ToString(v),
		})
	}
	if len(info.PartyList) > 0 {
		for _, v := range info.PartyList {
			tmpBodyStruct = append(tmpBodyStruct, util.BodyStruct{
				Key: "party_list[]",
				Val: cast.ToString(v),
			})
		}
	}
	if len(info.DispOrder) > 0 {
		for k, v := range info.DispOrder {
			tmpBodyStruct = append(tmpBodyStruct, util.BodyStruct{
				Key: "disp_order[" + cast.ToString(k) + "][disp_order]",
				Val: cast.ToString(v.DispOrder),
			}, util.BodyStruct{
				Key: "disp_order[" + cast.ToString(k) + "][depart_id]",
				Val: cast.ToString(v.DepartID),
			}, util.BodyStruct{
				Key: "disp_order[" + cast.ToString(k) + "][is_top]",
				Val: cast.ToString(v.IsTop),
			}, util.BodyStruct{
				Key: "disp_order[" + cast.ToString(k) + "][leader_rank]",
				Val: cast.ToString(v.LeaderRank),
			})
		}
	}
	if len(info.ExternalAttrs) > 0 {
		for k, v := range info.ExternalAttrs {
			tmpBodyStruct = append(tmpBodyStruct, util.BodyStruct{
				Key: "external_attrs[" + cast.ToString(k) + "][field_value]",
				Val: cast.ToString(v.FieldValue),
			}, util.BodyStruct{
				Key: "external_attrs[" + cast.ToString(k) + "][field_id]",
				Val: cast.ToString(v.FieldId),
			}, util.BodyStruct{
				Key: "external_attrs[" + cast.ToString(k) + "][field_name]",
				Val: cast.ToString(v.FieldName),
			}, util.BodyStruct{
				Key: "external_attrs[" + cast.ToString(k) + "][field_type]",
				Val: cast.ToString(v.FieldType),
			})
		}
	}
	fmt.Println(tmpBodyStruct)
	//return
	uri := fmt.Sprintf(saveMemberUrl, time.Now().UnixMilli())
	rspOrigin, err := util.PostFormEncodeWithHeaderUseSlice(uri, tmpBodyStruct, header)
	//fmt.Println(cast.ToString(rspOrigin))
	if err != nil {
		fmt.Println("saveMember req Err:", err)
		fmt.Println("saveMember req Err:", cast.ToString(rspOrigin))
	}
	err = json.Unmarshal(rspOrigin, &res)
	if err != nil {
		fmt.Println("saveMember解析出错:", err)
		fmt.Println("saveMember解析信息出错:", cast.ToString(rspOrigin))
	}
	return
}

type GetSingleMemberRes struct {
	Data GetSingleMemberData `json:"data"`
}
type GetSingleMemberDispOrder struct {
	DepartID   string        `json:"depart_id"`
	DispOrder  int           `json:"disp_order"`
	IsTop      bool          `json:"is_top"`
	LeaderRank int           `json:"leader_rank"`
	Pathids    []interface{} `json:"pathids"`
	PathNames  []interface{} `json:"path_names"`
}
type GetSingleMemberExternPositionInfo struct {
	BSynInnerPosition bool   `json:"b_syn_inner_position"`
	ExternPosition    string `json:"extern_position"`
}
type GetSingleMemberExternalWxfinder struct {
}
type GetSingleMemberSuperoirs struct {
	Vids []interface{} `json:"vids"`
}
type GetSingleMemberTag struct {
	List []interface{} `json:"list"`
}
type GetSingleMemberExtattr struct {
	Attrlist []interface{} `json:"attrlist"`
}
type GetSingleMemberData struct {
	CorpID              string                            `json:"corp_id"`
	Vid                 string                            `json:"vid"`
	Name                string                            `json:"name"`
	Mobile              string                            `json:"mobile"`
	Email               string                            `json:"email"`
	Gender              int                               `json:"gender"`
	Position            string                            `json:"position"`
	Wechat              string                            `json:"wechat"`
	Avatar              string                            `json:"avatar"`
	Account             string                            `json:"account"`
	ExtTel              string                            `json:"ext_tel"`
	DisableStat         int                               `json:"disable_stat"`
	IdentityStat        int                               `json:"identity_stat"`
	DeleteStat          int                               `json:"delete_stat"`
	ManageStat          int                               `json:"manage_stat"`
	JoinStat            int                               `json:"join_stat"`
	BindStat            int                               `json:"bind_stat"`
	EnglishName         string                            `json:"english_name"`
	LoginStat           int                               `json:"login_stat"`
	Alias               string                            `json:"alias"`
	ActiveBiz           bool                              `json:"active_biz"`
	Domainid            int                               `json:"domainid"`
	Gid                 string                            `json:"gid"`
	Pinyin              string                            `json:"pinyin"`
	MainpartyID         string                            `json:"mainparty_id"`
	DisableBiz          bool                              `json:"disable_biz"`
	VidBindGid          bool                              `json:"vid_bind_gid"`
	CountryCode         string                            `json:"country_code"`
	WxIDHash            string                            `json:"wx_id_hash"`
	DepartNames         []interface{}                     `json:"depart_names"`
	HideMobile          bool                              `json:"hide_mobile"`
	DispOrder           []GetSingleMemberDispOrder        `json:"disp_order"`
	Acctid              string                            `json:"acctid"`
	Attrs               []interface{}                     `json:"attrs"`
	IsQuit              bool                              `json:"is_quit"`
	IsWwBizmailVip      bool                              `json:"is_ww_bizmail_vip"`
	IsWwBizmail         bool                              `json:"is_ww_bizmail"`
	IsJoinQyh           bool                              `json:"is_join_qyh"`
	AcctidStat          int                               `json:"acctid_stat"`
	PstnExtensionNumber string                            `json:"pstn_extension_number"`
	ExternJobTitle      string                            `json:"extern_job_title"`
	WxNickName          string                            `json:"wx_nick_name"`
	ExternalAttrs       []SaveMemberExternalAttr          `json:"external_attrs"`
	Realname            string                            `json:"realname"`
	UserQuitTime        int                               `json:"user_quit_time"`
	IsReadyJoinAgain    bool                              `json:"is_ready_join_again"`
	XcxCorpAddress      string                            `json:"xcx_corp_address"`
	ExternPositionInfo  GetSingleMemberExternPositionInfo `json:"extern_position_info"`
	ExternalCorpInfo    string                            `json:"external_corp_info"`
	ExternalWxfinder    GetSingleMemberExternalWxfinder   `json:"external_wxfinder"`
	PersonalEmail       string                            `json:"personal_email"`
	Superoirs           GetSingleMemberSuperoirs          `json:"superoirs"`
	BizMail             string                            `json:"biz_mail"`
	SuperoirList        []interface{}                     `json:"superoirList"`
	Uin                 string                            `json:"uin"`
	Username            string                            `json:"username"`
	Nickname            string                            `json:"nickname"`
	Imgid               string                            `json:"imgid"`
	PartyList           []string                          `json:"party_list"`
	JoinStatus          string                            `json:"JoinStatus"`
	Domain              string                            `json:"domain"`
	Tag                 GetSingleMemberTag                `json:"tag"`
	Extattr             GetSingleMemberExtattr            `json:"extattr"`
	GenderStr           string                            `json:"gender_str"`
}

// GetSingleMember  获取管理员列表
func (r *Client) GetSingleMember(vid string) (ret GetSingleMemberRes, err error) {
	cookie := r.ctx.Config.Cookie
	var header = commonPlatForm
	header["cookie"] = cookie
	if cookie == "" {
		err = errors.New("cookie 缺失")
		return
	}
	uri := fmt.Sprintf(getSingleMemberUrl, vid, time.Now().UnixMilli())
	rspOrigin, err := util.GetWithHeader(uri, header)
	err = json.Unmarshal(rspOrigin, &ret)
	if ret.Data.Vid == "" {
		err = errors.New("请求出错:" + string(rspOrigin))
		return
	}
	if err != nil {
		return
	}
	return
}

type CorpEncryptDataAppInfoRes struct {
	Data CorpEncryptDataAppInfoData `json:"data"`
}
type CorpEncryptDataAppInfoData struct {
	BillInfo CorpEncryptDataAppInfoBillInfo `json:"billInfo"`
}

type CorpEncryptDataAppInfoBillInfo struct {
	InfoList []CorpEncryptDataAppInfoInfoList `json:"info_list"`
}

type CorpEncryptDataAppInfoInfoList struct {
	Type            int    `json:"type"`
	Licensecnt      string `json:"licensecnt"` // 名额数量
	Begintime       string `json:"begintime"`
	Endtime         string `json:"endtime"`
	Corptypeflag    int    `json:"corptypeflag"`
	Usecnt          string `json:"usecnt"` // 已使用数量
	ShrinkLicCnt    int    `json:"shrink_lic_cnt"`
	ShrinkBegintime int    `json:"shrink_begintime"`
	Name            string `json:"name"`
}

// GetCorpEncryptDataAppInfo  获取会话存档数量
func (r *Client) GetCorpEncryptDataAppInfo() (ret CorpEncryptDataAppInfoRes, err error) {
	cookie := r.ctx.Config.Cookie
	var header = commonPlatForm
	header["cookie"] = cookie
	if cookie == "" {
		err = errors.New("cookie 缺失")
		return
	}
	uri := fmt.Sprintf(getCorpEncryptDataAppInfoUrl, time.Now().UnixMilli())
	rspOrigin, err := util.GetWithHeader(uri, header)
	err = json.Unmarshal(rspOrigin, &ret)
	if err != nil {
		fmt.Println(string(rspOrigin))
		return
	}
	return
}

type GetApplyInfoResp struct {
	Data GetApplyInfoData `json:"data"`
}
type GetApplyInfoData struct {
	Application []GetApplyInfo `json:"application"`
	Total       string         `json:"total"`
}

type GetApplyInfo struct {
	Status     string `json:"status"`
	CreateTime int    `json:"create_time"`
	ApplyTime  int    `json:"apply_time"`
	Vid        string `json:"vid"`
	Name       string `json:"name"`
	Mobile     string `json:"mobile"`
	HosterVid  string `json:"hoster_vid"`
	Extra      Extra  `json:"extra"`
}

type Extra struct {
	Remark string `json:"remark"`
}

// GetApplyList  获取邀请客户进企业的列表
func (r *Client) GetApplyList(start, limit int) (ret GetApplyInfoResp, err error) {
	cookie := r.ctx.Config.Cookie
	var header = commonPlatForm
	header["cookie"] = cookie
	if cookie == "" {
		err = errors.New("cookie 缺失")
		return
	}
	uri := fmt.Sprintf(getApplyList, time.Now().UnixMilli(), limit, start)
	rspOrigin, err := util.GetWithHeader(uri, header)
	err = json.Unmarshal(rspOrigin, &ret)
	if ret.Data.Total == "0" {
		err = errors.New("请求出错:" + string(rspOrigin))
		return
	}
	if err != nil {
		return
	}
	return
}
