package no1

func twoSum(nums []int, target int) []int {
	res := make([]int, 0)

	numMap := make(map[int]int)

	for i, num := range nums {

		i2, exist := numMap[target-num]
		if exist {
			res = append(res, i2, i)
			break
		} else {
			numMap[num] = i
		}
	}

	return res
}
