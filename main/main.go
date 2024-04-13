package main

import (
	"fmt"

	"github.com/cjellesen/go-dsa/collections/linkedlist"
)

func main() {
	fmt.Println("Testing")
	test := linkedlist.LinkedList[int]{Head: nil, Tail: nil}
	fmt.Printf(
		"LinkedList has %d nodes",
		test.Count,
	)
}
