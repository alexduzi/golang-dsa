package main

import "cmp"

func main() {

}

type BinarySearchTreeSet[K cmp.Ordered] struct {
	root *Node[K]
	size int
}

func NewBinarySearchTreeSet[K cmp.Ordered]() *BinarySearchTreeSet[K] {
	sentinel := &Node[K]{sentinel: true}
	return &BinarySearchTreeSet[K]{root: sentinel}
}

func (tree *BinarySearchTreeSet[K]) Size() int {
	return tree.size
}

func (tree *BinarySearchTreeSet[K]) IsEmpty() bool {
	return tree.size == 0
}

func (tree *BinarySearchTreeSet[K]) Add() bool {
	return tree.size == 0
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

func (node *Node[K]) isSentinel() bool {
	return node.sentinel
}
