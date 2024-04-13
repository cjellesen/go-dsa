package stack

import (
	"testing"
)

func TestStackCreation(t *testing.T) {
	t.Log("Stack: Testing stack creation")
	stack := Stack[int]{}
	n := 5
	for i := range n {
		stack.Push(i + 1)
	}

	for !stack.IsEmpty() {
		peek, _ := stack.Peek()
		if peek != n {
			t.Fatalf("Expected stack peek value of %d, got %d", n, peek)
		}

		pop, _ := stack.Pop()
		if pop != n {
			t.Fatalf("Expected stack pop value of %d, got %d", n, pop)
		}

		n--
	}
}
