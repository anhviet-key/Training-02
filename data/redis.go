package data

import (
	"fmt"
	"log"

	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err := RedisClient.Ping().Result()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected succesfully to Redis on Port: 6379")
}
