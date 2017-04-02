package trie

import "errors"

type TrieNode struct {
	children  map[rune]*TrieNode
	isEndNode bool
}

type Trie struct {
	root *TrieNode
}

var (
	EmptyStringError = errors.New("Empty string inserted.")
	EmptyRootError   = errors.New("Trie has not been initialized.")
)

//NewNode creates a new TrieNode
func NewNode(isEndNode bool) *TrieNode {
	node := TrieNode{children: make(map[rune]*TrieNode), isEndNode: isEndNode}
	return &node
}

//NewTrie creates a new trie.
func NewTrie() *Trie {
	root := Trie{root: NewNode(false)}
	return &root
}

//Insert a string into trie
func (trie *Trie) Insert(input string) error {
	if input == "" {
		return EmptyStringError
	}
	current := trie.root
	var newNode *TrieNode
	for _, char := range input {
		node, ok := current.children[char]
		if ok {
			current = node
		} else {
			newNode = NewNode(false)
			current.children[char] = newNode
			current = newNode
		}
	}
	// Mark last node as true
	current.isEndNode = true
	return nil
}

//StartsWith checks if there is any input in trie which starts with given start string.
func (trie *Trie) Search(input string) (bool, error) {
	if input == "" {
		return false, EmptyStringError
	}

	current := trie.root

	for _, char := range input {
		next, ok := current.children[char]
		if !ok {
			return false, nil
		}
		current = next
	}
	return current.isEndNode, nil
}

//TODO: Implement delete function

// func deleteUtil(node *TrieNode)  {
// 	if node.
// }
//
// func (trie *Trie) Delete(input string) (bool, error) {
// 	current := trie.root
// 	deleteUtil(current, )
// }
