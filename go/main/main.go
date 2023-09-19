package main

import (
	col "collections"
	"fmt"
)

func main() {
	fmt.Println("Testing imprt")
	col.NewStack().Print()
	col.NewQueue().Print()
	col.NewPriorityQueue().Print()
	col.NewCircularQueue().Print()
}
