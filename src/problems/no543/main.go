package main

import (
	"fmt"
	"go-leetcode/src/common"
)

type TreeNode = common.TreeNode

var diameter int

/*
*543. 二叉树的直径
https://leetcode.cn/problems/diameter-of-binary-tree/?envType=featured-list&envId=2cktkvj
*/
func diameterOfBinaryTree(root *TreeNode) int {

	diameter = 0

	dfs(root)

	return diameter
}

func dfs(root *TreeNode) int {

	if root == nil {
		return 0
	}

	left := dfs(root.Left)
	right := dfs(root.Right)

	temp := left + right

	if diameter < temp {
		diameter = temp
	}

	if left > right {
		return left + 1
	} else {
		return right + 1
	}
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
	node4 := &TreeNode{
		Val: 4,
	}
	node5 := &TreeNode{
		Val: 5,
	}

	node1.Left = node2
	node1.Right = node3
	node2.Left = node4
	node2.Right = node5

	fmt.Printf("%v", diameterOfBinaryTree(node1))
}
