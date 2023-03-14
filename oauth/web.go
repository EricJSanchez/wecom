package oauth

import (
	"encoding/json"
	"fmt"
	"github.com/EricJSanchez/wecom/credential"
	"github.com/EricJSanchez/wecom/util"
	"net/url"
)

const (
	// 获取授权链接
	authorizeAddr = "https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&" +
		"response_type=code&scope=snsapi_base&state=%s#wechat_redirect"
	// 获取用户信息
	getuserinfoAddr = "https://qyapi.weixin.qq.com/cgi-bin/user/getuserinfo?access_token=%s&code=%s"
	// 获取用户敏感信息
	getUserDetailAddr = "https://qyapi.weixin.qq.com/cgi-bin/user/getuserdetail?access_token=%s"
)

// AuthorizeOptions 获取授权链接请求参数
type AuthorizeOptions struct {
	RedirectUri string `json:"redirect_uri"`
	CorpId      string `json:"corp_id"`
	//State       string `json:"state"`
}

// AuthorizeSchema 获取授权链接响应内容
type AuthorizeSchema struct {
	util.CommonError
	AuthorizeUrl string `json:"authorize_url"`
}

// Authorize 获取授权链接
func (r *Client) Authorize(options AuthorizeOptions) (info AuthorizeSchema, err error) {
	info = AuthorizeSchema{
		CommonError: util.CommonError{
			ErrCode: 0,
			ErrMsg:  "",
		},
		AuthorizeUrl: fmt.Sprintf(authorizeAddr, options.CorpId, url.QueryEscape(options.RedirectUri+"?corp_id="+options.CorpId), util.RandomStr(16)),
	}
	return info, nil
}

// GetuserinfoOptions 获取用户信息请求参数
type GetuserinfoOptions struct {
	Code string `json:"code"`
}

// GetuserinfoSchema 获取用户信息响应内容
type GetuserinfoSchema struct {
	util.CommonError
	UserId         string `json:"UserId"`
	UserTicket     string `json:"user_ticket"`
	DeviceId       string `json:"DeviceId"`
	OpenId         string `json:"OpenId"`
	ExternalUserid string `json:"external_userid"`
}

// Getuserinfo 获取用户信息
func (r *Client) Getuserinfo(options GetuserinfoOptions) (info GetuserinfoSchema, err error) {
	var (
		accessToken string
		data        []byte
	)
	_ = util.Record(r.cache, getuserinfoAddr)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.HTTPGet(fmt.Sprintf(getuserinfoAddr, accessToken, options.Code))
	if err != nil {
		return
	}
	fmt.Println("data", string(data))
	if err = json.Unmarshal(data, &info); err != nil {
		fmt.Println("Unmarshal err", err)
		return
	}
	if info.ErrCode != 0 {
		return info, NewSDKErr(info.ErrCode, info.ErrMsg)
	}
	return info, nil
}

type TicketSchema struct {
	util.CommonError
	Ticket string `json:"ticket"`
}

func (r *Client) GetTicket() (info TicketSchema, err error) {
	info = TicketSchema{
		CommonError: util.CommonError{
			ErrCode: 0,
			ErrMsg:  "",
		},
	}
	nwWeworkJsTicket := credential.NewWeworkJsTicket(r.ctx.AgentID, "", r.ctx.Cache)
	accessToken, err := r.ctx.GetAccessToken()
	fmt.Println("accessToken", accessToken)
	info.Ticket, err = nwWeworkJsTicket.GetTicket(accessToken)
	//info.Ticket, err = nwWeworkJsTicket.GetTicket("SRddpwR4OWFB-kKVFCJA3ZkOCSO_znRex67mfmUZ1sGQr-y5cC8oYyFBRxXYr5Snph3XPqv_vLvjuajsDeIJGDVcWCHrWE8_XENQY6UT7_3ayBt0ZcRMJ3OFXbr14Wo8uyFFomqlSW7b9Dr9hxJReX48pcsj9ypWYoAvTzy73VIG0VMQ_lrRwU6lwYxnxAhxY3s1FtZ_WulR5f505Unijg")
	//info.Ticket, err = nwWeworkJsTicket.GetTicket("azMcHk-SgteiUtw_sEHDX5NpEl4oQ85SkHqDXvTu2garIfW63zMydDhNzz9KE7eTKoGewS9tnrFhO8fR1v5K8DqWEr_M5QgbPXp5AOE57MbAr9GurpItgCPXUHx6hALsxVcqg2NaJYRBJr2L_n0BsJm0BqWJJZ2lc9ji26I24KcRJfpkxuH8R_7Kwhgdq2csbQfw8frjX08_jfXpxFal8g")
	//info.Ticket, err = nwWeworkJsTicket.GetTicket("0E6ReEMLFmISJJfMYqASHsyn7_pVS0gkJzOFpPwuXIbsvHEWSwOTYCkqz-eQtNzdpcs-pk0Q9yrGHQIxad3nLuTFDnYT29AlSOI-Ces_OzvapQV3BZMMGYw8J0Ymd3XMKVt82NQVx1_3_zEK5p2seD1c01eXG2h__uW-3-DqUfADm1IDJj-QlRHDcEj3ouqAKe4QmVkldsvPpWHkIXgMjQ")

	return info, err
}

type GetUserDetailOptions struct {
	UserTicket string `json:"user_ticket"`
}

type GetUserDetailSchema struct {
	util.CommonError
	Userid  string `json:"userid"`
	Gender  string `json:"gender"`
	Avatar  string `json:"avatar"`
	QrCode  string `json:"qr_code"`
	Mobile  string `json:"mobile"`
	Email   string `json:"email"`
	BizMail string `json:"biz_mail"`
	Address string `json:"address"`
}

func (r *Client) GetUserDetail(options GetUserDetailOptions) (info GetUserDetailSchema, err error) {
	var (
		accessToken string
		data        []byte
	)
	_ = util.Record(r.cache, getUserDetailAddr)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	res, err := json.Marshal(options)
	if err != nil {
		return
	}
	data, err = util.HTTPPost(fmt.Sprintf(getUserDetailAddr, accessToken), string(res))
	if err != nil {
		return
	}
	fmt.Println("data", string(data))
	if err = json.Unmarshal(data, &info); err != nil {
		fmt.Println("Unmarshal err", err)
		return
	}
	if info.ErrCode != 0 {
		return info, NewSDKErr(info.ErrCode, info.ErrMsg)
	}
	return info, nil
}
