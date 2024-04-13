package stack

import (
	"errors"

	"github.com/cjellesen/go-dsa/collections/linkedlist"
)

type Stack[T comparable] struct {
	linkedList linkedlist.LinkedList[T]
}

func (s *Stack[T]) Push(value T) {
	s.linkedList.AddHead(value)
}

func (s *Stack[T]) Pop() (T, error) {
	r, err := s.Peek()

	if err != nil {
		return r, err
	}

	s.linkedList.RemoveHead()
	return r, nil
}

func (s *Stack[T]) Peek() (T, error) {
	var r T
	if s.IsEmpty() {
		return r, errors.New("tried to peek into an empty stack")
	}

	return s.linkedList.Head.Value, nil
}

func (s *Stack[T]) IsEmpty() bool {
	return s.linkedList.Count() == 0
}
