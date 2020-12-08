package data_structures

type Queue struct {
	startNode *Node
	endNode *Node
}

func (q *Queue) Enqueue(val int) {
	node := &Node{Value: val}
	if q.startNode == nil {
		q.startNode = node
		q.endNode = node
		return
	}
	q.endNode.Next = node
	q.endNode = node
}
func (q *Queue) Dequeue() int {
	if q.startNode == nil {
		return 0
	}
	val := q.startNode
	q.startNode = val.Next
	return val.Value
}
