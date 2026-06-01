package main

import "fmt"

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
}

type Stack struct {
	Items []string
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

