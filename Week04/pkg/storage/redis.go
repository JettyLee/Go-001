package storage

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
)

func NewRedisDB() *redis.Client {

	var address, auth string
	address = "localhost:6379"
	auth = "root"
	redisDB := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: auth, // no password set
		DB:       0,    // use default DB
	})
	pong, err := redisDB.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal("redis连接错误")
	}
	if pong != "PONG" {
		log.Fatal("redis连接错误")
	} else {
		log.Println("redis 连接成功")
	}
	return redisDB
}
