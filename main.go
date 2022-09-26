package main

import (
	"book-api/app"
	"book-api/config"
)

func main() {
	config := config.GetConfig()

	app := app.App{}
	app.Initialize(config)
	app.Router.Run("localhost:3000")
}
