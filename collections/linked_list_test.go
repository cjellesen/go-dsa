package collections

import (
	"log"
	"testing"
)

func create_test_lists(n int) (LinkedList[int], LinkedList[int]) {
	sequential_add := LinkedList[int]{Head: nil, Tail: nil, Count: 0}
	reverse_add := LinkedList[int]{Head: nil, Tail: nil, Count: 0}
	for i := 0; i < n; i++ {
		sequential_add.AddHead(i)
		reverse_add.AddHead(n - 1 - i)
	}

	return sequential_add, reverse_add
}

func TestLinkedListAddHeadAndTail(t *testing.T) {
	t.Log("Testing head and tail adding")

	n := 10
	sequential_add, reverse_add := create_test_lists(n)

	if sequential_add.Count != n {
		t.Fatalf("Expected linked list to include %d nodes had %d", n, sequential_add.Count)
	}

	sequential_node := sequential_add.Head
	reverse_node := reverse_add.Head
	i := 0
	for i < n {
		log.Printf(
			"Iteration %d - Sequential node value: %d, Reverse node value %d",
			i,
			sequential_node.Value,
			reverse_node.Value,
		)

		if sequential_node.Value != n-1-i {
			t.Fatalf(
				"Ordering of sequential linked list node is not correct - Expected value %d got value %d",
				i,
				sequential_node.Value,
			)

		}

		if reverse_node.Value != i {
			t.Fatalf(
				"Ordering of reversed linked list node is not correct - Expected value %d got value %d",
				i,
				sequential_node.Value,
			)

		}

		if i < n {
			sequential_node = sequential_node.Next
			reverse_node = reverse_node.Next
		}

		if i == n {
			if sequential_node.Next != nil {
				t.Fatalf(
					"Expected the sequential tail node to have a next value of 'nil' got %d",
					sequential_node.Next.Value,
				)
			}

			if reverse_node.Next != nil {
				t.Fatalf(
					"Expected the reverse tail node to have a next value of 'nil' got %d",
					sequential_node.Next.Value,
				)
			}
		}

		i++
	}
}
