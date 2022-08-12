package conversation

import (
	"fmt"
	"github.com/spf13/cast"
	"time"
	"wecom/WeWorkFinanceSDK"
	"wecom/credential"
)

// GetMessageInstance 获取会话归档实例
func (r *Client) GetMessageInstance() (client *WeWorkFinanceSDK.Client, err error) {
	client, err = WeWorkFinanceSDK.NewClient(r.corpID, r.secret, r.rasPrivateKey)
	if err != nil {
		fmt.Println("SDK 初始化失败：", err)
	}
	return
}

// SetMessageSeq 设置消息seq
func (r *Client) SetMessageSeq(newSeq uint64) (seq uint64, err error) {
	r.seqLock.Lock()
	defer r.seqLock.Unlock()
	seqCacheKey := fmt.Sprintf("%s:seq:%s", credential.CacheKeyWorkPrefix, r.corpID)

	if val := r.cache.Get(seqCacheKey); val != nil {
		seq = cast.ToUint64(val)
	} else {
		seq = 0
	}

	if newSeq > seq {
		err = r.cache.Set(seqCacheKey, newSeq, time.Duration(720)*time.Hour)
		if err != nil {
			return
		}
		seq = newSeq
	}

	return
}

// GetMessageSeq 获取消息seq
func (r *Client) GetMessageSeq() (seq uint64, err error) {
	r.seqLock.RLock()
	defer r.seqLock.RUnlock()
	seqCacheKey := fmt.Sprintf("%s:seq:%s", credential.CacheKeyWorkPrefix, r.corpID)
	if val := r.cache.Get(seqCacheKey); val != nil {
		seq = cast.ToUint64(val)
	} else {
		seq = 0
	}
	return
}
