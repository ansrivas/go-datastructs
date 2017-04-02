package tree

import (
	"fmt"
	"math"
)

//Node represents a node in a binary tree
type BTNode struct {
	data  int
	left  *BTNode
	right *BTNode
}

type BinaryTree struct {
	root *BTNode
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{}
}

//NewNode creates a new node to be inserted in bst
func NewBTNode(data int) *BTNode {
	return &BTNode{
		data:  data,
		left:  nil,
		right: nil,
	}
}

func (bt *BinaryTree) InsertRoot(input int) *BTNode {
	(*bt).root = NewBTNode(input)
	return bt.root
}

//InsertLeft inserts a left child to this given node
func (node *BTNode) InsertLeft(input int) *BTNode {
	if node != nil {
		(*node).left = NewBTNode(input)
		return node.left
	}
	return nil
}

//InsertRight inserts a right child to this given node
func (node *BTNode) InsertRight(input int) *BTNode {
	if node != nil {
		(*node).right = NewBTNode(input)
		return node.right
	}
	return nil
}

func (node *BTNode) IsBSTUtil(min, max int) bool {
	if node.data <= min || node.data > max {
		return false
	}
	lsubtree := true
	rsubtree := true
	if node.left != nil {
		lsubtree = node.left.IsBSTUtil(min, node.data)
	}
	if node.right != nil {
		rsubtree = node.right.IsBSTUtil(node.data, max)
	}
	return lsubtree && rsubtree
}

//IsBST returns true/false if a given binary tree is BST.
func (bt *BinaryTree) IsBST() bool {
	min, max := math.MinInt16, math.MaxInt16
	return bt.root.IsBSTUtil(min, max)
}

func (node *BTNode) LCAUtil(inp1, inp2 int) *BTNode {

	if (node.data == inp1) || (node.data == inp2) {
		return node
	}
	var ltree, rtree *BTNode

	if node.left != nil {
		ltree = node.left.LCAUtil(inp1, inp2)
	}

	if node.right != nil {
		rtree = node.right.LCAUtil(inp1, inp2)
	}

	if rtree != nil && ltree != nil {
		return node
	}

	if ltree != nil {
		return ltree
	}
	return rtree
}

//LCA returns LCA of given two nodes
func (bt *BinaryTree) LCA(inp1, inp2 int) *BTNode {
	return bt.root.LCAUtil(inp1, inp2)
}

func (node *BTNode) PathWithGivenSumUtil(sumSoFar, sum int, path *([]int)) {

	sumSoFar += node.data
	(*path) = append((*path), node.data)
	if sumSoFar == sum {
		//TODO: Store path in a slice of slice or map of slices
		fmt.Println("Found path: ", *path)
	}
	if node.left != nil {
		node.left.PathWithGivenSumUtil(sumSoFar, sum, path)
	}
	if node.right != nil {
		node.right.PathWithGivenSumUtil(sumSoFar, sum, path)
	}
	*path = (*path)[0 : len(*path)-1]
}

//PathWithGivenSum will return a path with a given sum from ROOT -> any child ( not necessarily leaf )
func (bt *BinaryTree) PathWithGivenSum(sum int) {
	path := make([]int, 0)
	sumSoFar := 0
	bt.root.PathWithGivenSumUtil(sumSoFar, sum, &path)
}
