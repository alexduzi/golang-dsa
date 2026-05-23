package main

import "fmt"

func main() {
	linkedList := NewLinkedList()
	linkedList.AddAtEnd(1)
	linkedList.AddAtEnd(2)
	linkedList.AddAtEnd(3)
	linkedList.AddAtEnd(4)
	linkedList.AddAtEnd(5)

	linkedList.PrintLinkedList()

	fmt.Printf("Tamanho da lista: %v", linkedList.GetSize())
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
