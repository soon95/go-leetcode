package no19

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	cur := head
	length := 0
	for cur != nil {
		length++
		cur = cur.Next
	}

	if length == n {
		return head.Next
	}

	cur = head
	for i := 1; i < length-n; i++ {
		cur = cur.Next
	}

	if cur.Next != nil {
		cur.Next = cur.Next.Next
	}

	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}
