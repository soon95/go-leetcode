package no25

func reverseKGroup(head *ListNode, k int) *ListNode {

	length := 0
	cur := head
	for cur != nil {
		length++
		cur = cur.Next
	}

	if length < k {
		return head
	}

	cur = head
	for i := 1; i < k; i++ {
		cur = cur.Next
	}
	tail := cur

	tmp := reverseKGroup(tail.Next, k)

	newHead := reverse(head, tail)

	head.Next = tmp

	return newHead
}

func reverse(head, tail *ListNode) *ListNode {

	if head == tail {
		return head
	}

	newHead := reverse(head.Next, tail)

	head.Next.Next = head
	head.Next = nil

	return newHead
}

type ListNode struct {
	Val  int
	Next *ListNode
}
