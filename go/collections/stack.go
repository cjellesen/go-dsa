package collections

import (
	"fmt"
)

type Stack struct {
	Type CollectionType
}

func (s Stack) Print() {
	fmt.Println("I am a: ", s.Type)
}

func NewStack() *Stack {
	return &Stack{Type: StackEnum}
}
