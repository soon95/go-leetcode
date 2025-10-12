package no84

func largestRectangleArea(heights []int) int {

	stack := make([]int, 0)

	index := 0
	maxArea := 0

	for index < len(heights) {

		// 首先处理栈
		for len(stack) != 0 && heights[stack[len(stack)-1]] >= heights[index] {
			stack = stack[:len(stack)-1]
		}

		// 单独一根
		maxArea = max(maxArea, heights[index])

		if len(stack) == 0 {
			maxArea = max(maxArea, heights[index]*(index+1))
		} else {
			maxArea = max(maxArea, heights[index]*(index-stack[len(stack)-1]))

			for i, item := range stack {
				if i == 0 {
					maxArea = max(maxArea, heights[item]*(index+1))
				} else {
					maxArea = max(maxArea, heights[item]*(index-stack[i-1]))
				}

			}
		}

		stack = append(stack, index)

		index++
	}

	return maxArea
}
