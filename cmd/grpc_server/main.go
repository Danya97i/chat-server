package main

import (
	"context"

	"github.com/Danya97i/chat-server/internal/app"
)

func main() {
	ctx := context.Background()
	a, err := app.NewApp(ctx)
	if err != nil {
		panic(err)
	}
	err = a.Run()
	if err != nil {
		panic(err)
	}
}
