package no169

func majorityElement(nums []int) int {

	numCnt := make(map[int]int)

	for _, num := range nums {
		numCnt[num]++
	}

	for num, cnt := range numCnt {
		if cnt > len(nums)/2 {
			return num
		}
	}

	return 0
}
