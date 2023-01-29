package application

import (
	"encoding/xml"
	"fmt"
	"github.com/EricJSanchez/wecom/util"
	"github.com/spf13/cast"
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
	if options.Signature != util.Signature(r.ctx.Token, options.TimeStamp, options.Nonce, options.EchoStr, options.Encrypt) {
		return "", NewSDKErr(40015)
	}
	if options.EchoStr != "" {
		_, bData, err := util.DecryptMsg(r.corpID, options.EchoStr, r.encodingAESKey)
		fmt.Println(err)
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

// AddExternalContactCallbackMessage 添加企业客户事件
type AddExternalContactCallbackMessage struct {
	CallbackMessage
	UserID         string `json:"user_id"`
	ExternalUserID string `json:"external_user_id"`
	State          string `json:"state"`
	WelcomeCode    string `json:"welcome_code"`
}

// EditExternalContactCallbackMessage 编辑企业客户事件
type EditExternalContactCallbackMessage struct {
	CallbackMessage
	UserID         string `json:"user_id"`
	ExternalUserID string `json:"external_user_id"`
}

// AddHalfExternalContactCallbackMessage 外部联系人免验证添加成员事件
type AddHalfExternalContactCallbackMessage struct {
	CallbackMessage
	UserID         string `json:"user_id"`
	ExternalUserID string `json:"external_user_id"`
	State          string `json:"state"`
	WelcomeCode    string `json:"welcome_code"`
}

// DelExternalContactCallbackMessage 删除企业客户事件
type DelExternalContactCallbackMessage struct {
	CallbackMessage
	UserID         string `json:"user_id"`
	ExternalUserID string `json:"external_user_id"`
	Source         string `json:"source"`
}

// DelFollowUserCallbackMessage 删除跟进成员事件
type DelFollowUserCallbackMessage struct {
	CallbackMessage
	UserID         string `json:"user_id"`
	ExternalUserID string `json:"external_user_id"`
}

// TransferFailCallbackMessage 客户接替失败事件
type TransferFailCallbackMessage struct {
	CallbackMessage
	FailReason     string `json:"fail_reason"`
	UserID         string `json:"user_id"`
	ExternalUserID string `json:"external_user_id"`
}

// 客户同意进行聊天内容存档事件回调
type AuditApprovedCallbackMessage struct {
	CallbackMessage
	UserID         string `json:"user_id"`
	ExternalUserID string `json:"external_user_id"`
	WelcomeCode    string `json:"welcome_code"`
}

// CreateCallbackMessage 客户群创建事件
type CreateCallbackMessage struct {
	CallbackMessage
	ChatId string `json:"ChatId"`
}

// UpdateCallbackMessage 客户群变更事件
type UpdateCallbackMessage struct {
	CallbackMessage
	ChatId       string `json:"chat_id"`
	UpdateDetail string `json:"update_detail"`
	JoinScene    int    `json:"join_scene"`
	QuitScene    int    `json:"quit_scene"`
	MemChangeCnt int    `json:"mem_change_cnt"`
}

// DismissCallbackMessage 客户群解散事件
type DismissCallbackMessage struct {
	CallbackMessage
	ChatId string `json:"chat_id"`
}

// TagCreateCallbackMessage 企业客户标签创建事件
type TagCreateCallbackMessage struct {
	CallbackMessage
	Id      string `json:"id"`
	TagType string `json:"tag_type"`
}

// TagUpdateCallbackMessage 企业客户标签变更事件
type TagUpdateCallbackMessage struct {
	CallbackMessage
	Id      string `json:"id"`
	TagType string `json:"tag_type"`
}

// TagDeleteCallbackMessage 企业客户标签删除事件
type TagDeleteCallbackMessage struct {
	CallbackMessage
	Id      string `json:"id"`
	TagType string `json:"tag_type"`
}

// TagShuffleCallbackMessage 企业客户标签重排事件
type TagShuffleCallbackMessage struct {
	CallbackMessage
	Id         string `json:"id"`
	StrategyID string `json:"strategy_id"`
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
		fmt.Println("application GetCallbackMessage Unmarshal 1:", err, cast.ToString(encryptedMsg))
		return rawData, msg, err
	}
	signatureOptions.Encrypt = origin.Encrypt
	_, err = r.VerifyURL(signatureOptions)
	if err != nil {
		fmt.Println("GetCallbackMessage VerifyURL:", err)
		return rawData, msg, err
	}

	_, rawData, err = util.DecryptMsg(r.corpID, origin.Encrypt, r.encodingAESKey)
	if err != nil {
		fmt.Println("GetCallbackMessage DecryptMsg:", err)
		return rawData, msg, NewSDKErr(40016)
	}
	if err = xml.Unmarshal(rawData, &msg); err != nil {
		fmt.Println("application GetCallbackMessage Unmarshal 2:", err, cast.ToString(rawData))
		return rawData, msg, err
	}

	return rawData, msg, err
}
