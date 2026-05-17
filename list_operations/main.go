package main

import "fmt"

func main() {
	linkedList := NewLinkedList()
	linkedList.AddAtEnd(1)
	linkedList.AddAtEnd(2)
	linkedList.AddAtEnd(3)
	linkedList.AddAtEnd(4)
	linkedList.AddAtEnd(5)

	// fmt.Printf("%v\n", linkedList.Head.Next.Value)
	linkedList.PrintList()
}

type Node struct {
	Value int
	Next  *Node
}

func NewNode(value int) *Node {
	return &Node{
		Value: value,
	}
}

type LinkedList struct {
	Head *Node
	Size int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) AddAtEnd(value int) {
	node := NewNode(value)

	// tratar se a lista está vazia
	if l.Head == nil {
		l.Head = node
		l.Size++
		return
	}

	// pega o nó atual e percorre a lista até achar o final, que vai ser nulo
	current := l.Head
	for current.Next != nil {
		current = current.Next
	}

	current.Next = node
	l.Size++
}

func (l *LinkedList) PrintList() {
	if l.Head == nil {
		fmt.Println("list is empty")
		return
	}

	current := l.Head
	fmt.Println(current.Value)
	for current.Next != nil {
		current = current.Next
		fmt.Println(current.Value)
	}
}
