package main

import "fmt"

func main() {
	stack := new(Stack)

	for i := 0; i < 100; i++ {
		stack.Push(fmt.Sprintf("Item %v", i))
	}

	fmt.Printf("Stack isEmpty: %v\n", stack.IsEmpty())
	fmt.Printf("Stack length: %v\n", stack.Length())
	fmt.Printf("Stack peek: %v\n", stack.Peek())
	fmt.Printf("Stack pop: %v\n", stack.Pop())
	fmt.Printf("Stack length: %v\n", stack.Length())
	fmt.Printf("Stack pop: %v\n", stack.Pop())
	fmt.Printf("Stack pop: %v\n", stack.Pop())
	fmt.Printf("Stack length: %v\n", stack.Length())

	len := stack.Length()

	for i := 0; i < len; i++ {
		stack.Pop()
	}

	fmt.Printf("Stack isEmpty: %v\n", stack.IsEmpty())
	fmt.Printf("Stack length: %v\n", stack.Length())
}
