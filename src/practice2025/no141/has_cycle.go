package no141

func hasCycle(head *ListNode) bool {

	nodeMap := make(map[*ListNode]bool)

	for head != nil {

		_, exist := nodeMap[head]
		if exist {
			return true
		}

		nodeMap[head] = true
		head = head.Next
	}
	return false
}

type ListNode struct {
	Val  int
	Next *ListNode
}
