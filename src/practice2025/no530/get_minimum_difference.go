package no530

import "math"

func getMinimumDifference(root *TreeNode) int {

	minRes = math.MaxInt

	findMaxAndMin(root)

	return minRes
}

var minRes int

func findMaxAndMin(root *TreeNode) (int, int) {

	curNum, minNum, maxNum := root.Val, root.Val, root.Val

	if root.Left != nil {
		leftMin, leftMax := findMaxAndMin(root.Left)
		minRes = min(minRes, curNum-leftMax)
		minNum = leftMin
	}
	if root.Right != nil {
		rightMin, rightMax := findMaxAndMin(root.Right)
		minRes = min(minRes, rightMin-curNum)
		maxNum = rightMax
	}

	return minNum, maxNum
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
