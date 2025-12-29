package no530

import (
	"fmt"
	"testing"
)

func Test_getMinimumDifference(t *testing.T) {
	node1 := &TreeNode{Val: 4}
	node2 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 6}
	node4 := &TreeNode{Val: 1}
	node5 := &TreeNode{Val: 3}

	node1.Left = node2
	node1.Right = node3
	node2.Left = node4
	node2.Right = node5

	res := getMinimumDifference(node1)

	fmt.Println(res)

}
