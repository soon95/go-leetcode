package no283

func moveZeroes(nums []int) {

	notZeroIndex, index := 0, 0

	for index < len(nums) {

		if nums[index] != 0 {
			nums[notZeroIndex] = nums[index]
			notZeroIndex++
		}
		index++
	}

	for notZeroIndex < len(nums) {
		nums[notZeroIndex] = 0
		notZeroIndex++
	}

	return
}
