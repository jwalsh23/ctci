package data_structures

type TreeNodeQueue struct {
	startNode *QueueNode
	endNode   *QueueNode
}

type QueueNode struct {
	Val     *TreeNode
	Visited bool
	Next    *QueueNode
}

func (q *TreeNodeQueue) Enqueue(n *TreeNode) {
	node := &QueueNode{Val: n}
	if q.startNode == nil {
		q.startNode = node
		q.endNode = node
		return
	}
	q.endNode.Next = node
	q.endNode = node
}
func (q *TreeNodeQueue) Dequeue() *QueueNode {
	if q.startNode == nil {
		return nil
	}
	val := q.startNode
	q.startNode = val.Next
	return val
}

func (q *TreeNodeQueue) IsEmpty() bool {
	return q.startNode == nil
}
