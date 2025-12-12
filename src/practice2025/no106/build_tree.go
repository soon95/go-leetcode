package no106

func buildTree(inorder []int, postorder []int) *TreeNode {

	if len(postorder) == 0 {
		return nil
	}

	cur := postorder[len(postorder)-1]

	index := 0
	for i, val := range inorder {
		if val == cur {
			index = i
			break
		}
	}

	return &TreeNode{
		Val:   cur,
		Left:  buildTree(inorder[:index], postorder[:index]),
		Right: buildTree(inorder[index+1:], postorder[index:len(postorder)-1]),
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
