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
	next *Node
}

func (ll *LinkedList) printList() {
	sb := strings.Builder{}
	current := ll.head
	for current != nil {
		sb.WriteString(current.data)
		sb.WriteString("--> ")
		current = current.next
	}
	fmt.Println(sb.String())
}

func (ll *LinkedList) Length() (count int) {
	current := ll.head
	for current != nil {
		count++
		current = current.next
	}
	return
}

func (ll *LinkedList) insertBeginning(data string) {
	newNode := &Node{data, ll.head}
	ll.head = newNode
}

func (ll *LinkedList) insertEnding(data string) {
	newNode := &Node{data, nil}
	if ll.head.next == nil {
		ll.head.next = newNode
		return
	}

	// loop the entire list until reach at the end
	// when you reach the end the node will be nil
	current := ll.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func (ll *LinkedList) removeFirst() {
	if ll.head == nil {
		return
	}
	ll.head = ll.head.next
}

func (ll *LinkedList) removeLast() {
	if ll.head == nil || ll.head.next == nil {
		return
	}
	previous := ll.head
	for previous.next.next != nil {
		previous = previous.next
	}
	previous.next = nil
}

func (ll *LinkedList) insertAtPosition(data string, position int) {
	newNode := &Node{data: data}
	if position == 1 {
		newNode.next = ll.head
		ll.head = newNode
		return
	}

	previous := ll.head
	count := 1

	for count < position-1 {
		previous = previous.next
		count++
	}

	current := previous.next
	newNode.next = current
	previous.next = newNode
}

func (ll *LinkedList) removeAtPosition(position int) {
	if position <= 0 {
		return
	}

	if position == 1 {
		ll.head = ll.head.next
	}

	previous := ll.head
	count := 1

	for count < position-1 {
		previous = previous.next
		count++
	}

	current := previous.next
	previous.next = current.next
	current = nil
}

func (ll *LinkedList) checkLoop() bool {
	fastPointer := ll.head
	slowPointer := ll.head

	for fastPointer != nil && fastPointer.next != nil {
		fastPointer = fastPointer.next.next
		slowPointer = slowPointer.next
		if slowPointer == fastPointer {
			return true
		}
	}

	return false
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
	linkedList.insertAtPosition("insertingPos3", 3)
	linkedList.printList()
	linkedList.removeAtPosition(3)
	linkedList.printList()
}
