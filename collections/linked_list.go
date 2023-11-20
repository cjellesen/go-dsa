package collections

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

type LinkedList[T any] struct {
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
	if list.Count > 0 {

		if list.Count == 1 {
			list.Head = nil
			list.Tail = nil
		} else {
			curr := list.Head
			for curr.Next != list.Tail {
				curr = curr.Next
			}

			list.Tail = curr
			curr.Next = nil
		}
		list.Count--
	}
}
