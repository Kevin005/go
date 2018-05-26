package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
	"time"
)

func main() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()

	// 写入值10S后过期
	_, err = c.Do("SET", "password", "123456", "EX", "5")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}
	time.Sleep(6 * time.Second)
	password, err := redis.String(c.Do("GET", "password"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Got password %v \n", password)
	}
}
