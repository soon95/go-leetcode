package no61

import (
	"fmt"
	"testing"
)

func Test_rotateRight(t *testing.T) {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}
	node5 := &ListNode{Val: 5}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	res := rotateRight(node1, 0)

	fmt.Printf("%v", res)

}
