package no637

func averageOfLevels(root *TreeNode) []float64 {

	curLevel := make([]*TreeNode, 0)

	curLevel = append(curLevel, root)

	res := make([]float64, 0)
	for len(curLevel) > 0 {

		nextLevel := make([]*TreeNode, 0)

		sum := 0

		for _, node := range curLevel {
			sum += node.Val
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}

		res = append(res, float64(sum)/float64(len(curLevel)))
		curLevel = nextLevel
	}

	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
