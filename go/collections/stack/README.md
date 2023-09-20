# Stacks

A stack is a first-in-first-out (FIFO) datastructure and can take on different forms depending on its implementation.
There will be 2 examples of stack implementations in this repository:
1) A stack based on a linked list
2) A stack based on an array

For stack implementations based on linked lists it holds that these are quick (O(1)) to append to and to pop from (O(1)). However, since the implementation relies on linked lists they will be subject to allocations which will make incure a memory overhead.

For stack implementation based on arrays the same it holds that access is fast (O(1)) and insertion time grows with array size (O(n)). The part about insertions only holds though if the array bounds are not exceeded since this will trigger the need for capacity increase and require preallocating a new array, generally with length n*2, and copying over the content of the previous array. However, if it is assumed that in the case of needing an unbounded stack one would utility a linked list backed stack, then the array implementation can simply require a fixed array size and throw an exception if array bounds are exceeded.

For stacks backed with linked lists the following functions exists:
Push
Pop
IsEmpty
Peek

For stacks backed with arrays the following functions exists:
Push
Pop
IsEmpty
Peek
IsFull 