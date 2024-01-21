package httpapp

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	v1 "github.com/e99ui/microservice-example/internal/transport/http/v1"
	"github.com/go-chi/chi/v5"
)

const (
	defaultReadTimeout     = 5 * time.Second
	defaultWriteTimeout    = 5 * time.Second
	defaultAddr            = ":8000"
	defaultShutdownTimeout = 3 * time.Second
)

type App struct {
	log             *slog.Logger
	server          *http.Server
	shutdownTimeout time.Duration
}

func New(log *slog.Logger, opts ...Option) *App {
	handler := chi.NewRouter()
	v1.NewRouter(handler, nil)

	httpServer := &http.Server{
		Handler:      handler,
		ReadTimeout:  defaultReadTimeout,
		WriteTimeout: defaultReadTimeout,
		Addr:         defaultAddr,
	}

	a := &App{
		log:             log,
		server:          httpServer,
		shutdownTimeout: defaultShutdownTimeout,
	}

	for _, opt := range opts {
		opt(a)
	}

	return a
}

func (a *App) Run() error {
	const op = "httpapp.Run"

	log := a.log.With(
		slog.String("op", op),
	)

	log.Info("http server is running", slog.String("addr", a.server.Addr))

	if err := a.server.ListenAndServe(); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a *App) Stop() error {
	const op = "htppapp.Stop"

	a.log.With(slog.String("op", op)).Info("stopping http server")

	ctx, cancel := context.WithTimeout(context.Background(), a.shutdownTimeout)
	defer cancel()

	if err := a.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
