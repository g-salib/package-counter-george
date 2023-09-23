package calculation

import "strconv"

func CalculatePacks(orderQuantityStr string) map[int]int {
	orderQuantity, _ := strconv.Atoi(orderQuantityStr)
	packSizes := []int{PackSize5000, PackSize2000, PackSize1000, PackSize500, PackSize250}
	packs := make(map[int]int)

	for _, packSize := range packSizes {
		packCount := orderQuantity / packSize
		if packCount > 0 {
			packs[packSize] = packCount
			orderQuantity %= packSize
		}
	}

	return packs
}
