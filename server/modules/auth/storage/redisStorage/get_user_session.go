package redisStorage

import (
	"context"
	"fmt"
)

func (s *redisStorage) GetUserSession(ctx context.Context, email string) (map[string]string, error) {
	key := fmt.Sprintf("user_sessions:%s", email)
	return s.client.HGetAll(ctx, key).Result()
}
