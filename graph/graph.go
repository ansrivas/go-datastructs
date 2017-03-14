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

//AddVertex adds a new node to a graph
func (g *Graph) AddVertex(id int) error {
	ok := g.Contains(id)
	if ok {
		return errors.New("Node already exists")
	}
	(*g)[id] = newNode(id)
	return nil
}

//ListVertices gives a list of all the vertices in this graph
func (g *Graph) ListVertices() []*GNode {
	nodes := make([]*GNode, 0)
	for k := range *g {
		nodes = append(nodes, (*g)[k])
	}
	return nodes
}

//ListVerticesID gives a list of all the verticx-id (unique identifier of a node) in this graph
func (g *Graph) ListVerticesID() []int {
	nodes := make([]int, 0)
	for k := range *g {
		nodes = append(nodes, (*g)[k].id)
	}
	return nodes
}
