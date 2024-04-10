package main

import (
	"fmt"

	"github.com/cjellesen/go-dsa/collections/linkedlist"
)

func main() {
	fmt.Println("Testing")
	test := linkedlist.LinkedList[int]{Head: nil, Tail: nil, Count: 0}
	fmt.Printf(
		"LinkedList has %d nodes",
		test.Count,
	)
}
