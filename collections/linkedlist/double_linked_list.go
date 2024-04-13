package linkedlist

type BidirectionalNode[T comparable] struct {
	Value    T
	Previous *BidirectionalNode[T]
	Next     *BidirectionalNode[T]
}

type DoublyLinkedList[T comparable] struct {
	Head  *BidirectionalNode[T]
	Tail  *BidirectionalNode[T]
	Count int
}

func (list *DoublyLinkedList[T]) AddHead(value T) {
	temp := list.Head
	list.Head = &BidirectionalNode[T]{Value: value}
	list.Head.Next = temp
	list.Count++

	if list.Count == 1 {
		list.Tail = list.Head
	}
}
