package message

import (
	"wecom/cache"
	"wecom/config"
	"wecom/context"
	"wecom/credential"
)

// Client 通讯录实例
type Client struct {
	corpID         string // 企业ID
	secret         string // Secret是用于校验开发者身份的访问密钥
	token          string // 用于生成签名校验回调请求的合法性
	encodingAESKey string // 回调消息加解密参数是AES密钥的Base64编码，用于解密回调消息内容对应的密文
	cache          cache.Cache
	ctx            *context.Context
}

// NewClient 初始化通讯录实例
func NewClient(cfg *config.Config) (client *Client, err error) {
	if cfg.Cache == nil {
		return nil, NewSDKErr(50001)
	}

	//初始化 AccessToken Handle
	defaultAkHandle := credential.NewWorkAccessToken(
		cfg.CorpID,
		cfg.AgentSecret,
		credential.CacheKeyWorkPrefix+"application_"+cfg.AgentID+"_",
		cfg.Cache)
	ctx := &context.Context{
		Config:            cfg,
		AccessTokenHandle: defaultAkHandle,
	}

	client = &Client{
		corpID:         cfg.CorpID,
		secret:         cfg.AgentSecret,
		token:          cfg.Token,
		encodingAESKey: cfg.EncodingAESKey,
		cache:          cfg.Cache,
		ctx:            ctx,
	}

	return client, nil
}
