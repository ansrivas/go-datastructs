package graph

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GraphNodeCreation(t *testing.T) {
	assert := assert.New(t)
	graph := NewGraph()
	graph.AddVertex(1)
	graph.AddVertex(2)
	graph.AddVertex(3)

	assert.True(graph.Contains(1), "graph should contain node.id 1")
	assert.True(graph.Contains(2), "graph should contain node.id 2")
	assert.True(graph.Contains(3), "graph should contain node.id 3")
	assert.False(graph.Contains(5), "graph should NOT contain node.id 5")
	assert.False(graph.Contains(6), "graph should NOT contain node.id 6")

	err := graph.AddVertex(1)
	assert.NotNil(err, "Should not allow adding the same node again")
}

func Test_Vertices(t *testing.T) {
	assert := assert.New(t)

	graph := NewGraph()
	graph.AddVertex(1)
	graph.AddVertex(2)
	graph.AddVertex(3)

	vertices := graph.ListVerticesID()
	sort.Ints(vertices)
	actual := []int{1, 2, 3}
	assert.EqualValues(vertices, actual, "ListVerticesID should get 3 node-ids")

	vertexList := graph.ListVertices()

	expected := make([]int, 0)
	for _, node := range vertexList {
		expected = append(expected, node.id)
	}
	sort.Ints(expected)
	actual = []int{1, 2, 3}
	assert.EqualValues(expected, actual, "ListVertices should get 3 node-ids")
}
