package main

import (
	"context"
	"log"
	"net/http"

	"github.com/olezhek28/clean-architecture/internal/app"
)

func main() {
	ctx := context.Background()

	a, err := app.NewAppRest(ctx)
	if err != nil {
		log.Fatalf("failed to init app: %s", err.Error())
	}

	r := a.Run()
	a.ServiceProvider.UserImpl().BuildRouter(r)
	http.ListenAndServe(a.ServiceProvider.RESTConfig().Address(), r)

	if err != nil {
		log.Fatalf("failed to run app: %s", err.Error())
	}
}
