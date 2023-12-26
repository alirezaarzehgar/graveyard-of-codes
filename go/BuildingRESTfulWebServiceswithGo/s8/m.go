package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	ctx := context.Background()

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		rdb.Set(ctx, "key", "hi", 0)
	}
	fmt.Println(val)
	rdb.Del(ctx, "key")
}
