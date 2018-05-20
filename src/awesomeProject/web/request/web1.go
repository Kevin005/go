package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

func httpGet() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		//
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//
	}
	fmt.Println(string(body))
}

func main() {
	httpGet()
}
