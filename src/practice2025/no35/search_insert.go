package no35

func searchInsert(nums []int, target int) int {

	start, end := 0, len(nums)

	for start < end {
		mid := (start + end) / 2

		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			end = mid
		} else {
			start = mid + 1
		}
	}

	return start
}
