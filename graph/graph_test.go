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

	assert.Equal(0, len(graph.GetVertices()), "Empty vertex list")

	graph.AddVertex(1)
	graph.AddVertex(2)
	graph.AddVertex(3)

	vertices := graph.GetVerticesID()
	sort.Ints(vertices)
	actual := []int{1, 2, 3}
	assert.EqualValues(vertices, actual, "GetVerticesID should get 3 node-ids")

	vertexList := graph.GetVertices()

	expected := make([]int, 0)
	for _, node := range vertexList {
		expected = append(expected, node.id)
	}
	sort.Ints(expected)
	actual = []int{1, 2, 3}
	assert.EqualValues(expected, actual, "GetVertices should get 3 node-ids")
}

func Test_Edges(t *testing.T) {
	assert := assert.New(t)

	graph := NewGraph()

	graph.AddVertex(1)
	graph.AddVertex(2)
	graph.AddVertex(3)
	graph.AddVertex(4)

	graph.AddEdge(1, 2, true, 0)
	graph.AddEdge(1, 3, true, 0)
	graph.AddEdge(2, 3, true, 0)

	_, err := graph.GetEdges(9)
	assert.NotNil(err, "GetEdges should fail for non-existing nodes")
	edges, _ := graph.GetEdges(1)
	actualEdgeList := make([]int, 0)
	for _, edge := range edges {
		actualEdgeList = append(actualEdgeList, edge.target.id)
	}
	sort.Ints(actualEdgeList)
	expectedEdgeList := []int{2, 3}
	assert.Equal(expectedEdgeList, actualEdgeList, "GetEdges should return correct edge list")

	//Add edge and use `false` to declare it non-directed edge
	graph.AddEdge(3, 2, false, 0)
	expected, _ := graph.EdgeExists(2, 3)
	assert.Equal(expected, true, "Edges should exist between 3,2")

	//Test adding to non-existing vertex
	err = graph.AddEdge(4, 6, true, 0)
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

func Test_EdgesRemoval(t *testing.T) {
	assert := assert.New(t)

	graph := NewGraph()

	graph.AddVertex(1)
	graph.AddVertex(2)
	graph.AddVertex(3)
	graph.AddVertex(4)

	graph.AddEdge(1, 2, true, 0)
	graph.AddEdge(1, 3, true, 0)
	graph.AddEdge(2, 3, false, 0)
	graph.AddEdge(3, 4, true, 0)

	//Try with non-existent edges
	shouldBeNotNil := graph.RemoveEdge(6, 7, false)
	assert.NotNil(shouldBeNotNil, "Error should be thrown in case node is not found.")

	graph.RemoveEdge(1, 2, true)
	edgeListOne, _ := graph.GetEdges(1)
	actual := make([]int, 0)
	for _, edge := range edgeListOne {
		actual = append(actual, edge.target.id)
	}
	expected := []int{3}
	assert.Equal(actual, expected, "After removing edgelist of 1 should be only 3")

	//Now remove edge 3,4 to see wassup
	graph.RemoveEdge(3, 2, false)
	edgeListThree, _ := graph.GetEdges(3)
	actual = make([]int, 0)
	for _, edge := range edgeListThree {
		actual = append(actual, edge.target.id)
	}
	expected = []int{4}
	assert.Equal(actual, expected, "After removing edgelist of 3 should be only 4")

}

func Test_Neighbors(t *testing.T) {
	assert := assert.New(t)

	graph := NewGraph()

	graph.AddVertex(1)
	graph.AddVertex(2)
	graph.AddVertex(3)
	graph.AddVertex(4)

	graph.AddEdge(1, 2, true, 0)
	graph.AddEdge(1, 3, true, 0)
	// This will create edges from both 2->3 and 3->2
	graph.AddEdge(2, 3, false, 0)

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
