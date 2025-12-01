package no92

import "math"

func reverseBetween(head *ListNode, left int, right int) *ListNode {

	start := &ListNode{
		Val:  math.MinInt,
		Next: head,
	}
	cnt := 0
	cur := start
	for cur != nil {

		if cnt == left {
			leftNode = cur
		}
		if cnt == right {
			rightNode = cur
			afterRightNode = cur.Next
		}

		if cnt+1 == left {
			beforeLeftNode = cur
		}
		cnt++
		cur = cur.Next
	}

	doReverse(beforeLeftNode.Next)

	beforeLeftNode.Next = rightNode
	leftNode.Next = afterRightNode

	return start.Next
}

func doReverse(head *ListNode) *ListNode {

	if head == nil || head == rightNode {
		return head
	}

	newHead := doReverse(head.Next)
	head.Next.Next = head
	head.Next = nil

	return newHead
}

var (
	beforeLeftNode *ListNode
	leftNode       *ListNode
	rightNode      *ListNode
	afterRightNode *ListNode
)

type ListNode struct {
	Val  int
	Next *ListNode
}
