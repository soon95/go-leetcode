package no138

import (
	"fmt"
	"testing"
)

func Test_copyRandomList(t *testing.T) {
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node4 := &Node{Val: 4}
	node5 := &Node{Val: 5}
	node1.Next = node2
	node1.Random = nil
	node2.Next = node3
	node2.Random = node1
	node3.Next = node4
	node3.Random = node5
	node4.Next = node5
	node4.Random = node3
	node5.Random = node1

	res := copyRandomList(node1)

	fmt.Printf("%v", res)

}
