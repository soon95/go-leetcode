package no279

import "math"

func numSquares(n int) int {

	dp := make([]int, n+1)

	dp[0] = 0

	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt
		for j := 1; j <= int(math.Sqrt(float64(i))); j++ {

			if i-j*j >= 0 {
				dp[i] = min(dp[i], dp[i-j*j]+1)
			}
		}
	}

	return dp[n]
}
