package no34

func searchRange(nums []int, target int) []int {

	start, end := 0, len(nums)-1

	targetLeftIndex, targetRightIndex := -1, -1

	for start <= end {
		mid := (start + end) / 2

		if nums[mid] == target {
			targetLeftIndex = mid
			end = mid - 1
		} else if nums[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	start, end = 0, len(nums)-1
	for start <= end {
		mid := (start + end) / 2
		if nums[mid] == target {
			targetRightIndex = mid
			start = mid + 1
		} else if nums[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return []int{targetLeftIndex, targetRightIndex}
}
