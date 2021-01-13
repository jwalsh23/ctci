package chapter3_test

import (
	"testing"

	"github.com/jwalsh23/ctci/chapter3"
	"github.com/stretchr/testify/assert"
)

func TestMultiStack_Push(t *testing.T) {
	t.Run("pushes correctly", func(t *testing.T) {
		expectedStack1 := new(chapter3.Stack)
		expectedStack1.Push(1)
		expectedStack1.Push(2)

		expectedStack2 := new(chapter3.Stack)
		expectedStack2.Push(3)
		expectedStack2.Push(4)

		testObject := chapter3.NewMultiStack(2)
		testObject.Push(1)
		testObject.Push(2)
		testObject.Push(3)
		testObject.Push(4)

		assert.Equal(t, testObject.Stacks[0], expectedStack1)
		assert.Equal(t, testObject.Stacks[1], expectedStack2)
	})
}
func TestMultiStack_Pop(t *testing.T) {
	t.Run("pops as expected", func(t *testing.T) {
		testObject := chapter3.NewMultiStack(2)
		testObject.Push(1)
		testObject.Push(2)
		testObject.Push(3)
		testObject.Push(4)

		assert.Equal(t, 4, testObject.Pop())
		assert.Equal(t, 3, testObject.Pop())
		assert.Equal(t, 2, testObject.Pop())
		assert.Equal(t, 1, testObject.Pop())
	})
}

func TestMultiStack_PopAt(t *testing.T) {
	t.Run("pops as expected", func(t *testing.T) {
		testObject := chapter3.NewMultiStack(2)
		testObject.Push(1)
		testObject.Push(2)
		testObject.Push(3)
		testObject.Push(4)

		assert.Equal(t, 2, testObject.PopAt(0))
		assert.Equal(t, 1, testObject.PopAt(0))
		assert.Equal(t, 4, testObject.PopAt(1))
		assert.Equal(t, 3, testObject.PopAt(1))
	})
}
