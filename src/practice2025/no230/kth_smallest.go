package no230

func kthSmallest(root *TreeNode, k int) int {

	visited = make([]int, 0)

	_, res := dfs(root, k)

	return res
}

var visited []int

func dfs(root *TreeNode, k int) (bool, int) {

	if root == nil {
		return false, 0
	}

	find, res := dfs(root.Left, k)
	if find {
		return true, res
	}
	visited = append(visited, root.Val)

	if len(visited) == k {
		return true, visited[len(visited)-1]
	}

	return dfs(root.Right, k)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
