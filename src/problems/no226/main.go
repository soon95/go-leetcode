package main

import "go-leetcode/src/common"

type TreeNode = common.TreeNode

/*
226. 翻转二叉树
*https://leetcode.cn/problems/invert-binary-tree/
*/
func invertTree(root *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left

	invertTree(root.Left)
	invertTree(root.Right)

	return root
}

func main() {

}
