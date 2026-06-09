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

	tree.Remove(33)
	for _, key := range tree.Keys() {
		fmt.Println(key)
	}
	fmt.Println(tree.StringFormat())
}

type BinarySearchTreeSet[K cmp.Ordered] struct {
	root *Node[K]
	size int
}

func NewBinarySearchTreeSet[K cmp.Ordered]() *BinarySearchTreeSet[K] {
	sentinel := NewNodeSentinel[K](nil)
	return &BinarySearchTreeSet[K]{root: sentinel}
}

func NewBinarySearchTreeSetWithCollection[K cmp.Ordered](collection []K) (*BinarySearchTreeSet[K], error) {
	sentinel := NewNodeSentinel[K](nil)
	bTree := &BinarySearchTreeSet[K]{root: sentinel}
	err := bTree.AddAll(collection)
	if err != nil {
		return nil, err
	}
	return bTree, nil
}

func (tree *BinarySearchTreeSet[K]) Size() int {
	return tree.size
}

func (tree *BinarySearchTreeSet[K]) IsEmpty() bool {
	return tree.size == 0
}

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

func (tree *BinarySearchTreeSet[K]) Keys() []K {
	return tree.collectKeys(tree.root, make([]K, 0))
}

func (tree *BinarySearchTreeSet[K]) Contains(key K) bool {
	node := tree.findKeyLocation(tree.root, key)
	return !node.isSentinel()
}

func (tree *BinarySearchTreeSet[K]) StringFormat() string {
	return tree.stringFormatHelper(tree.root, 0, &strings.Builder{})
}

func (tree *BinarySearchTreeSet[K]) Remove(key K) (bool, error) {
	var zero K
	if key == zero {
		return false, errors.New("key cannot be null or empty")
	}

	nodeToRemove := tree.findKeyLocation(tree.root, key)

	if nodeToRemove.isSentinel() {
		return false, nil
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

// percurso interfixo, retorna os elementos ordenados
func (tree *BinarySearchTreeSet[K]) collectKeys(node *Node[K], keys []K) []K {
	if !node.isSentinel() {
		keys = tree.collectKeys(node.left, keys)
		keys = append(keys, node.key)
		keys = tree.collectKeys(node.right, keys)
	}
	return keys
}

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

type Node[K cmp.Ordered] struct {
	key      K
	parent   *Node[K]
	left     *Node[K]
	right    *Node[K]
	sentinel bool
}

func NewNode[K cmp.Ordered](key K, parent *Node[K]) *Node[K] {
	return &Node[K]{
		key:    key,
		parent: parent,
	}
}

func NewNodeSentinel[K cmp.Ordered](parent *Node[K]) *Node[K] {
	return &Node[K]{
		sentinel: true,
		parent:   parent,
	}
}

func (node *Node[K]) isSentinel() bool {
	return node.sentinel
}
