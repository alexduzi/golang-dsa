package main

func main() {

}

type Position[T any] interface {
	element() T
}

type Node[T any] struct {
	_element T
	parent   *Node[T]
	children []*Node[T]
}

func NewNode[T any](element T, parent *Node[T]) *Node[T] {
	return &Node[T]{
		_element: element,
		parent:   parent,
		children: []*Node[T]{},
	}
}

func (node *Node[T]) element() T {
	return node._element
}

func (node *Node[T]) getParent() *Node[T] {
	return node.parent
}

func (node *Node[T]) getChildren() []*Node[T] {
	return node.children
}

func (node *Node[T]) addChild(child *Node[T]) {
	node.children = append(node.children, child)
	child.parent = node
}

func (node *Node[T]) removeChild(child *Node[T]) {
	children := make([]*Node[T], 0, len(node.children)-1)
	for _, c := range node.children {
		if c != child {
			children = append(children, c)
		}
	}
	child.parent = nil
	node.children = children
}

func (node *Node[T]) setElement(element T) {
	node._element = element
}

func (node *Node[T]) isLeaf() bool {
	return len(node.children) == 0
}
