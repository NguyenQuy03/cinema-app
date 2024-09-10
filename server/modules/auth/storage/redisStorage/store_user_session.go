package redisStorage

import (
	"context"
	"fmt"
	"time"
)

func (s *redisStorage) StoreUserSession(ctx context.Context, key string, infors map[string]interface{}, expiration time.Duration) error {
	rKey := fmt.Sprintf("user_sessions:%s", key)

	pipe := s.client.Pipeline()

	// Update all non-empty fields in a single HMSet operation
	fieldsToSet := make(map[string]interface{})
	for field, value := range infors {
		if value != "" {
			fieldsToSet[field] = value
		}
	}

	if len(fieldsToSet) > 0 {
		pipe.HMSet(ctx, rKey, fieldsToSet)
	}

	if expiration > 0 {
		pipe.Expire(ctx, rKey, expiration)
	}

	_, err := pipe.Exec(ctx)
	return err
}
