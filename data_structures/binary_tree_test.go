package data_structures

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBST_Insert(t *testing.T) {
	t.Run("inserts node if doesnt exist", func(t *testing.T) {
		testObject := BST{}

		testObject.Insert(5)
		testObject.Insert(10)
		testObject.Insert(2)
		testObject.Insert(4)

		expected := BST{
			root: &treeNode{
				Value: 5,
				left:  &treeNode{
					Value: 2,
					left:  nil,
					right: &treeNode{Value: 4},
				},
				right: &treeNode{Value: 10},
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
			root: &treeNode{
				Value: 5,
				left:  &treeNode{
					Value: 2,
					left:  nil,
					right: &treeNode{Value: 4},
				},
				right: &treeNode{Value: 10},
			},
		}
		assert.Equal(t, expected, testObject)
	})
}
func TestBST_Delete(t *testing.T) {
	t.Run("node with only right node", func(t *testing.T) {
		t.Run("deletes child", func(t *testing.T) {
			testObject := BST{}

			testObject.Insert(5)
			testObject.Insert(10)
			testObject.Delete(10)

			expected := BST{
				root: &treeNode{
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
				root: &treeNode{
					Value: 10,
				},
			}
			assert.Equal(t, expected, testObject)
		})
	})
	t.Run("node with only left node", func(t *testing.T) {
		t.Run("deletes child", func(t *testing.T) {
			testObject := BST{}

			testObject.Insert(5)
			testObject.Insert(4)
			testObject.Delete(4)

			expected := BST{
				root: &treeNode{
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
				root: &treeNode{
					Value: 4,
				},
			}
			assert.Equal(t, expected, testObject)
		})
		t.Run("deletes parent with greatest on left side", func(t *testing.T) {
			testObject := BST{}

			testObject.Insert(5)
			testObject.Insert(3)
			testObject.Insert(4)
			testObject.Delete(5)

			expected := BST{
				root: &treeNode{
					Value: 3,
					right: &treeNode{Value: 4},
				},
			}
			assert.Equal(t, expected, testObject)
		})
	})
	t.Run("node with both left and right child", func(t *testing.T) {
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
				root: &treeNode{
					Value: 7,
					right: &treeNode{
						Value: 10,
						left: &treeNode{
							Value: 8,
							right: &treeNode{Value: 9},
						},
					},
					left: &treeNode{
						Value: 3,
						right: &treeNode{Value: 4},
					},
				},
			}
			assert.Equal(t, expected, testObject)
	})
}
