package tree

import "math"

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

func (node *BTNode) IsBST(min, max int) bool {
	if node.data <= min || node.data > max {
		return false
	}
	lsubtree := true
	rsubtree := true
	if node.left != nil {
		lsubtree = node.left.IsBST(min, node.data)
	}
	if node.right != nil {
		rsubtree = node.right.IsBST(node.data, max)
	}
	return lsubtree && rsubtree
}

//IsBST returns true/false if a given binary tree is BST.
func (bt *BinaryTree) IsBST() bool {
	min, max := math.MinInt16, math.MaxInt16
	return bt.root.IsBST(min, max)
}
