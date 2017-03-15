package graph

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
