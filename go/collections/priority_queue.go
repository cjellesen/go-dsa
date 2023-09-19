package collections

import (
	"fmt"
)

type PriorityQueue struct {
	Type CollectionType
}

func (pq PriorityQueue) Print() {
	fmt.Println("I am a: ", pq.Type)
}

func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{Type: PriorityQueueEnum}
}
