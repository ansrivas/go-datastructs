package graph

import "fmt"

//NewGStack creates a new stack which stores elements of type GNode
func NewGStack() *GStack {
	stack := make(GStack, 0)
	return &stack
}

//Push pushes a pointer to type GNode in the current stack
func (stack *GStack) Push(node *GNode) {
	*stack = append(*stack, node)
}

//Pop returns the node as LIFO, error if empty
func (stack *GStack) Pop() (*GNode, error) {
	length := stack.Len()
	if length == 0 {
		return nil, fmt.Errorf("Emtpy stack.")
	}
	node := (*stack)[length-1]

	(*stack) = (*stack)[0 : length-1]
	return node, nil
}

//Len returns the length of a stack
func (stack *GStack) Len() int {
	return len(*stack)
}
