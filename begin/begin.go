package main

import "fmt"
import "github.com/go-redis/redis"

func main() {
	fmt.Printf("hello, world\n")
	client := redis.NewClient(&redis.Options{
			Addr: "localhost:6379",
			Password: "",
			DB: 0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
}