package points

import (
	"context"

	"github.com/pedwards95/PDE-Fetch-receipt-processor-challenge/internal/errorhandler"
	"github.com/pedwards95/PDE-Fetch-receipt-processor-challenge/internal/localcache"
	"github.com/pedwards95/PDE-Fetch-receipt-processor-challenge/internal/models"

	"github.com/google/uuid"
)

// Manager ...
type Manager struct {
	pointscache  *localcache.LocalCache
	receiptcache *localcache.LocalCache
}

// New PointsManager object
func New(pcache *localcache.LocalCache, rcache *localcache.LocalCache) (*Manager, error) {
	return &Manager{
		pointscache:  pcache,
		receiptcache: rcache,
	}, nil
}

// CalculatePoints checks cache for points, then receipt, then calculates if cant find it
func (pm *Manager) CalculatePoints(ctx context.Context, id uuid.UUID) (*models.Points, error) {
	points := &models.Points{}
	var err error
	var ok bool

	cacheP := pm.pointscache.Get(id)
	if cacheP == nil {
		receipt := &models.Receipt{}
		cacheR := pm.receiptcache.Get(id)
		if cacheR == nil {
			return nil, errorhandler.ObjectNotFoundError("id", id.String())
		}
		receipt, ok = cacheR.(*models.Receipt)
		if !ok {
			return nil, errorhandler.InternalError("casting receipt")
		}
		points, err = pm.calculatePoints(receipt)
		if err != nil || points == nil {
			return nil, err
		}
	} else {
		points, ok = cacheP.(*models.Points)
		if !ok {
			return nil, errorhandler.InternalError("casting points")
		}
	}
	pm.pointscache.Add(id, points)
	return points, nil
}

/*
One point for every alphanumeric character in the retailer name.
50 points if the total is a round dollar amount with no cents.
25 points if the total is a multiple of 0.25.
5 points for every two items on the receipt.
If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned.
6 points if the day in the purchase date is odd.
10 points if the time of purchase is after 2:00pm and before 4:00pm.
*/

// calculatePoints function that encapsulates all point calculations
func (pm *Manager) calculatePoints(receipt *models.Receipt) (*models.Points, error) {
	total := int64(0)

	tAdd, err := pm.calculateRetailerPoints(receipt.Retailer)
	if err != nil {
		return nil, err
	}
	total += tAdd
	tAdd, err = pm.calculateTotalFieldPoints(receipt.Total)
	if err != nil {
		return nil, err
	}
	total += tAdd
	tAdd, err = pm.calculateItemsPoints(receipt.Items, receipt.Total)
	if err != nil {
		return nil, err
	}
	total += tAdd
	tAdd, err = pm.calculateDatePoints(receipt.PurchaseDate)
	if err != nil {
		return nil, err
	}
	total += tAdd
	tAdd, err = pm.calculateTimePoints(receipt.PurchaseTime)
	if err != nil {
		return nil, err
	}
	total += tAdd

	return &models.Points{Points: total}, nil
}
