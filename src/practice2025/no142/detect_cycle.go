package no142

func detectCycle(head *ListNode) *ListNode {

	nodeMap := make(map[*ListNode]bool)

	for head != nil {

		_, exist := nodeMap[head]
		if exist {
			return head
		}

		nodeMap[head] = true
		head = head.Next
	}
	return nil
}

type ListNode struct {
	Val  int
	Next *ListNode
}
