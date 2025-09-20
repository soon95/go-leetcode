package no234

import (
	"fmt"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 2}
	node4 := &ListNode{Val: 1}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4

	res := isPalindromeV2(node1)

	fmt.Printf("%v", res)

}
