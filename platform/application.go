package platform

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/EricJSanchez/wecom/util"
	"net/url"
	"strings"
	"time"
)

const (
	getCorpApplicationUrl = "https://work.weixin.qq.com/wework_admin/getCorpApplication?lang=zh_CN&f=json&ajax=1&timeZoneInfo[zone_offset]=-8&random=0.%d"
	saveOpenApiAppUrl     = "https://work.weixin.qq.com/wework_admin/apps/saveOpenApiApp?lang=zh_CN&f=json&ajax=1&timeZoneInfo[zone_offset]=-8&random=0.%d"
	saveIpConfigUrl       = "https://work.weixin.qq.com/wework_admin/apps/saveIpConfig?lang=zh_CN&f=json&ajax=1&timeZoneInfo[zone_offset]=-8&random=0.%d"
	getAppAdminInfoUrl    = "https://work.weixin.qq.com/wework_admin/apps/getAppAdminInfo?lang=zh_CN&f=json&ajax=1&timeZoneInfo[zone_offset]=-8&random=0.%d&app_id=%s&_d2st="
)

var commonPlatForm1 = map[string]string{
	"referer": "https://work.weixin.qq.com/wework_admin/frame",
	"origin":  "https://work.weixin.qq.com",
}

func Pr(val ...interface{}) {
	if len(val) > 1 {
		for _, v := range val {
			switch v.(type) {
			case []uint8:
				fmt.Println("[]uint8 ori: ", v)
				fmt.Printf("[]uint8 str: %s-\n\n", v)
				continue
			default:
				bytes, _ := json.MarshalIndent(v, "", "    ")
				fmt.Printf("%T : %s-\n", v, bytes)
			}
		}
	} else {
		bytes, _ := json.MarshalIndent(val, "", "    ")
		fmt.Printf("%T : %s-\n", val, bytes)
	}
}

type CorpApplicationRes struct {
	Data CorpApplicationResData `json:"data"`
}
type OpenapiApp struct {
	CorpID    string `json:"corp_id"`
	AppID     string `json:"app_id"`
	AppOpenID int    `json:"app_open_id"`
	Name      string `json:"name"`
	AppOpen   int    `json:"app_open"`
}
type CorpApplicationResData struct {
	OpenapiApp []OpenapiApp `json:"openapi_app"`
}

// GetCorpApplication 获取所有应用列表
func (r *Client) GetCorpApplication() (corpApp CorpApplicationRes, err error) {
	//var accessToken string
	cookie := r.ctx.Config.Cookie
	var header = commonPlatForm
	header["cookie"] = cookie
	if cookie == "" {
		err = errors.New("cookie 缺失")
		return
	}
	rspOrigin, err := util.PostFormEncodeWithHeader(fmt.Sprintf(getCorpApplicationUrl, time.Now().UnixMilli()), map[string]string{
		"app_type": "0",
		"_d2st":    "",
	}, header)
	err = json.Unmarshal(rspOrigin, &corpApp)
	if len(corpApp.Data.OpenapiApp) == 0 {
		err = errors.New("请求出错:" + string(rspOrigin))
		return
	}
	if err != nil {
		return
	}
	return
}

type SaveOpenApiAppRes struct {
	Data SaveOpenApiAppResData `json:"data"`
}
type SaveOpenApiAppResData struct {
	RejectSubadminIds []interface{} `json:"reject_subadmin_ids"`
}

// SettingApiCallback 设置应用API 接收
func (r *Client) SettingApiCallback(appId, callbackUrl, token, eAesKey string) (setRes SaveOpenApiAppRes, err error) {
	rTmp, err := url.Parse(callbackUrl)
	if err != nil {
		return
	}
	callBackHost := strings.Split(rTmp.Host, ":")
	if len(callBackHost) == 0 {
		err = errors.New("url解析错误")
	}
	//var accessToken string
	cookie := r.ctx.Config.Cookie
	var header = commonPlatForm
	header["cookie"] = cookie
	if cookie == "" {
		err = errors.New("cookie 缺失")
		return
	}
	postData := map[string]string{
		"callback_url":          callbackUrl,
		"url_token":             token,
		"callback_aeskey":       eAesKey,
		"report_approval_event": "true",
		"customer_event":        "true",
		"app_id":                appId,
		"callback_host":         callBackHost[0],
		"report_loc_flag":       "0",
		"is_report_enter":       "false",
		"_d2st":                 "",
	}
	rspOrigin, err := util.PostFormEncodeWithHeader(fmt.Sprintf(saveOpenApiAppUrl, time.Now().UnixMilli()), postData, header)
	Pr(postData, string(rspOrigin))
	if strings.Contains(string(rspOrigin), "errCode") {
		err = errors.New("请求出错:" + string(rspOrigin))
		return
	}
	err = json.Unmarshal(rspOrigin, &setRes)
	if err != nil {
		return
	}
	return
}

type SaveIpWhiteListRes struct {
	Data SaveIpWhiteListResData `json:"data"`
}
type SaveIpWhiteListResData struct {
	RejectSubadminIds []interface{} `json:"reject_subadmin_ids"`
}

// SaveIpWhiteList 设置应用API白名单
func (r *Client) SaveIpWhiteList(appId string, ipWhiteList []string) (setRes SaveIpWhiteListRes, err error) {
	cookie := r.ctx.Config.Cookie
	var header = commonPlatForm
	header["cookie"] = cookie
	if cookie == "" {
		err = errors.New("cookie 缺失")
		return
	}
	DataUrlVal := url.Values{}
	DataUrlVal.Add("app_id", appId)
	for _, ip := range ipWhiteList {
		DataUrlVal.Add("ipList[]", ip)
	}
	DataUrlVal.Add("_d2st", "")
	rspOrigin, err := util.PostFormEncodeStringWithHeader(fmt.Sprintf(saveIpConfigUrl, time.Now().UnixMilli()), DataUrlVal.Encode(), header)
	Pr(DataUrlVal.Encode(), string(rspOrigin))
	if strings.Contains(string(rspOrigin), "errCode") {
		err = errors.New("请求出错:" + string(rspOrigin))
		return
	}
	err = json.Unmarshal(rspOrigin, &setRes)
	if err != nil {
		return
	}
	return
}

type GetAppAdminInfoRes struct {
	Data GetAppAdminInfoData `json:"data"`
}
type GetAppAdminInfoData struct {
	Info []GetAppAdminInfoInfo `json:"info"`
}

type GetAppAdminInfoInfo struct {
	Vid      string               `json:"vid"`
	Flags    int                  `json:"flags"`
	RoleType int                  `json:"role_type"`
	Model    GetAppAdminInfoModel `json:"model,omitempty"`
}
type GetAppAdminInfoModel struct {
	CorpID string `json:"corp_id"`
	Vid    string `json:"vid"`
	Name   string `json:"name"`
	Acctid string `json:"acctid"`
	ID     string `json:"id"`
}

// GetAppAdminInfo 获取应用管理员
func (r *Client) GetAppAdminInfo(appId string) (adminRes GetAppAdminInfoRes, err error) {
	cookie := r.ctx.Config.Cookie
	var header = commonPlatForm
	header["cookie"] = cookie
	if cookie == "" {
		err = errors.New("cookie 缺失")
		return
	}

	rspOrigin, err := util.GetWithHeader(fmt.Sprintf(getAppAdminInfoUrl, time.Now().UnixMilli(), appId), header)
	Pr(string(rspOrigin))
	if strings.Contains(string(rspOrigin), "errCode") {
		err = errors.New("请求出错:" + string(rspOrigin))
		return
	}
	err = json.Unmarshal(rspOrigin, &adminRes)
	if err != nil {
		return
	}
	return
}
