package tree

import (
	"errors"
	"fmt"
)

//Node represents a node in a binary tree
type Node struct {
	data  int
	left  *Node
	right *Node
}

//BinarySearchTree is the initial point for storing BST
type BinarySearchTree struct {
	root *Node
}

var RootNilError = errors.New("Root is nil")

//Create a new BinarySearchTree representation
func NewBST() *BinarySearchTree {
	bst := BinarySearchTree{}
	return &bst
}

//NewNode creates a new node to be inserted in bst
func NewNode(data int) *Node {
	return &Node{
		data:  data,
		left:  nil,
		right: nil,
	}
}

func (node *Node) insertUtil(data int) {
	if data < node.data {
		if node.left == nil {
			node.left = NewNode(data)
			return
		}
		node.left.insertUtil(data)
	} else {
		if node.right == nil {
			node.right = NewNode(data)
			return
		}
		node.right.insertUtil(data)
	}
}

//Insert a node in a binary tree
func (bst *BinarySearchTree) Insert(data int) {
	if bst.root == nil {
		bst.root = NewNode(data)
	} else {
		bst.root.insertUtil(data)
	}
}

func (node *Node) inorderUtil() {
	if node.left != nil {
		node.left.inorderUtil()
	}
	fmt.Println("Current Node is: ", node.data)
	if node.right != nil {
		node.right.inorderUtil()
	}
}

//Inorder traversal of bst
func (bst *BinarySearchTree) Inorder() error {
	if bst.root == nil {
		return RootNilError
	}
	bst.root.inorderUtil()
	return nil
}

//Max returns max of two ints
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (node *Node) heightUtil() int {
	if node.left == nil && node.right == nil {
		return 1
	}

	lh := 0
	rh := 0
	if node.left != nil {
		lh = node.left.heightUtil()
	}
	if node.right != nil {
		rh = node.right.heightUtil()
	}

	return 1 + Max(lh, rh)
}

//Height gets the height of a tree, only root ==1
func (bst *BinarySearchTree) Height() (int, error) {
	if bst.root == nil {
		return 0, RootNilError
	}
	return bst.root.heightUtil(), nil
}

func (node *Node) diameterUtil() int {

	lh := 0
	rh := 0
	if node.left != nil {
		lh = node.left.heightUtil()
	}
	if node.right != nil {
		rh = node.right.heightUtil()
	}

	ldiam := 0
	rdiam := 0
	if node.left != nil {
		ldiam = node.left.diameterUtil()
	}
	if node.right != nil {
		rdiam = node.right.diameterUtil()
	}

	return Max(Max(ldiam, rdiam), lh+rh+1)
}

//Diameter returns diameter of the tree
func (bst *BinarySearchTree) Diameter() (int, error) {
	if bst.root == nil {
		return 0, RootNilError
	}
	return bst.root.diameterUtil(), nil
}

func (node *Node) sizeUtil() int {

	lsize := 0
	rsize := 0

	if node.left != nil {
		lsize = node.left.sizeUtil()
	}

	if node.right != nil {
		rsize = node.right.sizeUtil()
	}

	return lsize + rsize + 1
}

//Size returns the total number of nodes in a given binary tree.
func (bst *BinarySearchTree) Size() (int, error) {
	if bst.root == nil {
		return 0, RootNilError
	}

	return bst.root.sizeUtil(), nil
}

// func (node *Node) rootToLeafSum(path []int, sum int) ([]int, bool) {
//
// 	fmt.Println("node.data", node.data, sum)
// 	if sum == 0 && node.left == nil && node.right == nil {
// 		path := append(path, node.data)
// 		fmt.Println("path, ", path)
// 		return path, true
// 	}
// 	var rval, lval bool
// 	if node.left != nil {
// 		fmt.Println("Left leaf", path, sum-node.data)
// 		path, lval = node.left.rootToLeafSum(path, sum-node.data)
// 	}
//
// 	if node.right != nil {
// 		path, rval = node.right.rootToLeafSum(path, sum-node.data)
// 	}
// 	fmt.Println("path and rval", path, lval && rval)
// 	return path, lval && rval
// }

//RootToLeafSum returns path,true/false if a given sum exists in a tree
// func (bst *BinarySearchTree) RootToLeafSum(sum int) ([]int, bool) {
//
// 	path := make([]int, 0)
// 	return bst.root.rootToLeafSum(path, sum)
//
// }
