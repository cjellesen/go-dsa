package main

import (
	"fmt"

	"github.com/cjellesen/go-dsa/collections/stack/linked_stack"
)

func main() {
	fmt.Println("Testing imprt")
	asd := linked_stack.New()
	asd.IsEmpty()
	// fmt.Println("Is the stack empty: ", linked_stack.New().IsEmpty())
}
