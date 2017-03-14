package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewStack(t *testing.T) {
	assert := assert.New(t)

	stack := NewGStack()
	stack.Push(&GNode{id: 1})
	stack.Push(&GNode{id: 2})

	numElems := stack.Len()
	assert.Equal(numElems, 2, "Num of elements should be 2 in this stack")

	node1, _ := stack.Pop()
	assert.Equal(node1.id, 2, "Should pop from the last element")

	node2, _ := stack.Pop()
	assert.Equal(node2.id, 1, "Should pop from the last element")

	_, err := stack.Pop()
	assert.Equal(err.Error(), "Emtpy stack.", "Should return error if empty stack popped")
}
