package no105

func buildTree(preorder []int, inorder []int) *TreeNode {

	if len(preorder) == 0 {
		return nil
	}

	rootVal := preorder[0]

	findIndex := 0

	for index, num := range inorder {
		if num == rootVal {
			findIndex = index
			break
		}
	}

	return &TreeNode{
		Val:   rootVal,
		Left:  buildTree(preorder[1:findIndex+1], inorder[:findIndex]),
		Right: buildTree(preorder[findIndex+1:], inorder[findIndex+1:]),
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
