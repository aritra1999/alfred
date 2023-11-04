package cache

import (
	"context"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	rdb *redis.Client
	ctx = context.Background()
)

func RedisClient() {
	RedisHost := os.Getenv("REDIS_HOST")
	RedisPort := os.Getenv("REDIS_PORT")
	RedisPassword := os.Getenv("REDIS_PASSWORD")

	rdb = redis.NewClient(&redis.Options{
		Addr:     RedisHost + ":" + RedisPort,
		Password: RedisPassword,
		DB:       0,
	})

	if _, err := rdb.Ping(ctx).Result(); err != nil {
		panic(err)
	}
}

func Get(key string) (string, error) {
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func Set(key string, value interface{}, ttl time.Duration) error {
	if err := rdb.Set(ctx, key, value, ttl).Err(); err != nil {
		return err
	}
	return nil
}
