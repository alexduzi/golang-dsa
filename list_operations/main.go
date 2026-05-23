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

	fmt.Printf("Nó da posicao 3: %+v\n", linkedList.GetNode(3))

	linkedList.AddAtPosition(89, 2)
	linkedList.PrintLinkedList()
	fmt.Printf("Lista esta vazia?: %v\n", linkedList.IsEmpty())

	fmt.Printf("Nó da posicao 2: %+v\n", linkedList.GetNode(2))
	fmt.Printf("Tamanho da lista: %v\n", linkedList.GetSize())
	fmt.Printf("Posicao do elemento 89: %v\n", linkedList.IndexOf(89))
	fmt.Printf("Posicao do elemento 190: %v\n", linkedList.IndexOf(190))

	fmt.Printf("Removendo o elemento da posicao 4: %v\n", linkedList.RemoveAtPosition(4))
	linkedList.PrintLinkedList()
	fmt.Printf("Tamanho da lista: %v\n", linkedList.GetSize())
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
	if l.IsEmpty() {
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
	if l.IsEmpty() {
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

func (l *LinkedList) GetNode(index int) *Node {
	// verifica se o parametro index não ultrapassa a lista
	if index < 0 && index > l.GetSize() {
		return nil
	}

	current := l.head
	i := 0

	for i != index {
		current = current.next
		i++
	}

	if current != nil {
		return current
	}

	return nil
}

func (l *LinkedList) AddAtPosition(value, index int) {
	// se for primeira posicao
	if index == 0 {
		l.AddAtStart(value)
		return
	}

	// se nao encontrar a posicao adiciona no final
	if l.GetNode(index) == nil {
		l.AddAtEnd(value)
		return
	}

	node := NewNode(value)

	// encontra o elemento com a posicao
	aux := l.GetNode(index - 1)

	node.next = aux.next

	aux.next = node
	l.size++
}

func (l *LinkedList) IndexOf(value int) int {
	if l.IsEmpty() {
		return -1
	}

	current := l.head
	pos := 0
	for current != nil {
		if current.value == value {
			return pos
		}
		current = current.next
		pos++
	}

	return -1
}

func (l *LinkedList) Contains(value int) bool {
	index := l.IndexOf(value)
	if index != -1 {
		return true
	}

	return false
}

func (l *LinkedList) RemoveAtPosition(index int) int {
	if l.IsEmpty() || l.GetNode(index) == nil {
		return -1
	}

	var item int
	var aux *Node

	// remover do inicio
	if index == 0 {
		item = l.head.value
		l.head = l.head.next // o root vira o proximo
		l.size--
		return item
	}

	// remover do final
	if index == l.GetSize()-1 {
		aux = l.GetNode(index - 1)
		item = aux.next.value
		aux.next = nil
		l.size--
		return item
	}

	// remover no meio da lista
	aux = l.GetNode(index - 1)
	item = aux.next.value
	aux.next = aux.next.next
	l.size--

	return item
}
