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

	fmt.Printf("%v\n", list.ToArray())

	list.AddAtPosition(99, 3)
	fmt.Printf("%v\n", list.ToArray())

	list.RemoveFirst()
	fmt.Printf("%v\n", list.ToArray())

	list.RemoveLast()
	fmt.Printf("%v\n", list.ToArray())
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
	Size int
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

func (list *DoublyLinkedList) IsEmpty() bool {
	return list.Head == nil
}

func (list *DoublyLinkedList) GetSize() int {
	return list.Size
}

func (list *DoublyLinkedList) Clear() {
	list.Head = nil
	list.Tail = nil
	list.Size = 0
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

func (list *DoublyLinkedList) ToArray() (result []int) {
	current := list.Head
	result = make([]int, 0, list.Size)

	for current != nil {
		result = append(result, current.Value)
		current = current.Next
	}

	return
}

func (list *DoublyLinkedList) AddAtStart(value int) {
	node := NewNode(value)

	if list.IsEmpty() {
		list.Head = node
		list.Tail = node
		list.Size++
		return
	}

	aux := list.Head
	list.Head = node
	node.Next = aux
	list.Size++
}

func (list *DoublyLinkedList) GetNode(index int) *Node {
	if index < 0 || index > list.GetSize() {
		return nil
	}

	current := list.Head
	currIndex := 0

	for current != nil && currIndex < index {
		current = current.Next
		currIndex++
	}

	return current
}

func (list *DoublyLinkedList) GetValue(index int) int {
	current := list.GetNode(index)
	if current != nil {
		return current.Value
	}
	return -1
}

func (list *DoublyLinkedList) AddAtPosition(value, index int) {
	if index == 0 {
		list.AddAtStart(value)
		return
	}

	if list.GetNode(index) == nil || index == list.GetSize() {
		list.AddAtEnd(value)
		return
	}

	node := NewNode(value)

	// pega o anterior do desejado
	current := list.GetNode(index - 1)
	current.Next.Prev = node
	node.Next = current.Next
	node.Prev = current
	current.Next = node
	list.Size++
}

func (list *DoublyLinkedList) IndexOf(value int) int {
	current := list.Head
	index := 0

	for current != nil {
		if current.Value == value {
			return index
		}
		current = current.Next
		index++
	}

	return -1
}

func (list *DoublyLinkedList) RemoveFirst() *Node {
	aux := list.Head
	list.Head = list.Head.Next

	if list.Head == nil {
		list.Tail = nil
	} else {
		list.Head.Prev = nil
	}
	list.Size--
	return aux
}

func (list *DoublyLinkedList) RemoveLast() *Node {
	aux := list.Tail
	list.Tail = list.Tail.Prev

	if list.Tail == nil {
		list.Head = nil
	} else {
		list.Tail.Next = nil
	}
	list.Size--
	return aux
}

func (list *DoublyLinkedList) RemoveAtPosition(index int) int {
	if list.IsEmpty() || list.GetNode(index) != nil {
		return -1
	}

	if index == 0 {
		return list.RemoveFirst().Value
	}

	if index == list.GetSize()-1 {
		return list.RemoveLast().Value
	}

	current := list.GetNode(index)

	current.Prev.Next = current.Next
	current.Next.Prev = current.Prev
	list.Size--
	return current.Value
}
