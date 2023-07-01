package main

import (
	"fmt"
	"go-leetcode/src/common"
)

type TreeNode = common.TreeNode

/*
*105. 从前序与中序遍历序列构造二叉树
https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
*/
func buildTree(preorder []int, inorder []int) *TreeNode {

	if len(preorder) == 0 {
		return nil
	}

	rootVal := preorder[0]
	root := &TreeNode{
		Val: rootVal,
	}

	index := -1
	for i, val := range inorder {
		if rootVal == val {
			index = i
			break
		}
	}

	left := buildTree(preorder[1:index+1], inorder[0:index])

	right := buildTree(preorder[index+1:], inorder[index+1:])

	root.Left = left
	root.Right = right

	return root
}

func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}

	root := buildTree(preorder, inorder)

	fmt.Printf("%v", root)

}
