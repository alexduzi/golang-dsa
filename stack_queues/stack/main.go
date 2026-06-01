package main

import (
	"errors"
	"fmt"
)

// último que entra é o primeiro que sai
// LAST IN FIRST OUT
// LIFO
func main() {
	stack := NewStack(nil)

	stack.Push("teste1")
	stack.Push("teste2")
	stack.Push("teste3")
	stack.Push("teste4")

	fmt.Printf("%s\n", stack.Pop())
	fmt.Printf("%v\n", stack.Items)
	fmt.Printf("%s\n", stack.Peek())
	fmt.Printf("A pilha está vazia?: %v\n", stack.IsEmpty())
	fmt.Printf("Quantos elementos a pilha tem?: %d\n", stack.Count())

	fmt.Println()
	fmt.Println("---------------------------")
	fmt.Println("---------------------------")
	fmt.Println()

	stackArray := NewStackArray(nil)

	stackArray.Push("teste1")
	stackArray.Push("teste2")
	stackArray.Push("teste3")
	stackArray.Push("teste4")

	popElem, err := stackArray.Pop()
	if err != nil {
		fmt.Print("%w\n", err)
	}
	peekElem, err := stackArray.Pop()
	if err != nil {
		fmt.Print("%w\n", err)
	}
	fmt.Printf("%s\n", popElem)
	fmt.Printf("%s\n", peekElem)
	fmt.Printf("A pilha está vazia?: %v\n", stackArray.IsEmpty())
	fmt.Printf("Quantos elementos a pilha tem?: %d\n", stackArray.Count())
	fmt.Println("Limpando a pilha")
	stackArray.Clear()
	fmt.Printf("A pilha está vazia?: %v\n", stackArray.IsEmpty())
	fmt.Printf("Quantos elementos a pilha tem?: %d\n", stackArray.Count())
	stackArray.Push("teste900")
	stackArray.Push("teste950")
	peekElem, err = stackArray.Peek()
	if err != nil {
		fmt.Print("%w\n", err)
	}
	fmt.Printf("%s\n", peekElem)
	fmt.Printf("Quantos elementos a pilha tem?: %d\n", stackArray.Count())

	fmt.Println()
	fmt.Println("---------------------------")
	fmt.Println("---------------------------")
	fmt.Println()

	text := "((()))(("
	isBal, err := isBalanced(text)
	if err != nil {
		fmt.Print("%w\n", err)
	}
	fmt.Printf("O texto %s esta balanceado? -> %v\n", text, isBal)

	text = "((()))(((())))"
	isBal, err = isBalanced(text)
	if err != nil {
		fmt.Print("%w\n", err)
	}
	fmt.Printf("O texto %s esta balanceado? -> %v\n", text, isBal)

	text = "((())))))"
	isBal, err = isBalanced(text)
	if err != nil {
		fmt.Print("%w\n", err)
	}
	fmt.Printf("O texto %s esta balanceado? -> %v\n", text, isBal)

	text = "()([]{})"
	isValid, err := isValidParentheses(text)
	if err != nil {
		fmt.Print("%w\n", err)
	}
	fmt.Printf("O texto %s esta balanceado? -> %v\n", text, isValid)
}

// implementação com lista / slice
type Stack struct {
	Items   []string
	MaxSize *int
}

func NewStack(maxSize *int) *Stack {
	return &Stack{
		MaxSize: maxSize,
	}
}

// empilha
func (stack *Stack) Push(value string) {
	stack.Items = append(stack.Items, value)
}

// desempilha
func (stack *Stack) Pop() string {
	len := len(stack.Items)
	top := stack.Items[len-1]
	stack.Items = stack.Items[:len-1]
	return top
}

// retorna o primeiro elemento da pilha sem desempilha-lo
func (stack *Stack) Peek() string {
	len := len(stack.Items)
	return stack.Items[len-1]
}

// verifica se a pilha está vazia ou não
func (stack *Stack) IsEmpty() bool {
	return len(stack.Items) == 0
}

// retorna o número de elementos na pilha
func (stack *Stack) Count() int {
	return len(stack.Items)
}

// testa se a pilha está cheia (caso a pilha tenha limite de tamanho)
func (stack *Stack) IsFull() bool {
	if stack.MaxSize == nil {
		return false
	}

	return len(stack.Items) == *stack.MaxSize
}

func (stack *Stack) Clear() {
	if stack.MaxSize != nil {
		stack.Items = make([]string, 0)
	} else {
		stack.Items = make([]string, 100)
		stack.MaxSize = new(100)
	}
}

// implementação com array
type StackArray struct {
	Items   []string
	MaxSize *int
	Top     int
}

func NewStackArray(maxSize *int) *StackArray {
	stackArray := &StackArray{}
	stackArray.Top = -1

	if maxSize != nil {
		stackArray.Items = make([]string, 0, *maxSize)
		stackArray.MaxSize = maxSize
	} else {
		stackArray.Items = make([]string, 100)
		stackArray.MaxSize = new(100)
	}

	return stackArray
}

// empilha
func (stack *StackArray) Push(value string) error {
	if stack.IsFull() {
		return errors.New("Stack is full")
	}
	stack.Top++
	stack.Items[stack.Top] = value
	return nil
}

// desempilha
func (stack *StackArray) Pop() (string, error) {
	if stack.IsEmpty() {
		return "", errors.New("Stack is empty")
	}
	top := stack.Items[stack.Top]
	stack.Top--
	return top, nil
}

// retorna o primeiro elemento da pilha sem desempilha-lo
func (stack *StackArray) Peek() (string, error) {
	if stack.IsEmpty() {
		return "", errors.New("Stack is empty")
	}
	return stack.Items[stack.Top], nil
}

// verifica se a pilha está vazia ou não
func (stack *StackArray) IsEmpty() bool {
	return stack.Top == -1
}

// retorna o número de elementos na pilha
func (stack *StackArray) Count() int {
	return stack.Top + 1
}

// testa se a pilha está cheia (caso a pilha tenha limite de tamanho)
func (stack *StackArray) IsFull() bool {
	return stack.Top == *stack.MaxSize-1
}

func (stack *StackArray) Clear() {
	stack.Top = -1
}

// implementacao de uma stack genérica com array
type StackArrayGeneric[T any] struct {
	Items   []T
	MaxSize *int
	Top     int
}

func NewStackArrayGeneric[T any](maxSize *int) *StackArrayGeneric[T] {
	stackArray := &StackArrayGeneric[T]{}
	stackArray.Top = -1

	if maxSize != nil {
		stackArray.Items = make([]T, 0, *maxSize)
		stackArray.MaxSize = maxSize
	} else {
		stackArray.Items = make([]T, 100)
		stackArray.MaxSize = new(100)
	}

	return stackArray
}

// empilha
func (stack *StackArrayGeneric[T]) Push(value T) error {
	if stack.IsFull() {
		return errors.New("Stack is full")
	}
	stack.Top++
	stack.Items[stack.Top] = value
	return nil
}

// desempilha
func (stack *StackArrayGeneric[T]) Pop() (T, error) {
	if stack.IsEmpty() {
		var empty T
		return empty, errors.New("Stack is empty")
	}
	top := stack.Items[stack.Top]
	stack.Top--
	return top, nil
}

// retorna o primeiro elemento da pilha sem desempilha-lo
func (stack *StackArrayGeneric[T]) Peek() (T, error) {
	if stack.IsEmpty() {
		var empty T
		return empty, errors.New("Stack is empty")
	}
	return stack.Items[stack.Top], nil
}

// verifica se a pilha está vazia ou não
func (stack *StackArrayGeneric[T]) IsEmpty() bool {
	return stack.Top == -1
}

// retorna o número de elementos na pilha
func (stack *StackArrayGeneric[T]) Count() int {
	return stack.Top + 1
}

// testa se a pilha está cheia (caso a pilha tenha limite de tamanho)
func (stack *StackArrayGeneric[T]) IsFull() bool {
	return stack.Top == *stack.MaxSize-1
}

func (stack *StackArrayGeneric[T]) Clear() {
	stack.Top = -1
}

func isBalanced(text string) (bool, error) {
	stack := NewStackArrayGeneric[rune](nil)

	for _, char := range text {
		switch char {
		case '(':
			stack.Push(char)
		case ')':
			if stack.IsEmpty() {
				return false, nil
			}
			stack.Pop()
		}
	}

	return stack.IsEmpty(), nil
}

func isValidParentheses(text string) (bool, error) {
	stack := NewStackArrayGeneric[rune](nil)

	for _, char := range text {
		switch char {
		case '(':
			stack.Push(')')
		case '{':
			stack.Push('}')
		case '[':
			stack.Push(']')
		default:
			elem, _ := stack.Peek()
			if stack.IsEmpty() || elem != char {
				return false, nil
			}
			stack.Pop()
		}
	}

	return stack.IsEmpty(), nil
}
