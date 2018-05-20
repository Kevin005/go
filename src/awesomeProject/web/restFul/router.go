package main

import (
	"github.com/gorilla/mux"
)

/**
路由管理器
 */
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	//循环加载设置所有路由
	for _, route := range routes {
		//var handler http.Handler
		//handler = route.HandlerFunc
		//handler = Logger(handler, route.Name)
		//设置拦截器路由
		router.
			Methods(route.Method).//get post update delete
			Path(route.Pattern).// /xiaoshi/feedback
			Name(route.Name). //xiaoshi_feedback
			Handler(route.HandlerFunc) //处理函数
	}
	return router
}
