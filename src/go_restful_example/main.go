package main
import (
	"go_restful_example/app"
	"github.com/Kevin005/go_restful_example/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}