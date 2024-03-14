package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	redisHost = "localhost:6379"
	// redisPort     = 6379
	redisPassword = ""
	redisDB       = 0
)

func RedisConnect() (*redis.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	rdb := redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: redisPassword, // no password set
		DB:       redisDB,       // use default DB
	})

	status, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalln("redis connection was refused")
		return rdb, err
	}

	fmt.Println(status)
	fmt.Println("redis connection success!")
	return rdb, nil
}
