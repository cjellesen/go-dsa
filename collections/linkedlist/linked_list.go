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

func (list *LinkedList[T]) AddHead(value T) {
	temp := list.Head
	list.Head = &Node[T]{Value: value}
	list.Head.Next = temp
	list.Count++

	if list.Count == 1 {
		list.Tail = list.Head
	}
}

func (list *LinkedList[T]) AddTail(value T) {
	node := &Node[T]{Value: value}
	if list.Count == 0 {
		list.Head = node
		list.Tail = node
	} else {
		list.Tail.Next = node
	}

	list.Tail = node
	list.Count++
}

func (list *LinkedList[T]) Find(value T) *Node[T] {
	if list.Count == 0 {
		return nil
	}

	if list.Tail.Value == value {
		return nil
	}

	pointer := list.Head
	for pointer != list.Tail {
		if pointer.Value == value {
			return pointer
		}
		pointer = pointer.Next
	}

	return nil
}

func (list *LinkedList[T]) AddBefore(insertLoc *Node[T], value T) {
	panic("AddBefore has not been implemented yet")
}

func (list *LinkedList[T]) AddAfter(insertLoc *Node[T], value T) {
	panic("AddAfter has not been impemented yet")
}

func (list *LinkedList[T]) RemoveHead() {
	if list.Count == 0 {
		return
	}

	list.Head = list.Head.Next
	if list.Count == 0 {
		list.Tail = nil
	}
	list.Count--
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
	if list.Count == 0 || other.Count == 0 {
		return nil
	}

	if list.Head == other.Head {
		return list.Head
	}

	var p1 = list.Head
	var p2 = other.Head

	for p1 != p2 {
		p1 = p1.Next
		p2 = p2.Next

		if p1 == p2 {
			return p1
		}

		if p1 == nil {
			p1 = other.Head
		}

		if p2 == nil {
			p2 = list.Head
		}
	}

	return p1
}

// This function check whether a linked list has an internal loop
func (list *LinkedList[T]) HasLoop() bool {
	if list.Count == 0 {
		return false
	}

	slowPointer := list.Head
	fastPointer := list.Head
	for fastPointer != nil && fastPointer.Next != nil {
		slowPointer = slowPointer.Next
		fastPointer = fastPointer.Next.Next

		if slowPointer == fastPointer {
			return true
		}
	}
	return false
}

// This is necessary if we want to implement a HashSet based on the linked list
func (list *LinkedList[T]) Sort() {
	panic("Function has not been implemented yet")
}
