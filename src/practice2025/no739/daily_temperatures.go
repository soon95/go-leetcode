package no739

func dailyTemperatures(temperatures []int) []int {

	ans := make([]int, len(temperatures))
	stack := make([]int, 0)

	index := 0

	for index < len(temperatures) {

		if len(stack) == 0 {
			stack = append(stack, index)
			index++
		} else {
			top := stack[len(stack)-1]
			if temperatures[top] < temperatures[index] {
				ans[top] = index - top

				stack = stack[:len(stack)-1]

				continue
			} else {
				stack = append(stack, index)
				index++
			}
		}
	}

	return ans
}
