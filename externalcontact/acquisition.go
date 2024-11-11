package externalcontact

import (
	"encoding/json"
	"fmt"
	"github.com/EricJSanchez/wecom/util"
)

const (
	customerAcquisitionQuotaListLinkAddr = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/customer_acquisition/list_link?access_token=%s"
	customerAcquisitionQuotaGetAddr      = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/customer_acquisition/get?access_token=%s"
	// 查询剩余使用量
	customerAcquisitionQuotaAddr     = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/customer_acquisition_quota?access_token=%s"
	customerAcquisitionStatisticAddr = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/customer_acquisition/statistic?access_token=%s"
)

type AcquisitionQuotaLinkListOptions struct {
	Limit  int    `json:"limit"`
	Cursor string `json:"cursor"`
}

type AcquisitionQuotaLinkListSchema struct {
	util.CommonError
	LinkIDList []string `json:"link_id_list"`
	NextCursor string   `json:"next_cursor"`
}

func (r *Client) AcquisitionQuotaLinkList(options AcquisitionQuotaLinkListOptions) (info AcquisitionQuotaLinkListSchema, err error) {
	var (
		accessToken string
		data        []byte
	)
	_ = util.Record(r.cache, customerAcquisitionQuotaListLinkAddr)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	optionJson, err := json.Marshal(options)
	if err != nil {
		return
	}
	data, err = util.HTTPPost(fmt.Sprintf(customerAcquisitionQuotaListLinkAddr, accessToken), string(optionJson))
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

type AcquisitionQuotaGetOptions struct {
	LinkId string `json:"link_id"`
}

type AcquisitionQuotaGetSchema struct {
	util.CommonError
	Link           QuotaGetLink           `json:"link"`
	Range          QuotaGetRange          `json:"range"`
	PriorityOption QuotaGetPriorityOption `json:"priority_option"`
}
type QuotaGetLink struct {
	LinkName   string `json:"link_name"`
	URL        string `json:"url"`
	CreateTime int    `json:"create_time"`
	SkipVerify bool   `json:"skip_verify"`
}
type QuotaGetRange struct {
	UserList       []string `json:"user_list"`
	DepartmentList []int    `json:"department_list"`
}
type QuotaGetPriorityOption struct {
	PriorityType       int      `json:"priority_type"`
	PriorityUseridList []string `json:"priority_userid_list"`
}

func (r *Client) AcquisitionQuotaGet(options AcquisitionQuotaGetOptions) (info AcquisitionQuotaGetSchema, err error) {
	var (
		accessToken string
		data        []byte
	)
	_ = util.Record(r.cache, customerAcquisitionQuotaGetAddr)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	optionJson, err := json.Marshal(options)
	if err != nil {
		return
	}
	data, err = util.HTTPPost(fmt.Sprintf(customerAcquisitionQuotaGetAddr, accessToken), string(optionJson))
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

type AcquisitionQuotaSchema struct {
	util.CommonError
	Total     int         `json:"total"`
	Balance   int         `json:"balance"`
	QuotaList []QuotaList `json:"quota_list"`
}
type QuotaList struct {
	ExpireDate int `json:"expire_date"`
	Balance    int `json:"balance"`
}

func (r *Client) AcquisitionQuota() (list AcquisitionQuotaSchema, err error) {
	accessToken, err := r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err := util.HTTPGet(fmt.Sprintf(customerAcquisitionQuotaAddr, accessToken))
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

type AcquisitionQuotaStatisticOptions struct {
	LinkID    string `json:"link_id"`
	StartTime int    `json:"start_time"`
	EndTime   int    `json:"end_time"`
}

type AcquisitionQuotaStatisticSchema struct {
	util.CommonError
	ClickLinkCustomerCnt int `json:"click_link_customer_cnt"`
	NewCustomerCnt       int `json:"new_customer_cnt"`
}

func (r *Client) AcquisitionQuotaStatistic(options AcquisitionQuotaStatisticOptions) (info AcquisitionQuotaStatisticSchema, err error) {
	var (
		accessToken string
		data        []byte
	)
	_ = util.Record(r.cache, customerAcquisitionStatisticAddr)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	optionJson, err := json.Marshal(options)
	if err != nil {
		return
	}
	data, err = util.HTTPPost(fmt.Sprintf(customerAcquisitionStatisticAddr, accessToken), string(optionJson))
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
