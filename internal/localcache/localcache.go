package localcache

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/pedwards95/PDE-Fetch-receipt-processor-challenge/internal/logger"
)

// LocalCache ...
type LocalCache struct {
	logger         *logger.Logger
	cache          map[uuid.UUID]*CachedObject
	cleanerChannel chan struct{}
	doneChannel    chan struct{}
	cleaning       bool
}

// CachedObject ...
type CachedObject struct {
	expire time.Time
	object interface{}
}

// New local cache
func New(logger *logger.Logger) (*LocalCache, func(), error) {
	lc := &LocalCache{
		logger:         logger,
		cache:          make(map[uuid.UUID]*CachedObject),
		cleanerChannel: make(chan struct{}, 1),
		doneChannel:    make(chan struct{}, 1),
		cleaning:       false,
	}
	go lc.cleaner(30)
	return lc, lc.Stop, nil
}

// cleaner for cleaning out local cache
func (lc *LocalCache) cleaner(interval int32) {
	defer func() {
		lc.doneChannel <- struct{}{}
	}()

	ticker := int32(0)
	for {
		select {
		case <-lc.cleanerChannel:
			lc.logger.Infof(context.Background(), "Stopping LocalCache cleaner")
			return
		default:
			if ticker >= interval && !lc.cleaning {
				go lc.cleancache()
				lc.cleaning = true
				ticker = 0
			} else if ticker >= interval && lc.cleaning {
				//already cleaning
				ticker = 0
			} else {
				ticker++
			}
		}
		time.Sleep(1 * time.Second)
	}
}

// cleancache go function that does the actual cache cleaning
func (lc *LocalCache) cleancache() {
	lc.logger.Infof(context.Background(), "cleaning cache")
	for key, obj := range lc.cache {
		if time.Now().After(obj.expire) {
			lc.Remove(key)
		}
	}
	lc.cleaning = false
}

// Stop cache and cleaners
func (lc *LocalCache) Stop() {
	lc.cleanerChannel <- struct{}{}
	<-lc.doneChannel
	lc.logger.Infof(context.Background(), "LocalCache stopped")
}
