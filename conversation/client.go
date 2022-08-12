package conversation

import (
	"sync"
	"wecom/cache"
	"wecom/config"
	"wecom/context"
	"wecom/credential"
)

// Client 会话内容归档实例
type Client struct {
	corpID         string // 企业ID
	secret         string // Secret是用于校验开发者身份的访问密钥
	rasVersion     int    // 企业获取的会话内容将用此公钥加密，企业可用自行保存的私钥解开会话内容数据
	rasPrivateKey  string // 消息加密私钥，可以在企业微信管理端--管理工具--消息加密公钥查看对用公钥，私钥一般由自己保存
	token          string // 用于生成签名校验回调请求的合法性
	encodingAESKey string // 回调消息加解密参数是AES密钥的Base64编码，用于解密回调消息内容对应的密文
	cache          cache.Cache
	ctx            *context.Context
	seqLock        *sync.RWMutex
}

// NewClient 初始化会话内容归档实例实例
func NewClient(cfg *config.Config) (client *Client, err error) {
	if cfg.Cache == nil {
		return nil, NewSDKErr(50001)
	}

	//初始化 AccessToken Handle
	defaultAkHandle := credential.NewWorkAccessToken(
		cfg.CorpID,
		cfg.CorpSecret,
		credential.CacheKeyWorkPrefix+"conversation_",
		cfg.Cache)
	ctx := &context.Context{
		Config:            cfg,
		AccessTokenHandle: defaultAkHandle,
	}

	client = &Client{
		corpID:         cfg.CorpID,
		secret:         cfg.CustomerSecret,
		rasVersion:     cfg.RasVersion,
		rasPrivateKey:  cfg.RasPrivateKey,
		token:          cfg.Token,
		encodingAESKey: cfg.EncodingAESKey,
		cache:          cfg.Cache,
		ctx:            ctx,
		seqLock:        new(sync.RWMutex),
	}

	return client, nil
}
