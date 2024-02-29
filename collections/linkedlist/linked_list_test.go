package linkedlist

import (
	"math/rand/v2"
	"sort"
	"testing"
)

func create_sequential_and_reverse_test_data(n int) (LinkedList[int], LinkedList[int]) {
	sequential_values := LinkedList[int]{Head: nil, Tail: nil, Count: 0}
	reverse_values := LinkedList[int]{Head: nil, Tail: nil, Count: 0}
	for i := 0; i < n; i++ {
		sequential_values.AddTail(&Node[int]{Value: i})
		reverse_values.AddHead(&Node[int]{Value: i})
	}

	return sequential_values, reverse_values
}

func TestLinkedListAddHeadAndTail(t *testing.T) {
	t.Log("Testing head and tail adding")

	n := 10
	seq_list, rev_list := create_sequential_and_reverse_test_data(n)

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

func TestRemoveHeadAndTail(t *testing.T) {
	t.Log("Testing remove functions")

	test_seq, _ := create_sequential_and_reverse_test_data(3)
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

func create_test_data_with_dublicates(n int) ([]int, []int) {
	n_dublicates := (n * 10 / 100) + 2
	test_data_with_dublicates := make([]int, n+n_dublicates)
	test_data_without_dublicates := make([]int, n)

	for i := 0; i < n; i++ {
		test_data_with_dublicates[i+1] = i
		test_data_without_dublicates[i] = i
	}

	test_data_with_dublicates[0] = test_data_without_dublicates[0]
	test_data_with_dublicates[n+1] = test_data_without_dublicates[n-1]
	for i := 2; i < n_dublicates; i++ {
		test_data_with_dublicates[n+i] = rand.IntN(n)
	}

	sort.Ints(test_data_with_dublicates[:])
	return test_data_without_dublicates, test_data_with_dublicates
}

func TestRemoveDublicatesOnValueTypes(t *testing.T) {
	t.Log("Testing remove dublicates with values (int)")
	no_dub_sequence, dub_sequence := create_test_data_with_dublicates(10)
	t.Logf("Testing using data: %d", dub_sequence)
	linkedList := LinkedList[int]{Head: nil, Tail: nil, Count: 0}
	for i := 0; i < len(dub_sequence); i++ {
		linkedList.AddTail(&Node[int]{Value: dub_sequence[i]})
	}

	linkedList.RemoveDublicates()
	t.Logf("Expected result: %d", no_dub_sequence)
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
	t.Log("Testing remove dublicates using structs")
	no_dub_sequence, dub_sequence := create_test_data_with_dublicates(10)
	t.Logf("Testing using data: %d", dub_sequence)
	linkedList := LinkedList[ValueHolder]{Head: nil, Tail: nil, Count: 0}
	for i := 0; i < len(dub_sequence); i++ {
		linkedList.AddTail(&Node[ValueHolder]{Value: ValueHolder{Value: dub_sequence[i]}})
	}

	linkedList.RemoveDublicates()
	t.Logf("Expected result: %d", no_dub_sequence)
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
	t.Log("Testing remove dublicates using struct pointers")
	no_dub_sequence, dub_sequence := create_test_data_with_dublicates(10)
	t.Logf("Testing using data: %d", dub_sequence)
	reference_map := make(map[int]*ValueHolder)
	for i := 0; i < len(dub_sequence); i++ {
		_, ok := reference_map[dub_sequence[i]]
		if ok {
			continue
		}
		reference_map[dub_sequence[i]] = &ValueHolder{dub_sequence[i]}
	}

	linkedList := LinkedList[*ValueHolder]{Head: nil, Tail: nil, Count: 0}
	for i := 0; i < len(dub_sequence); i++ {
		linkedList.AddTail(&Node[*ValueHolder]{Value: reference_map[dub_sequence[i]]})
	}

	linkedList.RemoveDublicates()
	t.Logf("Expected result: %d", no_dub_sequence)
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

func create_arrays_with_offset(n int) ([]int, []int) {
	array := make([]int, n)
	offset_index := rand.IntN(n - 1)
	offset := make([]int, (n - offset_index))
	for i := 0; i < n; i++ {
		array[i] = i
		offset[i+offset_index] = i
	}

	return array, offset
}

func TestFindIntersect(t *testing.T) {
	t.Log("Testing remove dublicates using struct pointers")
	n := 10
	offset := rand.IntN(n - 1)

	list := LinkedList[int]{Head: nil, Tail: nil}
	offset_list := LinkedList[int]{Head: nil, Tail: nil}
	for i := 0; i < n; i++ {
		node := &Node[int]{Value: i}
		list.AddTail(node)
		if i >= offset {
			offset_list.AddTail(node)
		}
	}

	intersection := list.Intersects(&offset_list)
	if intersection.Value != offset {
		t.Fatalf(
			"Found an intersection point with value: %d, expected %d",
			intersection.Value,
			offset,
		)
	}
}
