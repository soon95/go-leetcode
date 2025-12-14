package no117

func connect(root *Node) *Node {

	if root == nil {
		return nil
	}

	visitNodes := make([]*Node, 0)

	visitNodes = append(visitNodes, root)

	for len(visitNodes) > 0 {

		nextVisitNodes := make([]*Node, 0)
		for i := 0; i < len(visitNodes); i++ {

			node := visitNodes[i]
			if node.Left != nil {
				nextVisitNodes = append(nextVisitNodes, node.Left)
			}

			if node.Right != nil {
				nextVisitNodes = append(nextVisitNodes, node.Right)
			}

			if i+1 < len(visitNodes) {
				node.Next = visitNodes[i+1]
			}
		}
		visitNodes = nextVisitNodes
	}

	return root
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}
