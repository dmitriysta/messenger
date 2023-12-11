//go:generate mockery

package interfaces

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisClient interface {
	Del(ctx context.Context, keys ...string) *redis.IntCmd
	Get(ctx context.Context, key string) *redis.StringCmd
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
}
