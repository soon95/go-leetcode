package no128

func longestConsecutive(nums []int) int {

	numMap := make(map[int]bool)

	for _, num := range nums {
		numMap[num] = true
	}

	res := 0
	for _, num := range nums {

		_, exist := numMap[num]
		if !exist {
			continue
		}

		temp := 1
		preNum := num - 1
		for {
			_, exist = numMap[preNum]
			if !exist {
				break
			}
			temp++

			delete(numMap, preNum)

			preNum--
		}

		postNum := num + 1
		for {
			_, exist = numMap[postNum]
			if !exist {
				break
			}

			temp++

			delete(numMap, postNum)

			postNum++
		}

		delete(numMap, num)

		if res <= temp {
			res = temp
		}
	}

	return res
}
