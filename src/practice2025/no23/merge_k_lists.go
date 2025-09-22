package no23

func mergeKLists(lists []*ListNode) *ListNode {

	minIndex := -1

	for i, node := range lists {

		if node != nil && (minIndex == -1 || node.Val < lists[minIndex].Val) {
			minIndex = i
		}
	}

	if minIndex == -1 {
		return nil
	}

	head := lists[minIndex]
	lists[minIndex] = head.Next
	head.Next = mergeKLists(lists)

	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}
