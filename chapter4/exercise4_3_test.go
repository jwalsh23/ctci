package chapter4_test

import (
	"testing"

	"github.com/jwalsh23/ctci/chapter4"
	"github.com/stretchr/testify/assert"

	"github.com/jwalsh23/ctci/data_structures"
)

func TestTreeToLists(t *testing.T) {
	t.Run("works as expected", func(t *testing.T) {
		tree := new(data_structures.BST)
		tree.Insert(10)
		tree.Insert(5)
		tree.Insert(3)
		tree.Insert(7)
		tree.Insert(20)
		tree.Insert(15)
		tree.Insert(30)
		tree.Insert(40)

		expected := []*data_structures.LinkedList{
			{Count: 1, N: &data_structures.Node{Value: 10}},
			{Count: 2, N: &data_structures.Node{Value: 5, Next: &data_structures.Node{Value: 20}}},
			{Count: 4,
				N: &data_structures.Node{
					Value: 3,
					Next: &data_structures.Node{
						Value: 7,
						Next: &data_structures.Node{
							Value: 15,
							Next:  &data_structures.Node{Value: 30},
						},
					},
				},
			},
			{Count: 1, N: &data_structures.Node{Value: 40}},
		}
		assert.Equal(t, expected, chapter4.TreeToLists(tree))
	})
}
