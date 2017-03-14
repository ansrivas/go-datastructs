package graph

import (
	"errors"
	"fmt"
)

//NewGraph creates a new graph
func NewGraph() *Graph {
	graph := make(Graph)
	return &graph
}

//NewNode creates a new node in a given graph
func newNode(id int) *GNode {
	node := &GNode{id: id, neighbors: make([]*GNode, 0)}
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

//EdgeExists returns true if an edge exists from src, target node
func (g *Graph) EdgeExists(src, target int) (bool, error) {
	if !g.Contains(src) || !g.Contains(target) {
		return false, fmt.Errorf("Either of nodes %d or %d does not exist, please add it first", src, target)
	}

	for _, nodes := range (*g)[src].neighbors {
		if target == nodes.id {
			return true, nil
		}
	}
	return false, nil

}

//AddEdge creates an edge between two nodes represented by id1, id2.
//`directed`=true creates a directed graph from id1->id2
func (g *Graph) AddEdge(id1, id2 int, directed bool) error {
	if !g.Contains(id1) || !g.Contains(id2) {
		return fmt.Errorf("Either of nodes %d or %d does not exist, please add it first", id1, id2)
	}

	node1, _ := (*g)[id1]
	node2, _ := (*g)[id2]

	node1.neighbors = append(node1.neighbors, node2)
	//undirected graph case
	if !directed {
		node2.neighbors = append(node2.neighbors, node1)
	}
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
