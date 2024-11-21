package points

import (
	"context"
	"testing"

	"github.com/pedwards95/PDE-Fetch-receipt-processor-challenge/internal/models"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCalculatePoints(t *testing.T) {
	pm, stop := MockPointsManager()
	defer stop()

	receipt := &models.Receipt{
		Retailer:     "Fetch",
		PurchaseDate: "2024-11-19",
		PurchaseTime: "20:25",
		Items: []*models.Item{
			{ShortDescription: "Items1", Price: "15.60"},
			{ShortDescription: "Item2", Price: "1.23"},
			{ShortDescription: "Item3", Price: "0.50"},
			{ShortDescription: "Items4", Price: "0.67"},
		},
		Total: "18.00",
	}

	id := uuid.New()
	pm.receiptcache.Add(id, receipt)

	ctx := context.Background()
	res, err := pm.CalculatePoints(ctx, id)
	assert.NoError(t, err)
	assert.Equal(t, int64(101), res.Points)
}
