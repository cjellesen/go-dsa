package linked_stack

import "fmt"

type node struct {
	previous_node *node
	value         interface{}
}

type LinkedStack struct {
	length   int
	top_node *node
}

// Creates a new Linked Stacks
func New() *LinkedStack {
	return &LinkedStack{top_node: nil, length: 0}
}

func NewFrom(value interface{}) *LinkedStack {
	ls := &LinkedStack{top_node: nil, length: 0}
	ls.Push(value)
	return ls
}

func PrintMe() {
	fmt.Println("Test")
}

// Creates a new node with the provided value and assignes the current
// top nodes in the linked list as this nodes previous value. After the
// node has been created the current top node pointer in the linked list
// is replaced with a pointer to the newly created node.
func (ls *LinkedStack) Push(value interface{}) {
	top_node := &node{previous_node: ls.top_node, value: value}
	ls.top_node = top_node
	ls.length++
}

// Graps the current top node of the linked list and replaces the top node
// pointer in the linked list with that of the previous node. Decrements
// the length of the Linked Stack and return the value to be poped.
func (ls *LinkedStack) Pop() interface{} {
	pop := ls.top_node
	ls.top_node = pop.previous_node
	ls.length--
	return pop.value
}

// Checks if the Linked Stack contains a top top node the return said node
// if there is one else it returns nil.
func (ls *LinkedStack) Peek() interface{} {
	if ls.length == 0 {
		return nil
	}

	return ls.top_node.value
}

// Checks if the Linked Stack is empty
func (ls *LinkedStack) IsEmpty() bool {
	return ls.length == 0
}
