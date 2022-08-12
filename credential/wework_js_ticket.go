package credential

import (
	"encoding/json"
	"fmt"
	"wecom/cache"
	"wecom/util"
	"sync"
	"time"
)

//微信获取ticket的url
const getWeworkTicketURL = "https://qyapi.weixin.qq.com/cgi-bin/ticket/get?access_token=%s&type=agent_config"

//WeworkJsTicket 默认获取js ticket方法
type WeworkJsTicket struct {
	appID          string
	cacheKeyPrefix string
	cache          cache.Cache
	//jsAPITicket 读写锁 同一个AppID一个
	jsAPITicketLock *sync.Mutex
}

//NewWeworkJsTicket new
func NewWeworkJsTicket(appID string, cacheKeyPrefix string, cache cache.Cache) JsTicketHandle {
	return &WeworkJsTicket{
		appID:           appID,
		cache:           cache,
		cacheKeyPrefix:  cacheKeyPrefix,
		jsAPITicketLock: new(sync.Mutex),
	}
}

// ResWeworkTicket 请求jsapi_tikcet返回结果
type ResWeworkTicket struct {
	util.CommonError
	Ticket    string `json:"ticket"`
	ExpiresIn int64  `json:"expires_in"`
}

//GetTicket 获取jsapi_ticket
func (js *WeworkJsTicket) GetTicket(accessToken string) (ticketStr string, err error) {
	//先从cache中取
	jsAPITicketCacheKey := fmt.Sprintf("%s_jsapi_ticket_%s", js.cacheKeyPrefix, js.appID)
	if val := js.cache.Get(jsAPITicketCacheKey); val != nil {
		return val.(string), nil
	}

	js.jsAPITicketLock.Lock()
	defer js.jsAPITicketLock.Unlock()

	// 双检，防止重复从微信服务器获取
	if val := js.cache.Get(jsAPITicketCacheKey); val != nil {
		return val.(string), nil
	}
	var ticket ResWeworkTicket
	ticket, err = GetWeworkTicketFromServer(accessToken)
	if err != nil {
		fmt.Println(err)
		return
	}
	expires := ticket.ExpiresIn - 1500
	err = js.cache.Set(jsAPITicketCacheKey, ticket.Ticket, time.Duration(expires)*time.Second)
	ticketStr = ticket.Ticket
	fmt.Println("er", err)
	fmt.Println("ticketStr", ticketStr)
	return
}

//GetWeworkTicketFromServer 从服务器中获取ticket
func GetWeworkTicketFromServer(accessToken string) (ticket ResWeworkTicket, err error) {
	var response []byte
	url := fmt.Sprintf(getWeworkTicketURL, accessToken)
	response, err = util.HTTPGet(url)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &ticket)
	if err != nil {
		return
	}
	if ticket.ErrCode != 0 {
		err = fmt.Errorf("getTicket Error : errcode=%d , errmsg=%s", ticket.ErrCode, ticket.ErrMsg)
		return
	}
	return
}

type JsTicket struct {
	Ticket string `json:"ticket"`
	err    error
}
