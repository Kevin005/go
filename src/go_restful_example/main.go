package main
//注释项目地址 https://github.com/mingrammer/go-todo-rest-api-example
import (
	"go_restful_example/app"
	"go_restful_example/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}