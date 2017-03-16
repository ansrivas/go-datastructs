package disjointsets

type Element struct {
	id interface{}
}

type Set struct {
	rank   int
	leader *Element
}

type DisjointSet map[interface{}]*Set
