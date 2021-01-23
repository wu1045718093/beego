package model

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var redisdb *redis.Client

func init() {
	redisdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := redisdb.Ping(context.TODO()).Result()
	if err != nil {
		fmt.Println("Redis 连接失败", err)
	} else {
		fmt.Println("Redis 连接成功")
	}
}

func RedisGet(email string, password string) (collect bool, ok bool) {
	val, err := redisdb.Get(context.TODO(), email).Result()
	switch {
	case err == redis.Nil:
		fmt.Println("key does not exist")
		return false, false
	case err != nil:
		fmt.Println("Get failed", err)
		return false, false
	}

	if val == password {
		fmt.Println("密码匹配成功", err)
		return true, true
	} else {
		fmt.Println("密码不匹配", err)
		return false, true
	}
}

func RedisSet(key string, val string) bool {
	err := redisdb.Set(context.TODO(), key, val, 3600*time.Second).Err()
	if err != nil {
		fmt.Printf("Redis set fail err: %v\n", err)
		return false
	} else {
		return true
	}
}
