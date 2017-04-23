package graph

import (
	"fmt"
	"sync"
)

var Lock = sync.RWMutex{}

//GQueue is a queue which holds GNode
type GQueue []*GNode

//GStack stores pointers to GNode
type GStack []*GNode

//GNode represent a node in a graph
type GNode struct {
	id        int
	neighbors []*GNode
	edges     []*Edge
}

//Edge stores edge information in a node
type Edge struct {
	source *GNode
	target *GNode
	weight int
}

//Graph is a map of nodes
type Graph map[int]*GNode

//NewGraph creates a new graph
func NewGraph() *Graph {
	graph := make(Graph)
	return &graph
}

//NewNode creates a new node in a given graph
func newNode(id int) *GNode {
	node := &GNode{
		id:        id,
		neighbors: make([]*GNode, 0),
		edges:     make([]*Edge, 0),
	}
	return node
}

//NewNode creates a new node in a given graph
func newEdge(source, target *GNode, weight int) *Edge {
	edge := &Edge{
		source: source,
		target: target,
		weight: weight,
	}
	return edge
}

func (g *Graph) getNodeFromId(id int) (node *GNode, ok bool) {
	Lock.RLock()
	node, ok = (*g)[id]
	Lock.RUnlock()
	return
}

//Contains checks if a node with given `id` exists in the graph
func (g *Graph) Contains(id int) bool {
	_, ok := g.getNodeFromId(id)
	return ok
}

//AddVertex adds a new node to a graph
func (g *Graph) AddVertex(id int) error {
	ok := g.Contains(id)
	if ok {
		return fmt.Errorf("Node already exists")
	}
	Lock.Lock()
	defer Lock.Unlock()
	(*g)[id] = newNode(id)
	return nil
}

//EdgeExists returns true if an edge exists from src, target node
func (g *Graph) EdgeExists(src, target int) (bool, error) {
	if !g.Contains(src) || !g.Contains(target) {
		return false, fmt.Errorf("Either of nodes %d or %d does not exist, please add it first", src, target)
	}

	node, ok := g.getNodeFromId(src)
	if ok {
		for _, nodes := range node.neighbors {
			if target == nodes.id {
				return true, nil
			}
		}
	}
	return false, nil

}

//AddEdge creates an edge between two nodes represented by id1, id2.
//`directed`=true creates a directed graph from id1->id2
func (g *Graph) AddEdge(id1, id2 int, directed bool, weight int) error {
	if !g.Contains(id1) || !g.Contains(id2) {
		return fmt.Errorf("Either of nodes %d or %d does not exist, please add it first", id1, id2)
	}

	node1, _ := g.getNodeFromId(id1)
	node2, _ := g.getNodeFromId(id2)

	node1.neighbors = append(node1.neighbors, node2)
	node1.edges = append(node1.edges, newEdge(node1, node2, weight))
	//undirected graph case
	if !directed {
		node2.neighbors = append(node2.neighbors, node1)
		node2.edges = append(node2.edges, newEdge(node2, node1, weight))
	}
	return nil
}

//GetVertices gives a list of all the vertices in this graph
func (g *Graph) GetVertices() []*GNode {
	nodes := make([]*GNode, 0)
	for k := range *g {
		vertex, ok := g.getNodeFromId(k)
		if ok {
			nodes = append(nodes, vertex)
		}
	}
	return nodes
}

//TODO: Get ride of this function later
//GetVerticesID gives a list of all the verticx-id (unique identifier of a node) in this graph
func (g *Graph) GetVerticesID() []int {
	nodes := make([]int, 0)
	for k := range *g {
		vertex, ok := g.getNodeFromId(k)
		if ok {
			nodes = append(nodes, vertex.id)
		}
	}
	return nodes
}

//GetNeighbors returns neighbors of a given node with identifier=id
func (g *Graph) GetNeighbors(id int) ([]*GNode, error) {
	node, ok := g.getNodeFromId(id)
	if !ok {
		return nil, fmt.Errorf("Node %d does not exist", id)
	}
	return node.neighbors, nil
}

//GetEdges returns the edges of a given node with identifier=id
func (g *Graph) GetEdges(id int) ([]*Edge, error) {
	node, ok := g.getNodeFromId(id)
	if !ok {
		return nil, fmt.Errorf("Node %d does not exist", id)
	}
	return node.edges, nil
}

//RemoveEdge removes an edge between two nodes, id1, id2.
//If `directed=false` both the edges will be removed.
func (g *Graph) RemoveEdge(id1, id2 int, directed bool) error {
	if !g.Contains(id1) || !g.Contains(id2) {
		return fmt.Errorf("Either of nodes %d or %d does not exist, please add it first", id1, id2)
	}

	node1, _ := g.getNodeFromId(id1)
	node2, _ := g.getNodeFromId(id2)

	//Removed the edge from node1.
	for idx, ed := range node1.edges {
		if ed.target.id == id2 {
			(node1.edges) = append(node1.edges[0:idx], node1.edges[idx+1:]...)
			break
		}
	}

	if !directed {
		for idx, ed := range node2.edges {
			if ed.target.id == id1 {
				(node2.edges) = append(node2.edges[0:idx], node2.edges[idx+1:]...)
				break
			}
		}
	}

	return nil
}
