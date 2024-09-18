package redis_cache

import (
	"context"

	"github.com/Meraiku/grpc_auth/internal/config"
	"github.com/redis/go-redis/v9"
)

type redisCache struct {
	r *redis.Client
}

func New(cfg *config.RedisConfig) (*redisCache, error) {
	opts := &redis.Options{
		Addr:     cfg.Address(),
		Password: cfg.Password,
		DB:       cfg.DBNum,
	}

	client := redis.NewClient(opts)
	status := client.Ping(context.TODO())
	if status.Err() != nil {
		return nil, status.Err()
	}

	return &redisCache{
		r: client,
	}, nil
}

func (r *redisCache) Write(ctx context.Context, key, value string) error {
	return nil
}
