package calculation

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/joho/godotenv"
)

type PackSize int

var packSizesMu sync.RWMutex
var packSizes []PackSize

func init() {
	// Load the pack sizes from the .env file
	if err := godotenv.Load(); err != nil {
		// Handle errors if the .env file is not found or cannot be loaded
		// You can log an error message or take appropriate action here.
	}

	// Retrieve the PACK_SIZES environment variable
	packSizesStr := os.Getenv("PACK_SIZES")

	// Split the comma-separated string into an array of strings
	packSizesArr := strings.Split(packSizesStr, ",")

	// Convert the array of strings to an array of PackSize
	for _, sizeStr := range packSizesArr {
		size, err := strconv.Atoi(sizeStr)
		if err != nil {
			// Handle errors if conversion fails
			// You can log an error message or take appropriate action here.
			continue
		}
		packSizes = append(packSizes, PackSize(size))
	}
}

func SetPackSizes(sizes []PackSize) {
	packSizesMu.Lock()
	defer packSizesMu.Unlock()

	packSizes = sizes
}

func GetPackSizes() []PackSize {
	packSizesMu.RLock()
	defer packSizesMu.RUnlock()

	return packSizes
}
func CalculatePacks(orderQuantity float64) []string {
	// Define your standard pack sizes in descending order
	packSizes := GetPackSizes()
	orderQuantity = math.Ceil(orderQuantity / float64(packSizes[len(packSizes)-1]))
	// Initialize a map to store the pack sizes and their counts
	packCounts := make(map[int]int)
	for i := 0; i < len(packSizes); {
		if orderQuantity >= float64(packSizes[i])/250 {
			orderQuantity -= (float64(packSizes[i]) / 250)
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
