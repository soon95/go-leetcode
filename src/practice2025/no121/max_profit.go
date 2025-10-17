package no121

import "math"

func maxProfit(prices []int) int {

	minPrice := math.MaxInt

	ans := math.MinInt
	for _, price := range prices {

		minPrice = min(minPrice, price)

		ans = max(ans, price-minPrice)
	}
	return ans
}
