package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/expram/orchestra/internal/app"
)

func main() {
	os.Exit(run())
}

func run() int {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	return app.New().Run(ctx)
}
