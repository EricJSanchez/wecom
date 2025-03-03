package platform

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/EricJSanchez/wecom/util"
	"strings"
	"time"
)

const (
	getDimissionUserInfoList = "https://work.weixin.qq.com/wework_admin/customer/pending/getDimissionUserInfoListV2?lang=zh_CN&f=json&ajax=1&timeZoneInfo[zone_offset]=-8&random=0.%d&limit=%d&last_id=%s&transfer_filter_type=3&time_since=0&time_before=%d&_d2st=a2631871"
)

type DimissionRes struct {
	Data DimissionResData `json:"data"`
}

type DimissionResData struct {
	List   []DimissionResDataList `json:"list"`
	Total  int                    `json:"total"`
	LastId string                 `json:"last_id"`
}

type DimissionResDataList struct {
	Vid              string `json:"vid"`
	Name             string `json:"name"`
	EnglishName      string `json:"english_name"`
	Partylist        string `json:"partylist"`
	PartylistName    string `json:"partylist_name"`
	CreateTime       int    `json:"create_time"`
	UpdateTimeUs     string `json:"update_time_us"`
	Flags            int    `json:"flags"`
	ExtendInfo       string `json:"extend_info"`
	RoomCnt          int    `json:"room_cnt"`
	TransferContact  int    `json:"transfer_contact"`
	TransferRoom     int    `json:"transfer_room"`
	IsWxExpired      bool   `json:"isWxExpired"`
	ExternDefaultCnt int    `json:"extern_default_cnt"`
	Image            string `json:"image"`
	Acctid           string `json:"acctid"`
	Gender           int    `json:"gender"`
	RoomDefaultCnt   int    `json:"room_default_cnt"`
	MainpartyId      string `json:"mainparty_id"`
}

// GetDimissionUserInfoListV2 获取所有待离职继承列表
func (r *Client) GetDimissionUserInfoListV2(last_id string, limit int) (demissionRes DimissionRes, err error) {
	cookie := r.ctx.Config.Cookie
	var header = commonPlatForm
	header["cookie"] = cookie
	if cookie == "" {
		err = errors.New("cookie 缺失")
		return
	}
	uri := fmt.Sprintf(getDimissionUserInfoList, time.Now().Nanosecond(), limit, last_id, time.Now().Unix())
	rspOrigin, err := util.GetWithHeader(uri, header)
	//Pr(uri, string(rspOrigin))
	if strings.Contains(string(rspOrigin), "errCode") {
		err = errors.New("请求出错:" + string(rspOrigin))
		return
	}
	err = json.Unmarshal(rspOrigin, &demissionRes)
	if err != nil {
		return
	}
	return
}
