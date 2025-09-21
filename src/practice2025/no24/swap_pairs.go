package no24

func swapPairs(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	tmp := swapPairs(head.Next.Next)

	newHead := head.Next

	newHead.Next = head
	head.Next = tmp

	return newHead
}

type ListNode struct {
	Val  int
	Next *ListNode
}
