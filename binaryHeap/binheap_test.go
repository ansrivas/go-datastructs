package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_HeapInsert(t *testing.T) {
	assert := assert.New(t)
	input := []int{12, 11, 13, 5, 6, 7}

	binheap := NewBinaryHeap(len(input))

	binheap.Initialize(input)

	expected := Heap{12, 11, 13, 5, 6, 7}
	assert.Equal(expected, *(binheap.GetHeap()), "Heap should match after initialization")

	binheap.HeapSort()
	expected = Heap{5, 6, 7, 11, 12, 13}
	assert.Equal(expected, *(binheap.GetHeap()), "Array should be sorted, after applying heapsort")

	//Swap out of index values
	input1 := []int{12, 11}
	binheap1 := NewBinaryHeap(len(input1))
	binheap1.Initialize(input1)
	//test out of index swaps
	outOfIndexErr := binheap1.Swap(12, 18)
	assert.NotNil(outOfIndexErr, "Should throw exception while trying to swap elements out of index")
}
