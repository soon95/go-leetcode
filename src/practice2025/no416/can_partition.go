package no416

func canPartition(nums []int) bool {

	sum := 0
	for _, num := range nums {
		sum += num
	}

	dp := make([]map[int]bool, len(nums))

	for i := 0; i < len(nums); i++ {
		dp[i] = make(map[int]bool)
	}

	dp[0] = map[int]bool{
		nums[0]: true,
	}

	for i := 1; i < len(nums); i++ {
		for key := range dp[i-1] {
			dp[i][key] = true
			dp[i][key+nums[i]] = true
		}
	}

	for key := range dp[len(nums)-1] {
		if key == sum-key {
			return true
		}
	}

	return false
}
