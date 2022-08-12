package context

import (
	"wecom/config"
	"wecom/credential"
)

// Context struct
type Context struct {
	*config.Config
	credential.AccessTokenHandle
}
