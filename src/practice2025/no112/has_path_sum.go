package no112

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	return visit(0, root, targetSum)
}

func visit(preSum int, cur *TreeNode, targetSum int) bool {

	if cur.Left == nil && cur.Right == nil {
		return preSum+cur.Val == targetSum
	} else if cur.Left == nil {

		return visit(preSum+cur.Val, cur.Right, targetSum)

	} else if cur.Right == nil {
		return visit(preSum+cur.Val, cur.Left, targetSum)
	} else {
		return visit(preSum+cur.Val, cur.Right, targetSum) || visit(preSum+cur.Val, cur.Left, targetSum)
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
