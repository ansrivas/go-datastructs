package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Queue(t *testing.T) {
	assert := assert.New(t)
	queue := NewGQueue()
	queue.Enqueue(&GNode{id: 1})
	queue.Enqueue(&GNode{id: 2})
	expected := queue.Len()
	assert.Equal(expected, 2, "Current two elements in queue")

	node, _ := queue.Dqueue()
	assert.Equal(node.id, 1, "Dqueue should return correct element")

	node2, err := queue.Dqueue()
	assert.Nil(err, "Nil error should be returned")
	assert.Equal(node2.id, 2, "Dqueue should return correct element")

	_, err = queue.Dqueue()
	assert.Equal(err.Error(), "Emtpy queue.", "Empty queue should throw error")
}
