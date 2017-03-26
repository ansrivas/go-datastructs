package binIndexedTree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Parent_Next(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(10, GetParent(11), "Parent of 11 should be 10")
	assert.Equal(0, GetParent(1), "Parent of 1 should be 0")
	assert.Equal(2, GetParent(3), "Parent of 3 should be 2")

	assert.Equal(2, GetNext(1), "Next of 1 should be 2")
	assert.Equal(4, GetNext(2), "Next of 2 should be 4")
	assert.Equal(8, GetNext(4), "Next of 4 should be 8")
}

func Test_Insert(t *testing.T) {
	assert := assert.New(t)

	input := []int{3, 2, -1, 6, 5, 4, -3, 3, 7, 2, 3}

	bit := NewBinaryIndexedTree(len(input))

	for idx, val := range input {
		bit.InsertTree(idx, val)
	}

	// This gives a sum from index [0,5] (inclusive)
	sum := bit.GetSum(5)
	assert.Equal(sum, 19, "Sum upto first 5 element should be 19")
	sum = bit.GetSum(7) - bit.GetSum(5)
	assert.Equal(sum, 0, "Sum between 7-5 should be 0")

	index := 3
	update := +3

	bit.UpdateTree(index, update)
	state := bit.GetTreeRepresentation()
	expectedState := BinIndTree{0, 3, 5, -1, 13, 5, 9, -3, 22, 7, 9, 3}
	assert.Equal(*state, expectedState, "After update the correct tree should be formed")

	index = 4
	update = +1
	bit.UpdateTree(index, update)
	expectedState = BinIndTree{0, 3, 5, -1, 13, 6, 10, -3, 23, 7, 9, 3}
	assert.Equal(*state, expectedState, "After update the correct tree should be formed")
}
