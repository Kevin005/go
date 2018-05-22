package route

import (
	"net/http"
	"awesomeProject/web/restFul/controller"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

/**
路由集
 */
var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		controller.Index,
	},
	Route{
		"TodoIndex",
		"GET",
		"/todos",
		controller.TodoIndex,
	},
	Route{
		"TodoShow",
		"GET",
		"/todos/{todoId}",
		controller.TodoShow,
	},
	Route{
		"TodoCreate",
		"POST",
		"/todos",
		controller.TodoCreate,
	},
}
