package main

import (
	"fmt"
	"go-leetcode/src/common"
)

type TreeNode = common.TreeNode

/*
*101. 对称二叉树
https://leetcode.cn/problems/symmetric-tree/
*/
func isSymmetric(root *TreeNode) bool {

	if root == nil {
		return true
	}

	return dfs(root.Left, root.Right)
}

func dfs(left, right *TreeNode) bool {

	if left == nil && right == nil {
		return true
	}

	if left == nil && right != nil {
		return false
	}

	if left != nil && right == nil {
		return false
	}

	if left.Val != right.Val {
		return false
	}

	return dfs(left.Left, right.Right) && dfs(left.Right, right.Left)
}

func main() {

	node1 := &TreeNode{
		Val: 1,
	}
	node2 := &TreeNode{
		Val: 2,
	}
	node3 := &TreeNode{
		Val: 2,
	}
	node4 := &TreeNode{
		Val: 3,
	}
	node5 := &TreeNode{
		Val: 3,
	}

	node1.Left = node2
	node1.Right = node3
	node2.Right = node4
	node3.Right = node5

	//node1 := &TreeNode{
	//	Val: 1,
	//}
	//node2 := &TreeNode{
	//	Val: 2,
	//}
	//node3 := &TreeNode{
	//	Val: 2,
	//}
	//node4 := &TreeNode{
	//	Val: 3,
	//}
	//node5 := &TreeNode{
	//	Val: 4,
	//}
	//node6 := &TreeNode{
	//	Val: 4,
	//}
	//node7 := &TreeNode{
	//	Val: 3,
	//}
	//node1.Left = node2
	//node1.Right = node3
	//node2.Left = node4
	//node2.Right = node5
	//node3.Left = node6
	//node3.Right = node7

	fmt.Printf("%v", isSymmetric(node1))

}
