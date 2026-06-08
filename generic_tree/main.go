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

	isExternal, _ := tree.IsExternal(node)
	isRoot, _ := tree.IsRoot(node)
	fmt.Printf("Aplicações node IsExternal: %v\n", isExternal)
	fmt.Printf("Aplicações node IsRoot: %v\n", isRoot)

	node1 := tree.Find("Livro Azul")
	isExternalNode1, _ := tree.IsExternal(node1)
	isRootNode1, _ := tree.IsRoot(node1)
	fmt.Printf("Livro Azul node IsExternal: %v\n", isExternalNode1)
	fmt.Printf("Livro Azul node IsRoot: %v\n", isRootNode1)

	fmt.Printf("Generic Tree size: %d\n", tree.Size())
	fmt.Printf("Generic Tree is empty?: %v\n", tree.IsEmpty())
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

// Position representa uma posição válida dentro da árvore,
// expondo apenas o elemento armazenado naquela posição.
type Position[T comparable] interface {
	element() T
}

// Node representa um nó da árvore genérica, armazenando um valor
// e referências para o nó pai e seus filhos.
type Node[T comparable] struct {
	value    T
	parent   *Node[T]
	children []*Node[T]
}

// NewNode cria e retorna um novo nó com o valor e pai informados.
func NewNode[T comparable](value T, parent *Node[T]) *Node[T] {
	return &Node[T]{
		value:    value,
		parent:   parent,
		children: make([]*Node[T], 0),
	}
}

// element retorna o valor armazenado no nó, satisfazendo a interface Position.
func (node *Node[T]) element() T {
	return node.value
}

// getParent retorna o nó pai, ou nil caso seja a raiz.
func (node *Node[T]) getParent() *Node[T] {
	return node.parent
}

// getChildren retorna a lista de filhos do nó.
func (node *Node[T]) getChildren() []*Node[T] {
	return node.children
}

// addChild adiciona um nó filho e atualiza o pai do filho para o nó atual.
func (node *Node[T]) addChild(child *Node[T]) {
	node.children = append(node.children, child)
	child.parent = node
}

// removeChild remove o filho informado e define seu pai como nil.
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

// setElement substitui o valor armazenado no nó.
func (node *Node[T]) setElement(value T) {
	node.value = value
}

// isLeaf retorna true se o nó não possui filhos.
func (node *Node[T]) isLeaf() bool {
	return len(node.children) == 0
}

// GenericTree é uma árvore genérica que armazena elementos do tipo T.
// Cada nó pode ter zero ou mais filhos.
type GenericTree[T comparable] struct {
	root *Node[T]
	size int
}

// NewGenericTree cria e retorna uma árvore genérica vazia.
func NewGenericTree[T comparable]() *GenericTree[T] {
	return &GenericTree[T]{}
}

// Root retorna o nó raiz da árvore.
func (tree *GenericTree[T]) Root() *Node[T] {
	return tree.root
}

// Size retorna a quantidade de nós presentes na árvore.
func (tree *GenericTree[T]) Size() int {
	return tree.size
}

// IsEmpty retorna true se a árvore não possui nenhum nó.
func (tree *GenericTree[T]) IsEmpty() bool {
	return tree.size == 0
}

// Elements retorna todos os valores armazenados na árvore em ordem DFS pré-ordem.
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

// Positions retorna todos os nós da árvore em ordem DFS pré-ordem.
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

// Find busca e retorna o primeiro nó cujo valor seja igual ao elemento informado.
// Retorna nil se o elemento não for encontrado.
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

// validate verifica se a posição informada é um *Node[T] válido e pertencente à árvore.
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

// Add insere um novo nó com o elemento informado como filho do nó pai.
// Se parent for nil, o novo nó é inserido como raiz (somente se a árvore estiver vazia).
func (tree *GenericTree[T]) Add(element T, parent *Node[T]) (*Node[T], error) {
	if !tree.IsEmpty() && parent == nil {
		return nil, errors.New("parent position can't be null for a non-empty generic tree")
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

// Children retorna uma cópia da lista de filhos do nó na posição informada.
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

// IsExternal retorna true se o nó na posição informada não possui filhos (é uma folha).
func (tree *GenericTree[T]) IsExternal(p Position[T]) (bool, error) {
	node, err := tree.validate(p)
	if err != nil {
		return false, err
	}
	return node.isLeaf(), nil
}

// IsRoot retorna true se o nó na posição informada é a raiz da árvore.
func (tree *GenericTree[T]) IsRoot(p Position[T]) (bool, error) {
	node, err := tree.validate(p)
	if err != nil {
		return false, err
	}
	return node == tree.root, nil
}

// Parent retorna o nó pai da posição informada.
// Retorna erro se a posição for a raiz, pois a raiz não possui pai.
func (tree *GenericTree[T]) Parent(p Position[T]) (*Node[T], error) {
	node, err := tree.validate(p)
	if err != nil {
		return nil, err
	}
	if node == tree.root {
		return nil, errors.New("the root has no parent")
	}
	return node.getParent(), nil
}

// Replace substitui o valor do nó na posição informada pelo novo elemento.
func (tree *GenericTree[T]) Replace(p Position[T], element T) error {
	node, err := tree.validate(p)
	if err != nil {
		return err
	}
	node.setElement(element)
	return nil
}

// Remove remove o nó na posição informada da árvore.
// Se for a raiz, a árvore fica vazia. Caso contrário, o nó é desvinculado do pai.
func (tree *GenericTree[T]) Remove(p Position[T]) error {
	node, err := tree.validate(p)
	if err != nil {
		return err
	}
	if node == tree.root {
		tree.root = nil
	} else {
		parent := node.getParent()
		parent.removeChild(node)
	}
	return nil
}
