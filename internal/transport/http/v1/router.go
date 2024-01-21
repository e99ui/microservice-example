package v1

import (
	"net/http"

	"github.com/e99ui/microservice-example/internal/transport/http/v1/cars"
	"github.com/go-chi/chi/v5"
)

func NewRouter(handler chi.Router, carService cars.CarService) {
	// K8s health check
	handler.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	handler.Route("/api/v1", func(r chi.Router) {
		handler.Mount("/cars", cars.NewCarRoutes(carService))
	})
}
