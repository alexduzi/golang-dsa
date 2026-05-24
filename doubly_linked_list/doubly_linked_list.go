package main

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
