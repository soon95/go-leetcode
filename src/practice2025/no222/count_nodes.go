package no222

func countNodes(root *TreeNode) int {

	cnt = 0

	visit(root)

	return cnt
}

func visit(root *TreeNode) {

	if root == nil {
		return
	}

	cnt++

	visit(root.Left)
	visit(root.Right)
}

var cnt int

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
