package no103

func zigzagLevelOrder(root *TreeNode) [][]int {

	if root == nil {
		return nil
	}

	fromLeft := true

	curList := make([]*TreeNode, 0)
	curList = append(curList, root)

	res := make([][]int, 0)

	for len(curList) > 0 {

		nextList := make([]*TreeNode, 0)

		tmp := make([]int, 0)

		for _, node := range curList {
			if node.Left != nil {
				nextList = append(nextList, node.Left)
			}
			if node.Right != nil {
				nextList = append(nextList, node.Right)
			}
		}

		if fromLeft {
			for i := 0; i < len(curList); i++ {
				tmp = append(tmp, curList[i].Val)
			}
		} else {
			for i := len(curList) - 1; i >= 0; i-- {
				tmp = append(tmp, curList[i].Val)
			}
		}

		curList = nextList
		res = append(res, tmp)
		fromLeft = !fromLeft
	}

	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
