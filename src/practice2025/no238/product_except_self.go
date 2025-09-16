package no238

func productExceptSelf(nums []int) []int {

	left, right := make([]int, len(nums)), make([]int, len(nums))

	tmp := 1
	for i := 0; i < len(nums)-1; i++ {
		tmp = tmp * nums[i]
		left[i] = tmp
	}

	tmp = 1
	for i := len(nums) - 1; i > 0; i-- {
		tmp = tmp * nums[i]
		right[i] = tmp
	}

	res := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		tmp = 1

		if i-1 >= 0 {
			tmp = tmp * left[i-1]
		}
		if i+1 < len(nums) {
			tmp = tmp * right[i+1]
		}

		res[i] = tmp
	}
	return res
}
