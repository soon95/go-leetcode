package no98

import "math"

func isValidBST(root *TreeNode) bool {

	return dfs(root, math.MinInt, math.MaxInt)
}

func dfs(root *TreeNode, from, to int) bool {

	if root == nil {
		return true
	}

	if root.Left != nil && (root.Left.Val >= root.Val || root.Left.Val <= from) {
		return false
	}

	if root.Right != nil && (root.Right.Val <= root.Val || root.Right.Val >= to) {
		return false
	}

	return dfs(root.Left, from, root.Val) && dfs(root.Right, root.Val, to)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
