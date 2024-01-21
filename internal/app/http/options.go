package httpapp

import (
	"net"
	"strconv"
	"time"
)

type Option func(a *App)

func Port(port int) Option {
	return func(a *App) {
		a.server.Addr = net.JoinHostPort("", strconv.Itoa(port))
	}
}

func ReadTimeout(timeout time.Duration) Option {
	return func(a *App) {
		a.server.ReadTimeout = timeout
	}
}

func WriteTimeout(timeout time.Duration) Option {
	return func(a *App) {
		a.server.WriteTimeout = timeout
	}
}

func ShutdownTimeout(timeout time.Duration) Option {
	return func(a *App) {
		a.shutdownTimeout = timeout
	}
}
