package main

import "fmt"

type Stack struct {
	items []string
}

func (s *Stack) Push(item string) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() string {
	if len(s.items) == 0 {
		return ""
	}

	lastIdx := len(s.items) - 1
	last := s.items[lastIdx]
	s.items = s.items[:lastIdx]
	return last
}

func (s *Stack) Peek() string {
	if len(s.items) == 0 {
		return ""
	}
	return s.items[len(s.items)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Length() int {
	return len(s.items)
}

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
