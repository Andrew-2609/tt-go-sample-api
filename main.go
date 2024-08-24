package main

import (
	"context"
	"tt-go-sample-api/server"
)

func main() {
	mainCtx := context.Background()

	app := server.NewApp()

	go app.Start(mainCtx)

	select {}
}
