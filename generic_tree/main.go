package main

import (
	"errors"
	"fmt"
	"slices"
	"strings"
)

func main() {
	tree := NewGenericTree[string]()

	root, err := tree.Add("Livro Azul", nil)
	if err != nil {
		panic(err)
	}

	intro, err := tree.Add("Introdução", root)
	if err != nil {
		panic(err)
	}
	tree.Add("Pra quem é este livro", intro)
	tree.Add("Agradecimentos", intro)

	cap1, err := tree.Add("Capítulo 1", root)
	if err != nil {
		panic(err)
	}
	tree.Add("Conceitos", cap1)
	tree.Add("Aplicações", cap1)

	cap2, err := tree.Add("Capítulo 2", root)
	if err != nil {
		panic(err)
	}
	metodos, _ := tree.Add("Métodos", cap2)
	tree.Add("Problema terreno", cap2)
	tree.Add("Problema carros", cap2)

	tree.Add("Método recursivo", metodos)
	tree.Add("Método imperativo", metodos)

	fmt.Println("PRINT DFS PRE ORDER:")
	printTree(tree)

	fmt.Println()

	fmt.Println("PRINT ELEMENTS:")
	for _, value := range tree.Elements() {
		fmt.Println(value)
	}

	fmt.Println()

	fmt.Println("PRINT POSITIONS:")
	for _, pos := range tree.Positions() {
		fmt.Println(pos.value)
	}

	fmt.Println()
	fmt.Println("PRINT FIND Aplicações:")
	node := tree.Find("Aplicações")
	fmt.Println(node.value)
}

func printTree[T comparable](tree *GenericTree[T]) {
	printTreeRecursive(tree.Root(), tree, 0)
}

func printTreeRecursive[T comparable](node *Node[T], tree *GenericTree[T], depth int) {
	spaces := "    "
	spaces = strings.Repeat(spaces, depth)
	fmt.Printf("%s%v\n", spaces, node.element())
	children, err := tree.Children(node)
	if err != nil {
		panic(err)
	}
	for _, child := range children {
		printTreeRecursive(child, tree, depth+1)
	}
}

type Position[T any] interface {
	element() T
}

type Node[T comparable] struct {
	value    T
	parent   *Node[T]
	children []*Node[T]
}

func NewNode[T comparable](value T, parent *Node[T]) *Node[T] {
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

type GenericTree[T comparable] struct {
	root *Node[T]
	size int
}

func NewGenericTree[T comparable]() *GenericTree[T] {
	return &GenericTree[T]{}
}

func (tree *GenericTree[T]) Root() *Node[T] {
	return tree.root
}

func (tree *GenericTree[T]) Size() int {
	return tree.size
}

func (tree *GenericTree[T]) IsEmpty() bool {
	return tree.size == 0
}

func (tree *GenericTree[T]) Elements() []T {
	return tree.collectElements(make([]T, 0), tree.root)
}

func (tree *GenericTree[T]) collectElements(list []T, node *Node[T]) []T {
	if node == nil {
		return list
	}

	list = append(list, node.value)

	for _, child := range node.getChildren() {
		list = tree.collectElements(list, child)
	}

	return list
}

func (tree *GenericTree[T]) Positions() []*Node[T] {
	return tree.collectPositions(make([]*Node[T], 0), tree.root)
}

func (tree *GenericTree[T]) collectPositions(list []*Node[T], node *Node[T]) []*Node[T] {
	if node == nil {
		return list
	}

	list = append(list, node)

	for _, child := range node.getChildren() {
		list = tree.collectPositions(list, child)
	}

	return list
}

func (tree *GenericTree[T]) Find(element T) *Node[T] {
	return tree.findRecursive(tree.root, element)
}

func (tree *GenericTree[T]) findRecursive(node *Node[T], target T) *Node[T] {
	if node == nil {
		return nil
	}
	if node.value == target {
		return node
	}
	for _, child := range node.getChildren() {
		found := tree.findRecursive(child, target)
		if found != nil {
			return found
		}
	}
	return nil
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

func (tree *GenericTree[T]) Add(element T, parent *Node[T]) (*Node[T], error) {
	if !tree.IsEmpty() && parent == nil {
		return nil, errors.New("Parent position can't be null for a non-empty generic tree")
	}

	if parent != nil {
		_, err := tree.validate(parent)
		if err != nil {
			return nil, err
		}
	}

	newNode := NewNode(element, parent)
	if parent == nil {
		tree.root = newNode
	} else {
		parent.addChild(newNode)
	}
	tree.size++
	return newNode, nil
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
