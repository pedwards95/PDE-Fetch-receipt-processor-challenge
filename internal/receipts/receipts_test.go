package receipts

import (
	"context"
	"testing"

	"github.com/pedwards95/PDE-Fetch-receipt-processor-challenge/internal/models"

	"github.com/stretchr/testify/assert"
)

func TestProcessReceipt(t *testing.T) {
	pm, stop := MockReceiptManager()
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

	ctx := context.Background()
	id, err := pm.ProcessReceipt(ctx, receipt)
	assert.NoError(t, err)

	resp, ok := pm.receiptcache.Get(id.ID).(*models.Receipt)
	assert.True(t, ok)
	assert.Equal(t, receipt.Retailer, resp.Retailer)

}
