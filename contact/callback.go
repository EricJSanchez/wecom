package contact

import (
	"encoding/xml"
	"fmt"
	"github.com/EricJSanchez/wecom/util"
)

// SignatureOptions 微信服务器验证参数
type SignatureOptions struct {
	Signature string `form:"msg_signature"`
	TimeStamp string `form:"timestamp"`
	Nonce     string `form:"nonce"`
	EchoStr   string `form:"echostr"`
	Encrypt   string `form:"encrypt"`
}

// VerifyURL 验证请求参数是否合法并返回解密后的消息内容
//  //Gin框架的使用示例
//	r.GET("/v1/event/callback", func(c *gin.Context) {
//		options := kf.SignatureOptions{}
//		//获取回调的的校验参数
//		if = c.ShouldBindQuery(&options); err != nil {
//			c.String(http.StatusUnauthorized, "参数解析失败")
//		}
//		// 调用VerifyURL方法校验当前请求，如果合法则把解密后的内容作为响应返回给微信服务器
//		echo, err := kfClient.VerifyURL(options)
//		if err == nil {
//			c.String(http.StatusOK, echo)
//		} else {
//			c.String(http.StatusUnauthorized, "非法请求来源")
//		}
//	})
func (r *Client) VerifyURL(options SignatureOptions) (string, error) {
	//fmt.Println("options.Signature:", options.Signature)
	//fmt.Println("util.Signature:", util.Signature(r.ctx.Token, options.TimeStamp, options.Nonce, options.EchoStr, options.Encrypt))
	//fmt.Println("r.ctx.Token:", r.ctx.Token)
	//fmt.Println("options.TimeStamp:", options.TimeStamp)
	//fmt.Println("options.Nonce:", options.Nonce)
	//fmt.Println("options.EchoStr:", options.EchoStr)
	//fmt.Println("options.Encrypt:", options.Encrypt)
	if options.Signature != util.Signature(r.ctx.Token, options.TimeStamp, options.Nonce, options.EchoStr, options.Encrypt) {
		return "", NewSDKErr(40015)
	}

	if options.EchoStr != "" {
		_, bData, err := util.DecryptMsg(r.corpID, options.EchoStr, r.encodingAESKey)
		if err != nil {
			return "", NewSDKErr(40016)
		}
		return string(bData), nil
	}

	return "", nil
}

// 原始回调消息内容
type callbackOriginMessage struct {
	ToUserName string // 企业微信的CorpID，当为第三方套件回调事件时，CorpID的内容为suiteid
	AgentID    string // 接收的应用id，可在应用的设置页面获取
	Encrypt    string // 消息结构体加密后的字符串
}

// CallbackMessage 基础回调消息
type CallbackMessage struct {
	ToUserName   string `json:"to_user_name"`   // 企业微信CorpID
	FromUserName string `json:"from_user_name"` // 此事件该值固定为sys，表示该消息由系统生成
	CreateTime   int    `json:"create_time"`    // 消息创建时间 （整型）
	MsgType      string `json:"msgtype"`        // 消息的类型，此时固定为event
	Event        string `json:"event"`          // 事件的类型，此时固定为change_contact
	ChangeType   string `json:"change_type"`    // 此时固定为delete_user
}

type CallbackMessageLog struct {
	UserID         string `json:"user_id"` // 成员UserID
	ExternalUserID string `json:"external_user_id"`
}

// CreateUserCallbackMessage 新增成员事件回调消息
type CreateUserCallbackMessage struct {
	CallbackMessage
	UserID         string                 `json:"user_id"`           // 成员UserID
	Name           string                 `json:"name"`              // 成员名称
	Department     string                 `json:"department"`        // 成员部门列表，仅返回该应用有查看权限的部门id
	MainDepartment int                    `json:"main_department"`   // 主部门
	IsLeaderInDept string                 `json:"is_leader_in_dept"` // 表示所在部门是否为上级，0-否，1-是，顺序与Department字段的部门逐一对应
	Mobile         string                 `json:"mobile"`            // 手机号码
	Position       string                 `json:"position"`          // 职位信息。长度为0~64个字节
	Gender         string                 `json:"gender"`            // 性别，1表示男性，2表示女性
	Email          string                 `json:"email"`             // 邮箱
	Status         int                    `json:"status"`            // 激活状态：1=已激活 2=已禁用 4=未激活 已激活代表已激活企业微信或已关注微工作台（原企业号）5=成员退出
	Avatar         string                 `json:"avatar"`            // 头像url。注：如果要获取小图将url最后的”/0”改成”/100”即可。
	Alias          string                 `json:"alias"`             // 成员别名
	Telephone      string                 `json:"telephone"`         // 座机
	Address        string                 `json:"address"`           // 地址
	ExtAttr        map[string]interface{} `json:"ext_attr"`          // 扩展属性
}

// UpdateUserCallbackMessage 更新成员事件回调消息
type UpdateUserCallbackMessage struct {
	CallbackMessage
	CreateUserCallbackMessage
	NewUserID string `json:"new_user_id"` // 新的UserID，变更时推送（userid由系统生成时可更改一次）
}

// DeleteUserCallbackMessage 删除成员事件回调消息
type DeleteUserCallbackMessage struct {
	CallbackMessage
	UserID string `json:"user_id"` // 成员UserID
}

// CreatePartyCallbackMessage 新增部门事件回调消息
type CreatePartyCallbackMessage struct {
	CallbackMessage
	Id       int    `json:"id"`        // 部门Id
	Name     string `json:"name"`      // 部门名称
	ParentId int    `json:"parent_id"` // 父部门id
	Order    int    `json:"order"`     // 部门排序
}

// UpdatePartyCallbackMessage 更新部门事件回调消息
type UpdatePartyCallbackMessage struct {
	CallbackMessage
	Id       int    `json:"id"`        // 部门Id
	Name     string `json:"name"`      // 部门名称，仅当该字段发生变更时传递
	ParentId int    `json:"parent_id"` // 父部门id，仅当该字段发生变更时传递
}

// DeletePartyCallbackMessage 删除部门事件回调消息
type DeletePartyCallbackMessage struct {
	CallbackMessage
	Id int `json:"id"` // 部门Id
}

// UpdateTagCallbackMessage 标签变更通知回调消息
type UpdateTagCallbackMessage struct {
	CallbackMessage
	TagId         int      `json:"tag_id"`          // 标签Id
	AddUserItems  []string `json:"add_user_items"`  // 标签中新增的成员userid列表，用逗号分隔
	DelUserItems  []string `json:"del_user_items"`  // 标签中删除的成员userid列表，用逗号分隔
	AddPartyItems []int    `json:"add_party_items"` // 标签中新增的部门id列表，用逗号分隔
	DelPartyItems []int    `json:"del_party_items"` // 标签中删除的部门id列表，用逗号分隔
}

// BatchJob 异步任务信息
type BatchJob struct {
	JobId   string `json:"job_id"`   // 异步任务id，最大长度为64字符
	JobType string `json:"job_type"` // 操作类型，字符串，目前分别有：sync_user(增量更新成员)、 replace_user(全量覆盖成员）、invite_user(邀请成员关注）、replace_party(全量覆盖部门)
	ErrCode int    `json:"err_code"` // 返回码
	ErrMsg  string `json:"err_msg"`  // 对返回码的文本描述内容
}

// BatchJobResultCallbackMessage 异步任务完成通知
type BatchJobResultCallbackMessage struct {
	CallbackMessage
	BatchJob BatchJob `json:"batch_job"` // 异步任务信息
}

// GetCallbackMessage 获取回调事件中的消息内容
//  //Gin框架的使用示例
//	r.POST("/v1/event/callback", func(c *gin.Context) {
//		var (
//			message kf.CallbackMessage
//			body []byte
//		)
//		// 读取原始消息内容
//		body, err = c.GetRawData()
//		if err != nil {
//			c.String(http.StatusInternalServerError, err.Error())
//			return
//		}
//		// 解析原始数据
//		message, err = kfClient.GetCallbackMessage(body)
//		if err != nil {
//			c.String(http.StatusInternalServerError, "消息获取失败")
//			return
//		}
//		fmt.Println(message)
//		c.String(200, "ok")
//	})
func (r *Client) GetCallbackMessage(signatureOptions SignatureOptions, encryptedMsg []byte) (rawData []byte, msg CallbackMessage, err error) {
	var origin callbackOriginMessage
	if err = xml.Unmarshal(encryptedMsg, &origin); err != nil {
		fmt.Println("GetCallbackMessage Unmarshal 1:", err)
		return rawData, msg, err
	}

	signatureOptions.Encrypt = origin.Encrypt
	_, err = r.VerifyURL(signatureOptions)
	if err != nil {
		return rawData, msg, err
	}

	_, rawData, err = util.DecryptMsg(r.corpID, origin.Encrypt, r.encodingAESKey)
	if err != nil {
		fmt.Println("GetCallbackMessage DecryptMsg:", err)
		return rawData, msg, NewSDKErr(40016)
	}
	if err = xml.Unmarshal(rawData, &msg); err != nil {
		fmt.Println("GetCallbackMessage Unmarshal 2:", err)
		return rawData, msg, err
	}

	return rawData, msg, err
}
