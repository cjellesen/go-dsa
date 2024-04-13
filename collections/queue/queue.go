package queue

import (
	"errors"

	"github.com/cjellesen/go-dsa/collections/linkedlist"
)

type Queue[T comparable] struct {
	linkedlist linkedlist.LinkedList[T]
}

func (q *Queue[T]) Enqueue(value T) {
	q.linkedlist.AddTail(value)
}

func (q *Queue[T]) Dequeue() (T, error) {
	r, err := q.Peek()

	if err != nil {
		return r, err
	}

	q.linkedlist.RemoveHead()
	return r, nil
}

func (q *Queue[T]) IsEmpty() bool {
	return q.linkedlist.Count() == 00
}

func (q *Queue[T]) Peek() (T, error) {
	var r T
	if q.IsEmpty() {
		return r, errors.New("tried to peek into an empty stack")
	}

	return q.linkedlist.Head.Value, nil
}
