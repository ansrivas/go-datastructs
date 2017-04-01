package commonProbs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ConstantTimeOps(t *testing.T) {
	//Should be able to do following operations in constant time.
	//Insert, Delete, Find, Get Random.
	assert := assert.New(t)

	ds := NewConstantDS()
	ds.Insert(1)
	ds.Insert(5)
	ds.Insert(3)

	assert.True(ds.Find(3), "3 should exist in DS")
	assert.True(ds.Find(1), "1 should exist in DS")
	assert.True(ds.Find(5), "5 should exist in DS")

	ds.Delete(5)
	assert.False(ds.Find(5), "After deletion 5 should not exist in here")

	_, err := ds.Delete(5)
	assert.NotNil(err, "Deletion of non-existing element should be null")

	val := ds.Random()
	expectedRange := []int{1, 3}
	var found bool
	for _, ex := range expectedRange {
		if ex == val {
			found = true
		}
	}
	assert.True(found, "Random should be returned from the inserted list")
}
