package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/pedwards95/PDE-Fetch-receipt-processor-challenge/internal/logger"
)

// App ...
type App struct {
	logger     *logger.Logger
	handler    http.Handler
	httpserver *http.Server

	donesMu sync.RWMutex
	dones   []chan os.Signal
}

// ...
const (
	SERVICE = "Fetch-receipt-processor-challenge"
)

// NewApp used to initialize a new application
func NewApp(logger *logger.Logger, handler http.Handler, httpserver *http.Server) (app *App, stop func(), err error) {
	return &App{
			logger:     logger,
			handler:    handler,
			httpserver: httpserver,
		}, func() {
			logger.Infof(context.Background(), "Stopping %s", SERVICE)
		}, nil
}

// Done is app done shutting down
func (app *App) Done() <-chan os.Signal {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	app.donesMu.Lock()
	app.dones = append(app.dones, signals)
	app.donesMu.Unlock()

	return signals
}

// Stop the application
func (app *App) Stop(stop func()) {
	done := <-app.Done()
	app.logger.Infof(context.Background(), "Received shutdown signal %s", done.String())
	stop()
	os.Exit(0)
}

func main() {
	app, stop, err := InitializeAndRun(context.Background())
	if err != nil {
		logger := logger.New()
		logger.Errorf(context.Background(), "Application initialization failed. Error: %+v", err)
		os.Exit(2)
	}
	app.Stop(stop)
}
