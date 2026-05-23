package main

import "fmt"

func main() {
	linkedList := NewLinkedList()
	fmt.Printf("Lista esta vazia?: %v\n", linkedList.IsEmpty())

	linkedList.AddAtEnd(1)
	linkedList.AddAtEnd(2)
	linkedList.AddAtEnd(3)
	linkedList.AddAtEnd(4)
	linkedList.AddAtEnd(5)

	linkedList.PrintLinkedList()

	fmt.Printf("Tamanho da lista: %v\n", linkedList.GetSize())

	fmt.Printf("Lista esta vazia?: %v\n", linkedList.IsEmpty())

	fmt.Printf("\n\nLimpando a lista\n\n")
	linkedList.Clear()

	fmt.Printf("Tamanho da lista: %v\n", linkedList.GetSize())

	fmt.Printf("Lista esta vazia?: %v\n", linkedList.IsEmpty())

	linkedList.AddAtEnd(1)
	linkedList.AddAtEnd(2)
	linkedList.AddAtEnd(3)
	linkedList.AddAtEnd(4)
	linkedList.AddAtEnd(5)

	linkedList.AddAtStart(90)
	linkedList.PrintLinkedList()
	fmt.Printf("Tamanho da lista: %v\n", linkedList.GetSize())

	fmt.Printf("Lista esta vazia?: %v\n", linkedList.IsEmpty())
}

type Node struct {
	value int
	next  *Node
}

func NewNode(value int) *Node {
	return &Node{
		value: value,
	}
}

type LinkedList struct {
	head *Node
	size int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) AddAtEnd(value int) {
	node := NewNode(value)

	// tratar se a lista está vazia
	if l.head == nil {
		l.head = node
		l.size++
		return
	}

	// pega o nó atual e percorre a lista até achar o final, que vai ser nulo
	current := l.head
	for current.next != nil {
		current = current.next
	}

	current.next = node
	l.size++
}

func (l *LinkedList) PrintLinkedList() {
	if l.head == nil {
		fmt.Printf("Lista vazia!\n")
		return
	}

	current := l.head

	for current != nil {
		fmt.Printf("%v\n", current.value)
		current = current.next
	}
}

func (l *LinkedList) GetSize() int {
	return l.size
}

func (l *LinkedList) IsEmpty() bool {
	return l.head == nil
}

func (l *LinkedList) Clear() {
	l.head = nil
	l.size = 0
}

func (l *LinkedList) AddAtStart(value int) {
	node := NewNode(value)

	// tratar se a lista está vazia
	// mesmo passo do método adicionar ao fim
	if l.IsEmpty() {
		l.head = node
		l.size++
		return
	}

	// guardar a referencia em uma var auxiliar
	aux := l.head

	l.head = node
	node.next = aux
	l.size++
}
