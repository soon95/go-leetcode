package no75

func sortColors(nums []int) {

	redCnt := 0
	whiteCnt := 0

	for _, num := range nums {
		if num == 0 {
			redCnt++
		} else if num == 1 {
			whiteCnt++
		}
	}

	for i := 0; i < len(nums); i++ {
		if i < redCnt {
			nums[i] = 0
		} else if i < redCnt+whiteCnt {
			nums[i] = 1
		} else {
			nums[i] = 2
		}
	}

}
