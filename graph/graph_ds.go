package graph

import (
	"fmt"
	"sort"
)

func (g *Graph) topoLogicalSortUtil(id int, visited map[int]bool, stack *GStack) {
	node, _ := g.getNodeFromId(id)

	_, ok := visited[id]
	if ok {
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

func (g *Graph) BreadthFirstSearch(startNode int) ([]*GNode, error) {
	visited := make(map[int]bool)
	queue := NewGQueue()
	bfs := make([]*GNode, 0)
	node, ok := g.getNodeFromId(startNode)
	if !ok {
		return nil, fmt.Errorf("Start node could not be found: %d", startNode)
	}

	//Start with the first node i.e. startNode
	queue.Enqueue(node)

	for queue.Len() != 0 {
		top, _ := queue.Dqueue()
		visited[top.id] = true
		bfs = append(bfs, top)
		for _, neighbor := range top.neighbors {
			_, ok := visited[neighbor.id]
			if ok {
				continue
			}
			visited[neighbor.id] = true
			queue.Enqueue(neighbor)
		}
	}
	return bfs, nil
}
