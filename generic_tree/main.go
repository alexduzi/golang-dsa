package main

import (
	"fmt"
	"slices"
)

func main() {
	tree := NewGenericTree[string]()

	root := tree.Add("Livro Azul", nil)

	intro := tree.Add("Introdução", root)

	tree.Add("Capítulo 1", root)
	tree.Add("Capítulo 2", root)

	tree.Add("Pra quem é este livro", intro)
	tree.Add("Agradecimentos", intro)

	printTree(tree)
}

func printTree[T any](tree *GenericTree[T]) {
	printTreeRecursive(tree.Root(), tree)
}

func printTreeRecursive[T any](node *Node[T], tree *GenericTree[T]) {
	fmt.Printf("%v\n", node.element())
	children, err := tree.Children(node)
	if err != nil {
		panic(err)
	}
	for _, child := range children {
		printTreeRecursive(child, tree)
	}
}

type Position[T any] interface {
	element() T
}

type Node[T any] struct {
	value    T
	parent   *Node[T]
	children []*Node[T]
}

func NewNode[T any](value T, parent *Node[T]) *Node[T] {
	return &Node[T]{
		value:    value,
		parent:   parent,
		children: make([]*Node[T], 0),
	}
}

func (node *Node[T]) element() T {
	return node.value
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

func (node *Node[T]) setElement(value T) {
	node.value = value
}

func (node *Node[T]) isLeaf() bool {
	return len(node.children) == 0
}

type GenericTree[T any] struct {
	root *Node[T]
	size int
}

func NewGenericTree[T any]() *GenericTree[T] {
	return &GenericTree[T]{}
}

func (tree *GenericTree[T]) Size() int {
	return tree.size
}

func (tree *GenericTree[T]) IsEmpty() bool {
	return tree.size == 0
}

func (tree *GenericTree[T]) Root() *Node[T] {
	return tree.root
}

func (tree *GenericTree[T]) validate(p Position[T]) (*Node[T], error) {
	node, ok := p.(*Node[T])
	if !ok {
		return nil, fmt.Errorf("invalid position type")
	}
	if node.parent == nil && node != tree.root {
		return nil, fmt.Errorf("position is no longer in the tree")
	}
	return node, nil
}

func (tree *GenericTree[T]) Add(element T, parent *Node[T]) *Node[T] {
	newNode := NewNode(element, parent)
	if parent == nil {
		tree.root = newNode
	} else {
		parent.addChild(newNode)
	}
	return newNode
}

func (tree *GenericTree[T]) Children(p Position[T]) ([]*Node[T], error) {
	node, err := tree.validate(p)
	if err != nil {
		return nil, err
	}
	// retorna uma cópia de children Go 1.20 e versões antigas
	// children := make([]*Node[T], 0, len(node.getChildren()))
	// copy(children, node.getChildren())
	return slices.Clone(node.getChildren()), nil
}
