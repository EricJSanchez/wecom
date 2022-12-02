package cache

import "time"

// Cache interface
type Cache interface {
	Get(key string) interface{}
	Set(key string, val interface{}, timeout time.Duration) error
	IsExist(key string) bool
	Delete(key string) error
	ZAdd(param []interface{}) (interface{}, error)
	Expire(key string, exp int) error
}
