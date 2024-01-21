package cars

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type CarService interface{}

type carRoutes struct {
	service CarService
}

func NewCarRoutes(service CarService) chi.Router {
	routes := &carRoutes{
		service: service,
	}

	handler := chi.NewRouter()
	handler.Get("/{id}", routes.Get)
	handler.Get("/count", routes.Count)
	handler.Post("/upload", routes.Upload)

	return handler
}

func (router *carRoutes) Get(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

func (router *carRoutes) Count(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

func (router *carRoutes) Upload(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}
