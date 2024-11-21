package receipts

import (
	"github.com/pedwards95/PDE-Fetch-receipt-processor-challenge/internal/localcache"
	"github.com/pedwards95/PDE-Fetch-receipt-processor-challenge/internal/logger"
)

// MockReceiptManager testing only
func MockReceiptManager() (*Manager, func()) {
	lg := logger.New()
	rc, rstop, _ := localcache.New(lg)
	pm, _ := New(rc)

	stop := func() {
		rstop()
	}

	return pm, stop
}
