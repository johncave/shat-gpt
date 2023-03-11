package main

import (
	"context"
)

func redisGet(key string) (string, error) {
	return RedisConn.Get(context.Background(), key).Result()
}

func redisSet(key string, val interface{}) error {
	return RedisConn.Set(context.Background(), key, val, 0).Err()
}

func redisIncrement(key string) (int64, error) {
	return RedisConn.Incr(context.Background(), key).Result()
}
