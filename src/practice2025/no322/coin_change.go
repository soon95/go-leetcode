package no322

import "math"

func coinChange(coins []int, amount int) int {

	dp := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		tmp := math.MaxInt
		for _, coin := range coins {

			if i-coin >= 0 && dp[i-coin] != -1 {
				tmp = min(tmp, dp[i-coin]+1)
			}
		}

		if tmp == math.MaxInt {
			dp[i] = -1
		} else {
			dp[i] = tmp
		}
	}

	return dp[amount]
}
