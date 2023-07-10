package cache

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type RedisCache struct {
	client *redis.Client
}

// NewRedis method creates a new redis client
func NewRedis() *RedisCache {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set (default)
		DB:       0,  // use default DB
	})

	if client == nil {
		fmt.Println("error could not connect to redis")
	}

	return &RedisCache{
		client: client,
	}

}

// Get method gets the value from redis
func (r RedisCache) Get() []byte {
	strCmd := r.client.Get(ctx, "cities")

	cacheBytes, err := strCmd.Bytes()

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return cacheBytes
}

// Put method puts the value to redis
func (r RedisCache) Put(value []byte) {
	err := r.client.Set(ctx, "cities", value, 0).Err()
	if err != nil {
		panic(err)
	}
}
