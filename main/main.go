package main

import (
	"fmt"

	"github.com/cjellesen/go-dsa/collections/linkedlist"
	"github.com/cjellesen/go-dsa/collections/stack"
)

func main() {
	fmt.Println("Testing")
	test := linkedlist.LinkedList[int]{Head: nil, Tail: nil}
	fmt.Printf(
		"LinkedList has %d nodes",
		test.Count(),
	)

	stack := stack.Stack[int]{}
	stack.IsEmpty()
}
