package no33

func search(nums []int, target int) int {

	start, end := 0, len(nums)-1

	for start <= end {

		mid := (start + end) / 2

		if nums[mid] == target {
			return mid
		} else {
			if nums[start] <= nums[end] {
				if nums[mid] > target {
					end = mid - 1
				} else {
					start = mid + 1
				}
			} else {

				if nums[start] <= nums[mid] {
					if nums[mid] > target && nums[start] <= target {
						end = mid - 1
					} else {
						start = mid + 1
					}
				} else {
					if nums[mid] < target && nums[end] >= target {
						start = mid + 1
					} else {
						end = mid - 1
					}
				}
			}
		}
	}

	return -1
}
