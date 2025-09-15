package no189

func rotate(nums []int, k int) {

	k = k % len(nums)

	tmp := make([]int, 0)
	for i := len(nums) - k; i < len(nums); i++ {
		tmp = append(tmp, nums[i])
	}

	for i := len(nums) - k - 1; i >= 0; i-- {
		nums[i+k] = nums[i]
	}

	for i := 0; i < len(tmp); i++ {
		nums[i] = tmp[i]
	}
}
