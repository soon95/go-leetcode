package main

import (
	"fmt"
	"go-leetcode/src/common"
)

type TreeNode = common.TreeNode

/*
*104. 二叉树的最大深度
https://leetcode.cn/problems/maximum-depth-of-binary-tree/
*/
func maxDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}

	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	curDepth := 0

	if leftDepth > rightDepth {
		curDepth = leftDepth + 1
	} else {
		curDepth = rightDepth + 1
	}

	return curDepth
}

func main() {

	//node1 := &TreeNode{
	//	Val: 3,
	//}
	//node2 := &TreeNode{
	//	Val: 9,
	//}
	//node3 := &TreeNode{
	//	Val: 20,
	//}
	//node4 := &TreeNode{
	//	Val: 15,
	//}
	//node5 := &TreeNode{
	//	Val: 7,
	//}
	//
	//node1.Left = node2
	//node1.Right = node3
	//node3.Left = node4
	//node3.Right = node5

	fmt.Printf("%v", maxDepth(nil))
}
