package no45

func jump(nums []int) int {

	if len(nums) == 1 {
		return 0
	}

	step := 0

	preMax := 0
	for {
		step++

		curMax := preMax
		for i := 0; i < len(nums) && i <= preMax; i++ {
			curMax = max(curMax, i+nums[i])
		}

		if curMax >= len(nums)-1 {
			break
		}

		preMax = curMax
	}

	return step
}
