package data_structures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntGraph_AddNode(t *testing.T) {
	testObject := IntGraph{}
	testObject.AddNode(1)
	testObject.AddNode(5)
	testObject.AddNode(6)

	expected := IntGraph{
		vertexEdge: map[int]*LinkedList{
			1: {},
			5: {},
			6: {},
		},
	}

	assert.Equal(t, expected, testObject)
}

func TestIntGraph_AddRelation(t *testing.T) {
	testObject := IntGraph{}
	testObject.AddNode(1)
	testObject.AddNode(5)
	testObject.AddNode(6)
	testObject.AddRelation(1, 7, true)
	testObject.AddRelation(5, 6, false)

	expected := IntGraph{
		vertexEdge: map[int]*LinkedList{
			1: {n: &Node{Value: 7}, count: 1},
			5: {n: &Node{Value: 6}, count: 1},
			6: {},
			7: {n: &Node{Value: 1}, count: 1},
		},
	}

	assert.Equal(t, expected, testObject)
}

func TestIntGraph_RelationExist(t *testing.T) {
	t.Run("relation does not exist", func(t *testing.T) {
		testObject := IntGraph{}
		testObject.AddNode(1)
		testObject.AddNode(5)
		testObject.AddNode(6)
		testObject.AddRelation(1, 7, true)
		testObject.AddRelation(5, 6, false)

		assert.False(t, testObject.RelationExist(6, 5))
	})
	t.Run("relation does exist", func(t *testing.T) {
		t.Run("first level", func(t *testing.T) {
			testObject := IntGraph{}
			testObject.AddNode(1)
			testObject.AddNode(5)
			testObject.AddNode(6)
			testObject.AddRelation(1, 7, true)
			testObject.AddRelation(5, 6, false)

			assert.True(t, testObject.RelationExist(7, 1))
		})
		t.Run("multi level", func(t *testing.T) {
			testObject := IntGraph{}
			testObject.AddNode(1)
			testObject.AddNode(5)
			testObject.AddNode(6)
			testObject.AddRelation(1, 7, true)
			testObject.AddRelation(5, 6, true)
			testObject.AddRelation(5, 8, true)
			testObject.AddRelation(5, 7, true)
			testObject.AddRelation(1, 6, true)

			assert.True(t, testObject.RelationExist(5, 1))
		})
	})
}
