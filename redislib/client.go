package redislib

import (
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis"
)

func Client() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		panic("Can't connect redis...")
	}
	return client
}

func SetKey(key string, value interface{}) {
	json, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}

	err = Client().Set(key, json, 0).Err()

	if err != nil {
		fmt.Println(err)
		panic("Can't not set redis key...")
	}

	fmt.Println("Set key successfully!")
}

func GetKey(key string) string {
	value := ""
	value, err := Client().Get(key).Result()
	if err != nil {
		fmt.Println(err)
		return value
	}

	// json.Unmarshal([]byte(jsonString), &book)
	return value
}
