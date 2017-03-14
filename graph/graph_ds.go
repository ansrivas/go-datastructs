package graph

import (
	"fmt"
	"sort"
)

// func (g *Graph) topoLogicalSortUtil(id int, visited map[int]bool, topoSortedList *[]int) {
func (g *Graph) topoLogicalSortUtil(id int, visited map[int]bool, stack *GStack) {
	fmt.Println("Fetching node: ", id)
	node, _ := g.getNodeFromId(id)
	fmt.Println("Current status of visited list: ", visited)

	_, ok := visited[id]
	if ok {
		fmt.Println("Node id was visited: ", id)
		return
	}

	// Mark this visited

	visited[id] = true

	// If this is a leaf node, just add it to topoSortedList and return
	neighbors := node.neighbors
	if len(neighbors) == 0 {
		stack.Push(node)
		return
	}

	//If this is not a leaf node just iterate over its neighbors
	for _, neighbor := range node.neighbors {
		g.topoLogicalSortUtil(neighbor.id, visited, stack)
	}

	//Done processing the children ? add it to topoSortedList.
	stack.Push(node)

}

//TopologicalSort does a TopologicalSorting of a given graph
// sortInput decides if the input provided to sorting method always processes the inputs (for testing purposes only)
func (g *Graph) TopologicalSort(sortInput ...bool) *GStack {
	visited := make(map[int]bool)
	stack := NewGStack()

	vertices := g.ListVerticesID()

	//Check if we need sorted inputs for testing purposes
	if len(sortInput) != 0 {
		sortedVertices := vertices
		sort.Ints(sortedVertices)
	}

	for _, vertex := range vertices {
		g.topoLogicalSortUtil(vertex, visited, stack)
	}
	return stack
}
