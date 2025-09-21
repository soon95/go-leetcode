package no138

func copyRandomList(head *Node) *Node {

	nodeMap := make(map[*Node]*Node)

	cur := head
	for cur != nil {
		newNode := &Node{
			Val: cur.Val,
		}
		nodeMap[cur] = newNode

		cur = cur.Next
	}

	cur = head

	for cur != nil {
		newCur := nodeMap[cur]

		newCur.Next = nodeMap[cur.Next]

		newCur.Random = nodeMap[cur.Random]

		cur = cur.Next
	}

	return nodeMap[head]
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}
