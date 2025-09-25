package no104

func maxDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}

	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
