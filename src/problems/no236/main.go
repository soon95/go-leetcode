package main

import "go-leetcode/src/common"

type TreeNode = common.TreeNode

var ans *TreeNode

/*
*
236. 二叉树的最近公共祖先
https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/
*/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	ans = nil

	hasNode(root, p, q)

	return ans
}

func hasNode(root, p, q *TreeNode) bool {

	if root == nil {
		return false
	}

	var rootHas bool
	if root.Val == p.Val || root.Val == q.Val {
		rootHas = true
	} else {
		rootHas = false
	}

	leftHas := hasNode(root.Left, p, q)
	rightHas := hasNode(root.Right, p, q)

	if rootHas {

		if leftHas || rightHas {
			ans = root
		}

	} else {
		if leftHas && rightHas {
			ans = root
		}
	}

	return leftHas || rightHas || rootHas
}

func main() {

	node1 := &TreeNode{
		Val: 3,
	}
	node2 := &TreeNode{
		Val: 5,
	}
	node3 := &TreeNode{
		Val: 1,
	}
	node4 := &TreeNode{
		Val: 6,
	}
	node5 := &TreeNode{
		Val: 2,
	}
	node6 := &TreeNode{
		Val: 0,
	}
	node7 := &TreeNode{
		Val: 8,
	}
	node8 := &TreeNode{
		Val: 7,
	}
	node9 := &TreeNode{
		Val: 4,
	}

	node1.Left = node2
	node1.Right = node3
	node2.Left = node4
	node2.Right = node5
	node3.Left = node6
	node3.Right = node7
	node5.Left = node8
	node5.Right = node9

	//println(lowestCommonAncestor(node1, node2, node3))
	println(lowestCommonAncestor(node1, node2, node9))
}
