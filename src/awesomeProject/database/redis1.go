package main

/*
地址  github.com/gomodule/redigo/redis
 */
import (
	"github.com/gomodule/redigo/redis"
	"fmt"
)

func main() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()

	c.Send("SET", "foo", "a")
	c.Send("GET", "foo")
	c.Flush()
	c.Receive()           // reply from SET
	v, err := c.Receive() // reply from GET
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(v)
}
