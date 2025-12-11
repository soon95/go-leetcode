package no100

import (
	"fmt"
	"testing"
)

func Test_isSameTree(t *testing.T) {

	nodeA1 := &TreeNode{Val: 1}
	nodeA2 := &TreeNode{Val: 2}
	nodeA3 := &TreeNode{Val: 3}
	nodeA1.Left = nodeA2
	nodeA1.Right = nodeA3

	nodeB1 := &TreeNode{Val: 1}
	nodeB2 := &TreeNode{Val: 2}
	nodeB3 := &TreeNode{Val: 3}
	nodeB1.Left = nodeB2
	nodeB1.Right = nodeB3

	res := isSameTree(nodeA1, nodeB1)

	fmt.Printf("%v", res)

}
