package db

import (
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

func InitRedisDB() (*redis.Options, error) {
	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		redisURL = "redis://@localhost:6379/0"
	}

	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse Redis URL: %w", err)
	}

	return opt, nil
}
