package main

import (
	"fmt"
	"go-leetcode/src/common"
)

type TreeNode = common.TreeNode

/*
*538. 把二叉搜索树转换为累加树
https://leetcode.cn/problems/convert-bst-to-greater-tree/
*/
func convertBST(root *TreeNode) *TreeNode {

	dfs(root, 0)

	return root
}

func dfs(root *TreeNode, sum int) int {

	if root == nil {
		return sum
	}

	sum = dfs(root.Right, sum)

	sum += root.Val

	root.Val = sum

	sum = dfs(root.Left, sum)

	return sum
}

func main() {

	node1 := &TreeNode{
		Val: 4,
	}
	node2 := &TreeNode{
		Val: 1,
	}
	node3 := &TreeNode{
		Val: 6,
	}
	node4 := &TreeNode{
		Val: 0,
	}
	node5 := &TreeNode{
		Val: 2,
	}
	node6 := &TreeNode{
		Val: 5,
	}
	node7 := &TreeNode{
		Val: 7,
	}
	node8 := &TreeNode{
		Val: 3,
	}
	node9 := &TreeNode{
		Val: 8,
	}

	node1.Left = node2
	node1.Right = node3
	node2.Left = node4
	node2.Right = node5
	node3.Left = node6
	node3.Right = node7
	node5.Right = node8
	node7.Right = node9

	ans := convertBST(node1)

	fmt.Printf("%v", ans)

}
