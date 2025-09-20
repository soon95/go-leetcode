package no21

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	var (
		cur  *ListNode
		next *ListNode
	)

	if list1.Val < list2.Val {
		cur = list1
		next = mergeTwoLists(list1.Next, list2)
	} else {
		cur = list2
		next = mergeTwoLists(list1, list2.Next)
	}

	cur.Next = next

	return cur
}

type ListNode struct {
	Val  int
	Next *ListNode
}
