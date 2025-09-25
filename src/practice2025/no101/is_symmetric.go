package no101

func isSymmetric(root *TreeNode) bool {

	if root == nil {
		return true
	}

	return mirror(root.Left, root.Right)
}

func mirror(left, right *TreeNode) bool {

	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	if left.Val != right.Val {
		return false
	}

	return mirror(left.Left, right.Right) && mirror(left.Right, right.Left)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
