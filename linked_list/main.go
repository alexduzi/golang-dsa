package main

import (
	"fmt"
	"strings"
)

type LinkedList struct {
	head *Node
}

type Node struct {
	data string
	node *Node
}

func (ll *LinkedList) printList() {
	sb := strings.Builder{}
	current := ll.head
	for current != nil {
		sb.WriteString(current.data)
		sb.WriteString("--> ")
		current = current.node
	}
	fmt.Println(sb.String())
}

func (ll *LinkedList) Length() (count int) {
	current := ll.head
	for current != nil {
		count++
		current = current.node
	}
	return
}

func (ll *LinkedList) insertBeginning(data string) {
	newNode := &Node{data, ll.head}
	ll.head = newNode
}

func (ll *LinkedList) insertEnding(data string) {
	newNode := &Node{data, nil}
	if ll.head.node == nil {
		ll.head.node = newNode
		return
	}

	// loop the entire list until reach at the end
	// when you reach the end the node will be nil
	current := ll.head
	for current.node != nil {
		current = current.node
	}
	current.node = newNode
}

func (ll *LinkedList) removeFirst() {
	if ll.head == nil {
		return
	}
	ll.head = ll.head.node
}

func (ll *LinkedList) removeLast() {
	if ll.head == nil || ll.head.node == nil {
		return
	}
	previous := ll.head
	for previous.node.node != nil {
		previous = previous.node
	}
	previous.node = nil
}

func main() {
	linkedList := &LinkedList{&Node{"node1", &Node{"node2", &Node{"node2", &Node{"node4", &Node{"node5", nil}}}}}}

	linkedList.printList()
	fmt.Printf("List length: %d\n", linkedList.Length())
	linkedList.insertBeginning("test1")
	linkedList.insertBeginning("test2")
	linkedList.printList()
	fmt.Printf("List length: %d\n", linkedList.Length())
	linkedList.insertEnding("ending")
	linkedList.printList()
	linkedList.insertBeginning("another one")
	linkedList.printList()
	linkedList.removeFirst()
	linkedList.printList()
	linkedList.removeLast()
	linkedList.printList()
}
