package no173

import (
	"fmt"
	"testing"
)

func TestBSTIterator(t *testing.T) {

	node1 := &TreeNode{Val: 7}
	node2 := &TreeNode{Val: 3}
	node3 := &TreeNode{Val: 15}
	node4 := &TreeNode{Val: 9}
	node5 := &TreeNode{Val: 20}
	node1.Left = node2
	node1.Right = node3
	node3.Left = node4
	node3.Right = node5

	constructor := Constructor(node1)

	next := constructor.Next()

	fmt.Printf("%v", next)

}
