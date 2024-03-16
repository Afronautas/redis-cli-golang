package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	context := context.Background()
	// err := redisClient.Set(context, "name", "Afronautas", 0).Err()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Operation executed successfully!")
	result, err := redisClient.Get(context, "name").Result()
	if err != nil {
		err := redisClient.Set(context, "name", "Afronautas has not been found...", 0).Err()
		if err != nil {
			panic(err)
		}
		fmt.Println("Operation executed successfully!")
	}
	fmt.Println("Result: ", result)

	// err = redisClient.Del(context, "name").Err()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Key has been deleted successfully!")
}
