package no129

import (
	"fmt"
	"testing"
)

func Test_sumNumbers(t *testing.T) {
	node1 := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 3}
	node1.Left = node2
	node1.Right = node3

	res := sumNumbers(node1)

	fmt.Printf("%v", res)

}
