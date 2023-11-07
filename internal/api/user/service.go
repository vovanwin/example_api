package user

import (
	"github.com/go-chi/chi/v5"
	"github.com/olezhek28/clean-architecture/internal/service"
	desc "github.com/olezhek28/clean-architecture/pkg/user_v1"
	"net/http"
)

type Implementation struct {
	desc.UnimplementedUserV1Server
	userService service.UserService
}

func NewImplementation(userService service.UserService) *Implementation {
	return &Implementation{
		userService: userService,
	}
}

func (i *Implementation) BuildRouter(r *chi.Mux) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	r.Get("/test", i.CreateREST)
}
