package main

import (
	"fmt"
	"go-leetcode/src/common"
)

type TreeNode = common.TreeNode

/*
*617. 合并二叉树
https://leetcode.cn/problems/merge-two-binary-trees/
*/
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {

	if root1 == nil {
		return root2
	}

	if root2 == nil {
		return root1
	}

	root := &TreeNode{
		Val: root1.Val + root2.Val,
	}

	root.Left = mergeTrees(root1.Left, root2.Left)
	root.Right = mergeTrees(root1.Right, root2.Right)

	return root
}

func main() {

	root1 := &TreeNode{
		Val: 1,
	}
	node2 := &TreeNode{
		Val: 3,
	}
	node3 := &TreeNode{
		Val: 2,
	}
	node4 := &TreeNode{
		Val: 5,
	}
	root1.Left = node2
	root1.Right = node3
	node2.Left = node4

	root2 := &TreeNode{
		Val: 2,
	}
	node5 := &TreeNode{
		Val: 1,
	}
	node6 := &TreeNode{
		Val: 3,
	}
	node7 := &TreeNode{
		Val: 4,
	}
	node8 := &TreeNode{
		Val: 7,
	}
	root2.Left = node5
	root2.Right = node6
	node5.Right = node7
	node6.Right = node8

	root := mergeTrees(root1, root2)

	fmt.Printf("%v", root)

}
