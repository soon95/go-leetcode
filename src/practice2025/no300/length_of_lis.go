package no300

func lengthOfLIS(nums []int) int {

	dp := make([]int, len(nums))

	dp[0] = 1
	ans := 1
	for i := 1; i < len(nums); i++ {
		tmp := 0
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				tmp = max(tmp, dp[j])
			}
		}
		dp[i] = tmp + 1

		ans = max(ans, dp[i])

	}

	return ans
}
