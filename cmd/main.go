package main

import (
	"context"

	"leather-shop/config"
	"leather-shop/internal/app"
)

func main() {
	ctx := context.Background()

	config := config.GetConfig()
	aplication := app.New(ctx, config)
	aplication.Run(ctx)
}
