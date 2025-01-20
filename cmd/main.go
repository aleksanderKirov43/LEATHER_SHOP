package main

import (
	"leather-shop/config"
	"leather-shop/internal/app"
)

func main() {
	config := config.GetConfig()
	aplication := app.New(config)
	aplication.Run()
}
