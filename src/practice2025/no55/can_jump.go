package no55

func canJump(nums []int) bool {

	dp := make([]bool, len(nums))
	dp[0] = true

	for i, num := range nums {

		if !dp[i] {
			break
		}

		for j := 1; i+j < len(nums) && j <= num; j++ {
			dp[i+j] = true
		}

	}

	return dp[len(dp)-1]
}
