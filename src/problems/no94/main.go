package main

import (
	"fmt"
	"go-leetcode/src/common"
)

type TreeNode = common.TreeNode

var visit []int

/*
*94. 二叉树的中序遍历
https://leetcode.cn/problems/binary-tree-inorder-traversal/
*/
func inorderTraversal(root *TreeNode) []int {

	visit = []int{}

	middleVisit(root)

	return visit
}

func middleVisit(node *TreeNode) {

	if node == nil {
		return
	}

	middleVisit(node.Left)
	visit = append(visit, node.Val)
	middleVisit(node.Right)

}

func main() {

	root := &TreeNode{
		Val: 1,
	}
	node2 := &TreeNode{
		Val: 2,
	}
	node3 := &TreeNode{
		Val: 3,
	}
	root.Right = node2
	node2.Left = node3

	traversal := inorderTraversal(root)

	fmt.Printf("%v", traversal)

}
