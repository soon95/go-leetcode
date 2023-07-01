package main

import (
	"fmt"
	"go-leetcode/src/common"
)

type TreeNode = common.TreeNode

var ans [][]int

/*
*102. 二叉树的层序遍历
https://leetcode.cn/problems/binary-tree-level-order-traversal/
*/
func levelOrder(root *TreeNode) [][]int {

	ans = [][]int{}

	if root == nil {
		return ans
	}

	nodes := []*TreeNode{
		root,
	}

	dfs(nodes)

	return ans
}

func dfs(nodes []*TreeNode) {

	if len(nodes) == 0 {
		return
	}

	temp := []int{}

	var children []*TreeNode

	for _, node := range nodes {

		temp = append(temp, node.Val)

		if node.Left != nil {
			children = append(children, node.Left)
		}

		if node.Right != nil {
			children = append(children, node.Right)
		}
	}

	ans = append(ans, temp)

	dfs(children)
}

func main() {

	node1 := &TreeNode{
		Val: 3,
	}
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

	fmt.Printf("%v", levelOrder(node1))

}
