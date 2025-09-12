package no11

func maxArea(height []int) int {

	startStack, endStack := make([]int, 0), make([]int, 0)

	startStack = append(startStack, 0)
	endStack = append(endStack, len(height)-1)

	for i := 1; i < len(height); i++ {
		if height[i] > height[startStack[len(startStack)-1]] {
			startStack = append(startStack, i)
		}
	}
	for i := len(height) - 2; i >= 0; i-- {
		if height[i] > height[endStack[len(endStack)-1]] {
			endStack = append(endStack, i)
		}
	}

	start, end := 0, 0
	res := 0

	for start < len(startStack) && end < len(endStack) {

		if height[startStack[start]] < height[endStack[end]] {

			tmp := height[startStack[start]] * (endStack[end] - startStack[start])

			if res < tmp {
				res = tmp
			}

			start++
		} else {
			tmp := height[endStack[end]] * (endStack[end] - startStack[start])

			if res < tmp {
				res = tmp
			}

			end++
		}
	}

	return res
}
