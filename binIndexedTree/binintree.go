package binIndexedTree

// Space complexity: O(n)
// To search: 0(logn)
// To update elements : O(logn)
// To create for the first time: O(nlogn)

type BinIndTree []int

//NewBinaryIndexedTree returns a new binary indexed tree
// This function can be optimized by passing the length of input array
func NewBinaryIndexedTree(inputLength int) *BinIndTree {
	bit := make(BinIndTree, inputLength+1)
	return &bit
}

//getParent will get the parent of current index
//Currently calculated by taking 2's compliment of a number, and'ing it with the number
//and then subtracting it from actual number.
// Basically trying to find the rightmost setbit and flipping it.
//For 2's compliment either use `-num` or take 1's compliment(~ operator) and add 1 to it
func GetParent(input int) int {
	return input - (input & -input)
}

//GetNext returns the next node in the tree, which is affected whenever there is an update
//in a given node. Similar explanation as above.
func GetNext(input int) int {
	return input + (input & -input)
}

//InsertTree will insert a node into tree ( takes in total nlogn time to build it up)
func (bit *BinIndTree) InsertTree(index, val int) {
	next := index + 1
	for {
		(*bit)[next] += val
		next = GetNext(next)
		if next >= len(*bit) {
			break
		}
	}
}

//UpdateTree will update a value in the given tree, with some other value
//To update a value, give the index and a number with sign which would represent an add or sum
//For eg. to update an index with +3 pass 3, with -3 pass -3
//from current example, to update value 6 at index 3, to new value 9, "3" should be passed as updatedVal
//O(logn) time to update the elements
func (bit *BinIndTree) UpdateTree(index, update int) {
	next := index + 1
	for {
		(*bit)[next] += update
		next = GetNext(next)
		if next >= len(*bit) {
			break
		}
	}
}

//GetSum returns the sum of a prefix
//For eg. to get sum between 0,5 GetSum(5) needs to be called.
func (bit *BinIndTree) GetSum(index int) int {
	next := index + 1
	value := 0
	for {
		if next == 0 {
			break
		}
		value += (*bit)[next]
		next = GetParent(next)
	}
	return value
}

///GetTreeRepresentation returns the state of tree at any given time.
//Index 0 is always 0 as it stores everything as 1 index based.
func (bit *BinIndTree) GetTreeRepresentation() *BinIndTree {
	return bit
}
