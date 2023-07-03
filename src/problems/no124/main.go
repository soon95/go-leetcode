package main

import (
	"fmt"
	"go-leetcode/src/common"
	"math"
)

type TreeNode = common.TreeNode

var maxPath int

/*
*124. 二叉树中的最大路径和
https://leetcode.cn/problems/binary-tree-maximum-path-sum/
*/
func maxPathSum(root *TreeNode) int {

	maxPath = math.MinInt

	dfs(root)

	return maxPath
}

func dfs(root *TreeNode) int {

	if root == nil {
		return math.MinInt
	}

	leftPath := dfs(root.Left)
	rightPath := dfs(root.Right)

	if leftPath < 0 {
		leftPath = 0
	}
	if rightPath < 0 {
		rightPath = 0
	}

	temp := root.Val + leftPath + rightPath

	if maxPath < temp {
		maxPath = temp
	}

	if leftPath > rightPath {
		return root.Val + leftPath
	} else {
		return root.Val + rightPath
	}
}

func main() {

	node1 := &TreeNode{
		Val: -10,
	}
	node2 := &TreeNode{
		Val: 9,
	}
	node3 := &TreeNode{
		Val: 20,
	}
	node4 := &TreeNode{
		Val: 15,
	}
	node5 := &TreeNode{
		Val: 7,
	}
	node1.Left = node2
	node1.Right = node3
	node3.Left = node4
	node3.Right = node5

	fmt.Printf("%v", maxPathSum(node1))

}
