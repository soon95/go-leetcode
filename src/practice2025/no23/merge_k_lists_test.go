package no23

import (
	"fmt"
	"testing"
)

func Test_mergeKLists(t *testing.T) {
	nodeA1 := &ListNode{Val: 1}
	nodeA2 := &ListNode{Val: 4}
	nodeA3 := &ListNode{Val: 5}

	nodeA1.Next = nodeA2
	nodeA2.Next = nodeA3

	nodeB1 := &ListNode{Val: 1}
	nodeB2 := &ListNode{Val: 3}
	nodeB3 := &ListNode{Val: 4}

	nodeB1.Next = nodeB2
	nodeB2.Next = nodeB3

	nodeC1 := &ListNode{Val: 2}
	nodeC2 := &ListNode{Val: 6}

	nodeC1.Next = nodeC2

	res := mergeKLists([]*ListNode{nodeA1, nodeB1, nodeC1})

	fmt.Printf("%v", res)

}
