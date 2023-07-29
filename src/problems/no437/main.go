package main

import (
	"fmt"
	"go-leetcode/src/common"
)

type TreeNode = common.TreeNode

var ans int

/*
*437. 路径总和 III
https://leetcode.cn/problems/path-sum-iii/
*/
func pathSum(root *TreeNode, targetSum int) int {

	ans = 0

	doPathSum(root, targetSum)

	return ans
}

func doPathSum(root *TreeNode, targetSum int) {

	if root == nil {
		return
	}

	dfs(root, 0, targetSum)

	doPathSum(root.Left, targetSum)
	doPathSum(root.Right, targetSum)

}

func dfs(root *TreeNode, preSum, targetSum int) {

	if root == nil {
		return
	}

	curSum := root.Val + preSum

	if curSum == targetSum {
		ans++
	}

	dfs(root.Left, curSum, targetSum)

	dfs(root.Right, curSum, targetSum)
}

func main() {

	node1 := &TreeNode{
		Val: 10,
	}
	node2 := &TreeNode{
		Val: 5,
	}
	node3 := &TreeNode{
		Val: -3,
	}
	node4 := &TreeNode{
		Val: 3,
	}
	node5 := &TreeNode{
		Val: 2,
	}
	node6 := &TreeNode{
		Val: 11,
	}
	node7 := &TreeNode{
		Val: 3,
	}
	node8 := &TreeNode{
		Val: -2,
	}
	node9 := &TreeNode{
		Val: 1,
	}
	node1.Left = node2
	node1.Right = node3
	node2.Left = node4
	node2.Right = node5
	node3.Right = node6
	node4.Left = node7
	node4.Right = node8
	node5.Right = node9

	target := 8

	fmt.Printf("%v", pathSum(node1, target))

}
