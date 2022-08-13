package wecom

import (
	"github.com/EricJSanchez/wecom/application"
	"github.com/EricJSanchez/wecom/config"
	"github.com/EricJSanchez/wecom/contact"
	"github.com/EricJSanchez/wecom/context"
	"github.com/EricJSanchez/wecom/conversation"
	"github.com/EricJSanchez/wecom/credential"
	"github.com/EricJSanchez/wecom/externalcontact"
	"github.com/EricJSanchez/wecom/material"
	"github.com/EricJSanchez/wecom/message"
	"github.com/EricJSanchez/wecom/oauth"
)

type Wecom struct {
	ctx *context.Context
}

func NewWecom(cfg *config.Config) *Wecom {
	defaultAkHandle := credential.NewWorkAccessToken(cfg.CorpID, cfg.CorpSecret, credential.CacheKeyWorkPrefix, cfg.Cache)
	ctx := &context.Context{
		Config:            cfg,
		AccessTokenHandle: defaultAkHandle,
	}
	return &Wecom{ctx: ctx}
}

// GetContext get Context
func (w *Wecom) GetContext() *context.Context {
	return w.ctx
}

// GetContact get contact
func (w *Wecom) GetContact() (*contact.Client, error) {
	return contact.NewClient(w.ctx.Config)
}

// GetApplication get application
func (w *Wecom) GetApplication() (*application.Client, error) {
	return application.NewClient(w.ctx.Config)
}

// GetExternalContact get external contact
func (w *Wecom) GetExternalContact() (*externalcontact.Client, error) {
	return externalcontact.NewClient(w.ctx.Config)
}

// GetConversation get conversation
func (w *Wecom) GetConversation() (*conversation.Client, error) {
	return conversation.NewClient(w.ctx.Config)
}

// GetOAuth get oauth
func (w *Wecom) GetOAuth() (*oauth.Client, error) {
	return oauth.NewClient(w.ctx.Config)
}

// GetMassage get message
func (w *Wecom) GetMassage() (*message.Client, error) {
	return message.NewClient(w.ctx.Config)
}

// GetMaterial get material
func (w *Wecom) GetMaterial() (*material.Client, error) {
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
