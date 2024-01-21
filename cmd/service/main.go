package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/e99ui/microservice-example/internal/app"
	"github.com/e99ui/microservice-example/internal/config"
	"github.com/e99ui/slog-extra/handlers/pretty"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	sl := setupLogger(cfg.Env)

	if err := app.Run(cfg, sl); err != nil {
		log.Fatalln(err)
	}
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		opts := pretty.PrettyHandlerOptions{
			SlogOpts: &slog.HandlerOptions{
				Level: slog.LevelDebug,
			},
		}

		log = pretty.NewPrettyLogger(opts)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
