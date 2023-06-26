package main

import "fmt"

/*
84. 柱状图中最大的矩形
*https://leetcode.cn/problems/largest-rectangle-in-histogram/

从左到右 从右到左，各找出第一个比当前元素小的元素
用单调栈
*/
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

	//heights := []int{2, 1, 5, 6, 2, 3}
	//heights := []int{2, 4}
	heights := []int{1, 1}

	fmt.Printf("%v", largestRectangleArea(heights))

}
