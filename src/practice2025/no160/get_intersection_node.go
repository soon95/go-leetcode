package no160

func getIntersectionNode(headA, headB *ListNode) *ListNode {

	nodeMap := make(map[*ListNode]bool)

	for headA != nil {
		nodeMap[headA] = true
		headA = headA.Next
	}

	for headB != nil {
		_, exist := nodeMap[headB]
		if exist {
			return headB
		}
		headB = headB.Next
	}

	return nil
}

type ListNode struct {
	Val  int
	Next *ListNode
}
