package httpserver

import (
	"context"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/pedwards95/PDE-Fetch-receipt-processor-challenge/internal/logger"
)

// New httpserver
func New(ctx context.Context, lg *logger.Logger, handler http.Handler) (*http.Server, func(), error) {
	addr := ""
	lg.Infof(ctx, "ENV: %+v", os.Getenv("docker"))
	if os.Getenv("docker") == "true" {
		addr = "0.0.0.0:8080"
	} else {
		addr = "localhost:8080"
	}
	server := &http.Server{
		Addr:         addr,
		Handler:      handler,
		ReadTimeout:  time.Duration(60) * time.Second,
		WriteTimeout: time.Duration(60) * time.Second,
	}
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		lg.Errorf(ctx, "failed to create HTTP listener at address %s: %+v", addr, err)
		return nil, nil, err
	}

	go func() {
		lg.Infof(ctx, "Starting HTTP server at address %s", addr)
		if err := server.Serve(listener); err != http.ErrServerClosed {
			lg.Errorf(ctx, "failed to run HTTP server at address %s: %+v", addr, err)
		}
	}()

	stop := func() {
		defer func() {
			if err := recover(); err != nil {
				lg.Errorf(ctx, "panic while shutting down HTTP server")
			}
		}()
		if server != nil {
			lg.Infof(ctx, "Stopping HTTP server.")
			context, cancel := context.WithTimeout(context.Background(), time.Duration(120)*time.Second)
			defer cancel()
			err := server.Shutdown(context)
			if err != nil {
				lg.Errorf(ctx, "Error while shutting down HTTP server: %+v", err)
			}
		}
	}
	return server, stop, nil
}
