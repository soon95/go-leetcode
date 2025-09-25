package no543

func diameterOfBinaryTree(root *TreeNode) int {

	diameter = 0
	_ = depth(root)
	return diameter
}

var diameter int

func depth(root *TreeNode) int {

	if root == nil {
		return 0
	}

	depthL := depth(root.Left)
	depthR := depth(root.Right)

	diameter = max(diameter, depthL+depthR)

	return max(depthL, depthR) + 1
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
