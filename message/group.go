package message

import (
	"encoding/json"
	"fmt"
	"wecom/util"
)

const (
	// 创建群聊会话
	createGroupAddr = "https://qyapi.weixin.qq.com/cgi-bin/appchat/create?access_token=%s"
	// 修改群聊会话
	updateGroupAddr = "https://qyapi.weixin.qq.com/cgi-bin/appchat/update?access_token=%s"
	// 获取群聊会话
	getGroupAddr = "https://qyapi.weixin.qq.com/cgi-bin/appchat/get?access_token=%s&chatid=%s"
	// 发送群组应用消息
	sendGroupMessageAddr = "https://qyapi.weixin.qq.com/cgi-bin/appchat/send?access_token=%s"
)

// CreateGroupOptions 请求参数
type CreateGroupOptions struct {
	Name     string   `json:"name"`
	Owner    string   `json:"owner"`
	Userlist []string `json:"userlist"`
	Chatid   string   `json:"chatid"`
}

// CreateGroupSchema 创建群聊会话响应内容
type CreateGroupSchema struct {
	util.CommonError
	Chatid string `json:"chatid"`
}

// CreateGroup 创建群聊会话
func (r *Client) CreateGroup(options CreateGroupOptions) (info CreateGroupSchema, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(createGroupAddr, accessToken), options)
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

// UpdateGroupOptions 修改群聊会话请求参数
type UpdateGroupOptions struct {
	Chatid      string   `json:"chatid"`
	Name        string   `json:"name"`
	Owner       string   `json:"owner"`
	AddUserList []string `json:"add_user_list"`
	DelUserList []string `json:"del_user_list"`
}

// UpdateGroup 修改群聊会话
func (r *Client) UpdateGroup(options UpdateGroupOptions) (info util.CommonError, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(updateGroupAddr, accessToken), options)
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

// GetGroupOptions 请求参数
type GetGroupOptions struct {
	Chatid string `json:"chatid"`
}

// GetGroupSchema 获取群聊会话响应内容
type GetGroupSchema struct {
	util.CommonError
	ChatInfo ChatInfo `json:"chat_info"`
}
type ChatInfo struct {
	Chatid   string   `json:"chatid"`
	Name     string   `json:"name"`
	Owner    string   `json:"owner"`
	Userlist []string `json:"userlist"`
}

// GetGroup 获取群聊会话
func (r *Client) GetGroup(options GetGroupOptions) (info GetGroupSchema, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.HTTPGet(fmt.Sprintf(getGroupAddr, accessToken, options.Chatid))
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

type BaseSendGroupMessage struct {
	Chatid  string `json:"chatid"`
	Msgtype string `json:"msgtype"`
	Safe    int    `json:"safe"`
}

// SendGroupTextMessageOptions 文本消息
type SendGroupTextMessageOptions struct {
	BaseSendGroupMessage
	Text SendGroupTextMessage `json:"text"`
}
type SendGroupTextMessage struct {
	Content string `json:"content"`
}

// SendGroupImageMessageOptions 图片消息
type SendGroupImageMessageOptions struct {
	BaseSendGroupMessage
	Image SendGroupImageMessage `json:"image"`
}
type SendGroupImageMessage struct {
	MediaID string `json:"media_id"`
}

// SendGroupVoiceMessageOptions 语音消息
type SendGroupVoiceMessageOptions struct {
	Chatid  string                `json:"chatid"`
	Msgtype string                `json:"msgtype"`
	Voice   SendGroupVoiceMessage `json:"voice"`
}
type SendGroupVoiceMessage struct {
	MediaID string `json:"media_id"`
}

// SendGroupVideoMessageOptions 视频消息
type SendGroupVideoMessageOptions struct {
	BaseSendGroupMessage
	Video Video `json:"video"`
}
type SendGroupVideoMessage struct {
	MediaID     string `json:"media_id"`
	Description string `json:"description"`
	Title       string `json:"title"`
}

// SendGroupFileMessageOptions 文件消息
type SendGroupFileMessageOptions struct {
	Chatid  string `json:"chatid"`
	Msgtype string `json:"msgtype"`
	File    File   `json:"file"`
	Safe    int    `json:"safe"`
}
type SendGroupFileMessage struct {
	MediaID string `json:"media_id"`
}

// SendGroupTextcardMessageOptions 文本卡片消息
type SendGroupTextcardMessageOptions struct {
	BaseSendGroupMessage
	Textcard Textcard `json:"textcard"`
}
type SendGroupTextcardMessage struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Btntxt      string `json:"btntxt"`
}

// SendGroupNewsMessageOptions 图文消息
type SendGroupNewsMessageOptions struct {
	Chatid  string               `json:"chatid"`
	Msgtype string               `json:"msgtype"`
	News    SendGroupNewsMessage `json:"news"`
	Safe    int                  `json:"safe"`
}
type SendGroupNewsArticlesMessage struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Picurl      string `json:"picurl"`
}
type SendGroupNewsMessage struct {
	Articles []SendGroupNewsArticlesMessage `json:"articles"`
}

// SendGroupMpnewsMessageOptions 图文消息（mpnews）
type SendGroupMpnewsMessageOptions struct {
	Chatid  string                 `json:"chatid"`
	Msgtype string                 `json:"msgtype"`
	Mpnews  SendGroupMpnewsMessage `json:"mpnews"`
	Safe    int                    `json:"safe"`
}
type SendGroupMpnewsArticlesMessage struct {
	Title            string `json:"title"`
	ThumbMediaID     string `json:"thumb_media_id"`
	Author           string `json:"author"`
	ContentSourceURL string `json:"content_source_url"`
	Content          string `json:"content"`
	Digest           string `json:"digest"`
}
type SendGroupMpnewsMessage struct {
	Articles []SendGroupMpnewsArticlesMessage `json:"articles"`
}

// SendGroupMarkdownMessageOptions markdown消息
type SendGroupMarkdownMessageOptions struct {
	Chatid   string                   `json:"chatid"`
	Msgtype  string                   `json:"msgtype"`
	Markdown SendGroupMarkdownMessage `json:"markdown"`
}
type SendGroupMarkdownMessage struct {
	Content string `json:"content"`
}

// SendGroupMessage 发送群组应用消息
func (r *Client) SendGroupMessage(options interface{}) (info util.CommonError, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(createGroupAddr, accessToken), options)
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
