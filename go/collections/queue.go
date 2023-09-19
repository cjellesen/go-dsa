package collections

import (
	"fmt"
)

type Queue struct {
	Type CollectionType
}

func (q Queue) Print() {
	fmt.Println("I am a: ", q.Type)
}

func NewQueue() *Queue {
	return &Queue{Type: QueueEnum}
}
