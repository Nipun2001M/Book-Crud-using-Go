package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()
var Client *redis.Client

func RedisConnection() {
	Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	

	
	pong, err := Client.Ping(Ctx).Result()
	if err != nil {
		fmt.Println("eror connectiong to redis", err)
	}
	fmt.Println("sucessfully connected redis db", pong)

}
