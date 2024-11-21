package points

import (
	"testing"

	"github.com/pedwards95/PDE-Fetch-receipt-processor-challenge/internal/models"

	"github.com/stretchr/testify/assert"
)

// One point for every alphanumeric character in the retailer name.
func TestCalculateRetailerPoints(t *testing.T) {
	pm, stop := MockPointsManager()
	defer stop()

	res, err := pm.calculateRetailerPoints("Walmart")
	assert.NoError(t, err)
	assert.Equal(t, int64(7), res)

	res, err = pm.calculateRetailerPoints("")
	assert.Error(t, err)
	assert.Equal(t, int64(0), res)

	res, err = pm.calculateRetailerPoints("      ")
	assert.NoError(t, err)
	assert.Equal(t, int64(0), res)

	res, err = pm.calculateRetailerPoints("123--abc")
	assert.NoError(t, err)
	assert.Equal(t, int64(6), res)

	res, err = pm.calculateRetailerPoints("No Name Monkies")
	assert.NoError(t, err)
	assert.Equal(t, int64(13), res)

	res, err = pm.calculateRetailerPoints("🤡🤡🤡🤡")
	assert.Error(t, err)
	assert.Equal(t, int64(0), res)
}

// 50 points if the total is a round dollar amount with no cents.
// 25 points if the total is a multiple of 0.25
func TestCalculateTotalFieldPoints(t *testing.T) {
	pm, stop := MockPointsManager()
	defer stop()

	res, err := pm.calculateTotalFieldPoints("1.50")
	assert.NoError(t, err)
	assert.Equal(t, int64(25), res)

	res, err = pm.calculateTotalFieldPoints("0")
	assert.NoError(t, err)
	assert.Equal(t, int64(0), res)

	res, err = pm.calculateTotalFieldPoints("100.00")
	assert.NoError(t, err)
	assert.Equal(t, int64(75), res)

	res, err = pm.calculateTotalFieldPoints("11.11")
	assert.NoError(t, err)
	assert.Equal(t, int64(0), res)

	res, err = pm.calculateTotalFieldPoints("100.001")
	assert.NoError(t, err)
	assert.Equal(t, int64(0), res)

	res, err = pm.calculateTotalFieldPoints("100")
	assert.NoError(t, err)
	assert.Equal(t, int64(75), res)

	res, err = pm.calculateTotalFieldPoints("-50.50")
	assert.Error(t, err)
	assert.Equal(t, int64(0), res)
}

// 5 points for every two items on the receipt.
// If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned.
func TestCalculateItemsPoints(t *testing.T) {
	pm, stop := MockPointsManager()
	defer stop()

	items := []*models.Item{}
	res, err := pm.calculateItemsPoints(items, "0")
	assert.Error(t, err)
	assert.Equal(t, int64(0), res)

	newItem := &models.Item{
		ShortDescription: "Item1",
		Price:            "1.50",
	}
	items = append(items, newItem)

	res, err = pm.calculateItemsPoints(items, "1.50")
	assert.NoError(t, err)
	assert.Equal(t, int64(0), res)

	newItem = &models.Item{
		ShortDescription: "🤡🤡🤡123",
		Price:            "0.87",
	}
	items = append(items, newItem)

	res, err = pm.calculateItemsPoints(items, "2.37")
	assert.Error(t, err)
	assert.Equal(t, int64(0), res)

	items = items[:len(items)-1]

	newItem = &models.Item{
		ShortDescription: "Item22abc",
		Price:            "11.99",
	}
	items = append(items, newItem)

	res, err = pm.calculateItemsPoints(items, "13.49")
	assert.NoError(t, err)
	assert.Equal(t, int64(8), res)

	newItem = &models.Item{
		ShortDescription: "",
		Price:            "0",
	}
	items = append(items, newItem)

	res, err = pm.calculateItemsPoints(items, "13.49")
	assert.Error(t, err)
	assert.Equal(t, int64(0), res)

	items = items[:len(items)-1]

	newItem = &models.Item{
		ShortDescription: "badite",
		Price:            "-1.50",
	}
	items = append(items, newItem)

	res, err = pm.calculateItemsPoints(items, "11.99")
	assert.Error(t, err)
	assert.Equal(t, int64(0), res)
}

// 6 points if the day in the purchase date is odd.
func TestCalculateDatePoints(t *testing.T) {
	pm, stop := MockPointsManager()
	defer stop()

	res, err := pm.calculateDatePoints("2000-10-04")
	assert.NoError(t, err)
	assert.Equal(t, int64(0), res)

	res, err = pm.calculateDatePoints("2000-10-05")
	assert.NoError(t, err)
	assert.Equal(t, int64(6), res)

	res, err = pm.calculateDatePoints("bad string")
	assert.Error(t, err)
	assert.Equal(t, int64(0), res)
}

// 10 points if the time of purchase is after 2:00pm and before 4:00pm.
func TestCalculateTimePoints(t *testing.T) {
	pm, stop := MockPointsManager()
	defer stop()

	res, err := pm.calculateTimePoints("1:00")
	assert.NoError(t, err)
	assert.Equal(t, int64(0), res)

	res, err = pm.calculateTimePoints("14:00")
	assert.NoError(t, err)
	assert.Equal(t, int64(0), res)

	res, err = pm.calculateTimePoints("16:00")
	assert.NoError(t, err)
	assert.Equal(t, int64(0), res)

	res, err = pm.calculateTimePoints("15:30")
	assert.NoError(t, err)
	assert.Equal(t, int64(10), res)

	res, err = pm.calculateTimePoints("13:120")
	assert.Error(t, err)
	assert.Equal(t, int64(0), res)

	res, err = pm.calculateTimePoints("bad string")
	assert.Error(t, err)
	assert.Equal(t, int64(0), res)
}
