package linkedlist

import (
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
	seq_list, rev_list := create_test_lists(n)

	if seq_list.Count != n {
		t.Fatalf("Expected linked list to include %d nodes had %d", n, seq_list.Count)
	}

	seq_node := seq_list.Head
	rev_node := rev_list.Head
	i := 0
	for i < n {
		t.Logf(
			"Iteration %d - Sequential node value: %d, Reverse node value %d",
			i,
			seq_node.Value,
			rev_node.Value,
		)

		if seq_node.Value != n-1-i {
			t.Fatalf(
				"Ordering of sequential linked list node is not correct - Expected value %d got value %d",
				i,
				seq_node.Value,
			)

		}

		if rev_node.Value != i {
			t.Fatalf(
				"Ordering of reversed linked list node is not correct - Expected value %d got value %d",
				i,
				seq_node.Value,
			)

		}

		if i < n {
			seq_node = seq_node.Next
			rev_node = rev_node.Next
		}

		if i == n {
			if seq_node.Next != nil {
				t.Fatalf(
					"Expected the sequential tail node to have a next value of 'nil' got %d",
					seq_node.Next.Value,
				)
			}

			if rev_node.Next != nil {
				t.Fatalf(
					"Expected the reverse tail node to have a next value of 'nil' got %d",
					seq_node.Next.Value,
				)
			}
		}

		i++
	}
}

func TestRemove(t *testing.T) {
	t.Log("Testing remove functions")

	test_seq, _ := create_test_lists(3)
	second_last := test_seq.Head
	for second_last.Next != test_seq.Tail {
		second_last = second_last.Next
	}

	test_seq.RemoveTail()
	if test_seq.Tail != second_last {
		t.Fatalf(
			"Failed to correctly remove tail, expected %d after removal got %d",
			second_last.Value,
			test_seq.Tail.Value,
		)
	}

	second := test_seq.Head.Next
	test_seq.RemoveHead()

	if test_seq.Head != second {
		t.Fatalf(
			"Failed to correctly remove head, expected %d after removal got %d",
			second.Value,
			test_seq.Head.Value,
		)
	}
}
