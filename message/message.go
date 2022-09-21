package message

import (
	"encoding/json"
	"fmt"
	"github.com/EricJSanchez/wecom/util"
)

const (
	// 发送消息
	sendMessageAddr = "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%s"
	// 更新模版卡片消息
	updateTemplateCardAddr = "https://qyapi.weixin.qq.com/cgi-bin/message/update_template_card?access_token=%s"
	// 撤回应用消息
	recallMessageAddr = "https://qyapi.weixin.qq.com/cgi-bin/message/recall?access_token=%s"
)

type BaseMessageOptions struct {
	Touser                 string `json:"touser"`
	Toparty                string `json:"toparty"`
	Totag                  string `json:"totag"`
	Msgtype                string `json:"msgtype"`
	Agentid                int    `json:"agentid"`
	Safe                   int    `json:"safe"`
	EnableIDTrans          int    `json:"enable_id_trans"`
	EnableDuplicateCheck   int    `json:"enable_duplicate_check"`
	DuplicateCheckInterval int    `json:"duplicate_check_interval"`
}

type BaseMessageSchema struct {
	util.CommonError
	Invaliduser  string `json:"invaliduser"`
	Invalidparty string `json:"invalidparty"`
	Invalidtag   string `json:"invalidtag"`
	Msgid        string `json:"msgid"`
	ResponseCode string `json:"response_code"`
}

// 文本消息
type TextMessageOptions struct {
	BaseMessageOptions
	Text Text `json:"text"`
}

type Text struct {
	Content string `json:"content"`
}

// 图片消息
type ImageMessageOptions struct {
	BaseMessageOptions
	Image Image `json:"image"`
}
type Image struct {
	MediaID string `json:"media_id"`
}

// 语音消息
type VoiceMessageOptions struct {
	BaseMessageOptions
	Voice Voice `json:"voice"`
}
type Voice struct {
	MediaID string `json:"media_id"`
	Name    string `json:"name"`
	Size    int    `json:"size"`
}

// 视频消息
type VideoMessageOptions struct {
	BaseMessageOptions
	Video Video `json:"video"`
}
type Video struct {
	MediaID     string `json:"media_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// 文件消息
type FileMessageOptions struct {
	BaseMessageOptions
	File File `json:"file"`
}
type File struct {
	MediaID string `json:"media_id"`
	Name    string `json:"name"`
	Size    int    `json:"size"`
}

// 文本卡片消息
type TextcardMessageOptions struct {
	BaseMessageOptions
	Textcard Textcard `json:"textcard"`
}
type Textcard struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Btntxt      string `json:"btntxt"`
}

// 图文消息
type NewsMessageOptions struct {
	BaseMessageOptions
	News News `json:"news"`
}
type Articles struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Picurl      string `json:"picurl"`
	Appid       string `json:"appid"`
	Pagepath    string `json:"pagepath"`
}
type News struct {
	Articles []Articles `json:"articles"`
}

// 图文消息（mpnews）
type MpnewsMessageOptions struct {
	BaseMessageOptions
	Mpnews Mpnews `json:"mpnews"`
}
type MpnewsArticles struct {
	Title            string `json:"title"`
	ThumbMediaID     string `json:"thumb_media_id"`
	Author           string `json:"author"`
	ContentSourceURL string `json:"content_source_url"`
	Content          string `json:"content"`
	Digest           string `json:"digest"`
}
type Mpnews struct {
	Articles []MpnewsArticles `json:"articles"`
}

// markdown消息
type MarkdownMessageOptions struct {
	BaseMessageOptions
	Markdown Markdown `json:"markdown"`
}
type Markdown struct {
	Content string `json:"content"`
}

// 小程序通知消息
type MiniprogramNoticeMessageOptions struct {
	BaseMessageOptions
	MiniprogramNotice MiniprogramNotice `json:"miniprogram_notice"`
}
type ContentItem struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
type MiniprogramNotice struct {
	Appid             string        `json:"appid"`
	Page              string        `json:"page"`
	Title             string        `json:"title"`
	Description       string        `json:"description"`
	EmphasisFirstItem bool          `json:"emphasis_first_item"`
	ContentItem       []ContentItem `json:"content_item"`
}

// 模板卡片消息
type TemplateCardMessageOptions struct {
	BaseMessageOptions
	TemplateCard TemplateCard `json:"template_card"`
}
type Source struct {
	IconURL   string `json:"icon_url"`
	Desc      string `json:"desc"`
	DescColor int    `json:"desc_color"`
}
type ActionList struct {
	Text string `json:"text"`
	Key  string `json:"key"`
}
type ActionMenu struct {
	Desc       string       `json:"desc"`
	ActionList []ActionList `json:"action_list"`
}
type MainTitle struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}
type QuoteArea struct {
	Type      int    `json:"type"`
	URL       string `json:"url"`
	Title     string `json:"title"`
	QuoteText string `json:"quote_text"`
}
type EmphasisContent struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}
type HorizontalContentList struct {
	Keyname string `json:"keyname"`
	Value   string `json:"value"`
	Type    int    `json:"type,omitempty"`
	URL     string `json:"url,omitempty"`
	MediaID string `json:"media_id,omitempty"`
	Userid  string `json:"userid,omitempty"`
}
type JumpList struct {
	Type     int    `json:"type"`
	Title    string `json:"title"`
	URL      string `json:"url,omitempty"`
	Appid    string `json:"appid,omitempty"`
	Pagepath string `json:"pagepath,omitempty"`
}
type CardAction struct {
	Type     int    `json:"type"`
	URL      string `json:"url"`
	Appid    string `json:"appid"`
	Pagepath string `json:"pagepath"`
}
type TemplateCard struct {
	CardType              string                  `json:"card_type"`
	Source                Source                  `json:"source"`
	ActionMenu            ActionMenu              `json:"action_menu"`
	TaskID                string                  `json:"task_id"`
	MainTitle             MainTitle               `json:"main_title"`
	QuoteArea             QuoteArea               `json:"quote_area"`
	EmphasisContent       EmphasisContent         `json:"emphasis_content"`
	SubTitleText          string                  `json:"sub_title_text"`
	HorizontalContentList []HorizontalContentList `json:"horizontal_content_list"`
	JumpList              []JumpList              `json:"jump_list"`
	CardAction            CardAction              `json:"card_action"`
}

// 图文展示型
type TemplateCardImageMessageOptions struct {
	BaseMessageOptions
	TemplateCard TemplateCardImage `json:"template_card"`
}
type ImageTextArea struct {
	Type     int    `json:"type"`
	URL      string `json:"url"`
	Title    string `json:"title"`
	Desc     string `json:"desc"`
	ImageURL string `json:"image_url"`
}
type CardImage struct {
	URL         string  `json:"url"`
	AspectRatio float64 `json:"aspect_ratio"`
}
type VerticalContentList struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}
type TemplateCardImage struct {
	CardType              string                  `json:"card_type"`
	Source                Source                  `json:"source"`
	ActionMenu            ActionMenu              `json:"action_menu"`
	TaskID                string                  `json:"task_id"`
	MainTitle             MainTitle               `json:"main_title"`
	QuoteArea             QuoteArea               `json:"quote_area"`
	ImageTextArea         ImageTextArea           `json:"image_text_area"`
	CardImage             CardImage               `json:"card_image"`
	VerticalContentList   []VerticalContentList   `json:"vertical_content_list"`
	HorizontalContentList []HorizontalContentList `json:"horizontal_content_list"`
	JumpList              []JumpList              `json:"jump_list"`
	CardAction            CardAction              `json:"card_action"`
}

// 按钮交互型
type TemplateCardButtonMessageOptions struct {
	BaseMessageOptions
	TemplateCard TemplateCardButton `json:"template_card"`
}
type OptionList struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}
type ButtonSelection struct {
	QuestionKey string       `json:"question_key"`
	Title       string       `json:"title"`
	OptionList  []OptionList `json:"option_list"`
	SelectedID  string       `json:"selected_id"`
}
type ButtonList struct {
	Text  string `json:"text"`
	Style int    `json:"style"`
	Key   string `json:"key"`
}
type TemplateCardButton struct {
	CardType              string                  `json:"card_type"`
	Source                Source                  `json:"source"`
	ActionMenu            ActionMenu              `json:"action_menu"`
	MainTitle             MainTitle               `json:"main_title"`
	QuoteArea             QuoteArea               `json:"quote_area"`
	SubTitleText          string                  `json:"sub_title_text"`
	HorizontalContentList []HorizontalContentList `json:"horizontal_content_list"`
	CardAction            CardAction              `json:"card_action"`
	TaskID                string                  `json:"task_id"`
	ButtonSelection       ButtonSelection         `json:"button_selection"`
	ButtonList            []ButtonList            `json:"button_list"`
}

// 投票选择型
type TemplateCardVoteMessageOptions struct {
	BaseMessageOptions
	TemplateCard TemplateCardVote `json:"template_card"`
}
type SourceVote struct {
	IconURL string `json:"icon_url"`
	Desc    string `json:"desc"`
}
type OptionListVote struct {
	ID        string `json:"id"`
	Text      string `json:"text"`
	IsChecked bool   `json:"is_checked"`
}
type Checkbox struct {
	QuestionKey string           `json:"question_key"`
	OptionList  []OptionListVote `json:"option_list"`
	Mode        int              `json:"mode"`
}
type SubmitButton struct {
	Text string `json:"text"`
	Key  string `json:"key"`
}
type TemplateCardVote struct {
	CardType     string       `json:"card_type"`
	Source       SourceVote   `json:"source"`
	MainTitle    MainTitle    `json:"main_title"`
	TaskID       string       `json:"task_id"`
	Checkbox     Checkbox     `json:"checkbox"`
	SubmitButton SubmitButton `json:"submit_button"`
}

// 多项选择型
type TemplateCardFunctionMessageOptions struct {
	BaseMessageOptions
	TemplateCard TemplateCardFunction `json:"template_card"`
}
type SourceFunction struct {
	IconURL string `json:"icon_url"`
	Desc    string `json:"desc"`
}
type SelectList struct {
	QuestionKey string       `json:"question_key"`
	Title       string       `json:"title"`
	SelectedID  string       `json:"selected_id"`
	OptionList  []OptionList `json:"option_list"`
}
type TemplateCardFunction struct {
	CardType     string         `json:"card_type"`
	Source       SourceFunction `json:"source"`
	MainTitle    MainTitle      `json:"main_title"`
	TaskID       string         `json:"task_id"`
	SelectList   []SelectList   `json:"select_list"`
	SubmitButton SubmitButton   `json:"submit_button"`
}

// SendMessage 发送应用消息
func (r *Client) SendMessage(options interface{}) (info BaseMessageSchema, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(sendMessageAddr, accessToken), options)
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

// 更新模版卡片消息
type BaseTemplateCardUpdateMessage struct {
	Userids      []string `json:"userids"`
	Partyids     []int    `json:"partyids"`
	Agentid      int      `json:"agentid"`
	ResponseCode string   `json:"response_code"`
}

// 更新按钮为不可点击状态
type TemplateCardUpdateMessageOptions struct {
	BaseTemplateCardUpdateMessage
	Tagids       []int        `json:"tagids"`
	Atall        int          `json:"atall"`
	TemplateCard TemplateCard `json:"template_card"`
}

type Button struct {
	ReplaceName string `json:"replace_name"`
}

// 更新为新的卡片 - 文本通知型
type TemplateCardUpdateTextMessageOptions struct {
	Userids      []string               `json:"userids"`
	Partyids     []int                  `json:"partyids"`
	Agentid      int                    `json:"agentid"`
	ResponseCode string                 `json:"response_code"`
	TemplateCard TemplateCardUpdateText `json:"template_card"`
}

type TemplateCardUpdateText struct {
	CardType              string                  `json:"card_type"`
	Source                Source                  `json:"source"`
	ActionMenu            ActionMenu              `json:"action_menu"`
	MainTitle             MainTitle               `json:"main_title"`
	QuoteArea             QuoteArea               `json:"quote_area"`
	EmphasisContent       EmphasisContent         `json:"emphasis_content"`
	SubTitleText          string                  `json:"sub_title_text"`
	HorizontalContentList []HorizontalContentList `json:"horizontal_content_list"`
	JumpList              []JumpList              `json:"jump_list"`
	CardAction            CardAction              `json:"card_action"`
}

// 更新为新的卡片 - 图文展示型
type TemplateCardUpdateImageMessageOptions struct {
	Userids      []string                `json:"userids"`
	Partyids     []int                   `json:"partyids"`
	Agentid      int                     `json:"agentid"`
	ResponseCode string                  `json:"response_code"`
	TemplateCard TemplateCardUpdateImage `json:"template_card"`
}

type TemplateCardUpdateImage struct {
	CardType              string                  `json:"card_type"`
	Source                Source                  `json:"source"`
	ActionMenu            ActionMenu              `json:"action_menu"`
	MainTitle             MainTitle               `json:"main_title"`
	QuoteArea             QuoteArea               `json:"quote_area"`
	ImageTextArea         ImageTextArea           `json:"image_text_area"`
	CardImage             CardImage               `json:"card_image"`
	VerticalContentList   []VerticalContentList   `json:"vertical_content_list"`
	HorizontalContentList []HorizontalContentList `json:"horizontal_content_list"`
	JumpList              []JumpList              `json:"jump_list"`
	CardAction            CardAction              `json:"card_action"`
}

// 更新为新的卡片 - 按钮交互型
type TemplateCardUpdateButtonMessageOptions struct {
	Userids      []string                 `json:"userids"`
	Partyids     []int                    `json:"partyids"`
	Agentid      int                      `json:"agentid"`
	ResponseCode string                   `json:"response_code"`
	TemplateCard TemplateCardUpdateButton `json:"template_card"`
}

type TemplateCardUpdateButton struct {
	CardType              string                  `json:"card_type"`
	Source                Source                  `json:"source"`
	ActionMenu            ActionMenu              `json:"action_menu"`
	MainTitle             MainTitle               `json:"main_title"`
	QuoteArea             QuoteArea               `json:"quote_area"`
	SubTitleText          string                  `json:"sub_title_text"`
	HorizontalContentList []HorizontalContentList `json:"horizontal_content_list"`
	CardAction            CardAction              `json:"card_action"`
	ButtonSelection       ButtonSelection         `json:"button_selection"`
	ButtonList            []ButtonList            `json:"button_list"`
	ReplaceText           string                  `json:"replace_text"`
}

// 更新为新的卡片 - 投票选择型
type TemplateCardUpdateVoteMessageOptions struct {
	Userids      []string               `json:"userids"`
	Partyids     []int                  `json:"partyids"`
	Agentid      int                    `json:"agentid"`
	ResponseCode string                 `json:"response_code"`
	TemplateCard TemplateCardUpdateVote `json:"template_card"`
}

type CheckboxUpdateVote struct {
	QuestionKey string           `json:"question_key"`
	OptionList  []OptionListVote `json:"option_list"`
	Disable     bool             `json:"disable"`
	Mode        int              `json:"mode"`
}

type TemplateCardUpdateVote struct {
	CardType     string             `json:"card_type"`
	Source       SourceVote         `json:"source"`
	MainTitle    MainTitle          `json:"main_title"`
	Checkbox     CheckboxUpdateVote `json:"checkbox"`
	SubmitButton SubmitButton       `json:"submit_button"`
	ReplaceText  string             `json:"replace_text"`
}

// 更新为新的卡片 - 多项选择型
type TemplateCardUpdateFunctionMessageOptions struct {
	Userids      []string                   `json:"userids"`
	Partyids     []int                      `json:"partyids"`
	Tagids       []int                      `json:"tagids"`
	Atall        int                        `json:"atall"`
	Agentid      int                        `json:"agentid"`
	ResponseCode string                     `json:"response_code"`
	TemplateCard TemplateCardUpdateFunction `json:"template_card"`
}
type SelectListUpdateFunction struct {
	QuestionKey string       `json:"question_key"`
	Title       string       `json:"title"`
	SelectedID  string       `json:"selected_id"`
	Disable     bool         `json:"disable"`
	OptionList  []OptionList `json:"option_list"`
}

type TemplateCardUpdateFunction struct {
	CardType     string                     `json:"card_type"`
	Source       SourceFunction             `json:"source"`
	MainTitle    MainTitle                  `json:"main_title"`
	SelectList   []SelectListUpdateFunction `json:"select_list"`
	SubmitButton SubmitButton               `json:"submit_button"`
	ReplaceText  string                     `json:"replace_text"`
}

type BaseTemplateCardUpdateSchema struct {
	util.CommonError
	Invaliduser []string `json:"invaliduser"`
}

// UpdateTemplateCard 更新模版卡片消息
func (r *Client) UpdateTemplateCard(options interface{}) (info BaseTemplateCardUpdateSchema, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(updateTemplateCardAddr, accessToken), options)
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

type RecallMessageOptions struct {
	Msgid string `json:"msgid"`
}

// RecallMessage 撤回应用消息
func (r *Client) RecallMessage(options interface{}) (info util.CommonError, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(recallMessageAddr, accessToken), options)
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
