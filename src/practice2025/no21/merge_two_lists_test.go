package no21

import (
	"fmt"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	nodeA1 := &ListNode{Val: 1}
	nodeA2 := &ListNode{Val: 2}
	nodeA3 := &ListNode{Val: 4}

	nodeA1.Next = nodeA2
	nodeA2.Next = nodeA3

	nodeB1 := &ListNode{Val: 1}
	nodeB2 := &ListNode{Val: 3}
	nodeB3 := &ListNode{Val: 4}

	nodeB1.Next = nodeB2
	nodeB2.Next = nodeB3

	res := mergeTwoLists(nodeA1, nodeB1)

	fmt.Printf("%v", res)

}
