package data

import "time"

type RedisCache interface {
	Set(key string, value interface{}, expiration time.Duration) error
	Get(key string) (interface{}, error)
	Delete(key string) error
}
