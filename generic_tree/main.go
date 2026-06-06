package main

func main() {

}

type Position[T any] interface {
	element()
}

type Node[T any] struct {
	element  T
	parent   *Node[T]
	children []*Node[T]
}

func NewNode[T any](element T, parent *Node[T]) *Node[T] {
	return &Node[T]{
		element:  element,
		parent:   parent,
		children: []*Node[T]{},
	}
}

func (node *Node[T]) Element() T {
	return node.element
}

func (node *Node[T]) isLeaf() bool {
	return len(node.children) == 0
}
