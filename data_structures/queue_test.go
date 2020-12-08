package data_structures

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueue_Enqueue(t *testing.T) {
	t.Run("enqueues elements", func(t *testing.T) {
		endNode := &Node{Value: 15}
		startNode := &Node{Value: 5, Next: &Node{Value: 10, Next: endNode}}
		expected := Queue{startNode: startNode, endNode: endNode}

		testObject := Queue{}
		testObject.Enqueue(5)
		testObject.Enqueue(10)
		testObject.Enqueue(15)

		assert.Equal(t, expected, testObject)
	})
}
func TestQueue_Dequeue(t *testing.T) {
	t.Run("dequeues elements", func(t *testing.T) {
		testObject := Queue{}
		testObject.Enqueue(5)
		testObject.Enqueue(10)

		assert.Equal(t, 5, testObject.Dequeue())
		assert.Equal(t, 10, testObject.Dequeue())
		assert.Equal(t, 0, testObject.Dequeue())
		assert.Equal(t, 0, testObject.Dequeue())
	})
}
