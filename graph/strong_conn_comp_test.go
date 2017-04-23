package graph

//
// import (
// 	"testing"
//
// 	"github.com/stretchr/testify/assert"
// )
//
// func Test_StronglyConnComp(t *testing.T) {
// 	assert := assert.New(t)
//
// 	graph := NewGraph()
//
// 	for i := 1; i < 11; i++ {
// 		graph.AddVertex(i)
// 	}
// 	graph.AddEdge(2, 4, true, 0)
// 	graph.AddEdge(2, 3, true, 0)
// 	graph.AddEdge(3, 1, true, 0)
// 	graph.AddEdge(1, 2, true, 0)
//
// 	graph.AddEdge(4, 5, true, 0)
// 	graph.AddEdge(5, 6, true, 0)
// 	graph.AddEdge(6, 4, true, 0)
//
// 	graph.AddEdge(7, 6, true, 0)
// 	graph.AddEdge(7, 8, true, 0)
//
// 	graph.AddEdge(8, 9, true, 0)
// 	graph.AddEdge(9, 10, true, 0)
// 	graph.AddEdge(10, 7, true, 0)
//
// 	graph.AddEdge(10, 11, true, 0)
//
//     firstComponent := []int{1,2,3}
// 	expected := []int{}
// 	assert.Equal(expected, actual, msgAndArgs)
// }
