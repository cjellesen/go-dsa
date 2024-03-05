package linkedlist

type BidirectionalNode[T comparable] struct {
	Value    T
	Previous *BidirectionalNode[T]
	Next     *BidirectionalNode[T]
}

type DoublyLinkedList[T comparable] struct {
	Head  *Node[T]
	Count int
}

func (list *DoublyLinkedList[T]) AddHead(value T) {
}
