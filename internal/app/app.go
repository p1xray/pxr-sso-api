package app

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"pxr-sso-api/pkg/logger/sl"
	"syscall"
	"time"
)

const logTag = "[pxr-sso-api:app]"

// App is the application structure.
type App struct {
	di         *diContainer
	httpServer *http.Server
}

// New creates a new application and initializes all dependencies through the DI container.
func New() *App {
	a := &App{
		di: newDIContainer(),
	}

	a.initDeps()

	return a
}

// Start launches the application.
func (a *App) Start() {
	log := a.di.Logger()

	log.Info(logTag + " starting application")
	log.Debug(logTag+" application configuration", slog.Any("config", a.di.Config()))

	if err := a.httpServer.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		panic(err)
	}
}

// GracefulStop gracefully stops the application.
func (a *App) GracefulStop() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	s := <-stop

	log := a.di.Logger()
	log.Debug(logTag + " signal received from OS: " + s.String())
	log.Info(logTag + " stopping application...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := a.httpServer.Shutdown(ctx); err != nil {
		log.Error(logTag+" application stop error", sl.Err(err))
	}

	log.Info(logTag + " application stopped")
}

// initDeps sequentially calls initialization functions.
// If you need to add a new step (migration, metrics, etc.),
// you must add the function to the inits slice.
func (a *App) initDeps() {
	inits := []func(){
		a.initHTTPServer,
	}

	for _, fn := range inits {
		fn()
	}
}

// initGRPCServer initializes the HTTP server.
func (a *App) initHTTPServer() {
	cfg := a.di.Config()

	a.httpServer = &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Port),
		Handler: a.di.Handler().Routes(),
	}
}
