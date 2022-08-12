package message

import (
	"encoding/xml"
	"fmt"
	"wecom/util"
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
//
//	 //Gin框架的使用示例
//		r.GET("/v1/event/callback", func(c *gin.Context) {
//			options := kf.SignatureOptions{}
//			//获取回调的的校验参数
//			if = c.ShouldBindQuery(&options); err != nil {
//				c.String(http.StatusUnauthorized, "参数解析失败")
//			}
//			// 调用VerifyURL方法校验当前请求，如果合法则把解密后的内容作为响应返回给微信服务器
//			echo, err := kfClient.VerifyURL(options)
//			if err == nil {
//				c.String(http.StatusOK, echo)
//			} else {
//				c.String(http.StatusUnauthorized, "非法请求来源")
//			}
//		})
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

// callbackOriginMessage 原始回调消息内容
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

// ResponseMessage 基础响应消息
type ResponseMessage struct {
	ToUserName   string `json:"to_user_name"`
	FromUserName string `json:"from_user_name"`
	CreateTime   int    `json:"create_time"`
	MsgType      string `json:"msg_type"`
}

// SubscribeCallbackMessage 成员关注及取消关注事件
type SubscribeCallbackMessage struct {
	CallbackMessage
	AgentID int `json:"agent_id"`
}

// EnterAgentCallbackMessage 进入应用
type EnterAgentCallbackMessage struct {
	CallbackMessage
	EventKey string `json:"event_key"`
	AgentID  int    `json:"agent_id"`
}

// LocationCallbackMessage 上报地理位置
type LocationCallbackMessage struct {
	CallbackMessage
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Precision int     `json:"precision"`
	AgentID   int     `json:"agent_id"`
	AppType   string  `json:"app_type"`
}

type BatchJobApplication struct {
	JobID   string `json:"job_id"`
	JobType string `json:"job_type"`
	ErrCode int    `json:"err_code"`
	ErrMsg  string `json:"err_msg"`
}

// BatchJobResultApplicationCallbackMessage 异步任务完成事件推送
type BatchJobResultApplicationCallbackMessage struct {
	CallbackMessage
	BatchJob BatchJobApplication `json:"batch_job"`
}

// CreateUserCallbackMessage 新增成员事件回调消息
type CreateUserCallbackMessage struct {
	CallbackMessage
	UserID         string                 `json:"user_id"`           // 成员UserID
	Name           string                 `json:"name"`              // 成员名称
	Department     []int                  `json:"department"`        // 成员部门列表，仅返回该应用有查看权限的部门id
	MainDepartment int                    `json:"main_department"`   // 主部门
	IsLeaderInDept []int                  `json:"is_leader_in_dept"` // 表示所在部门是否为上级，0-否，1-是，顺序与Department字段的部门逐一对应
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

// ClickCallbackMessage 点击菜单拉取消息的事件推送
type ClickCallbackMessage struct {
	CallbackMessage
	EventKey string `json:"event_key"`
	AgentID  int    `json:"agent_id"`
}

// ViewCallbackMessage 点击菜单跳转链接的事件推送
type ViewCallbackMessage struct {
	CallbackMessage
	EventKey string `json:"event_key"`
	AgentID  int    `json:"agent_id"`
}

// ScancodeCallbackMessage 扫码推事件的事件推送 / 扫码推事件且弹出“消息接收中”提示框的事件推送
type ScancodeCallbackMessage struct {
	CallbackMessage
	EventKey     string       `json:"event_key"`
	ScanCodeInfo ScanCodeInfo `json:"scan_code_info"`
	AgentID      int          `json:"agent_id"`
}
type ScanCodeInfo struct {
	ScanType   string `json:"scan_type"`
	ScanResult string `json:"scan_result"`
}

// PicSysphotoCallbackMessage 弹出系统拍照发图的事件推送 / 弹出系统拍照发图的事件推送
type PicSysphotoCallbackMessage struct {
	CallbackMessage
	EventKey     string       `json:"event_key"`
	SendPicsInfo SendPicsInfo `json:"send_pics_info"`
	AgentID      int          `json:"agent_id"`
}
type Item struct {
	PicMd5Sum string `json:"pic_md_5_sum"`
}
type PicList struct {
	Item Item `json:"item"`
}
type SendPicsInfo struct {
	Count   int     `json:"count"`
	PicList PicList `json:"pic_list"`
}

// LocationSelectCallbackMessage 弹出地理位置选择器的事件推送
type LocationSelectCallbackMessage struct {
	CallbackMessage
	EventKey         string           `json:"event_key"`
	SendLocationInfo SendLocationInfo `json:"send_location_info"`
	AgentID          int              `json:"agent_id"`
	AppType          string           `json:"app_type"`
}
type SendLocationInfo struct {
	LocationX string `json:"location_x"`
	LocationY string `json:"location_y"`
	Scale     string `json:"scale"`
	Label     string `json:"label"`
	Poiname   string `json:"poiname"`
}

// OpenApprovalChangeCallbackMessage 审批状态通知事件
type OpenApprovalChangeCallbackMessage struct {
	CallbackMessage
	AgentID      int          `json:"agent_id"`
	ApprovalInfo ApprovalInfo `json:"approval_info"`
}
type OpenApprovalItem struct {
	ItemName   string `json:"item_name"`
	ItemUserID string `json:"item_user_id"`
	ItemImage  string `json:"item_image"`
	ItemStatus int    `json:"item_status"`
	ItemSpeech string `json:"item_speech"`
	ItemOpTime int    `json:"item_op_time"`
}
type Items struct {
	Item OpenApprovalItem `json:"item"`
}
type ApprovalNode struct {
	NodeStatus int   `json:"node_status"`
	NodeAttr   int   `json:"node_attr"`
	NodeType   int   `json:"node_type"`
	Items      Items `json:"items"`
}
type ApprovalNodes struct {
	ApprovalNode ApprovalNode `json:"approval_node"`
}
type NotifyNode struct {
	ItemName   string `json:"item_name"`
	ItemUserID string `json:"item_user_id"`
	ItemImage  string `json:"item_image"`
}
type NotifyNodes struct {
	NotifyNode NotifyNode `json:"notify_node"`
}
type ApprovalInfo struct {
	ThirdNo        string        `json:"third_no"`
	OpenSpName     string        `json:"open_sp_name"`
	OpenTemplateID string        `json:"open_template_id"`
	OpenSpStatus   int           `json:"open_sp_status"`
	ApplyTime      int           `json:"apply_time"`
	ApplyUserName  string        `json:"apply_user_name"`
	ApplyUserID    string        `json:"apply_user_id"`
	ApplyUserParty string        `json:"apply_user_party"`
	ApplyUserImage string        `json:"apply_user_image"`
	ApprovalNodes  ApprovalNodes `json:"approval_nodes"`
	NotifyNodes    NotifyNodes   `json:"notify_nodes"`
	Approverstep   int           `json:"approverstep"`
}

// ShareAgentChangeCallbackMessage 共享应用事件回调
type ShareAgentChangeCallbackMessage struct {
	CallbackMessage
	AgentID int `json:"agent_id"`
}

// TemplateCardEventCallbackMessage 模板卡片事件推送
type TemplateCardEventCallbackMessage struct {
	CallbackMessage
	EventKey      string        `json:"event_key"`
	TaskID        string        `json:"task_id"`
	CardType      string        `json:"card_type"`
	ResponseCode  string        `json:"response_code"`
	AgentID       int           `json:"agent_id"`
	SelectedItems SelectedItems `json:"selected_items"`
}
type OptionIds struct {
	OptionID []string `json:"option_id"`
}
type SelectedItem struct {
	QuestionKey string    `json:"question_key"`
	OptionIds   OptionIds `json:"option_ids"`
}
type SelectedItems struct {
	SelectedItem []SelectedItem `json:"selected_item"`
}

// TemplateCardMenuEventCallbackMessage 通用模板卡片右上角菜单事件推送
type TemplateCardMenuEventCallbackMessage struct {
	CallbackMessage
	EventKey     string `json:"event_key"`
	TaskID       string `json:"task_id"`
	CardType     string `json:"card_type"`
	ResponseCode string `json:"response_code"`
	AgentID      int    `json:"agent_id"`
}

// GetCallbackMessage 获取回调事件中的消息内容
//
//	 //Gin框架的使用示例
//		r.POST("/v1/event/callback", func(c *gin.Context) {
//			var (
//				message kf.CallbackMessage
//				body []byte
//			)
//			// 读取原始消息内容
//			body, err = c.GetRawData()
//			if err != nil {
//				c.String(http.StatusInternalServerError, err.Error())
//				return
//			}
//			// 解析原始数据
//			message, err = kfClient.GetCallbackMessage(body)
//			if err != nil {
//				c.String(http.StatusInternalServerError, "消息获取失败")
//				return
//			}
//			fmt.Println(message)
//			c.String(200, "ok")
//		})
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

// TextResponseMessage 文本消息
type TextResponseMessage struct {
	ResponseMessage
	Content string `json:"content"`
}

// ImageResponseMessage 图片消息
type ImageResponseMessage struct {
	ResponseMessage
	Image ResponseImage `json:"image"`
}
type ResponseImage struct {
	MediaID string `json:"media_id"`
}

// VoiceResponseMessage 语音消息
type VoiceResponseMessage struct {
	ResponseMessage
	Voice ResponseVoice `json:"voice"`
}
type ResponseVoice struct {
	MediaID string `json:"media_id"`
}

// VideoResponseMessage 视频消息
type VideoResponseMessage struct {
	ResponseMessage
	Video ResponseVideo `json:"video"`
}
type ResponseVideo struct {
	MediaID     string `json:"media_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// ArticleResponseMessage 图文消息
type ArticleResponseMessage struct {
	ResponseMessage
	ArticleCount int              `json:"article_count"`
	Articles     ResponseArticles `json:"articles"`
}
type ResponseItem struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	PicURL      string `json:"pic_url"`
	URL         string `json:"url"`
}
type ResponseArticles struct {
	Item []ResponseItem `json:"item"`
}

// ButtonResponseMessage 更新点击用户的按钮文案
type ButtonResponseMessage struct {
	ResponseMessage
	Button ResponseButton `json:"button"`
}
type ResponseButton struct {
	ReplaceName string `json:"replace_name"`
}

// TemplateCardTextResponseMessage 更新点击用户的整张卡片 - 文本通知型
type TemplateCardTextResponseMessage struct {
	ResponseMessage
	TemplateCard TextTemplateCard `json:"template_card"`
}
type TextSource struct {
	IconURL   string `json:"icon_url"`
	Desc      string `json:"desc"`
	DescColor int    `json:"desc_color"`
}
type TextMainTitle struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}
type TextHorizontalContentList struct {
	KeyName string `json:"key_name"`
	Value   string `json:"value"`
	Type    int    `json:"type,omitempty"`
	URL     string `json:"url,omitempty"`
}
type TextJumpList struct {
	Title string `json:"title"`
	Type  int    `json:"type"`
	URL   string `json:"url"`
}
type TextCardAction struct {
	Title string `json:"title"`
	Type  int    `json:"type"`
	URL   string `json:"url"`
}
type TextEmphasisContent struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}
type TextActionList struct {
	Text string `json:"text"`
	Key  string `json:"key"`
}
type TextActionMenu struct {
	Desc       string           `json:"desc"`
	ActionList []TextActionList `json:"action_list"`
}
type TextQuoteArea struct {
	Type      int    `json:"type"`
	URL       string `json:"url"`
	Title     string `json:"title"`
	QuoteText string `json:"quote_text"`
}
type TextTemplateCard struct {
	CardType              string                      `json:"card_type"`
	Source                TextSource                  `json:"source"`
	MainTitle             TextMainTitle               `json:"main_title"`
	SubTitleText          string                      `json:"sub_title_text"`
	HorizontalContentList []TextHorizontalContentList `json:"horizontal_content_list"`
	JumpList              TextJumpList                `json:"jump_list"`
	CardAction            TextCardAction              `json:"card_action"`
	EmphasisContent       TextEmphasisContent         `json:"emphasis_content"`
	ActionMenu            TextActionMenu              `json:"action_menu"`
	QuoteArea             TextQuoteArea               `json:"quote_area"`
}

// TemplateCardImageTextResponseMessage 更新点击用户的整张卡片 - 图文展示型
type TemplateCardImageTextResponseMessage struct {
	ResponseMessage
	TemplateCard ImageTextTemplateCard `json:"TemplateCard"`
}
type ImageTextSource struct {
	IconURL   string `json:"IconUrl"`
	Desc      string `json:"Desc"`
	DescColor int    `json:"DescColor"`
}
type ImageTextMainTitle struct {
	Title string `json:"Title"`
	Desc  string `json:"Desc"`
}
type ImageTextHorizontalContentList struct {
	KeyName string `json:"KeyName"`
	Value   string `json:"Value"`
	Type    int    `json:"Type,omitempty"`
	URL     string `json:"Url,omitempty"`
}
type ImageTextJumpList struct {
	Title string `json:"Title"`
	Type  int    `json:"Type"`
	URL   string `json:"Url"`
}
type ImageTextCardAction struct {
	Title string `json:"Title"`
	Type  int    `json:"Type"`
	URL   string `json:"Url"`
}
type ImageTextCardImage struct {
	URL         string  `json:"Url"`
	AspectRatio float64 `json:"AspectRatio"`
}
type ImageTextVerticalContentList struct {
	Title string `json:"Title"`
	Desc  string `json:"Desc"`
}
type ImageTextActionList struct {
	Text string `json:"Text"`
	Key  string `json:"Key"`
}
type ImageTextActionMenu struct {
	Desc       string                `json:"Desc"`
	ActionList []ImageTextActionList `json:"ActionList"`
}
type ImageTextQuoteArea struct {
	Type      int    `json:"Type"`
	URL       string `json:"Url"`
	Title     string `json:"Title"`
	QuoteText string `json:"QuoteText"`
}
type ImageTextImageTextArea struct {
	Type     int    `json:"Type"`
	URL      string `json:"Url"`
	Title    string `json:"Title"`
	Desc     string `json:"Desc"`
	ImageURL string `json:"ImageUrl"`
}
type ImageTextTemplateCard struct {
	CardType              string                           `json:"CardType"`
	Source                ImageTextSource                  `json:"Source"`
	MainTitle             ImageTextMainTitle               `json:"MainTitle"`
	HorizontalContentList []ImageTextHorizontalContentList `json:"HorizontalContentList"`
	JumpList              ImageTextJumpList                `json:"JumpList"`
	CardAction            ImageTextCardAction              `json:"CardAction"`
	CardImage             ImageTextCardImage               `json:"CardImage"`
	VerticalContentList   []ImageTextVerticalContentList   `json:"VerticalContentList"`
	ActionMenu            ImageTextActionMenu              `json:"ActionMenu"`
	QuoteArea             ImageTextQuoteArea               `json:"QuoteArea"`
	ImageTextArea         ImageTextImageTextArea           `json:"ImageTextArea"`
}

// TemplateCardButtonResponseMessage 更新点击用户的整张卡片 - 按钮交互型
type TemplateCardButtonResponseMessage struct {
	ToUserName   string             `json:"ToUserName"`
	FromUserName string             `json:"FromUserName"`
	CreateTime   int                `json:"CreateTime"`
	MsgType      string             `json:"MsgType"`
	TemplateCard ButtonTemplateCard `json:"TemplateCard"`
}
type ButtonSource struct {
	IconURL   string `json:"IconUrl"`
	Desc      string `json:"Desc"`
	DescColor int    `json:"DescColor"`
}
type ButtonMainTitle struct {
	Title string `json:"Title"`
	Desc  string `json:"Desc"`
}
type ButtonHorizontalContentList struct {
	KeyName string `json:"KeyName"`
	Value   string `json:"Value"`
	Type    int    `json:"Type,omitempty"`
	URL     string `json:"Url,omitempty"`
}
type ButtonJumpList struct {
	Title string `json:"Title"`
	Type  int    `json:"Type"`
	URL   string `json:"Url"`
}
type ButtonCardAction struct {
	Title string `json:"Title"`
	Type  int    `json:"Type"`
	URL   string `json:"Url"`
}
type ButtonButtonList struct {
	Text  string `json:"Text"`
	Style int    `json:"Style"`
	Key   string `json:"Key"`
}
type ButtonActionList struct {
	Text string `json:"Text"`
	Key  string `json:"Key"`
}
type ButtonActionMenu struct {
	Desc       string             `json:"Desc"`
	ActionList []ButtonActionList `json:"ActionList"`
}
type ButtonQuoteArea struct {
	Type      int    `json:"Type"`
	URL       string `json:"Url"`
	Title     string `json:"Title"`
	QuoteText string `json:"QuoteText"`
}
type ButtonOptionList struct {
	ID   string `json:"Id"`
	Text string `json:"Text"`
}
type ButtonButtonSelection struct {
	QuestionKey string             `json:"QuestionKey"`
	Title       string             `json:"Title"`
	SelectedID  string             `json:"SelectedId"`
	Disable     bool               `json:"Disable"`
	OptionList  []ButtonOptionList `json:"OptionList"`
}
type ButtonTemplateCard struct {
	CardType              string                        `json:"CardType"`
	Source                ButtonSource                  `json:"Source"`
	MainTitle             ButtonMainTitle               `json:"MainTitle"`
	SubTitleText          string                        `json:"SubTitleText"`
	HorizontalContentList []ButtonHorizontalContentList `json:"HorizontalContentList"`
	JumpList              ButtonJumpList                `json:"JumpList"`
	CardAction            ButtonCardAction              `json:"CardAction"`
	ButtonList            []ButtonButtonList            `json:"ButtonList"`
	ReplaceText           string                        `json:"ReplaceText"`
	ActionMenu            ButtonActionMenu              `json:"ActionMenu"`
	QuoteArea             ButtonQuoteArea               `json:"QuoteArea"`
	ButtonSelection       ButtonButtonSelection         `json:"ButtonSelection"`
}

// TemplateCardVoteResponseMessage 更新点击用户的整张卡片 - 投票选择型
type TemplateCardVoteResponseMessage struct {
	ToUserName   string           `json:"ToUserName"`
	FromUserName string           `json:"FromUserName"`
	CreateTime   int              `json:"CreateTime"`
	MsgType      string           `json:"MsgType"`
	TemplateCard VoteTemplateCard `json:"TemplateCard"`
}
type VoteSource struct {
	IconURL string `json:"IconUrl"`
	Desc    string `json:"Desc"`
}
type VoteMainTitle struct {
	Title string `json:"Title"`
	Desc  string `json:"Desc"`
}
type VoteOptionList struct {
	ID        string `json:"Id"`
	Text      string `json:"Text"`
	IsChecked bool   `json:"IsChecked"`
}
type VoteCheckBox struct {
	QuestionKey string           `json:"QuestionKey"`
	OptionList  []VoteOptionList `json:"OptionList"`
	Disable     bool             `json:"Disable"`
	Mode        int              `json:"Mode"`
}
type VoteSubmitButton struct {
	Text string `json:"Text"`
	Key  string `json:"Key"`
}
type VoteTemplateCard struct {
	CardType     string           `json:"CardType"`
	Source       VoteSource       `json:"Source"`
	MainTitle    VoteMainTitle    `json:"MainTitle"`
	CheckBox     VoteCheckBox     `json:"CheckBox"`
	SubmitButton VoteSubmitButton `json:"SubmitButton"`
	ReplaceText  string           `json:"ReplaceText"`
}

// TemplateCardSelectResponseMessage 更新点击用户的整张卡片 - 多项选择型
type TemplateCardSelectResponseMessage struct {
	ToUserName   string             `json:"ToUserName"`
	FromUserName string             `json:"FromUserName"`
	CreateTime   int                `json:"CreateTime"`
	MsgType      string             `json:"MsgType"`
	TemplateCard SelectTemplateCard `json:"TemplateCard"`
}
type SelectSource struct {
	IconURL string `json:"IconUrl"`
	Desc    string `json:"Desc"`
}
type SelectMainTitle struct {
	Title string `json:"Title"`
	Desc  string `json:"Desc"`
}
type SelectOptionList struct {
	ID   string `json:"Id"`
	Text string `json:"Text"`
}
type SelectSelectList struct {
	QuestionKey string             `json:"QuestionKey"`
	Title       string             `json:"Title"`
	SelectedID  string             `json:"SelectedId"`
	Disable     bool               `json:"Disable"`
	OptionList  []SelectOptionList `json:"OptionList"`
}
type SelectSubmitButton struct {
	Text string `json:"Text"`
	Key  string `json:"Key"`
}
type SelectTemplateCard struct {
	CardType     string             `json:"CardType"`
	Source       SelectSource       `json:"Source"`
	MainTitle    SelectMainTitle    `json:"MainTitle"`
	SelectList   SelectSelectList   `json:"SelectList"`
	SubmitButton SelectSubmitButton `json:"SubmitButton"`
	ReplaceText  string             `json:"ReplaceText"`
}
