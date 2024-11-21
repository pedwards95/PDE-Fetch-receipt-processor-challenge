package main

import (
	"context"

	"github.com/pedwards95/PDE-Fetch-receipt-processor-challenge/internal/handler"
	"github.com/pedwards95/PDE-Fetch-receipt-processor-challenge/internal/httpserver"
	"github.com/pedwards95/PDE-Fetch-receipt-processor-challenge/internal/localcache"
	"github.com/pedwards95/PDE-Fetch-receipt-processor-challenge/internal/logger"
	"github.com/pedwards95/PDE-Fetch-receipt-processor-challenge/internal/points"
	"github.com/pedwards95/PDE-Fetch-receipt-processor-challenge/internal/receipts"
)

// InitializeAndRun dependency injection and configs
func InitializeAndRun(ctx context.Context) (*App, func(), error) {

	logger := logger.New()

	pcache, pCacheStop, err := localcache.New(logger)
	if err != nil {
		return nil, nil, err
	}

	rcache, rCacheStop, err := localcache.New(logger)
	if err != nil {
		pCacheStop()
		return nil, nil, err
	}

	rm, err := receipts.New(rcache)
	if err != nil {
		pCacheStop()
		rCacheStop()
		return nil, nil, err
	}

	pm, err := points.New(pcache, rcache)
	if err != nil {
		pCacheStop()
		rCacheStop()
		return nil, nil, err
	}

	handler, err := handler.New(logger, pm, rm)
	if err != nil {
		pCacheStop()
		rCacheStop()
		return nil, nil, err
	}

	httpserver, serverStop, err := httpserver.New(ctx, logger, handler)
	if err != nil {
		pCacheStop()
		rCacheStop()
		return nil, nil, err
	}

	app, appStop, err := NewApp(logger, handler, httpserver)
	if err != nil {
		pCacheStop()
		rCacheStop()
		serverStop()
		return nil, nil, err
	}

	stopFunc := func() {
		serverStop()
		pCacheStop()
		rCacheStop()
		serverStop()
		appStop()
	}

	return app, stopFunc, nil

}
