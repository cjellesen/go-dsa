package collections

import (
	"fmt"
)

type CircularQueue struct {
	Type CollectionType
}

func (cq CircularQueue) Print() {
	fmt.Println("I am a: ", cq.Type)
}

func NewCircularQueue() *CircularQueue {
	return &CircularQueue{Type: CircularQueueEnum}
}
