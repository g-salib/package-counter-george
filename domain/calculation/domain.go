package calculation

import (
	"fmt"
	"math"
)

func CalculatePacks(orderQuantity float64) []string {
	// Define your standard pack sizes in descending order
	packSizes := []float64{5000, 2000, 1000, 500, 250}
	orderQuantity = math.Ceil(orderQuantity / float64(packSizes[len(packSizes)-1]))
	// Initialize a map to store the pack sizes and their counts
	packCounts := make(map[int]int)
	for i := 0; i < len(packSizes); {
		if orderQuantity >= packSizes[i]/250 {
			orderQuantity -= (packSizes[i] / 250)
			packCounts[int(packSizes[i])]++
		} else {
			i++
		}

	}

	// Convert the pack counts to strings in the desired format
	result := []string{}
	for packSize, packCount := range packCounts {
		if packCount > 0 {
			result = append(result, fmt.Sprintf("%dx%d", packCount, packSize))
		}
	}

	return result
}
