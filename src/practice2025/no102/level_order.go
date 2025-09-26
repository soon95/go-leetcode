package no102

func levelOrder(root *TreeNode) [][]int {

	pre, tmp := make([]*TreeNode, 0), make([]*TreeNode, 0)

	res := make([][]int, 0)
	if root == nil {
		return res
	}

	pre = append(pre, root)

	for len(pre) != 0 {

		tmp = make([]*TreeNode, 0)

		curRes := make([]int, 0)

		for _, node := range pre {

			curRes = append(curRes, node.Val)

			if node.Left != nil {
				tmp = append(tmp, node.Left)
			}
			if node.Right != nil {
				tmp = append(tmp, node.Right)
			}
		}
		res = append(res, curRes)

		pre = tmp
	}

	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
