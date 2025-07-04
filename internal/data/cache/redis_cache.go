package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(client *redis.Client) *RedisCache {
	return &RedisCache{client: client}
}

func (c *RedisCache) Set(key string, value interface{}, expiration time.Duration) error {
	return c.client.Set(context.Background(), key, value, expiration).Err()
}

func (c *RedisCache) Get(key string) (interface{}, error) {
	return c.client.Get(context.Background(), key).Result()
}

func (c *RedisCache) Delete(key string) error {
	return c.client.Del(context.Background(), key).Err()
}
