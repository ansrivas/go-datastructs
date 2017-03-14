package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_TopologicalSort(t *testing.T) {
	assert := assert.New(t)
	graph := NewGraph()

	for i := 1; i < 9; i++ {
		graph.AddVertex(i)
	}
	graph.AddEdge(1, 3, true)
	graph.AddEdge(2, 3, true)
	graph.AddEdge(3, 5, true)
	graph.AddEdge(5, 8, true)
	graph.AddEdge(5, 6, true)
	graph.AddEdge(6, 7, true)
	graph.AddEdge(2, 4, true)
	graph.AddEdge(4, 6, true)

	nodeList := graph.TopologicalSort(true)

	//Get the node list and publish the id
	actual := []int{}
	for nodeList.Len() > 0 {
		node, err := nodeList.Pop()
		if err != nil {
			continue
		}
		actual = append(actual, node.id)
	}
	expected := []int{2, 4, 1, 3, 5, 6, 7, 8}

	assert.EqualValues(expected, actual, "Topological sorted list")
}
