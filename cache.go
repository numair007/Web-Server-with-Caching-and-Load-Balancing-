// cache.go
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func setCache(client *redis.Client, key string, value string, expiration time.Duration) error {
	err := client.Set(ctx, key, value, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

func getCache(client *redis.Client, key string) (string, error) {
	val, err := client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address
		Password: "",               // No password
		DB:       0,                // Default DB
	})

	err := setCache(client, "example_key", "example_value", 10*time.Second)
	if err != nil {
		log.Fatal(err)
	}

	cachedValue, err := getCache(client, "example_key")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Cached Value:", cachedValue)
}
