package wecom

import (
	"wecom/application"
	"wecom/config"
	"wecom/contact"
	"wecom/context"
	"wecom/conversation"
	"wecom/credential"
	"wecom/externalcontact"
	"wecom/material"
	"wecom/message"
	"wecom/oauth"
)

type Wework struct {
	ctx *context.Context
}

func NewWework(cfg *config.Config) *Wework {
	defaultAkHandle := credential.NewWorkAccessToken(cfg.CorpID, cfg.CorpSecret, credential.CacheKeyWorkPrefix, cfg.Cache)
	ctx := &context.Context{
		Config:            cfg,
		AccessTokenHandle: defaultAkHandle,
	}
	return &Wework{ctx: ctx}
}

// GetContext get Context
func (w *Wework) GetContext() *context.Context {
	return w.ctx
}

// GetContact get contact
func (w *Wework) GetContact() (*contact.Client, error) {
	return contact.NewClient(w.ctx.Config)
}

// GetApplication get application
func (w *Wework) GetApplication() (*application.Client, error) {
	return application.NewClient(w.ctx.Config)
}

// GetExternalContact get external contact
func (w *Wework) GetExternalContact() (*externalcontact.Client, error) {
	return externalcontact.NewClient(w.ctx.Config)
}

// GetConversation get conversation
func (w *Wework) GetConversation() (*conversation.Client, error) {
	return conversation.NewClient(w.ctx.Config)
}

// GetOAuth get oauth
func (w *Wework) GetOAuth() (*oauth.Client, error) {
	return oauth.NewClient(w.ctx.Config)
}

// GetMassage get message
func (w *Wework) GetMassage() (*message.Client, error) {
	return message.NewClient(w.ctx.Config)
}

// GetMaterial get material
func (w *Wework) GetMaterial() (*material.Client, error) {
	return material.NewClient(w.ctx.Config)
}

//
//func (w *Wework) GetTicket() (*credential.WeworkTicket, error) {
//	return credential.NewTicket(w.ctx.Config)
//}

//func (w *Wework) GetCredential() (ticket string, err error) {
//	accessToken, err := w.ctx.GetAccessToken()
//	return credential.NewWeworkJsTicket().GetTicket(accessToken)
//}
