package no437

func pathSum(root *TreeNode, targetSum int) int {

	newRoot := rebuildTree(root, 0)

	pathCnt = 0

	visit(newRoot, targetSum, nil)

	return pathCnt
}

var pathCnt int

func visit(root *TreeNode, targetSum int, pre []int) {

	if root == nil {
		return
	}
	curVal := root.Val
	if curVal == targetSum {
		pathCnt++
	}

	for _, preVal := range pre {
		if curVal-preVal == targetSum {
			pathCnt++
		}
	}

	cur := make([]int, 0)
	cur = append(cur, pre...)
	cur = append(cur, root.Val)

	visit(root.Left, targetSum, cur)
	visit(root.Right, targetSum, cur)
}

func rebuildTree(root *TreeNode, preSum int) *TreeNode {
	if root == nil {
		return nil
	}

	sum := root.Val + preSum
	newRoot := &TreeNode{
		Val: sum,
	}

	newRoot.Left = rebuildTree(root.Left, sum)
	newRoot.Right = rebuildTree(root.Right, sum)

	return newRoot
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
