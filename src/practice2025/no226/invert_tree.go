package no226

func invertTree(root *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)

	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
