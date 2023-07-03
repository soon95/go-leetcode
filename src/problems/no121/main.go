package main

import (
	"fmt"
	"math"
)

/*
*121. 买卖股票的最佳时机
https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/
*/
func maxProfit(prices []int) int {

	ans := 0

	min := math.MaxInt

	for _, price := range prices {

		if min > price {
			min = price
		}

		temp := price - min

		if ans < temp {
			ans = temp
		}

	}

	return ans
}

func main() {

	prices := []int{7}

	fmt.Printf("%v", maxProfit(prices))

}
