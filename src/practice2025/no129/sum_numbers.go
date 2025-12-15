package no129

func sumNumbers(root *TreeNode) int {

	sum = 0

	visit(root, 0)

	return sum
}

var sum int

func visit(root *TreeNode, preSum int) {

	if root == nil {
		return
	}

	curSum := preSum*10 + root.Val

	if root.Left == nil && root.Right == nil {
		sum += curSum
		return
	}

	visit(root.Left, curSum)
	visit(root.Right, curSum)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
