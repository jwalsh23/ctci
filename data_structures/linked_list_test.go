package data_structures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedList_Add(t *testing.T) {
	t.Run("adds element when list is not nil", func(t *testing.T) {
		expected := LinkedList{N: &Node{Value: 5}, Count: 1}

		var testObject LinkedList
		testObject.Add(5)
		assert.Equal(t, expected, testObject)
	})
	t.Run("adds multiple elements", func(t *testing.T) {
		expected := LinkedList{N: &Node{Value: 5, Next: &Node{Value: 10}}, Count: 2}

		var testObject LinkedList
		testObject.Add(5)
		testObject.Add(10)
		assert.Equal(t, expected, testObject)
	})
}

func TestLinkedList_Empty(t *testing.T) {
	t.Run("is empty when l is nil", func(t *testing.T) {
		var testObject *LinkedList
		assert.True(t, testObject.Empty())
	})
	t.Run("is empty when l has no elements", func(t *testing.T) {
		testObject := LinkedList{}
		assert.True(t, testObject.Empty())
	})
	t.Run("is not empty after adding element", func(t *testing.T) {
		testObject := LinkedList{}
		testObject.Add(5)
		assert.False(t, testObject.Empty())
	})
}
func TestLinkedList_DeleteValue(t *testing.T) {
	t.Run("does not deleteNode when cannot find value", func(t *testing.T) {
		expected := LinkedList{N: &Node{Value: 5, Next: &Node{Value: 10}}, Count: 2}

		var testObject LinkedList
		testObject.Add(5)
		testObject.Add(10)

		testObject.DeleteValue(20)
		assert.Equal(t, expected, testObject)
	})
	t.Run("deletes all values that occur", func(t *testing.T) {
		expected := LinkedList{N: &Node{Value: 5}, Count: 1}

		var testObject LinkedList
		testObject.Add(10)
		testObject.Add(5)
		testObject.Add(10)

		testObject.DeleteValue(10)
		assert.Equal(t, expected, testObject)
	})
	t.Run("deletes all values that occur", func(t *testing.T) {
		expected := LinkedList{}

		var testObject LinkedList
		testObject.Add(10)
		testObject.Add(10)
		testObject.Add(10)

		testObject.DeleteValue(10)
		assert.Equal(t, expected, testObject)
	})
}
func TestLinkedList_DeleteIndex(t *testing.T) {
	t.Run("does not deleteNode when cannot find index", func(t *testing.T) {
		expected := LinkedList{N: &Node{Value: 5, Next: &Node{Value: 10}}, Count: 2}

		var testObject LinkedList
		testObject.Add(5)
		testObject.Add(10)

		testObject.DeleteIndex(5)
		assert.Equal(t, expected, testObject)
	})
	t.Run("deletes middle index", func(t *testing.T) {
		expected := LinkedList{N: &Node{Value: 10, Next: &Node{Value: 10}}, Count: 2}

		var testObject LinkedList
		testObject.Add(10)
		testObject.Add(5)
		testObject.Add(10)

		testObject.DeleteIndex(1)
		assert.Equal(t, expected, testObject)
	})
	t.Run("deletes first index", func(t *testing.T) {
		expected := LinkedList{N: &Node{Value: 5, Next: &Node{Value: 10}}, Count: 2}

		var testObject LinkedList
		testObject.Add(10)
		testObject.Add(5)
		testObject.Add(10)

		testObject.DeleteIndex(0)
		assert.Equal(t, expected, testObject)
	})
	t.Run("deletes last index", func(t *testing.T) {
		expected := LinkedList{N: &Node{Value: 10, Next: &Node{Value: 5}}, Count: 2}

		var testObject LinkedList
		testObject.Add(10)
		testObject.Add(5)
		testObject.Add(10)

		testObject.DeleteIndex(2)
		assert.Equal(t, expected, testObject)
	})
}
