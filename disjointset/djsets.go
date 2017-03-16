package disjointset

import (
	"errors"
	"fmt"
	"sync"
)

var Lock = sync.RWMutex{}

type Element struct {
	id interface{}
}

type Set struct {
	rank   int
	leader *Element
}

type DisjointSet map[interface{}]*Set

//NewDisjointSet to make it more golangish :)
func NewDisjointSet() *DisjointSet {
	djset := make(DisjointSet)
	return &djset
}

// Makeset creates a set
func (dj *DisjointSet) MakeSet(id interface{}) {
	// Create an element and put itself as its parent
	element := &Element{id: id}
	(*dj)[id] = &Set{rank: 0, leader: element}
}

func (dj *DisjointSet) getNodeFromId(id interface{}) (node *Set, ok bool) {
	Lock.RLock()
	node, ok = (*dj)[id]
	Lock.RUnlock()
	return
}

//Find returns the leader of the current set
func (dj *DisjointSet) Find(id interface{}) (interface{}, error) {
	set, ok := dj.getNodeFromId(id)
	if !ok {
		return nil, fmt.Errorf("Unable to find element %v in the set", id)
	}
	return set.leader.id, nil
}

//Union creates a union of two sets.
//if the two sets have same rank, then id2 will be chosen as the leader always.
func (dj *DisjointSet) Union(id1, id2 interface{}) error {
	set1, ok1 := dj.getNodeFromId(id1)
	set2, ok2 := dj.getNodeFromId(id2)

	if !(ok1 && ok2) {
		return errors.New(fmt.Sprintf("Unable to find elements %v or %v in the set", id1, id2))
	}

	//Get rank of two sets take smaller one and make its parent point to the set with bigger rank (i.e. depth)
	if set1.rank > set2.rank {
		(*set2).leader = set1.leader
	} else {
		(*set1).leader = set2.leader
		if set1.rank == set2.rank {
			(*set2).rank += 1
			(*set1).rank += 1
		}
	}

	return nil
}
