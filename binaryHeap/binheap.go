package heap

import (
	"errors"
	"fmt"
)

type Heap []int

var (
	EmptyHeapError = errors.New("Heap is empty, please use `Initialize`.")
)

//NewBinaryHeap creates a new binary heap and returns a pointer to the heap object
func NewBinaryHeap(numElements int) *Heap {
	heap := make(Heap, numElements)
	return &heap
}

func (heap *Heap) Initialize(input []int) {
	copy(*heap, input)
}

//GetHeapArr returns the heap object at any point in time during program execution.
func (heap *Heap) GetHeap() *Heap {
	return heap
}

//Swap swaps two elements of a given heap object
func (heap *Heap) Swap(idx1, idx2 int) error {
	if (idx1 < len(*heap)) && (idx2 < len(*heap)) {
		(*heap)[idx1], (*heap)[idx2] = (*heap)[idx2], (*heap)[idx1]
		return nil
	}
	return fmt.Errorf("Index out of range, received %d and %d", idx1, idx2)
}

//Heapify  min-max heapifies the current array.
//idx is the root node to apply heapify function to.
//maxLen is the subset of heap array, on which heapify should be applied to.
func (heap *Heap) Heapify(idx, maxLen int) {
	lchild := 2*idx + 1
	rchild := 2*idx + 2
	largest := idx
	heapLen := maxLen

	if (lchild < heapLen) && ((*heap)[lchild] > (*heap)[largest]) {
		largest = lchild
	}

	if (rchild < heapLen) && ((*heap)[rchild] > (*heap)[largest]) {
		largest = rchild
	}
	if largest != idx {
		heap.Swap(largest, idx)
		heap.Heapify(largest, heapLen)
	}
}

//HeapSort does a heap sort, who would say, right ? :P
func (heap *Heap) HeapSort() error {
	heapLen := len(*heap)

	for i := heapLen/2 - 1; i >= 0; i-- {
		heap.Heapify(i, len(*heap))
	}

	for i := len(*heap) - 1; i >= 0; i-- {
		heap.Swap(0, i)
		heap.Heapify(0, i)
	}
	return nil
}
