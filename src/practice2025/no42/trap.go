package no42

func trap(height []int) int {

	leftHeight, rightHeight := make([]int, len(height)), make([]int, len(height))

	leftHeight[0] = height[0]
	for i := 1; i < len(height); i++ {
		if leftHeight[i-1] < height[i] {
			leftHeight[i] = height[i]
		} else {
			leftHeight[i] = leftHeight[i-1]
		}
	}
	rightHeight[len(height)-1] = height[len(height)-1]

	for i := len(height) - 2; i >= 0; i-- {
		if rightHeight[i+1] < height[i] {
			rightHeight[i] = height[i]
		} else {
			rightHeight[i] = rightHeight[i+1]
		}
	}

	res := 0

	for i := 0; i < len(height); i++ {

		tmp := leftHeight[i]
		if tmp > rightHeight[i] {
			tmp = rightHeight[i]
		}

		res = res + tmp - height[i]
	}

	return res
}
