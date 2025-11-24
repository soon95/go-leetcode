package no219

func containsNearbyDuplicate(nums []int, k int) bool {

	tmp := make(map[int]int)

	for i, num := range nums {

		if pre, exist := tmp[num]; exist && i-pre <= k {
			return true
		}

		tmp[num] = i

	}
	return false
}
