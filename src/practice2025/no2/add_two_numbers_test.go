package no2

import (
	"fmt"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	nodeA1 := &ListNode{Val: 2}
	nodeA2 := &ListNode{Val: 4}
	nodeA3 := &ListNode{Val: 3}

	nodeA1.Next = nodeA2
	nodeA2.Next = nodeA3

	nodeB1 := &ListNode{Val: 5}
	nodeB2 := &ListNode{Val: 6}
	nodeB3 := &ListNode{Val: 4}

	nodeB1.Next = nodeB2
	nodeB2.Next = nodeB3

	res := addTwoNumbers(nodeA1, nodeB1)

	fmt.Printf("%v", res)

}
