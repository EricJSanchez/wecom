package util

import (
	"errors"
	"fmt"
	"github.com/EricJSanchez/wecom/cache"
	"math"
	url2 "net/url"
	"time"
)

// Record 记录接口调用频率
func Record(cache cache.Cache, url string) (err error) {
	// 提取路径
	urlInfo, err := url2.Parse(url)
	if err != nil {
		return
	}
	if urlInfo.Path == "" {
		err = errors.New("url路径有误")
		return
	}
	// 生成 key
	todayKey := fmt.Sprintf("wx:wecom:api:count:%s", time.Now().Format("20060102"))
	// zadd wx:wecom:api:count:%s incr 1 /path/api/contact##1440
	exist := cache.IsExist(todayKey)
	_, err = cache.ZAdd([]interface{}{todayKey, "incr", 1, urlInfo.Path + "##" + getTodayMin()})
	if err != nil {
		return
	}
	if !exist {
		cache.Expire(todayKey, getTodayLeftSec()+60)
	}

	return
}

// 获取现在是今天的第几分钟
func getTodayMin() string {
	t2 := time.Now()
	zeroTime := time.Date(t2.Year(), t2.Month(), t2.Day(), 0, 0, 0, 0, t2.Location())
	min := math.Ceil(time.Now().Sub(zeroTime).Minutes())
	minString := fmt.Sprintf("%.0f", min)
	return minString
}

// 获取今天还剩余多少秒
func getTodayLeftSec() int {
	t2 := time.Now()
	nextDayTime := time.Date(t2.Year(), t2.Month(), t2.Day(), 0, 0, 0, 0, t2.Location()).AddDate(0, 0, 1)
	sec := math.Ceil(nextDayTime.Sub(time.Now()).Seconds())
	return int(sec)
}
