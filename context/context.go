package context

import (
	"github.com/EricJSanchez/wecom/config"
	"github.com/EricJSanchez/wecom/credential"
)

// Context struct
type Context struct {
	*config.Config
	credential.AccessTokenHandle
}
