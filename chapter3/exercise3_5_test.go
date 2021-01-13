package chapter3_test

import (
	"testing"

	"github.com/jwalsh23/ctci/chapter3"
	"github.com/stretchr/testify/assert"
)

func TestSortStack_IsEmpty(t *testing.T) {
	t.Run("is empty", func(t *testing.T) {

		testObject := chapter3.NewSortStack()
		testObject.Push(10)
		testObject.Push(8)
		testObject.Push(6)

		testObject.Pop()
		testObject.Pop()
		testObject.Pop()

		assert.True(t, testObject.IsEmpty())

	})
	t.Run("is not empty", func(t *testing.T) {
		testObject := chapter3.NewSortStack()
		testObject.Push(10)
		testObject.Push(8)
		testObject.Push(6)

		assert.False(t, testObject.IsEmpty())
	})
}

func TestSortStack_Pop(t *testing.T) {
	testObject := chapter3.NewSortStack()
	testObject.Push(10)
	testObject.Push(8)
	testObject.Push(6)
	testObject.Push(20)

	assert.Equal(t, 6, testObject.Pop())
	assert.Equal(t, 8, testObject.Pop())
	assert.Equal(t, 10, testObject.Pop())
	assert.Equal(t, 20, testObject.Pop())
}

func TestSortStack_Peek(t *testing.T) {
	testObject := chapter3.NewSortStack()
	testObject.Push(10)
	testObject.Push(8)
	testObject.Push(6)
	testObject.Push(20)

	assert.Equal(t, 6, testObject.Peek())
}
