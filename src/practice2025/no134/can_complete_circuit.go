package no134

import "math"

func canCompleteCircuit(gas []int, cost []int) int {

	dp := make([]int, len(gas))
	for i := 0; i < len(gas); i++ {
		dp[i] = gas[i] - cost[i]
	}

	preSum := math.MaxInt
	preIndex := 0
	tmp := 0
	for i := 0; i < len(dp); i++ {
		tmp += dp[i]
		if preSum > tmp {
			preSum = tmp
			preIndex = i
		}
	}
	if tmp < 0 {
		return -1
	}

	tmp = 0
	for i := preIndex + 1; i < len(dp); i++ {
		tmp += dp[i]
		if tmp < 0 {
			return -1
		}
	}

	return (preIndex + 1) % len(dp)
}
