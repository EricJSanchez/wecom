package contact

import (
	"encoding/json"
	"fmt"
	"github.com/EricJSanchez/wecom/util"
)

const (
	// 创建标签
	tagCreateAddr = "https://qyapi.weixin.qq.com/cgi-bin/tag/create?access_token=%s"
	// 更新标签名字
	tagUpdateAddr = "https://qyapi.weixin.qq.com/cgi-bin/tag/update?access_token=%s"
	// 删除标签
	tagDeleteAddr = "https://qyapi.weixin.qq.com/cgi-bin/tag/delete?access_token=%s&tagid=%d"
	// 获取标签成员
	tagGetAddr = "https://qyapi.weixin.qq.com/cgi-bin/tag/get?access_token=%s&tagid=%d"
	// 增加标签成员
	tagAddtagusersAddr = "https://qyapi.weixin.qq.com/cgi-bin/tag/addtagusers?access_token=%s"
	// 删除标签成员
	tagDeltagusersAddr = "https://qyapi.weixin.qq.com/cgi-bin/tag/deltagusers?access_token=%s"
	// 获取标签列表
	tagListAddr = "https://qyapi.weixin.qq.com/cgi-bin/tag/list?access_token=%s"
)

type TagCommonError struct {
	CommonError CommonError `json:"common_error"`
	TagId       int         `json:"tagid"`
}
type CommonError struct {
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type CreateTagOption struct {
	TagId   int    `json:"tagid"`
	TagName string `json:"tagname"`
}

func (r *Client) CreateTag(options CreateTagOption) (info TagCommonError, err error) {
	var (
		accessToken string
		data        []byte
	)
	_ = util.Record(r.cache, tagCreateAddr)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(tagCreateAddr, accessToken), options)
	if err != nil {
		return
	}
	if err = json.Unmarshal(data, &info); err != nil {
		return
	}
	if info.CommonError.ErrCode != 0 {
		return info, NewSDKErr(info.CommonError.ErrCode, info.CommonError.ErrMsg)
	}
	return info, err
}
