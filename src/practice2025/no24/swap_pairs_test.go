package no24

import (
	"fmt"
	"testing"
)

func Test_swapPairs(t *testing.T) {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4

	res := swapPairs(node1)

	fmt.Printf("%v", res)

}
