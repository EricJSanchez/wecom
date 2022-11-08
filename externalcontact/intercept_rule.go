package externalcontact

import (
	"encoding/json"
	"fmt"
	"github.com/EricJSanchez/wecom/util"
)

const (
	//获取客户群详情
	addInterceptRuleAddr    = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/add_intercept_rule?access_token=%s"
	updateInterceptRuleAddr = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/update_intercept_rule?access_token=%s"
	delInterceptRuleAddr    = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/del_intercept_rule?access_token=%s"
)

type AddInterceptRuleOptions struct {
	RuleName        string          `json:"rule_name"`
	WordList        []string        `json:"word_list"`
	SemanticsList   []int           `json:"semantics_list"`
	InterceptType   int             `json:"intercept_type"`
	ApplicableRange ApplicableRange `json:"applicable_range"`
}

type ApplicableRange struct {
	UserList       []string `json:"user_list"`
	DepartmentList []int    `json:"department_list"`
}

type InterceptRuleSchema struct {
	util.CommonError
	RuleId string `json:"rule_id"`
}

func (r *Client) AddInterceptRule(options AddInterceptRuleOptions) (info InterceptRuleSchema, err error) {
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
	data, err = util.HTTPPost(fmt.Sprintf(addInterceptRuleAddr, accessToken), string(optionJson))
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

type UpdateInterceptRuleOptions struct {
	RuleID                string                `json:"rule_id"`
	RuleName              string                `json:"rule_name"`
	WordList              []string              `json:"word_list"`
	ExtraRule             ExtraRule             `json:"extra_rule"`
	InterceptType         int                   `json:"intercept_type"`
	AddApplicableRange    AddApplicableRange    `json:"add_applicable_range"`
	RemoveApplicableRange RemoveApplicableRange `json:"remove_applicable_range"`
}
type ExtraRule struct {
	SemanticsList []int `json:"semantics_list"`
}
type AddApplicableRange struct {
	UserList       []string `json:"user_list"`
	DepartmentList []int    `json:"department_list"`
}
type RemoveApplicableRange struct {
	UserList       []string `json:"user_list"`
	DepartmentList []int    `json:"department_list"`
}

func (r *Client) UpdateInterceptRule(options UpdateInterceptRuleOptions) (info util.CommonError, err error) {
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
	data, err = util.HTTPPost(fmt.Sprintf(updateInterceptRuleAddr, accessToken), string(optionJson))
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

type DelInterceptRuleOptions struct {
	RuleID string `json:"rule_id"`
}

func (r *Client) DelInterceptRule(options DelInterceptRuleOptions) (info util.CommonError, err error) {
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
	data, err = util.HTTPPost(fmt.Sprintf(delInterceptRuleAddr, accessToken), string(optionJson))
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
