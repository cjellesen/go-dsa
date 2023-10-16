package stack

import "errors"

type node[T any] struct {
	previous_node *node[T]
	value         *T
}

type LinkedStack[T any] struct {
	length   int
	top_node *node[T]
}

// Creates a new Linked Stacks
func New[T any]() *LinkedStack[T] {
	return &LinkedStack[T]{top_node: nil, length: 0}
}

// Creates a new node with the provided value and assignes the current
// top nodes in the linked list as this nodes previous value. After the
// node has been created the current top node pointer in the linked list
// is replaced with a pointer to the newly created node. The generics in
// go apparently allow nil to match [T any] (and [T comparable]), as such
// there is nothing that prevent adding nil to the stack. Do not do this
// as a panic will be thrown.
func (ls *LinkedStack[T]) Push(value *T) {
	if value == nil {
		panic("Tried to add the value 'nil' to the stack")
	}

	top_node := &node[T]{previous_node: ls.top_node, value: value}
	ls.top_node = top_node
	ls.length++
}

// Graps the current top node of the linked list and replaces the top node
// pointer in the linked list with that of the previous node. Decrements
// the length of the Linked Stack and return the value to be poped.
func (ls *LinkedStack[T]) Pop() (*T, error) {
	if ls.length == 0 {
		return nil, errors.New("Tried to Pop from an empty stack")
	}

	pop := ls.top_node
	ls.top_node = pop.previous_node
	ls.length--
	return pop.value, nil
}

// Checks if the Linked Stack contains a top top node the return said node
// if there is one else it returns nil.
func (ls *LinkedStack[T]) Peek() (*T, error) {
	if ls.length == 0 {
		return nil, errors.New("Tried to Peek into an empty stack")
	}

	return ls.top_node.value, nil
}

// Checks if the Linked Stack is empty
func (ls *LinkedStack[T]) IsEmpty() bool {
	return ls.length == 0
}
