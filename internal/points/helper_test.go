package points

import (
	"github.com/pedwards95/PDE-Fetch-receipt-processor-challenge/internal/localcache"
	"github.com/pedwards95/PDE-Fetch-receipt-processor-challenge/internal/logger"
)

// MockPointsManager testing only
func MockPointsManager() (*Manager, func()) {
	lg := logger.New()
	pc, pstop, _ := localcache.New(lg)
	rc, rstop, _ := localcache.New(lg)
	pm, _ := New(pc, rc)

	stop := func() {
		pstop()
		rstop()
	}

	return pm, stop
}
