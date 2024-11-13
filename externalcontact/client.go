package externalcontact

import (
	"github.com/EricJSanchez/wecom/cache"
	"github.com/EricJSanchez/wecom/config"
	"github.com/EricJSanchez/wecom/context"
	"github.com/EricJSanchez/wecom/credential"
)

// Client 外部联系人实例
type Client struct {
	corpID         string // 企业ID
	agentId        string // 应用
	secret         string // Secret是用于校验开发者身份的访问密钥
	token          string // 用于生成签名校验回调请求的合法性
	encodingAESKey string // 回调消息加解密参数是AES密钥的Base64编码，用于解密回调消息内容对应的密文
	cache          cache.Cache
	ctx            *context.Context
}

// NewClient 初始化联系人实例
func NewClient(cfg *config.Config) (client *Client, err error) {
	if cfg.Cache == nil {
		return nil, NewSDKErr(50001)
	}
	var cacheKey string
	var secret string
	if cfg.AgentID != "" && cfg.AgentID != "0" {
		cacheKey = credential.CacheKeyWorkPrefix + "externalcontact:" + cfg.CorpID + ":" + cfg.AgentID
		secret = cfg.AgentSecret
	} else {
		cacheKey = credential.CacheKeyWorkPrefix + "externalcontact:" + cfg.CorpID
		secret = cfg.CustomerSecret
	}
	//初始化 AccessToken Handle
	defaultAkHandle := credential.NewWorkAccessToken(
		cfg.CorpID,
		secret,
		cacheKey,
		cfg.Cache)
	ctx := &context.Context{
		Config:            cfg,
		AccessTokenHandle: defaultAkHandle,
	}

	client = &Client{
		corpID:         cfg.CorpID,
		agentId:        cfg.AgentID,
		secret:         secret,
		token:          cfg.Token,
		encodingAESKey: cfg.EncodingAESKey,
		cache:          cfg.Cache,
		ctx:            ctx,
	}

	return client, nil
}
