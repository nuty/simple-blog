package providers

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	// "github.com/nuty/simple-blog/config"
	"log"
)

var Rdb *redis.Client

func InitRedis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",         
		DB:       0, 
	})
	pong, err := Rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
	fmt.Println("Connected to Redis! Response:", pong)
}
