package no560

func subarraySum(nums []int, k int) int {

	res := 0

	for i := 0; i < len(nums); i++ {

		tmp := 0
		for j := i; j < len(nums); j++ {
			tmp = tmp + nums[j]
			if k == tmp {
				res++
			}
		}
	}

	return res
}
