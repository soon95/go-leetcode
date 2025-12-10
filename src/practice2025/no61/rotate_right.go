package no61

func rotateRight(head *ListNode, k int) *ListNode {

	if head == nil {
		return head
	}

	length := 0

	cur := head
	var tail *ListNode
	for cur != nil {
		length++
		tail = cur
		cur = cur.Next
	}

	move := length - k%length

	preLength := 0
	cur = head
	var preTail *ListNode
	for preLength < move {
		preLength++
		preTail = cur
		cur = cur.Next
	}
	postHead := cur

	if postHead == nil {
		return head
	}

	preTail.Next = nil

	tail.Next = head

	return postHead
}

type ListNode struct {
	Val  int
	Next *ListNode
}
