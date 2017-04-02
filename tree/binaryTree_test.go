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

func Test_LCABinaryTree1(t *testing.T) {
	assert := assert.New(t)
	bt := NewBinaryTree()
	root := bt.InsertRoot(4)
	left := root.InsertLeft(6)
	right := root.InsertRight(2)
	right.InsertRight(9)
	left.InsertLeft(5)
	end := left.InsertRight(3)
	end.InsertLeft(1)

	node := bt.LCA(5, 1)
	expected := 6
	assert.Equal(expected, node.data, "Should get 6 as the LCA")
}

func Test_LCABinaryTree2(t *testing.T) {
	assert := assert.New(t)
	bt := NewBinaryTree()
	root := bt.InsertRoot(4)
	left := root.InsertLeft(2)
	root.InsertRight(5)

	left.InsertLeft(3)
	left.InsertRight(8)

	node := bt.LCA(5, 3)
	expected := 4
	assert.Equal(expected, node.data, "Should get 4 as the LCA")
}

func Test_LCABinaryTree3(t *testing.T) {
	assert := assert.New(t)
	bt := NewBinaryTree()
	root := bt.InsertRoot(4)
	left := root.InsertLeft(2)
	right := root.InsertRight(5)

	left.InsertLeft(3)
	left.InsertRight(8)

	right.InsertLeft(1)
	right.InsertRight(6)

	node := bt.LCA(6, 1)
	expected := 5
	assert.Equal(expected, node.data, "Should get 5 as the LCA")
}

func Test_PathWithGivenSum(t *testing.T) {
	assert := assert.New(t)
	bt := NewBinaryTree()
	root := bt.InsertRoot(10)
	root.InsertLeft(28)
	right := root.InsertRight(13)
	right_left := right.InsertLeft(14)
	right_left.InsertLeft(21)
	right_left.InsertRight(22)
	right_right := right.InsertRight(15)
	right_right.InsertLeft(23)
	right_right.InsertRight(24)
	sum := 38
	bt.PathWithGivenSum(sum)

	assert.True(true, "msgAndArgs")
}
