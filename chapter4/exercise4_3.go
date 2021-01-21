package chapter4

import (
	"math"

	"github.com/jwalsh23/ctci/data_structures"
)

func TreeToLists(bst *data_structures.BST) []*data_structures.LinkedList {
	if bst.Root == nil {
		return nil
	}
	count := 0
	var returnSlc []*data_structures.LinkedList
	queue := new(data_structures.TreeNodeQueue)
	queue.Enqueue(bst.Root)
	depth := 0
	for !queue.IsEmpty() {
		queueNode := queue.Dequeue()
		count++
		level := math.Log2(float64(count))
		if level == float64(depth) {
			depth++
		}
		if queueNode.Val == nil {
			continue
		}
		if depth > len(returnSlc) {
			returnSlc = append(returnSlc, new(data_structures.LinkedList))
		}
		returnSlc[depth-1].Add(queueNode.Val.Value)
		queue.Enqueue(queueNode.Val.Left)
		queue.Enqueue(queueNode.Val.Right)
	}

	return returnSlc
}
