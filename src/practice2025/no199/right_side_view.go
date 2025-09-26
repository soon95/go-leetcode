package no199

func rightSideView(root *TreeNode) []int {

	pre, tmp := make([]*TreeNode, 0), make([]*TreeNode, 0)

	layerList := make([][]int, 0)
	if root == nil {
		return nil
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
		layerList = append(layerList, curRes)

		pre = tmp
	}

	res := make([]int, 0)
	for _, items := range layerList {
		res = append(res, items[len(items)-1])
	}

	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
