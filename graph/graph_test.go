package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GraphNodeCreation(t *testing.T) {
	assert := assert.New(t)
	graph := NewGraph()
	graph.AddNode(1)
	graph.AddNode(2)
	graph.AddNode(3)

	assert.True(graph.Contains(1), "graph should contain node.id 1")
	assert.True(graph.Contains(2), "graph should contain node.id 2")
	assert.True(graph.Contains(3), "graph should contain node.id 3")
	assert.False(graph.Contains(5), "graph should NOT contain node.id 5")
	assert.False(graph.Contains(6), "graph should NOT contain node.id 6")

	err := graph.AddNode(1)
	assert.NotNil(err, "Should not allow adding the same node again")

}
