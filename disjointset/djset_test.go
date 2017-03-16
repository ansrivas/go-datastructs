package disjointset

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DisjointSet(t *testing.T) {
	assert := assert.New(t)

	djset := NewDisjointSet()

	djset.MakeSet(1)
	djset.MakeSet(2)

	leader, _ := djset.Find(2)
	assert.Equal(2, leader, "2's leader should be 2")
	leader, _ = djset.Find(1)
	assert.Equal(1, leader, "1's leader should be 1")

	_, err := djset.Find(4)
	assert.NotNil(err, "Non existing elements should throw error")

	//Taking union of 1,2
	djset.Union(1, 2)
	leader1, _ := djset.Find(2)
	leader2, _ := djset.Find(1)
	assert.Equal(leader1, leader2, "Leader for both the sets should be same")

	err = djset.Union(1, 4)
	assert.NotNil(err, "Non existent nodes should throw error")

	//Now joining two sets
	djset.MakeSet(3)
	djset.MakeSet(5)
	djset.MakeSet(6)
	djset.Union(3, 5)
	djset.Union(3, 6)

	leader, _ = djset.Find(6)
	assert.Equal(leader, 5, "Leader should be  5")

	//In case of equal rank, the second element's leader will become the global leader
	djset.Union(1, 3)
	leader, _ = djset.Find(1)
	assert.Equal(5, leader, "Leader should be  5")

}
