package oauth

import (
	"fmt"
	"github.com/EricJSanchez/wecom/util"
	"net/url"
)

const (
	// 获取二维码链接
	qrConnectAddr = "https://open.work.weixin.qq.com/wwopen/sso/qrConnect?appid=%s&agentid=%s&" +
		"redirect_uri=%s&state=%s"
)

// QrConnectOptions 获取二维码链接请求参数
type QrConnectOptions struct {
	RedirectUri string `json:"redirect_uri"`
	CorpId      string `json:"corp_id"`
	AgentId     string `json:"agent_id"`
}

// QrConnectSchema 获取二维码链接响应内容
type QrConnectSchema struct {
	util.CommonError
	QrConnectUrl string `json:"qr_connect_url"`
}

// QrConnect 获取二维码链接
func (r *Client) QrConnect(options QrConnectOptions) (info QrConnectSchema, err error) {
	info = QrConnectSchema{
		CommonError: util.CommonError{
			ErrCode: 0,
			ErrMsg:  "",
		},
		QrConnectUrl: fmt.Sprintf(qrConnectAddr, options.CorpId, options.AgentId,
			url.QueryEscape(fmt.Sprintf(options.RedirectUri, options.CorpId)), util.RandomStr(16)),
	}
	return info, nil
}
