package linked_stack

import (
	"testing"
)

type testObjects struct {
	Value int
}

func TestLinkedStackPush(t *testing.T) {
	linked_stack := New[testObjects]()
	iter_test_size := 3
	for i := 0; i < iter_test_size; i++ {
		linked_stack.Push(&testObjects{i})
	}
	count := iter_test_size - 1
	for !linked_stack.IsEmpty() {
		stack_element, err := linked_stack.Pop()
		if err != nil {
			t.Fatalf("Testing failed - Tried to Pop from an empty stack")
		}

		if stack_element.Value != count {
			t.Fatalf("Linked Stack Push test failed: Expected value %d got value %d", count, stack_element.Value)
		}
		count--
	}
}
