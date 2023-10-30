package main

import (
	"context"
	"log"

	"github.com/olezhek28/clean-architecture/internal/app"
)

func main() {
	ctx := context.Background()

	a, err := app.NewAppRest(ctx)
	if err != nil {
		log.Fatalf("failed to init app: %s", err.Error())
	}

	err = a.Run()
	if err != nil {
		log.Fatalf("failed to run app: %s", err.Error())
	}
}
