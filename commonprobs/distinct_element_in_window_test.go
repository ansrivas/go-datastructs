package commonprobs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewDistinctWindowed(t *testing.T) {
	assert := assert.New(t)

	input := []int{2, 1, 1, 3, 5, 2, 6, 6, 7}
	expected := []int{2, 2, 3, 3, 3, 2, 2}
	windowLen := 3

	ds := NewDistinctWindowed(windowLen, input)
	nodes, _ := ds.GetDistElemInWindow()
	assert.Equal(expected, nodes, "Count of distinct elements should be same.")

	input = []int{2, 1}
	ds = NewDistinctWindowed(windowLen, input)
	_, err := ds.GetDistElemInWindow()
	assert.NotNil(err, "Passing window length larger than input length should return error")
	assert.True(true, "true")
}
