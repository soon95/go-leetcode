package main

import (
	"fmt"
)

/*
85. 最大矩形
*https://leetcode.cn/problems/maximal-rectangle/submissions/

竟然需要复用 no.84 最大面积的思想
*/
func maximalRectangle(matrix [][]byte) int {

	length := len(matrix)
	if length == 0 {
		return 0
	}
	width := len(matrix[0])

	dp := make([][]int, length)

	for i := 0; i < length; i++ {
		dp[i] = make([]int, width)
	}

	for i := 0; i < length; i++ {
		if matrix[i][width-1] == '0' {
			dp[i][width-1] = 0
		} else {
			dp[i][width-1] = 1
		}
	}

	for i := 0; i < length; i++ {

		for j := width - 2; j >= 0; j-- {

			if matrix[i][j] == '0' {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i][j+1] + 1
			}
		}
	}

	ans := 0

	for i := 0; i < width; i++ {

		heights := []int{}
		for _, row := range dp {
			heights = append(heights, row[i])
		}

		temp := largestRectangleArea(heights)
		if temp > ans {
			ans = temp
		}

	}

	return ans
}

func largestRectangleArea(heights []int) int {

	length := len(heights)

	leftDp := make([]int, length)
	leftMinStack := []int{-1}
	for i := 0; i < length; i++ {

		for len(leftMinStack) != 1 && heights[leftMinStack[len(leftMinStack)-1]] >= heights[i] {
			leftMinStack = leftMinStack[0 : len(leftMinStack)-1]
		}

		leftDp[i] = leftMinStack[len(leftMinStack)-1]

		leftMinStack = append(leftMinStack, i)

	}

	rightDp := make([]int, length)
	rightMinStack := []int{length}
	rightDp[length-1] = length
	for i := length - 1; i >= 0; i-- {

		for len(rightMinStack) != 1 && heights[rightMinStack[len(rightMinStack)-1]] >= heights[i] {
			rightMinStack = rightMinStack[0 : len(rightMinStack)-1]
		}

		rightDp[i] = rightMinStack[len(rightMinStack)-1]

		rightMinStack = append(rightMinStack, i)
	}

	ans := 0

	for i := 0; i < length; i++ {

		temp := heights[i] * (rightDp[i] - leftDp[i] - 1)

		if ans < temp {
			ans = temp
		}
	}

	return ans
}

func main() {

	//matrix := [][]byte{
	//	{'1', '0', '1', '0', '0'},
	//	{'1', '0', '1', '1', '1'},
	//	{'1', '1', '1', '1', '1'},
	//	{'1', '0', '0', '1', '0'},
	//}
	matrix := [][]byte{{'0', '1'}}

	fmt.Printf("%v", maximalRectangle(matrix))

}
