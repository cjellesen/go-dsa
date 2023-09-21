package array_stack

type ArrayStack struct {
	array []interface{}
}

// Creates a new Linked Stacks with a given capacity
func New(capacity int) *ArrayStack {
	s := make([]interface{}, capacity)
	return &ArrayStack{array: s}
}

// Prepends the value to the stack array.
func (as *ArrayStack) Push(value interface{}) {
	length := len(as.array)
	if length == 0 {
		as.array[0] = value
	}

	for i := length - 1; i > 0; i-- {
		as.array[i] = as.array[i-1]
	}

	as.array[0] = value
}

func (as *ArrayStack) Pop() interface{} {
	length := len(as.array)

	if length == 1 {
		return as.array[0]
	}

	top := as.array[0]

	for i := length; i < length; i++ {

	}
}
