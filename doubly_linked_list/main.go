package main

import "fmt"

func main() {
	list := NewDoublyLinkedList()
	list.AddAtEnd(20)
	list.AddAtEnd(9)
	list.AddAtEnd(86)
	list.AddAtEnd(-2)
	list.AddAtEnd(16)
	list.AddAtEnd(23)

	fmt.Printf("%v\n", list)
}

type Node struct {
	Value int
	Next  *Node
	Prev  *Node
}

func NewNode(value int) *Node {
	return &Node{
		Value: value,
	}
}

type DoublyLinkedList struct {
	Head *Node
	Tail *Node
	Size int32
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

func (list *DoublyLinkedList) IsEmpty() bool {
	return list.Head == nil
}

func (list *DoublyLinkedList) AddAtEnd(value int) {
	node := NewNode(value)

	if list.IsEmpty() {
		list.Head = node
		list.Tail = node
		list.Size++
		return
	}

	node.Prev = list.Tail
	list.Tail.Next = node
	list.Tail = node
	list.Size++
}
