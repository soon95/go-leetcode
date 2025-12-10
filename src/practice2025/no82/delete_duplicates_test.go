package no82

import (
	"fmt"
	"testing"
)

func Test_deleteDuplicates(t *testing.T) {

	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 3}
	node5 := &ListNode{Val: 4}
	node6 := &ListNode{Val: 4}
	node7 := &ListNode{Val: 5}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	node6.Next = node7

	res := deleteDuplicates(node1)

	fmt.Printf("%v", res)

}
