package queue

import (
	"testing"
)

func TestQueueCreation(t *testing.T) {
	t.Log("Queue: Testing queue creation")
	queue := Queue[int]{}
	n := 5
	for i := range n {
		queue.Enqueue(i + 1)
	}

	n = 1
	for !queue.IsEmpty() {
		peek, _ := queue.Peek()
		if peek != n {
			t.Fatalf("Expected queue peek value of %d, got %d", n, peek)
		}

		pop, _ := queue.Dequeue()
		if pop != n {
			t.Fatalf("Expected queue pop value of %d, got %d", n, pop)
		}

		n++
	}
}
