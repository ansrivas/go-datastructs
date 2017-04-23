package commonprobs

import "errors"

// You are given an array of size n.
// Now print the number of distinct elements in the array for every frame of size k.

type DistinctWindowedArr struct {
	windowLen int
	input     []int
	dict      map[int]int
}

//NewDistinctWindowed returns a given window len + inputy values
func NewDistinctWindowed(windowLen int, input []int) *DistinctWindowedArr {
	return &DistinctWindowedArr{
		windowLen: windowLen,
		input:     input,
		dict:      make(map[int]int),
	}
}

func (ds *DistinctWindowedArr) GetDistElemInWindow() ([]int, error) {
	length := len(ds.input)
	output := make([]int, 0)
	if length < ds.windowLen {
		return nil, errors.New("Window is larger than actual array")
	}

	startIdx := 0
	//For the first initialization
	for i := startIdx; i < ds.windowLen+startIdx; i++ {
		key := ds.input[i]
		node, ok := (*ds).dict[key]
		if ok {
			(*ds).dict[key] = node + 1
		} else {
			(*ds).dict[key] = 1
		}
	}

	for {
		output = append(output, len(ds.dict))

		//Case if the previous value occurs more than one time.
		key := ds.input[startIdx]
		node, _ := (*ds).dict[key]
		if node == 1 {
			delete((*ds).dict, key)
		} else {
			(*ds).dict[key] = node - 1
		}

		if startIdx+ds.windowLen < length {
			//case where next value is already present or needs to be inserted.
			nextKey := ds.input[startIdx+ds.windowLen]
			node, ok := (*ds).dict[nextKey]
			if ok {
				(*ds).dict[nextKey] = node + 1
			} else {
				(*ds).dict[nextKey] = 1
			}
			startIdx++
		} else {
			break
		}
	}

	return output, nil
}
