package chapter4

type GraphNode struct {
	Value    string
	adjacent map[*GraphNode]struct{}
}

func (n *GraphNode) AddEdge(node *GraphNode) {
	if n.adjacent == nil {
		n.adjacent = map[*GraphNode]struct{}{
			node: {},
		}
		return
	}
	n.adjacent[node] = struct{}{}
}

func (n *GraphNode) IsConnected(other *GraphNode) bool {
	return breadthFirstSearch(n, other)
}

func breadthFirstSearch(n1, n2 *GraphNode) bool {
	queue := new(GraphQueue)
	queue.Enqueue(n1)
	for !queue.IsEmpty() {
		node := queue.Dequeue()
		if node.Visited {
			continue
		}
		node.Visited = true
		if _, ok := node.Val.adjacent[n2]; ok {
			return true
		}
		for n := range node.Val.adjacent {
			queue.Enqueue(n)
		}
	}
	return false
}
