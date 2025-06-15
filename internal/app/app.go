package app

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	httpapp "pxr-sso-api/internal/app/http"
	"pxr-sso-api/internal/config"
	"pxr-sso-api/internal/lib/logger/sl"
	"time"
)

// App is an application.
type App struct {
	log     *slog.Logger
	httpApp *httpapp.App
}

// New creates new instance of application.
func New(
	log *slog.Logger,
	cfg *config.Config,
) *App {
	httpapp := httpapp.New(log, cfg.Server.Port)

	return &App{
		log:     log,
		httpApp: httpapp,
	}
}

// MustRun runs the application and panics if an error occurs.
func (a *App) MustRun() {
	if err := a.httpApp.Run(); !errors.Is(err, http.ErrServerClosed) {
		panic(err)
	}
}

// GracefulStop stops the application.
func (a *App) GracefulStop() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := a.httpApp.Stop(ctx); err != nil {
		a.log.Error("HTTP app stop error", sl.Err(err))
	}
}
