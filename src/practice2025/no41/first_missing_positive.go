package no41

import "math"

func firstMissingPositive(nums []int) int {

	leftMap, rightMap := make(map[int]int), make(map[int]int)

	numSet := make(map[int]bool)

	for _, num := range nums {

		if num <= 0 {
			continue
		}

		if _, exist := numSet[num]; exist {
			continue
		}
		numSet[num] = true

		right, rightExist := leftMap[num+1]
		left, leftExist := rightMap[num-1]

		if rightExist && leftExist {

			delete(leftMap, num+1)
			delete(rightMap, right)
			delete(rightMap, num-1)
			delete(leftMap, left)

			leftMap[left] = right
			rightMap[right] = left
		} else if rightExist {
			delete(leftMap, num+1)
			delete(rightMap, right)

			leftMap[num] = right
			rightMap[right] = num

		} else if leftExist {
			delete(rightMap, num-1)
			delete(leftMap, left)

			leftMap[left] = num
			rightMap[num] = left

		} else {

			leftMap[num] = num
			rightMap[num] = num
		}
	}

	res := math.MaxInt
	fromOne := false
	for left, right := range leftMap {
		if left == 1 {
			fromOne = true
		}
		if left-1 > 0 {
			res = min(res, left-1)
		}
		if right+1 > 0 {
			res = min(res, right+1)
		}
	}
	if fromOne {
		return res
	} else {
		return 1
	}
}
