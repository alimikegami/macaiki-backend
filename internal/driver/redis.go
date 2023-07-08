package driver

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var RDB *redis.Client

func CreateRedisConnection(address, password string) {
	RDB = redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       0,
	})

	_, err := RDB.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
}
