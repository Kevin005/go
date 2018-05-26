package handler

/**
教程 https://blog.csdn.net/wangshubo1989/article/details/75050024
获取包 go get github.com/garyburd/redigo/redis
 */
import (
	"github.com/garyburd/redigo/redis"

	"fmt"
)

/**
获取redis token
 */
func getTokenOfCache(rToken string) string {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer c.Close()
	token, err := redis.String(c.Do("GET", rToken))
	if err != nil {
		fmt.Println(err)
	}
	return token
}

/**
设置redis token
 */
func setTokenToCache(rToken string) {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()
	_, err = c.Do("SET", rToken, rToken)
	if err != nil {
		fmt.Println(err)
	}
}
