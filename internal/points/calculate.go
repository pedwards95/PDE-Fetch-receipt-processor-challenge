package points

import (
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/pedwards95/PDE-Fetch-receipt-processor-challenge/internal/errorhandler"
	"github.com/pedwards95/PDE-Fetch-receipt-processor-challenge/internal/models"
)

// ...
const (
	DATE_LAYOUT = "2006-01-02"
	TIME_LAYOUT = "15:04"
)

// calculateRetailerPoints calculation function for retailer name
func (pm *Manager) calculateRetailerPoints(name string) (int64, error) {
	// One point for every alphanumeric character in the retailer name.
	patternName := regexp.MustCompile(`^[a-zA-Z0-9_ \-&]+$`)
	if !patternName.MatchString(name) {
		return 0, errorhandler.ValidationError("Retailer")
	}
	return countAlpha(name), nil
}

// calculateTotalFieldPoints calculation function for receipt total
func (pm *Manager) calculateTotalFieldPoints(totalString string) (int64, error) {
	points := int64(0)

	patternFloat := regexp.MustCompile(`^[+-]?([0-9]+([.][0-9]*)?|[.][0-9]+)$`)
	if !patternFloat.MatchString(totalString) {
		return 0, errorhandler.ValidationError("Total")
	}

	total, err := strconv.ParseFloat(totalString, 64)
	if err != nil {
		return 0, errorhandler.ValidationError("total")
	}

	if total == 0 {
		return 0, nil
	}

	if total < 0 {
		return 0, errorhandler.ValidationError("total", "Total is negative")
	}

	// 50 points if the total is a round dollar amount with no cents.
	if total == math.Trunc(total) {
		points += 50
	}

	// 25 points if the total is a multiple of 0.25.
	divd := total / 0.25
	if (divd) == math.Trunc(divd) {
		points += 25
	}

	return points, nil
}

// calculateItemsPoints calculation function for receipt items
func (pm *Manager) calculateItemsPoints(items []*models.Item, totalString string) (int64, error) {
	points := int64(0)
	patternName := regexp.MustCompile(`^[a-zA-Z0-9_ \-&]+$`)
	patternFloat := regexp.MustCompile(`^[+-]?([0-9]+([.][0-9]*)?|[.][0-9]+)$`)

	total, err := strconv.ParseFloat(totalString, 64)
	if err != nil {
		return 0, errorhandler.ValidationError("total")
	}
	sum := float64(0)

	// 5 points for every two items on the receipt.
	points += (int64(math.Floor((float64(len(items)) / 2)))) * 5

	if len(items) == 0 {
		return 0, errorhandler.ValidationError("items", "No items on receipt")
	}

	// If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned.
	// Note: Spaces count as a character
	for _, item := range items {
		if !patternName.MatchString(item.ShortDescription) {
			return 0, errorhandler.ValidationError("Item Name", "%+v", item.ShortDescription)
		}
		if !patternFloat.MatchString(item.Price) {
			return 0, errorhandler.ValidationError("Item Price", "%+v", item.Price)
		}
		p, err := strconv.ParseFloat(item.Price, 64)
		if err != nil {
			return 0, errorhandler.ValidationError("item price")
		}
		if p < 0 {
			return 0, errorhandler.ValidationError("item.price", "Negative price")
		}
		if len(strings.Trim(item.ShortDescription, " "))%3 == 0 {
			points += int64(math.Ceil(p * 0.2))
		}
		sum += p
	}

	if sum != total {
		return 0, errorhandler.ValidationError("Sum != Total", "Invalid receipt")
	}
	return points, nil
}

// calculateDatePoints calculation function for receipt date
func (pm *Manager) calculateDatePoints(dateStr string) (int64, error) {
	points := int64(0)

	// 6 points if the day in the purchase date is odd.
	date, err := time.Parse(DATE_LAYOUT, dateStr)
	if err != nil {
		return 0, errorhandler.ValidationError("date")
	}

	if date.Day()%2 == 1 {
		points += 6
	}

	return points, nil
}

// calculateTimePoints calculation function for receipt time
func (pm *Manager) calculateTimePoints(timeStr string) (int64, error) {
	points := int64(0)

	parts := strings.Split(timeStr, ":")
	if len(parts) != 2 {
		return 0, errorhandler.ValidationError("time")
	}

	hour, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil || hour < 0 || hour > 24 {
		return 0, errorhandler.ValidationError("time")
	}
	min, err := strconv.ParseInt(parts[1], 10, 32)
	if err != nil || min < 0 || min > 59 {
		return 0, errorhandler.ValidationError("time")
	}

	// 10 points if the time of purchase is after 2:00pm and before 4:00pm.
	if hour >= 14 && hour <= 16 {
		if (hour == 14 || hour == 16) && min == 0 {
			return 0, nil
		}
		points += 10
	}

	return points, nil
}
