package platform

import (
	"github.com/EricJSanchez/wecom/cache"
	"github.com/EricJSanchez/wecom/config"
	"github.com/EricJSanchez/wecom/context"
	"github.com/EricJSanchez/wecom/credential"
)

// Client 管理后台实例
type Client struct {
	corpID string // 企业ID
	cookie string // cookie
	cache  cache.Cache
	ctx    *context.Context
}

// NewClient 初始化Platform
func NewClient(cfg *config.Config) (client *Client, err error) {
	if cfg.Cache == nil {
		return nil, NewSDKErr(50001)
	}

	//初始化 AccessToken Handle
	defaultAkHandle := credential.NewWorkAccessToken(
		cfg.CorpID,
		cfg.AgentSecret,
		credential.CacheKeyWorkPrefix+"application:"+cfg.CorpID+":"+cfg.AgentID,
		cfg.Cache)
	ctx := &context.Context{
		Config:            cfg,
		AccessTokenHandle: defaultAkHandle,
	}

	client = &Client{
		corpID: cfg.CorpID,
		cookie: "",
		cache:  cfg.Cache,
		ctx:    ctx,
	}

	return client, nil
}
