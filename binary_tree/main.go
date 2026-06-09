package main

import (
	"cmp"
	"errors"
	"fmt"
)

func main() {
	tree := NewBinarySearchTreeSet[int]()
	tree.Add(53)
	tree.Add(67)

	fmt.Println(tree.Size())
	fmt.Println(tree.IsEmpty())
}

type BinarySearchTreeSet[K cmp.Ordered] struct {
	root *Node[K]
	size int
}

func NewBinarySearchTreeSet[K cmp.Ordered]() *BinarySearchTreeSet[K] {
	sentinel := NewNodeSentinel[K](nil)
	return &BinarySearchTreeSet[K]{root: sentinel}
}

func (tree *BinarySearchTreeSet[K]) Size() int {
	return tree.size
}

func (tree *BinarySearchTreeSet[K]) IsEmpty() bool {
	return tree.size == 0
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
