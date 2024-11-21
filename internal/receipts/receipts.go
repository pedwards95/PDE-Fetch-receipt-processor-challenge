package receipts

import (
	"context"

	"github.com/pedwards95/PDE-Fetch-receipt-processor-challenge/internal/localcache"
	"github.com/pedwards95/PDE-Fetch-receipt-processor-challenge/internal/models"

	"github.com/google/uuid"
)

// Manager ...
type Manager struct {
	receiptcache *localcache.LocalCache
}

// New ReceiptsManager object
func New(rcache *localcache.LocalCache) (*Manager, error) {
	return &Manager{
		receiptcache: rcache,
	}, nil
}

// ProcessReceipt returns a new id and caches item
func (rm *Manager) ProcessReceipt(ctx context.Context, receipt *models.Receipt) (*models.ID, error) {
	id := uuid.New()
	rm.receiptcache.Add(id, receipt)
	return &models.ID{ID: id}, nil
}
