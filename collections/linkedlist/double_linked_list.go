package linkedlist

type BidirectionalNode[T comparable] struct {
	Value    T
	Previous *BidirectionalNode[T]
	Next     *BidirectionalNode[T]
}

type DoublyLinkedList[T comparable] struct {
	Head  *BidirectionalNode[T]
	Tail  *BidirectionalNode[T]
	count int
}

func (list *DoublyLinkedList[T]) Count() int {
	return list.count
}

func (list *DoublyLinkedList[T]) AddHead(value T) {
	node := &BidirectionalNode[T]{Value: value}
	if list.count == 0 {
		list.Head = node
		list.Tail = node
	} else {
		node.Next = list.Head
		list.Head.Previous = node
		list.Head = node
	}
	list.count++
}

func (list *DoublyLinkedList[T]) AddTail(value T) {
	node := &BidirectionalNode[T]{Value: value}
	if list.count == 0 {
		list.Head = node
		list.Tail = node
	} else {
		node.Previous = list.Tail
		list.Tail.Next = node
		list.Tail = node
	}
	list.count++
}

func (list *DoublyLinkedList[T]) Find(value T) *BidirectionalNode[T] {
	if list.count == 0 {
		return nil
	}

	if list.Tail.Value == value {
		return list.Tail
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

func (list *DoublyLinkedList[T]) AddBefore(insertLoc *BidirectionalNode[T], value T) error {
	panic("DoublyLinkedList: AddBefore is not implemented yet")
}

func (list *DoublyLinkedList[T]) AddAfter(insertLoc *BidirectionalNode[T], value T) error {
	panic("DoublyLinkedList: AddAfter is not implemeted yet")
}

func (list *DoublyLinkedList[T]) RemoveHead() {
	panic("DoublyLinkedList: RemoveHead is not implemented yet")
}

func (list *DoublyLinkedList[T]) RemoveTail() {
	panic("DoublyLinkedList: RemioveTail is not implemented yet")
}

func (list *DoublyLinkedList[T]) RemoveDublicates() {
	panic("DoublyLinkedList: RemoveDublicates is not implemented yet")
}

func (list *DoublyLinkedList[T]) Intersects(other *DoublyLinkedList[T]) *BidirectionalNode[T] {
	panic("DoublyLinkedList: Intersects is not implemented yet")
}

func (list *DoublyLinkedList[T]) HasLoop() bool {
	panic("DoublyLinkedList: HasLoop is not implemented yet")
}

func (list *DoublyLinkedList[T]) Sort() {
	panic("DoublyLinkedList: Sort is not implemented yet")
}

func (list *DoublyLinkedList[T]) ToArray() {
	panic("DoublyLinkedList: ToArray is not implemented yet")
}
