package main

import (
	"log"
	"net/http"
	"awesomeProject/web/restFul/route"
)

func main() {
	router := route.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
