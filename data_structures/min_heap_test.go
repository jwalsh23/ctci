package data_structures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntMinHeap_Add(t *testing.T) {
	t.Run("add items to the heap in order", func(t *testing.T) {
		testObject := IntMinHeap{}
		testObject.Add(5)
		testObject.Add(10)
		testObject.Add(20)
		testObject.Add(15)
		testObject.Add(16)
		testObject.Add(30)
		testObject.Add(32)

		expected := IntMinHeap{
			slc: []int{0, 5, 10, 20, 15, 16, 30, 32},
		}

		assert.Equal(t, expected, testObject)
	})
	t.Run("add items to the heap in order", func(t *testing.T) {
		testObject := IntMinHeap{}
		testObject.Add(5)
		testObject.Add(10)
		testObject.Add(20)
		testObject.Add(15)
		testObject.Add(16)
		testObject.Add(30)
		testObject.Add(32)
		testObject.Add(4)
		testObject.Add(19)

		expected := IntMinHeap{
			slc: []int{0, 4, 5, 20, 10, 16, 30, 32, 15, 19},
		}

		assert.Equal(t, expected, testObject)
	})
}

func TestIntMinHeap_Remove(t *testing.T) {
	t.Run("add items to the heap in order", func(t *testing.T) {
		testObject := IntMinHeap{}
		testObject.Add(5)
		testObject.Add(10)
		testObject.Add(20)
		testObject.Add(15)
		testObject.Add(16)
		testObject.Add(30)
		testObject.Add(32)
		testObject.Add(4)
		testObject.Add(19)

		assert.Equal(t, 4, testObject.Remove())
		assert.Equal(t, 5, testObject.Remove())
		assert.Equal(t, 10, testObject.Remove())
		assert.Equal(t, 15, testObject.Remove())
		assert.Equal(t, 16, testObject.Remove())
		assert.Equal(t, 19, testObject.Remove())
		assert.Equal(t, 20, testObject.Remove())
		assert.Equal(t, 30, testObject.Remove())
		assert.Equal(t, 32, testObject.Remove())
	})
}
