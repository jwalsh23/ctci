package chapter4

type GraphQueue struct {
	startNode *QueueNode
	endNode   *QueueNode
}

type QueueNode struct {
	Val     *GraphNode
	Visited bool
	Next    *QueueNode
}

func (q *GraphQueue) Enqueue(n *GraphNode) {
	node := &QueueNode{Val: n}
	if q.startNode == nil {
		q.startNode = node
		q.endNode = node
		return
	}
	q.endNode.Next = node
	q.endNode = node
}
func (q *GraphQueue) Dequeue() *QueueNode {
	if q.startNode == nil {
		return nil
	}
	val := q.startNode
	q.startNode = val.Next
	return val
}

func (q *GraphQueue) IsEmpty() bool {
	return q.startNode == nil
}
