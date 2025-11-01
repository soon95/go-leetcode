package no80

func removeDuplicates(nums []int) int {

	if len(nums) < 2 {
		return len(nums)
	}

	index := 2
	pre := nums[0]
	for i := 2; i < len(nums); i++ {
		if nums[i] != pre {
			pre = nums[i-1]
			nums[index] = nums[i]
			index++
		} else {
			pre = nums[i-1]
		}
	}

	return index

}
