package linkedlist

import (
	"testing"
)

func create_test_lists(n int) (LinkedList[int], LinkedList[int]) {
	sequential_values := LinkedList[int]{Head: nil, Tail: nil, Count: 0}
	reverse_values := LinkedList[int]{Head: nil, Tail: nil, Count: 0}
	for i := 0; i < n; i++ {
		sequential_values.AddTail(i)
		reverse_values.AddHead(i)
	}

	return sequential_values, reverse_values
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

		if seq_node.Value != i {
			t.Fatalf(
				"Ordering of sequential linked list node is not correct - Expected value %d got value %d",
				i,
				seq_node.Value,
			)

		}

		if rev_node.Value != n-1-i {
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

func TestRemoveDublicatesOnValueTypes(t *testing.T) {
	sequence := [10]int{1, 1, 2, 3, 4, 4, 5, 6, 7, 7}
	linkedList := LinkedList[int]{Head: nil, Tail: nil, Count: 0}
	for i := 0; i < len(sequence); i++ {
		linkedList.AddTail(sequence[i])
	}

	linkedList.RemoveDublicates()
	no_dub_sequence := [7]int{1, 2, 3, 4, 5, 6, 7}
	i := 0
	curr := linkedList.Head
	for curr.Next != nil {
		t.Logf("Found value %d, expected value %d", curr.Value, no_dub_sequence[i])

		if curr.Value != no_dub_sequence[i] {
			t.Fatalf(
				"Failed to correctly remove dublicates, expected value %d got %d",
				no_dub_sequence[i],
				curr.Value,
			)
		}
		curr = curr.Next
		i++
	}
}

type ValueHolder struct {
	Value int
}

func TestRemoveDublicatesOnValueStructTypes(t *testing.T) {
	sequence := [10]int{1, 1, 2, 3, 4, 4, 5, 6, 7, 7}
	linkedList := LinkedList[ValueHolder]{Head: nil, Tail: nil, Count: 0}
	for i := 0; i < len(sequence); i++ {
		value := ValueHolder{Value: sequence[i]}
		linkedList.AddTail(value)
	}

	linkedList.RemoveDublicates()
	no_dub_sequence := [7]int{1, 2, 3, 4, 5, 6, 7}
	i := 0
	curr := linkedList.Head
	for curr.Next != nil {
		t.Logf("Found value %d, expected value %d", curr.Value.Value, no_dub_sequence[i])

		if curr.Value.Value != no_dub_sequence[i] {
			t.Fatalf(
				"Failed to correctly remove dublicates, expected value %d got %d",
				no_dub_sequence[i],
				curr.Value.Value,
			)
		}
		curr = curr.Next
		i++
	}
}

func TestRemoveDublicatesOnReferenceStructTypes(t *testing.T) {
	sequence := [10]int{1, 1, 2, 3, 4, 4, 5, 6, 7, 7}
	reference_map := make(map[int]*ValueHolder)
	for i := 0; i < len(sequence); i++ {
		_, ok := reference_map[sequence[i]]
		if ok {
			continue
		}
		reference_map[sequence[i]] = &ValueHolder{sequence[i]}
	}

	linkedList := LinkedList[*ValueHolder]{Head: nil, Tail: nil, Count: 0}
	for i := 0; i < len(sequence); i++ {
		linkedList.AddTail(reference_map[sequence[i]])
	}

	linkedList.RemoveDublicates()
	no_dub_sequence := [7]int{1, 2, 3, 4, 5, 6, 7}
	i := 0
	curr := linkedList.Head
	for curr.Next != nil {
		t.Logf("Found value %d, expected value %d", curr.Value.Value, no_dub_sequence[i])

		if curr.Value.Value != no_dub_sequence[i] {
			t.Fatalf(
				"Failed to correctly remove dublicates, expected value %d got %d",
				no_dub_sequence[i],
				curr.Value.Value,
			)
		}
		curr = curr.Next
		i++
	}
}
