package stack

import (
	"collections/linkedlist"
	"errors"
)

type Stack[T comparable] struct {
	LinkedList linkedlist.LinkedList[T]
}

// Creates a new node with the provided value and assignes the current
// top nodes in the linked list as this nodes previous value. After the
// node has been created the current top node pointer in the linked list
// is replaced with a pointer to the newly created node. The generics in
// go apparently allow nil to match [T any] (and [T comparable]), as such
// there is nothing that prevent adding nil to the stack. Do not do this
// as a panic will be thrown.
func (s *Stack[T]) Push(value T) {
	s.LinkedList.AddHead(value)
}

// Graps the current top node of the linked list and replaces the top node
// pointer in the linked list with that of the previous node. Decrements
// the length of the Linked Stack and return the value to be poped.
func (s *Stack[T]) Pop() (T, error) {
	r, err := s.Peek()

	if err != nil {
		return r, err
	}

	s.LinkedList.RemoveHead()
	return r, nil
}

// Checks if the Linked Stack contains a top top node the return said node
// if there is one else it returns nil.
func (s *Stack[T]) Peek() (T, error) {
	var r T
	if s.IsEmpty() {
		return r, errors.New("tried to peek into an empty stack")
	}

	return s.LinkedList.Head.Value, nil
}

// Checks if the Linked Stack is empty
func (s *Stack[T]) IsEmpty() bool {
	return s.LinkedList.Count == 0
}
