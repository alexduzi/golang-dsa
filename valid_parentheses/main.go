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
	fmt.Println(isValid("()"))
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

func isValid(s string) bool {
	stack := []rune{} // Stack for opening brackets
	hash := map[rune]rune{')': '(', ']': '[', '}': '{'}

	for _, char := range s {
		if match, found := hash[char]; found {
			// Check if stack is non-empty and matches
			if len(stack) > 0 && stack[len(stack)-1] == match {
				stack = stack[:len(stack)-1] // Pop
			} else {
				return false // Invalid
			}
		} else {
			stack = append(stack, char) // Push opening bracket
		}
	}
	return len(stack) == 0 // Valid if stack is empty
}
