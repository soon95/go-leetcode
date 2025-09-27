package no236

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	_, pPath := findPath(root, p, nil)
	_, qPath := findPath(root, q, nil)

	pPathMap := make(map[*TreeNode]bool)
	for _, node := range pPath {
		pPathMap[node] = true
	}

	for i := len(qPath) - 1; i >= 0; i-- {
		_, exist := pPathMap[qPath[i]]

		if exist {
			return qPath[i]
		}

	}

	return nil
}

func findPath(root, target *TreeNode, prePath []*TreeNode) (bool, []*TreeNode) {

	if root == nil {
		return false, nil
	}

	path := make([]*TreeNode, 0)
	path = append(prePath, root)

	if root == target {
		return true, path
	}

	lFind, lPath := findPath(root.Left, target, path)
	if lFind {
		return true, lPath
	}

	rFind, rPath := findPath(root.Right, target, path)
	if rFind {
		return true, rPath
	}

	return false, nil
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
