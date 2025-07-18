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
	fmt.Println(isValidParentheses("((({{{}}}[])))"))
}

func isValidParentheses(s string) bool {
	stack := Stack{}

	for _, charPos := range s {
		char := fmt.Sprintf("%c", charPos)
		if char == "(" || char == "{" || char == "[" {
			stack.Push(char)
		} else {
			if stack.IsEmpty() {
				return false
			}

			top := stack.Peek()

			if char == ")" && top == "(" ||
				char == "}" && top == "{" ||
				char == "]" && top == "[" {
				stack.Pop()
			} else {
				return false
			}
		}
	}
	return true
}
