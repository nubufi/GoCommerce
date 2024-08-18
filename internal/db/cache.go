package db

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"GoCommerce/internal/models"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

// ConnectToRedis connects to the redis database
func ConnectToRedis() {
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: "",
		DB:       0,
	})
}

func SetCache(key string, value interface{}) {
	ctx := context.Background()
	// Serialize the orders to JSON
	jsonOrders, err := json.Marshal(value)
	if err != nil {
		fmt.Println("Error serializing orders:", err)
		return
	}
	RedisClient.Set(ctx, key, jsonOrders, 0)
}

func GetCache[T models.Order | models.Product | models.CartItem](key string, object []T) ([]T, error) {
	ctx := context.Background()
	val := RedisClient.Get(ctx, key).Val()

	if val == "" {
		return nil, errors.New("cache miss")
	}

	err := json.Unmarshal([]byte(val), &object)
	if err != nil {
		return nil, err
	}
	return object, nil
}

func ClearCache(key string) {
	ctx := context.Background()
	RedisClient.Del(ctx, key)
}
