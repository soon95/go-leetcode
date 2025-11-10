package no167

func twoSum(numbers []int, target int) []int {

	left, right := 0, len(numbers)-1

	for left < right {

		if numbers[left]+numbers[right] > target {
			right--
		} else if numbers[left]+numbers[right] < target {
			left++
		} else {
			break
		}
	}

	return []int{left + 1, right + 1}
}
