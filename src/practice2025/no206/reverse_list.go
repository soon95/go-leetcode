package no206

func reverseList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	res := reverseList(head.Next)

	head.Next.Next = head
	head.Next = nil

	return res
}

type ListNode struct {
	Val  int
	Next *ListNode
}
