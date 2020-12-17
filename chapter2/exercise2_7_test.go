package chapter2_test

import (
	"testing"

	"github.com/jwalsh23/ctci/chapter2"
	"github.com/jwalsh23/ctci/data_structures"
	"github.com/stretchr/testify/assert"
)

func TestIntersection(t *testing.T) {
	t.Run("is not intersection", func(t *testing.T) {
		list1 := data_structures.LinkedList{}
		list2 := data_structures.LinkedList{}

		list1.Add(1)
		list1.Add(2)
		list1.Add(3)

		list2.Add(4)
		list2.Add(2)
		list2.Add(5)

		assert.False(t, chapter2.Intersection(list1, list2))
	})
	t.Run("is intersection", func(t *testing.T) {
		intersectionVal := data_structures.Node{Value: 5, Next: &data_structures.Node{Value: 4}}

		firstListStart := &data_structures.Node{Value: 10, Next: &intersectionVal}
		secondListStart := &data_structures.Node{
			Value: 12,
			Next:  &data_structures.Node{Value: 50, Next: &intersectionVal},
		}

		list1 := data_structures.LinkedList{N: firstListStart, Count: 3}
		list2 := data_structures.LinkedList{N: secondListStart, Count: 4}

		assert.True(t, chapter2.Intersection(list1, list2))
	})
	t.Run("is intersection", func(t *testing.T) {
		intersectionVal := data_structures.Node{Value: 5, Next: &data_structures.Node{Value: 4}}

		firstListStart := &data_structures.Node{Value: 10, Next: &intersectionVal}
		secondListStart := &data_structures.Node{
			Value: 12,
			Next:  &data_structures.Node{Value: 50, Next: &intersectionVal},
		}

		list1 := data_structures.LinkedList{N: firstListStart, Count: 3}
		list2 := data_structures.LinkedList{N: secondListStart, Count: 4}

		assert.True(t, chapter2.Intersection(list2, list1))
	})
}

func TestIntersectionNoExtraStructures(t *testing.T) {
	t.Run("is not intersection", func(t *testing.T) {
		list1 := data_structures.LinkedList{}
		list2 := data_structures.LinkedList{}

		list1.Add(1)
		list1.Add(2)
		list1.Add(3)

		list2.Add(4)
		list2.Add(2)
		list2.Add(5)

		assert.Nil(t, chapter2.IntersectionNoExtraStructures(list1, list2))
	})
	t.Run("is intersection", func(t *testing.T) {
		intersectionVal := data_structures.Node{Value: 5, Next: &data_structures.Node{Value: 4}}

		firstListStart := &data_structures.Node{Value: 10, Next: &intersectionVal}
		secondListStart := &data_structures.Node{
			Value: 12,
			Next:  &data_structures.Node{Value: 50, Next: &intersectionVal},
		}

		list1 := data_structures.LinkedList{N: firstListStart, Count: 3}
		list2 := data_structures.LinkedList{N: secondListStart, Count: 4}

		assert.Equal(t, &intersectionVal, chapter2.IntersectionNoExtraStructures(list1, list2))
	})
	t.Run("is intersection", func(t *testing.T) {
		intersectionVal := data_structures.Node{Value: 5, Next: &data_structures.Node{Value: 4}}

		firstListStart := &data_structures.Node{Value: 10, Next: &intersectionVal}
		secondListStart := &data_structures.Node{
			Value: 12,
			Next:  &data_structures.Node{Value: 50, Next: &intersectionVal},
		}

		list1 := data_structures.LinkedList{N: firstListStart, Count: 3}
		list2 := data_structures.LinkedList{N: secondListStart, Count: 4}

		assert.Equal(t, &intersectionVal, chapter2.IntersectionNoExtraStructures(list2, list1))
	})
}
