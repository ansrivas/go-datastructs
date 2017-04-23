package commonprobs

import (
	"errors"
	"math/rand"
	"time"
)

type ConstantDS struct {
	arr  []int
	dict map[int]int
}

//NewConstantDS creates a new DS which returns each and every operation in constant time
func NewConstantDS() *ConstantDS {
	return &ConstantDS{
		arr:  make([]int, 0),
		dict: make(map[int]int)}
}

//Insert inserts a value in our constant time DS
func (ds *ConstantDS) Insert(input int) {
	_, ok := ds.dict[input]
	if !ok {
		idx := len(ds.arr)
		(*ds).arr = append((*ds).arr, input)
		ds.dict[input] = idx
	}
}

//Find returns true or false if a given input is found in the given DS
func (ds *ConstantDS) Find(input int) bool {
	_, ok := ds.dict[input]
	return ok
}

//Delete the given value in constant time, if not present return error
func (ds *ConstantDS) Delete(input int) (bool, error) {
	idx, ok := ds.dict[input]
	if !ok {
		return false, errors.New("Element not found")
	}

	(*ds).arr[idx], (*ds).arr[len(ds.arr)-1] = (*ds).arr[len(ds.arr)-1], (*ds).arr[idx]
	(*ds).arr = (*ds).arr[:len(ds.arr)-1]

	ds.dict[ds.arr[idx]] = idx
	delete(ds.dict, input)
	return true, nil
}

//Random returns a ranom selected from this DS
func (ds *ConstantDS) Random() int {
	rand.Seed(time.Now().Unix())
	idx := rand.Intn(len(ds.arr))
	return ds.arr[idx]
}
