package no148

func sortList(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}

	cur := head
	for cur.Next != nil {
		cur = cur.Next
	}

	return doSort(head, cur)
}

func doSort(head, tail *ListNode) *ListNode {

	if head == tail {
		head.Next = nil
		return head
	} else if head.Next == tail {

		if head.Val < tail.Val {
			tail.Next = nil
			return head
		} else {
			tail.Next = head
			head.Next = nil
			return tail
		}
	}

	// 找到中点
	lowNode, fastNode := head, head
	for fastNode.Next != nil && fastNode.Next.Next != nil && fastNode.Next != tail && fastNode.Next.Next != tail {
		lowNode = lowNode.Next
		fastNode = fastNode.Next.Next
	}
	half := lowNode.Next
	newHead1 := doSort(head, lowNode)
	newHead2 := doSort(half, tail)

	// 归并

	return mergeTwoLists(newHead1, newHead2)
}

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
