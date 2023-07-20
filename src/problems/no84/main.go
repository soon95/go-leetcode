package main

import "fmt"

/*
84. 柱状图中最大的矩形
*https://leetcode.cn/problems/largest-rectangle-in-histogram/

从左到右 从右到左，各找出第一个比当前元素小的元素
用单调栈
*/
//func largestRectangleArea(heights []int) int {
//
//	length := len(heights)
//
//	leftDp := make([]int, length)
//	leftMinStack := []int{-1}
//	for i := 0; i < length; i++ {
//
//		for len(leftMinStack) != 1 && heights[leftMinStack[len(leftMinStack)-1]] >= heights[i] {
//			leftMinStack = leftMinStack[0 : len(leftMinStack)-1]
//		}
//
//		leftDp[i] = leftMinStack[len(leftMinStack)-1]
//
//		leftMinStack = append(leftMinStack, i)
//
//	}
//
//	rightDp := make([]int, length)
//	rightMinStack := []int{length}
//	rightDp[length-1] = length
//	for i := length - 1; i >= 0; i-- {
//
//		for len(rightMinStack) != 1 && heights[rightMinStack[len(rightMinStack)-1]] >= heights[i] {
//			rightMinStack = rightMinStack[0 : len(rightMinStack)-1]
//		}
//
//		rightDp[i] = rightMinStack[len(rightMinStack)-1]
//
//		rightMinStack = append(rightMinStack, i)
//	}
//
//	ans := 0
//
//	for i := 0; i < length; i++ {
//
//		temp := heights[i] * (rightDp[i] - leftDp[i] - 1)
//
//		if ans < temp {
//			ans = temp
//		}
//	}
//
//	return ans
//}

/*
*
使用单调栈
*/
func largestRectangleAreaV2(heights []int) int {

	ans := 0

	nums := make([]int, len(heights)+2)

	for i := 0; i < len(heights); i++ {
		nums[i+1] = heights[i]
	}

	length := len(nums)

	// 定义一个单调栈
	indexStack := []int{}

	for i := 0; i < length; i++ {

		for len(indexStack) != 0 && nums[indexStack[len(indexStack)-1]] > nums[i] {

			// 当前栈顶的元素 以当前栈顶元素构成矩形
			height := nums[indexStack[len(indexStack)-1]]

			indexStack = indexStack[0 : len(indexStack)-1]

			leftIndex := indexStack[len(indexStack)-1]
			rightIndex := i

			width := rightIndex - leftIndex - 1

			temp := height * width

			ans = max(ans, temp)
		}

		indexStack = append(indexStack, i)
	}
	return ans
}

func max(x, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}

func main() {

	//heights := []int{2, 1, 5, 100, 2, 3}
	//heights := []int{2, 4}
	//heights := []int{1, 1}
	heights := []int{3, 1, 3, 2, 2}

	fmt.Printf("%v", largestRectangleAreaV2(heights))

}
