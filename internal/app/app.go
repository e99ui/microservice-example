package app

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	httpapp "github.com/e99ui/microservice-example/internal/app/http"
	"github.com/e99ui/microservice-example/internal/config"
	"github.com/e99ui/slog-extra/sl"
)

func Run(cfg *config.Config, log *slog.Logger) error {

	log.Info("starting application", slog.Any("config", cfg))
	httpapp := httpapp.New(
		log,
		httpapp.Port(cfg.Http.Port),
		httpapp.ReadTimeout(cfg.Http.ReadTimeout),
		httpapp.WriteTimeout(cfg.Http.WriteTimeout),
	)

	go func() {
		err := httpapp.Run()
		if err != nil {
			log.With(sl.Err(err))
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	sign := <-stop

	log.Info("stopping application", slog.String("signal", sign.String()))

	if err := httpapp.Stop(); err != nil {
		log.With(sl.Err(err))
	}

	log.Info("application stopped")

	return nil
}
