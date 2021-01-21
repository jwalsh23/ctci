package data_structures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBST_Insert(t *testing.T) {
	t.Run("inserts node if doesnt exist", func(t *testing.T) {
		testObject := BST{}

		testObject.Insert(5)
		testObject.Insert(10)
		testObject.Insert(2)
		testObject.Insert(4)

		expected := BST{
			Root: &TreeNode{
				Value: 5,
				Left: &TreeNode{
					Value: 2,
					Left:  nil,
					Right: &TreeNode{Value: 4},
				},
				Right: &TreeNode{Value: 10},
			},
		}
		assert.Equal(t, expected, testObject)
	})
	t.Run("does not insert node if already exist", func(t *testing.T) {
		testObject := BST{}

		testObject.Insert(5)
		testObject.Insert(10)
		testObject.Insert(2)
		testObject.Insert(4)
		testObject.Insert(5)
		testObject.Insert(10)

		expected := BST{
			Root: &TreeNode{
				Value: 5,
				Left: &TreeNode{
					Value: 2,
					Left:  nil,
					Right: &TreeNode{Value: 4},
				},
				Right: &TreeNode{Value: 10},
			},
		}
		assert.Equal(t, expected, testObject)
	})
}
func TestBST_Delete(t *testing.T) {
	t.Run("node with only Right node", func(t *testing.T) {
		t.Run("deletes child", func(t *testing.T) {
			testObject := BST{}

			testObject.Insert(5)
			testObject.Insert(10)
			testObject.Delete(10)

			expected := BST{
				Root: &TreeNode{
					Value: 5,
				},
			}
			assert.Equal(t, expected, testObject)
		})
		t.Run("deletes parent", func(t *testing.T) {
			testObject := BST{}

			testObject.Insert(5)
			testObject.Insert(10)
			testObject.Delete(5)

			expected := BST{
				Root: &TreeNode{
					Value: 10,
				},
			}
			assert.Equal(t, expected, testObject)
		})
	})
	t.Run("node with only Left node", func(t *testing.T) {
		t.Run("deletes child", func(t *testing.T) {
			testObject := BST{}

			testObject.Insert(5)
			testObject.Insert(4)
			testObject.Delete(4)

			expected := BST{
				Root: &TreeNode{
					Value: 5,
				},
			}
			assert.Equal(t, expected, testObject)
		})
		t.Run("deletes parent", func(t *testing.T) {
			testObject := BST{}

			testObject.Insert(5)
			testObject.Insert(4)
			testObject.Delete(5)

			expected := BST{
				Root: &TreeNode{
					Value: 4,
				},
			}
			assert.Equal(t, expected, testObject)
		})
		t.Run("deletes parent with greatest on Left side", func(t *testing.T) {
			testObject := BST{}

			testObject.Insert(5)
			testObject.Insert(3)
			testObject.Insert(4)
			testObject.Delete(5)

			expected := BST{
				Root: &TreeNode{
					Value: 3,
					Right: &TreeNode{Value: 4},
				},
			}
			assert.Equal(t, expected, testObject)
		})
	})
	t.Run("node with both Left and Right child", func(t *testing.T) {
		testObject := BST{}

		testObject.Insert(5)
		testObject.Insert(3)
		testObject.Insert(4)
		testObject.Insert(10)
		testObject.Insert(8)
		testObject.Insert(9)
		testObject.Insert(7)
		testObject.Delete(5)

		expected := BST{
			Root: &TreeNode{
				Value: 7,
				Right: &TreeNode{
					Value: 10,
					Left: &TreeNode{
						Value: 8,
						Right: &TreeNode{Value: 9},
					},
				},
				Left: &TreeNode{
					Value: 3,
					Right: &TreeNode{Value: 4},
				},
			},
		}
		assert.Equal(t, expected, testObject)
	})
}
