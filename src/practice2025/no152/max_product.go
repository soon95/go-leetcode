package no152

func maxProduct(nums []int) int {

	minDp := make([]int, len(nums))
	maxDp := make([]int, len(nums))
	minDp[0], maxDp[0] = nums[0], nums[0]

	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		minDp[i] = min(nums[i], min(maxDp[i-1]*nums[i], minDp[i-1]*nums[i]))
		maxDp[i] = max(nums[i], max(maxDp[i-1]*nums[i], minDp[i-1]*nums[i]))
		ans = max(ans, maxDp[i])
	}
	return ans
}
