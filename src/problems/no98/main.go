package main

import (
	"fmt"
	"go-leetcode/src/common"
	"math"
)

type TreeNode = common.TreeNode

/*
*98. 验证二叉搜索树
https://leetcode.cn/problems/validate-binary-search-tree/
*/
func isValidBST(root *TreeNode) bool {

	return dfs(root, math.MinInt, math.MaxInt)
}

func dfs(root *TreeNode, min, max int) bool {

	if root == nil {
		return true
	}

	if root.Val <= min || root.Val >= max {
		return false
	}

	if root.Left != nil && root.Left.Val >= root.Val {
		return false
	}

	if root.Right != nil && root.Right.Val <= root.Val {
		return false
	}

	return dfs(root.Left, min, root.Val) && dfs(root.Right, root.Val, max)

}

func main() {

	node1 := &TreeNode{
		Val: 1,
	}
	node2 := &TreeNode{
		Val: 2,
	}
	node3 := &TreeNode{
		Val: 3,
	}
	node2.Left = node1
	node2.Right = node3

	//node1 := &TreeNode{
	//	Val: 5,
	//}
	//node2 := &TreeNode{
	//	Val: 1,
	//}
	//node3 := &TreeNode{
	//	Val: 7,
	//}
	//node4 := &TreeNode{
	//	Val: 3,
	//}
	//node5 := &TreeNode{
	//	Val: 8,
	//}
	//node1.Left = node2
	//node1.Right = node3
	//node3.Left = node4
	//node3.Right = node5

	fmt.Printf("%v", isValidBST(node1))

}
