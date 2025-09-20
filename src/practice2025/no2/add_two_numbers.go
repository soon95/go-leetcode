package no2

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return doAdd(l1, l2, 0)
}

func doAdd(l1 *ListNode, l2 *ListNode, pre int) *ListNode {

	if pre == 0 && l1 == nil && l2 == nil {
		return nil
	}

	sum := pre

	if l1 != nil {
		sum = sum + l1.Val
		l1 = l1.Next
	}

	if l2 != nil {
		sum = sum + l2.Val
		l2 = l2.Next
	}

	val := sum % 10
	pre = sum / 10

	cur := &ListNode{
		Val: val,
	}

	cur.Next = doAdd(l1, l2, pre)

	return cur
}

type ListNode struct {
	Val  int
	Next *ListNode
}
