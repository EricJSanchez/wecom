package platform

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/EricJSanchez/wecom/util"
	"time"
)

const (
	deptCacheUrl       = "https://work.weixin.qq.com/wework_admin/contacts/party/cache?lang=zh_CN&f=json&ajax=1&timeZoneInfo[zone_offset]=-8&random=0.%d"
	deptStaffListUrl   = "https://work.weixin.qq.com/wework_admin/contacts/getDepartMember?lang=zh_CN&f=json&ajax=1&timeZoneInfo[zone_offset]=-8&random=0.%d&action=getpartycontacts&partyid=%s&page=%d&limit=%d&joinstatus=0&fetchchild=1&preFetch=false&use_corp_cache=0&_d2st=a2300212"
	searchStaffListUrl = "https://work.weixin.qq.com/wework_admin/search_contact?lang=zh_CN&f=json&ajax=1&timeZoneInfo[zone_offset]=-8&random=0.%d"
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
	ContactList ContactList `json:"contact_list"`
	MemberCount int         `json:"member_count"`
	DisableCnt  int         `json:"disable_cnt"`
	PageCount   int         `json:"page_count"`
	Status      string      `json:"status"`
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
	err = json.Unmarshal(rspOrigin, &searchStaff)
	return
}
