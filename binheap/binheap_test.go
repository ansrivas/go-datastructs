package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_HeapInsert(t *testing.T) {
	assert := assert.New(t)
	input := []int{12, 11, 13, 5, 6, 7}

	binheap := NewBinaryHeap(input)

	expected := Heap{12, 11, 13, 5, 6, 7}
	assert.Equal(expected, *(binheap.GetHeap()), "Heap should match after initialization")

	//Let's heapify it on 0 and get extractMax, should return 13
	binheap.Heapify(0, binheap.Len())
	maxval := binheap.ExtractMax()
	assert.Equal(13, maxval, "Extract min should work in constant time.")

	binheap.HeapSort()
	expected = Heap{5, 6, 7, 11, 12, 13}
	assert.Equal(expected, *(binheap.GetHeap()), "Array should be sorted, after applying heapsort")

	//Swap out of index values
	input1 := []int{12, 11}
	binheap1 := NewBinaryHeap(input1)
	//test out of index swaps
	outOfIndexErr := binheap1.Swap(12, 18)
	assert.NotNil(outOfIndexErr, "Should throw exception while trying to swap elements out of index")
}
