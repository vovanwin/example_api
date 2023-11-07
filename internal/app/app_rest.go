package app

import (
	"context"
	"flag"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/olezhek28/clean-architecture/internal/api/user"
	desc "github.com/olezhek28/clean-architecture/pkg/user_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"log"
	"net/http"

	"github.com/olezhek28/clean-architecture/internal/config"
)

type AppRest struct {
	ServiceProvider *serviceProvider
	grpcServer      *grpc.Server
}

func NewAppRest(ctx context.Context) (*AppRest, error) {
	a := &AppRest{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *AppRest) Run() *chi.Mux {
	log.Printf("REST server is running on %s", a.ServiceProvider.RESTConfig().Address())

	flag.Parse()

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	return r
}

func (a *AppRest) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *AppRest) initConfig(_ context.Context) error {
	err := config.Load(".env")
	if err != nil {
		return err
	}

	return nil
}

func (a *AppRest) initServiceProvider(_ context.Context) error {
	a.ServiceProvider = newServiceProvider()
	return nil
}

func (a *AppRest) initGRPCServer(_ context.Context) error {
	a.grpcServer = grpc.NewServer(grpc.Creds(insecure.NewCredentials()))

	reflection.Register(a.grpcServer)

	desc.RegisterUserV1Server(a.grpcServer, a.ServiceProvider.UserImpl())

	return nil
}

func (a *AppRest) runRESTServer() error {
	log.Printf("REST server is running on %s", a.ServiceProvider.RESTConfig().Address())

	flag.Parse()

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	user.NewImplementation(a.ServiceProvider.userService).BuildRouter(r)

	http.ListenAndServe(a.ServiceProvider.RESTConfig().Address(), r)

	return nil
}
