package no117

import (
	"fmt"
	"testing"
)

func Test_connect(t *testing.T) {
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node4 := &Node{Val: 4}
	node5 := &Node{Val: 5}
	node6 := &Node{Val: 7}

	node1.Left = node2
	node1.Right = node3
	node2.Left = node4
	node2.Right = node5
	node3.Right = node6

	res := connect(node1)

	fmt.Printf("%v", res)

}
