package main

import (
	"cRoutes/comrouting"
	"fmt"
)

func main() {
	router := comrouting.CreateCRouter()
	comrouting.AddCRouteItem(router, "192.168.0.0", 16)
	comrouting.AddCRouteItem(router, "192.168.1.0", 24)
	comrouting.PrintCRoute(router)
	ret := comrouting.SearchCRoute(router, "192.168.1.0")
	fmt.Printf("ret=%v\n\n", ret)

	comrouting.DelCRouteItem(router, "192.168.1.0", 24)
	comrouting.PrintCRoute(router)
	ret = comrouting.SearchCRoute(router, "192.168.1.0")
	fmt.Printf("ret=%v\n", ret)
}
