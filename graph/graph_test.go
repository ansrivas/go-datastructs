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

	assert.Equal(0, len(graph.ListVertices()), "Empty vertex list")

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

func Test_Edges(t *testing.T) {
	assert := assert.New(t)

	graph := NewGraph()

	graph.AddVertex(1)
	graph.AddVertex(2)
	graph.AddVertex(3)
	graph.AddVertex(4)

	graph.AddEdge(1, 2, true)
	graph.AddEdge(2, 3, true)

	//Add edge and use `false` to declare it non-directed edge
	graph.AddEdge(3, 2, false)
	expected, _ := graph.EdgeExists(2, 3)
	assert.Equal(expected, true, "Edges should exist between 3,2")

	//Test adding to non-existing vertex
	err := graph.AddEdge(4, 6, true)
	assert.NotNil(err, "Adding an edge to non-existing vertex should throw error")

	expected, _ = graph.EdgeExists(1, 2)
	assert.Equal(expected, true, "Edges should exist between 1,2")
	expected, _ = graph.EdgeExists(2, 3)
	assert.Equal(expected, true, "Edges should exist between 2,3")

	expected, _ = graph.EdgeExists(3, 4)
	assert.Equal(expected, false, "Edges should NOT exist between 3,4")

	expected, err = graph.EdgeExists(3, 5)
	assert.NotNil(err, "Edges should NOT exist between non existing nodes")
	assert.Equal(expected, false, "Edges should NOT exist between non existing nodes")
}

func Test_Neighbors(t *testing.T) {
	assert := assert.New(t)

	graph := NewGraph()

	graph.AddVertex(1)
	graph.AddVertex(2)
	graph.AddVertex(3)
	graph.AddVertex(4)

	graph.AddEdge(1, 2, true)
	graph.AddEdge(1, 3, true)
	// This will create edges from both 2->3 and 3->2
	graph.AddEdge(2, 3, false)

	//This should return two objects node.id=2 and node.id=3
	neighbors, err := graph.GetNeighbors(1)

	identifiers := make([]int, 0)
	for _, neighbor := range neighbors {
		identifiers = append(identifiers, neighbor.id)
	}
	sort.Ints(identifiers)

	assert.Nil(err, "Should not throw error")
	assert.Equal(identifiers, []int{2, 3}, "Neighbors of node1 should be node2 and node3")

	_, err = graph.GetNeighbors(5)
	assert.NotNil(err, "Non existent node should not throw error")

}
