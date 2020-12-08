package data_structures

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack_Push(t *testing.T) {
	t.Run("pushes elements", func(t *testing.T) {
		expected := Stack{top: &stackNode{Value: 10, prev: &stackNode{Value: 5}}}

		testObject := Stack{}
		testObject.Push(5)
		testObject.Push(10)

		assert.Equal(t, expected, testObject)
	})
}
func TestStack_Pop(t *testing.T) {
	t.Run("pops elements", func(t *testing.T) {
		testObject := Stack{}
		testObject.Push(5)
		testObject.Push(10)

		assert.Equal(t, 10, testObject.Pop())
		assert.Equal(t, 5, testObject.Pop())
		assert.Equal(t, 0, testObject.Pop())
		assert.Equal(t, 0, testObject.Pop())
	})
}
