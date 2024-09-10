package redisStorage

import "github.com/redis/go-redis/v9"

type redisStorage struct {
	client *redis.Client
}

func NewRedisStorage(client *redis.Client) *redisStorage {
	return &redisStorage{client: client}
}
