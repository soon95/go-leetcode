package no82

func deleteDuplicates(head *ListNode) *ListNode {

	numCnt := make(map[int]int)

	cur := head
	for cur != nil {
		numCnt[cur.Val]++
		cur = cur.Next
	}

	var newHead *ListNode
	var newPre *ListNode
	cur = head

	for cur != nil {

		if numCnt[cur.Val] == 1 {
			if newHead == nil {
				newHead = cur
			} else {
				newPre.Next = cur
			}
			newPre = cur

			cur = cur.Next
			newPre.Next = nil

		} else {
			cur = cur.Next
		}
	}

	return newHead
}

type ListNode struct {
	Val  int
	Next *ListNode
}
