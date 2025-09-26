package no114

func flatten(root *TreeNode) {
	doFlatten(root)
}

func doFlatten(root *TreeNode) (*TreeNode, *TreeNode) {

	if root == nil {
		return nil, nil
	}

	lHead, lTail := doFlatten(root.Left)
	rHead, rTail := doFlatten(root.Right)

	tail := root
	root.Left = nil
	if lHead != nil && lTail != nil {
		root.Right = lHead

		tail = lTail
	}

	if rHead != nil && rTail != nil {

		tail.Right = rHead

		tail = rTail
	}

	return root, tail
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
