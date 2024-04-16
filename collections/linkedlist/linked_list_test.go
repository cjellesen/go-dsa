package linkedlist

import (
	"math/rand/v2"
	"sort"
	"testing"
)

func TestCount(t *testing.T) {
	t.Log("LinkedList: Testing Count() function in uninitialized array")
	list := LinkedList[int]{
		Head: nil,
		Tail: nil,
	}

	if list.Count() != 0 {
		t.Fatalf("Count of initialized linked list is different from 0")
	}
}

func test_arrays_for_equality_by_element[T comparable](array1 *[]T, array2 *[]T, t *testing.T) {
	if len(*array1) != len(*array2) {
		t.Fatalf(
			"LinkedList - ToArray() test failed - Returned array and expected array does not have identical length",
		)
	}

	for i := range len(*array1) {
		if (*array1)[i] != (*array2)[i] {
			t.Fatalf(
				"LinkedList - ToArray() test failed - Value at nth index %d was not of equal value",
				i,
			)
		}
	}
}

func TestAddHead(t *testing.T) {
	t.Log("LinkedList: Testing AddHead() function")
	list := LinkedList[int]{
		Head: nil,
		Tail: nil,
	}

	for i := range 5 {
		list.AddHead(i + 1)
	}

	expected_array := []int{5, 4, 3, 2, 1}
	array := list.ToArray()
	t.Logf("Testing using data: %d", array)
	t.Logf("Expected Result: %d", expected_array)
	test_arrays_for_equality_by_element(&array, &expected_array, t)
}

func TestRemoveHeadFromListOf1(t *testing.T) {
	t.Log("LinkedList: Testing RemoveHead() from list of count = 1")
	list := LinkedList[int]{
		Head: nil,
		Tail: nil,
	}

	list.AddHead(1)
	list.RemoveHead()

	if list.count != 0 && (list.Head != nil || list.Tail != nil) {
		t.Logf("Failed to correctly reset linked list")
	}
}

func TestRemoveHead(t *testing.T) {
	t.Log("LinkedList: Testing RemoveHead() function")
	list := LinkedList[int]{
		Head: nil,
		Tail: nil,
	}

	for i := range 5 {
		list.AddHead(i + 1)
	}

	list.RemoveHead()
	expected_array := []int{4, 3, 2, 1}
	array := list.ToArray()
	t.Logf("Testing using data: %d", array)
	t.Logf("Expected Result: %d", expected_array)
	test_arrays_for_equality_by_element(&array, &expected_array, t)
}

func TestAddTail(t *testing.T) {
	t.Log("LinkedList: Testing AddTail() function")
	list := LinkedList[int]{
		Head: nil,
		Tail: nil,
	}

	for i := range 5 {
		list.AddTail(i + 1)
	}

	array := list.ToArray()
	expected_array := []int{1, 2, 3, 4, 5}
	t.Logf("Testing using data: %d", array)
	t.Logf("Expected Result: %d", expected_array)
	test_arrays_for_equality_by_element(&array, &expected_array, t)
}

func TestRemoveTailFromListOf1(t *testing.T) {
	t.Log("LinkedList: Testing RemoveTail() from list of count = 1")
	list := LinkedList[int]{
		Head: nil,
		Tail: nil,
	}

	list.AddHead(1)
	list.RemoveTail()
	if list.count != 0 && (list.Head != nil || list.Tail != nil) {
		t.Logf("Failed to correctly reset linked list")
	}
}

func TestRemoveTail(t *testing.T) {
	t.Log("LinkedList: Testing RemoveTail() function")
	list := LinkedList[int]{
		Head: nil,
		Tail: nil,
	}

	for i := range 5 {
		list.AddHead(i + 1)
	}

	list.RemoveTail()
	expected_array := []int{5, 4, 3, 2}
	array := list.ToArray()
	if len(expected_array) != len(array) {
		t.Fatalf("Arrays does not match")
	}
	t.Logf("Testing using data: %d", array)
	t.Logf("Expected Result: %d", expected_array)
	test_arrays_for_equality_by_element(&array, &expected_array, t)
}

func TestAddHeadAddTailAlteration(t *testing.T) {
	t.Log("LinkedList: Testing alternating AddHead(), AddTail()")
	list := LinkedList[int]{
		Head:  nil,
		Tail:  nil,
		count: 0,
	}
	for i := range 5 {
		if i%2 == 0 {
			list.AddHead(i)
		} else {
			list.AddTail(i)
		}
	}
	expected_array := []int{4, 2, 0, 1, 3}
	array := list.ToArray()
	if len(expected_array) != len(array) {
		t.Fatalf("Arrays does not match")
	}
	t.Logf("Testing using data: %d", array)
	t.Logf("Expected Result: %d", expected_array)
	test_arrays_for_equality_by_element(&array, &expected_array, t)
}

func TestFind(t *testing.T) {
	t.Log("LinkedList: Testing Find() function using value reference")
	n := 5
	linked_list := LinkedList[int]{Head: nil, Tail: nil}
	for i := range n {
		linked_list.AddTail(i)
	}

	valid_value := rand.IntN(n-2) + 1
	valid_node := linked_list.Find(valid_value)
	if valid_node.Value != valid_value {
		t.Fatalf("Could not find value %d in linked list", valid_value)
	}

	t.Logf("Successfully found node %d", valid_node.Value)
}

// Will create an array with dublicates, where there will always be dublicates at the array edges as well as somewhere
// in between.
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
	t.Log("LinkedList: Testing RemoveDublicates() function")
	no_dub_sequence, dub_sequence := create_test_data_with_dublicates(10)
	t.Logf("Testing using data: %d", dub_sequence)
	linkedList := LinkedList[int]{Head: nil, Tail: nil}
	for i := 0; i < len(dub_sequence); i++ {
		linkedList.AddTail(dub_sequence[i])
	}

	linkedList.RemoveDublicates()
	t.Logf("Expected Result: %d", no_dub_sequence)
	i := 0
	curr := linkedList.Head
	for curr.Next != nil {
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
	t.Log("LinkedList: Testing Intersects() function")
	n := 10
	diffLength := rand.IntN(n / 2)
	intersectIndex := rand.IntN(n-diffLength-1) + diffLength

	list1 := LinkedList[int]{Head: nil, Tail: nil}
	list2 := LinkedList[int]{Head: nil, Tail: nil}
	for i := 0; i < n; i++ {
		if i < intersectIndex {
			list1.AddTail(i)
			list2.AddTail(i)
		} else if i == intersectIndex {
			t.Logf("Intersecting lists at index: %d", i)
			list1.AddTail(i)
			list2.Tail.Next = list1.Tail
		} else {
			list1.AddTail(i)
		}
	}

	intersection := list1.Intersects(&list2)
	if intersection == nil {
		t.Fatalf("No intersect found, one should have been present at index: %d", intersectIndex)
	}
	if intersection.Value != intersectIndex {
		t.Fatalf(
			"Found an intersection point with value: %d, expected %d",
			intersection.Value,
			intersectIndex,
		)
	}

	t.Logf("Found an intersect at index: %d", intersection.Value)
}

func TestLoopDetection(t *testing.T) {
	t.Log("LinkedList: Testing HasLoop() function")
	n := 10
	diffLength := rand.IntN(n / 2)
	loopIndex := rand.IntN(n-diffLength-1) + diffLength

	var loopNode *Node[int]
	list_with_loop := LinkedList[int]{Head: nil, Tail: nil}
	list_without_loop := LinkedList[int]{Head: nil, Tail: nil}
	for i := 0; i < n; i++ {
		list_without_loop.AddTail(i)
		list_with_loop.AddTail(i)
		if i == loopIndex {
			loopNode = list_with_loop.Tail
		}
	}

	list_with_loop.Tail.Next = loopNode
	if !list_with_loop.HasLoop() {
		t.Fatalf("Failed to a loop in the linked list with a loop")
	}

	if list_without_loop.HasLoop() {
		t.Fatal("Detected a loop in the list without a loop")
	}
}

func TestInsertionAfterNode(t *testing.T) {
	t.Log("LinkedList: Testing AddAfter() function")
	list := LinkedList[int]{
		Head: nil,
		Tail: nil,
	}

	for i := range 5 {
		list.AddTail(i + 1)
	}

	insertion_node := list.Find(3)
	list.AddAfter(insertion_node, 20)
	array := list.ToArray()
	expected_array := []int{1, 2, 3, 20, 4, 5}
	t.Logf("Testing using data: %d", array)
	t.Logf("Expected Result: %d", expected_array)
	test_arrays_for_equality_by_element(&array, &expected_array, t)
}

func TestInsertionBeforeNode(t *testing.T) {
	t.Log("LinkedList: Testing AddBefore() function")
	list := LinkedList[int]{
		Head: nil,
		Tail: nil,
	}

	for i := range 5 {
		list.AddTail(i + 1)
	}

	insertion_node := list.Find(3)
	list.AddBefore(insertion_node, 20)
	array := list.ToArray()
	expected_array := []int{1, 2, 20, 3, 4, 5}
	t.Logf("Testing using data: %d", array)
	t.Logf("Expected Result: %d", expected_array)
	test_arrays_for_equality_by_element(&array, &expected_array, t)
}
