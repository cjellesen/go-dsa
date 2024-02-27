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
	list.Head = &Node[T]{Value: value, Next: nil}
	list.Head.Next = temp
	list.Count++

	if list.Count == 1 {
		list.Tail = list.Head
	}
}

func (list *LinkedList[T]) AddTail(value T) {
	node := &Node[T]{Value: value, Next: nil}
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
	for curr != nil {
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

func (list *LinkedList[T]) Intersects(other *LinkedList[T]) bool {
	panic("Function has not been implemented yet")
}

func (list *LinkedList[T]) GetIntersection(other *LinkedList[T]) LinkedList[T] {
	panic("Function has not been implemented yet")
}

func (list *LinkedList[T]) HasLoop() {
	panic("Function has not been implemented yet")
}
