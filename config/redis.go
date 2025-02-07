package config

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var Rdb *redis.Client

func ConnectRedis() error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	ctx := context.Background()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return err
	}
	Rdb = rdb
	return nil
}
