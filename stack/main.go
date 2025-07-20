package main

import "fmt"

type Stack[T int | string] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() T {
	if s.IsEmpty() {
		var initial T
		return initial
	}

	lastIdx := len(s.items) - 1
	last := s.items[lastIdx]
	s.items = s.items[:lastIdx]
	return last
}

func (s *Stack[T]) Peek() T {
	if s.IsEmpty() {
		var initial T
		return initial
	}
	return s.items[len(s.items)-1]
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack[T]) Length() int {
	return len(s.items)
}

func main() {
	arr := []int{1, 4, 5, 6, 1}

	reverseArr := nextGreaterElement(arr)

	for _, value := range reverseArr {
		fmt.Printf("%v\t", value)
	}

}

func nextGreaterElement(arr []int) []int {
	length := len(arr)
	result := make([]int, length)
	stack := Stack[int]{}

	for i := length - 1; i >= 0; i-- {
		if !stack.IsEmpty() {
			for !stack.IsEmpty() && stack.Peek() <= arr[i] {
				stack.Pop()
			}
		}

		if stack.IsEmpty() {
			result[i] = -1
		} else {
			result[i] = stack.Peek()
		}

		stack.Push(arr[i])
	}

	return result
}

func StackOperations() {
	stack := Stack[string]{}

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
