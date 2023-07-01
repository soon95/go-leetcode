package main

import (
	"fmt"
	"go-leetcode/src/common"
)

type TreeNode = common.TreeNode

/*
*114. 二叉树展开为链表
https://leetcode.cn/problems/flatten-binary-tree-to-linked-list/
*/
func flatten(root *TreeNode) {

	if root == nil {
		return
	}

	dfs(root)

}

func dfs(root *TreeNode) (head, tail *TreeNode) {

	left := root.Left
	right := root.Right

	root.Left = nil
	root.Right = nil
	curHead, curTail := root, root

	if left != nil {
		leftHead, leftTail := dfs(left)

		curHead.Right = leftHead

		curTail = leftTail

	}

	if right != nil {
		rightHead, rightTail := dfs(right)

		curTail.Right = rightHead

		curTail = rightTail

	}

	return curHead, curTail
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
	node6 := &TreeNode{
		Val: 6,
	}

	node1.Left = node2
	node2.Left = node3
	node2.Right = node4
	node1.Right = node5
	node5.Right = node6

	flatten(node6)

	fmt.Printf("%v", node1)

}
