package cache

import (
	"os"
	"github.com/go-redis/redis/v8"
)

func GetRedisClient() *redis.Client { 
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_DSN"), 
		Password: "",                     
		DB:       0,                      
	})

	return client 
}