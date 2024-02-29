package linkedlist

type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

type LinkedList[T comparable] struct {
	Head  *Node[T]
	Tail  *Node[T]
	Count int
}

func (list *LinkedList[T]) AddHead(node *Node[T]) {
	temp := list.Head
	list.Head = node
	list.Head.Next = temp
	list.Count++

	if list.Count == 1 {
		list.Tail = list.Head
	}
}

func (list *LinkedList[T]) AddTail(node *Node[T]) {
	if list.Count == 0 {
		list.Head = node
		list.Tail = node
	} else {
		list.Tail.Next = node
	}

	list.Tail = node
	list.Count++
}

func (list *LinkedList[T]) RemoveHead() {
	if list.Count > 0 {
		list.Head = list.Head.Next
		if list.Count == 0 {
			list.Tail = nil
		}
		list.Count--
	}
}

func (list *LinkedList[T]) RemoveTail() {
	if list.Count == 0 {
		return
	}

	if list.Count == 1 {
		list.Head = nil
		list.Tail = nil
		list.Count--
		return
	}

	curr := list.Head
	for curr.Next != list.Tail {
		curr = curr.Next
	}

	list.Tail = curr
	curr.Next = nil

	list.Count--
}

func (list *LinkedList[T]) RemoveDublicates() {
	if list.Count == 0 {
		return
	}

	set := make(map[T]struct{})
	set[list.Head.Value] = struct{}{}

	curr := list.Head.Next
	prev := list.Head
	for curr.Next != nil {
		_, ok := set[curr.Value]
		if ok {
			prev.Next = curr.Next
			list.Count--
		} else {
			set[curr.Value] = struct{}{}
			prev = curr
		}

		curr = curr.Next
	}

	if list.Tail != prev {
		list.Tail = prev
	}
}

// This function checks whether 2 linked lists intersects.
// If an intersection exists the point of intersection is return, otherwise a nil is returned.
func (list *LinkedList[T]) Intersects(other *LinkedList[T]) *Node[T] {
	var short *LinkedList[T]
	var long *LinkedList[T]
	if list.Count > other.Count {
		short = other
		long = list
	} else {
		short = list
		long = other
	}

	pointer := long.Head
	for pointer != short.Head {
		pointer = pointer.Next
	}

	if pointer != short.Head {
		return nil
	} else {
		return pointer
	}
}

// This function check whether a linked list has an internal loop
func (list *LinkedList[T]) HasLoop() bool {
	panic("Function has not been implemented yet")
}

// This is necessary if we want to implement a HashSet based on the linked list
func (list *LinkedList[T]) Sort() {
	panic("Function has not been implemented yet")
}
