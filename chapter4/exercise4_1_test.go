package chapter4_test

import (
	"testing"

	"github.com/jwalsh23/ctci/chapter4"
	"github.com/stretchr/testify/assert"
)

func TestGraphNode_IsConnected(t *testing.T) {
	t.Run("is connected", func(t *testing.T) {
		startNode := &chapter4.GraphNode{Value: "S"}
		firstEdge := &chapter4.GraphNode{Value: "X"}
		firstEdge.AddEdge(&chapter4.GraphNode{Value: "Z"})
		secondEdge := &chapter4.GraphNode{Value: "T"}
		firstEdge.AddEdge(secondEdge)
		thirdEdge := &chapter4.GraphNode{Value: "A"}
		startNode.AddEdge(firstEdge)
		startNode.AddEdge(thirdEdge)

		assert.True(t, startNode.IsConnected(secondEdge))
	})
	t.Run("is not connected", func(t *testing.T) {
		startNode := &chapter4.GraphNode{Value: "S"}
		firstEdge := &chapter4.GraphNode{Value: "X"}
		firstEdge.AddEdge(&chapter4.GraphNode{Value: "Z"})
		secondEdge := &chapter4.GraphNode{Value: "T"}
		thirdEdge := &chapter4.GraphNode{Value: "A"}
		thirdEdge.AddEdge(firstEdge)
		startNode.AddEdge(firstEdge)

		assert.False(t, startNode.IsConnected(secondEdge))
	})
}
