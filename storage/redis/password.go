package redis

import (
	"context"
	"log"
	"time"

	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
)

func ConnectDB() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	err := rdb.Ping(context.Background()).Err()
	if err != nil {
		log.Fatalf("could not connect to redis: %v", err)
	}

	return rdb
}

func StoreCode(ctx context.Context, id string, code string) error {
	rdb := ConnectDB()
	err := rdb.Set(ctx, id, code, 5*time.Minute).Err()
	if err != nil {
		return errors.Wrap(err, "code storing failure")
	}
	return nil
}

func GetCode(ctx context.Context, id string) (string, error) {
	rdb := ConnectDB()
	code, err := rdb.Get(ctx, id).Result()
	if err != nil {
		return "", errors.Wrap(err, "code getting failure")
	}
	return code, nil
}

func DeleteCode(ctx context.Context, id string) error {
	rdb := ConnectDB()
	err := rdb.Del(ctx, id).Err()
	if err != nil {
		return errors.Wrap(err, "code deleting failure")
	}
	return nil
}
