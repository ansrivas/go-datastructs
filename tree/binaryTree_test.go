package tree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_InsertRoot(t *testing.T) {
	assert := assert.New(t)
	bt := NewBinaryTree()
	root := bt.InsertRoot(4)
	assert.Equal(bt.root.data, 4, "Insert into root node should be successful")
	root.InsertLeft(3)
	root.InsertRight(5)
	assert.Equal(bt.root.left.data, 3, "Insert into root.left node should be 3")
	assert.Equal(bt.root.right.data, 5, "Insert into root.right node should be 5")

	var testNode *BTNode
	insertLeft := testNode.InsertLeft(4)
	insertRight := testNode.InsertRight(4)
	assert.Nil(insertLeft, "Insert into an uninitialized node should return nil")
	assert.Nil(insertRight, "Insert into an uninitialized node should return nil")

}

func Test_BinaryTree(t *testing.T) {
	assert := assert.New(t)

	bt := NewBinaryTree()
	root := bt.InsertRoot(4)
	root.InsertLeft(2)
	left := root.InsertLeft(2)
	root.InsertRight(5)
	left.InsertLeft(3)
	left.InsertRight(8)
	fmt.Println(bt.IsBST())

	assert.False(bt.IsBST(), "This binary tree is not a BST")
}

func Test_IsBST(t *testing.T) {
	assert := assert.New(t)

	bt := NewBinaryTree()
	root := bt.InsertRoot(10)
	left := root.InsertLeft(10)
	left.InsertLeft(-5)
	right := root.InsertRight(19)
	right.InsertLeft(17)
	right.InsertRight(21)
	fmt.Println(bt.IsBST())

	assert.True(bt.IsBST(), "This binary tree is a BST")
}
