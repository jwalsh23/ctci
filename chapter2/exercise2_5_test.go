package chapter2_test

import (
	"testing"

	"github.com/jwalsh23/ctci/chapter2"
	"github.com/jwalsh23/ctci/data_structures"
	"github.com/stretchr/testify/assert"
)

func TestSumListBackward(t *testing.T) {
	t.Run("equal length", func(t *testing.T) {
		var list1, list2 data_structures.LinkedList
		list1.Add(7)
		list1.Add(1)
		list1.Add(6)
		// list 2
		list2.Add(5)
		list2.Add(9)
		list2.Add(2)

		expected := data_structures.LinkedList{}
		expected.Add(2)
		expected.Add(1)
		expected.Add(9)

		assert.Equal(t, expected, chapter2.SumListBackward(list1, list2))
	})
	t.Run("unequal length", func(t *testing.T) {
		var list1, list2 data_structures.LinkedList
		list1.Add(7)
		list1.Add(1)
		list1.Add(6)
		// list 2
		list2.Add(2)

		expected := data_structures.LinkedList{}
		expected.Add(9)
		expected.Add(1)
		expected.Add(6)

		assert.Equal(t, expected, chapter2.SumListBackward(list1, list2))
	})
}
