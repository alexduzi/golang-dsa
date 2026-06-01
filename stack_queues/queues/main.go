package main

import (
	"errors"
	"fmt"
)

func main() {
	queue := NewQueue[string]()

	queue.Add("test1")
	queue.Add("test2")
	queue.Add("test3")

	fmt.Printf("A fila está vazia: %v\n", queue.IsEmpty())
	fmt.Printf("A fila tem %v items\n", queue.Count())
	first, err := queue.Remove()
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("O primeiro item é: %s\n", first)

	fmt.Println()
	fmt.Println("---------------------------")
	fmt.Println("---------------------------")
	fmt.Println()

	fmt.Printf("Problema dos tickets timeRequiredToBuy([]int{2, 3, 2}, 2)): %v\n", timeRequiredToBuy([]int{2, 3, 2}, 2))
}

// Definição de Fila
// Uma fila é uma estrutura de dados linear que segue a ordem de
// operação FIFO, onde o primeiro elemento inserido é o primeiro a ser removido
// FIFO: first in, first out
// exemplos de aplicações:
// jobs de impressão
// eventos em sistemas distribuídos
// envio de pacotes em redes
// processamento de requisições
// sistemas de mensageria
// processos em um sistema operacional
type Queue[T any] struct {
	Items []T
	Index int
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		Items: make([]T, 100),
		Index: -1,
	}
}

func (queue *Queue[T]) Add(value T) {
	queue.Index++
	queue.Items[queue.Index] = value
}

func (queue *Queue[T]) Remove() (T, error) {
	var empty T
	if queue.IsEmpty() {
		return empty, errors.New("Queue is empty")
	}
	value := queue.Items[0]
	// para assegurar que não terá vazamento de memória
	queue.Items[0] = empty
	queue.Items = queue.Items[1:]
	queue.Index--
	return value, nil
}

func (queue *Queue[T]) Peek() (T, error) {
	var empty T
	if queue.IsEmpty() {
		return empty, errors.New("Queue is empty")
	}
	value := queue.Items[0]
	return value, nil
}

func (queue *Queue[T]) IsEmpty() bool {
	return queue.Index == -1
}

func (queue *Queue[T]) Count() int {
	return queue.Index + 1
}

// problema tickets leetcode
// https://leetcode.com/problems/time-needed-to-buy-tickets/description/
func timeRequiredToBuy(tickets []int, k int) int {
	queue := []int{}

	for i := 0; i < len(tickets); i++ {
		queue = append(queue, i)
	}

	time := 0
	for true {
		// tira da fila
		i := queue[0]
		queue = queue[1:]

		// decrementa, é a mesma coisa que comprar o ticket
		tickets[i]--

		// conta +1 tempo
		time++

		// se tiver ticket para comprar volta para o final da fila
		if tickets[i] > 0 {
			queue = append(queue, i)
		} else if i == k {
			return time
		}
	}

	return -1
}
