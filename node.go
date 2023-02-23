package main

// This file implements the Nodes of a filesystem tree.

import (
	"fmt"
	"io"
)

// Node in the Library's file tree.
type Node struct {
	payload  interface{}
	parent   *Node
	children []*Node
}

// Walks through the tree and applies function f to each Node.
func (this *Node) Walk(f func(*Node)) {
	f(this)
	for _, c := range this.children {
		c.Walk(f)
	}
}

// Construct new node with given parent
// and link the parent-child pointers.
func (parent *Node) NewChild(file string) (child *Node) {
	child = &Node{file, parent, nil}
	parent.children = append(parent.children, child)
	return
}

// Returns full path represented by this node.
func (n *Node) Strin