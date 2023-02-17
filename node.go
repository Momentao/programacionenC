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

// Construct n