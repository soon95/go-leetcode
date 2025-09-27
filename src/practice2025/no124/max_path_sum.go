package no124

import "math"

func maxPathSum(root *TreeNode) int {

	maxSum = math.MinInt

	doFind(root)

	return maxSum
}

var maxSum int

func doFind(root *TreeNode) int {

	if root == nil {
		return 0
	}

	lSum := doFind(root.Left)
	rSum := doFind(root.Right)

	curSum := root.Val
	if lSum > 0 {
		curSum += lSum
	}
	if rSum > 0 {
		curSum += rSum
	}

	maxSum = max(maxSum, curSum)

	preSum := max(lSum, rSum)

	if preSum < 0 {
		return root.Val
	} else {
		return root.Val + preSum
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
