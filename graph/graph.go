package graph

import "errors"

//NewGraph creates a new graph
func NewGraph() *Graph {
	graph := make(Graph)
	return &graph
}

//NewNode creates a new node in a given graph
func newNode(id int) *GNode {
	node := &GNode{id: id, edges: make([]*GNode, 0)}
	return node
}

//Contains checks if a node with given `id` exists in the graph
func (g *Graph) Contains(id int) bool {
	_, ok := (*g)[id]
	return ok
}

//AddNode adds a new node to a graph
func (g *Graph) AddNode(id int) error {
	ok := g.Contains(id)
	if ok {
		return errors.New("Node already exists")
	}
	(*g)[id] = newNode(id)
	return nil
}
