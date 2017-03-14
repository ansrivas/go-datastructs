package graph

import "fmt"

//NewGQueue creates a new queue
func NewGQueue() *GQueue {
	queue := make(GQueue, 0)
	return &queue
}

//Enqueue appends a new node at the end of the current queue
func (queue *GQueue) Enqueue(node *GNode) {
	*queue = append(*queue, node)
}

//Dqueue returns the node in the front of queue else error
func (queue *GQueue) Dqueue() (*GNode, error) {
	if len(*queue) == 0 {
		return nil, fmt.Errorf("Emtpy queue.")
	}

	top := (*queue)[0]
	*queue = (*queue)[1:]
	return top, nil
}

func (queue *GQueue) Len() int {
	return len(*queue)
}
