package main

import (
	"cmp"
	"errors"
	"fmt"
	"strings"
)

func main() {
	// tree := NewBinarySearchTreeSet[int]()
	// tree.Add(52)
	// tree.Add(17)
	// tree.Add(67)

	collection := []int{52, 17, 67, 11, 33, 55, 83, 14, 31, 46, 23, 26}
	tree, err := NewBinarySearchTreeSetWithCollection(collection)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Size: %d\n", tree.Size())
	fmt.Printf("IsEmpty: %v\n", tree.IsEmpty())
	fmt.Printf("Contains(11): %v\n", tree.Contains(11))
	fmt.Println()

	for _, key := range tree.Keys() {
		fmt.Println(key)
	}

	fmt.Println()

	fmt.Println(tree.StringFormat())

	fmt.Println()

	fmt.Println(tree.StringFormat())

	tree.Remove(17)
	fmt.Println(tree.StringFormat())

	tree.Remove(67)
	fmt.Println(tree.StringFormat())
}

// BinarySearchTreeSet é uma árvore binária de busca que funciona como um conjunto,
// ou seja, não permite chaves duplicadas.
// O constraint cmp.Ordered garante que o tipo K suporta os operadores <, <=, >, >=,
// necessários para navegar e comparar chaves na árvore sem precisar de uma função de comparação externa.
type BinarySearchTreeSet[K cmp.Ordered] struct {
	root *Node[K]
	size int
}

// NewBinarySearchTreeSet cria e retorna uma nova árvore binária de busca vazia.
func NewBinarySearchTreeSet[K cmp.Ordered]() *BinarySearchTreeSet[K] {
	sentinel := NewNodeSentinel[K](nil)
	return &BinarySearchTreeSet[K]{root: sentinel}
}

// NewBinarySearchTreeSetWithCollection cria uma nova árvore binária de busca e insere todos os elementos da coleção fornecida.
func NewBinarySearchTreeSetWithCollection[K cmp.Ordered](collection []K) (*BinarySearchTreeSet[K], error) {
	sentinel := NewNodeSentinel[K](nil)
	bTree := &BinarySearchTreeSet[K]{root: sentinel}
	err := bTree.AddAll(collection)
	if err != nil {
		return nil, err
	}
	return bTree, nil
}

// Size retorna o número de elementos presentes na árvore.
func (tree *BinarySearchTreeSet[K]) Size() int {
	return tree.size
}

// IsEmpty retorna verdadeiro se a árvore não contiver nenhum elemento.
func (tree *BinarySearchTreeSet[K]) IsEmpty() bool {
	return tree.size == 0
}

// AddAll insere todos os elementos da coleção na árvore. Retorna erro se algum elemento for inválido.
func (tree *BinarySearchTreeSet[K]) AddAll(collection []K) error {
	var err error = nil
	for _, value := range collection {
		err = tree.Add(value)
		if err != nil {
			return err
		}
	}
	return err
}

// Add insere uma nova chave na árvore. Chaves duplicadas ou zero são ignoradas/rejeitadas.
func (tree *BinarySearchTreeSet[K]) Add(key K) error {
	var zero K
	if key == zero {
		return errors.New("key cannot be null or empty")
	}

	if tree.IsEmpty() {
		tree.root = NewNode(key, nil)
		tree.root.left = NewNodeSentinel(tree.root)
		tree.root.right = NewNodeSentinel(tree.root)
		tree.size++
		return nil
	}

	node := tree.findKeyLocation(tree.root, key)
	if node.isSentinel() {
		parent := node.parent
		newNode := NewNode(key, parent)
		newNode.left = NewNodeSentinel(newNode)
		newNode.right = NewNodeSentinel(newNode)

		if node == parent.left {
			parent.left = newNode
		} else if node == parent.right {
			parent.right = newNode
		}
		tree.size++
	}

	return nil
}

// Keys retorna todas as chaves da árvore em ordem crescente (percurso em-ordem).
func (tree *BinarySearchTreeSet[K]) Keys() []K {
	return tree.collectKeys(tree.root, make([]K, 0))
}

// Contains retorna verdadeiro se a chave fornecida estiver presente na árvore.
func (tree *BinarySearchTreeSet[K]) Contains(key K) bool {
	node := tree.findKeyLocation(tree.root, key)
	return !node.isSentinel()
}

// StringFormat retorna uma representação textual da árvore rotacionada 90° para a esquerda.
func (tree *BinarySearchTreeSet[K]) StringFormat() string {
	return tree.stringFormatHelper(tree.root, 0, &strings.Builder{})
}

// Remove exclui a chave da árvore. Retorna verdadeiro se encontrada e removida, falso se não existir.
func (tree *BinarySearchTreeSet[K]) Remove(key K) (bool, error) {
	var zero K
	if key == zero {
		return false, errors.New("key cannot be null or empty")
	}

	nodeToRemove := tree.findKeyLocation(tree.root, key)

	if nodeToRemove.isSentinel() {
		return false, nil
	}

	if !nodeToRemove.left.isSentinel() && !nodeToRemove.right.isSentinel() {
		sucessor := tree.findMin(nodeToRemove.right)
		nodeToRemove.key = sucessor.key
		nodeToRemove = sucessor
	}

	var child *Node[K]

	if nodeToRemove.left.isSentinel() {
		child = nodeToRemove.right
	} else {
		child = nodeToRemove.left
	}

	child.parent = nodeToRemove.parent

	if nodeToRemove.parent == nil {
		tree.root = child
	} else if nodeToRemove == nodeToRemove.parent.left {
		nodeToRemove.parent.left = child
	} else {
		nodeToRemove.parent.right = child
	}
	tree.size--
	return true, nil
}

// Union retorna uma nova árvore contendo todos os elementos de ambas as árvores (sem duplicatas).
func (tree *BinarySearchTreeSet[K]) Union(other *BinarySearchTreeSet[K]) *BinarySearchTreeSet[K] {
	bTreeResult := NewBinarySearchTreeSet[K]()
	for _, value := range other.Keys() {
		bTreeResult.Add(value)
	}
	for _, value := range tree.Keys() {
		bTreeResult.Add(value)
	}
	return bTreeResult
}

// Intersection retorna uma nova árvore contendo apenas os elementos presentes em ambas as árvores.
func (tree *BinarySearchTreeSet[K]) Intersection(other *BinarySearchTreeSet[K]) *BinarySearchTreeSet[K] {
	bTreeResult := NewBinarySearchTreeSet[K]()
	for _, value := range tree.Keys() {
		if other.Contains(value) {
			bTreeResult.Add(value)
		}
	}
	return bTreeResult
}

// Difference retorna uma nova árvore com os elementos presentes em tree mas ausentes em other.
func (tree *BinarySearchTreeSet[K]) Difference(other *BinarySearchTreeSet[K]) *BinarySearchTreeSet[K] {
	bTreeResult := NewBinarySearchTreeSet[K]()
	for _, value := range tree.Keys() {
		if !other.Contains(value) {
			bTreeResult.Add(value)
		}
	}
	return bTreeResult
}

// findMin retorna o nó com a menor chave na subárvore enraizada em node.
func (tree *BinarySearchTreeSet[K]) findMin(node *Node[K]) *Node[K] {
	for !node.left.isSentinel() {
		node = node.left
	}
	return node
}

// stringFormatHelper constrói recursivamente a representação textual da árvore com indentação por profundidade.
func (tree *BinarySearchTreeSet[K]) stringFormatHelper(node *Node[K], depth int, sb *strings.Builder) string {
	if !node.isSentinel() {
		tree.stringFormatHelper(node.right, depth+1, sb)
		spaces := strings.Repeat("        ", depth)
		parent := ""
		if depth > 0 {
			parent = fmt.Sprintf("%v", node.parent.key)
		}
		sb.WriteString(fmt.Sprintf("%s(%v)%s\n", spaces, node.key, parent))
		tree.stringFormatHelper(node.left, depth+1, sb)
	}
	return sb.String()
}

// collectKeys percorre a árvore em ordem (esquerda, raiz, direita) e acumula as chaves ordenadas.
func (tree *BinarySearchTreeSet[K]) collectKeys(node *Node[K], keys []K) []K {
	if !node.isSentinel() {
		keys = tree.collectKeys(node.left, keys)
		keys = append(keys, node.key)
		keys = tree.collectKeys(node.right, keys)
	}
	return keys
}

// findKeyLocation busca a posição da chave na árvore. Retorna o nó se encontrado ou o sentinela onde deveria estar.
func (tree *BinarySearchTreeSet[K]) findKeyLocation(node *Node[K], key K) *Node[K] {
	for !node.isSentinel() {
		if key == node.key {
			return node
		} else if key < node.key {
			node = node.left
		} else {
			node = node.right
		}
	}
	return node
}

// Node representa um nó da árvore binária de busca, podendo ser um nó real ou um sentinela (folha nula).
// Usa o mesmo constraint cmp.Ordered da árvore para manter compatibilidade de tipos com BinarySearchTreeSet.
type Node[K cmp.Ordered] struct {
	key      K
	parent   *Node[K]
	left     *Node[K]
	right    *Node[K]
	sentinel bool
}

// NewNode cria um novo nó com a chave e o pai fornecidos.
func NewNode[K cmp.Ordered](key K, parent *Node[K]) *Node[K] {
	return &Node[K]{
		key:    key,
		parent: parent,
	}
}

// NewNodeSentinel cria um nó sentinela (nulo) usado como marcador de folha na árvore.
func NewNodeSentinel[K cmp.Ordered](parent *Node[K]) *Node[K] {
	return &Node[K]{
		sentinel: true,
		parent:   parent,
	}
}

// isSentinel retorna verdadeiro se o nó for um sentinela (nó nulo/folha).
func (node *Node[K]) isSentinel() bool {
	return node.sentinel
}
